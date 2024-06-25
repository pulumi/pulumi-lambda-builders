# Copyright 2016-2024, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from enum import Enum
import os
from typing import Any, Optional
import tempfile
from aws_lambda_builders.builder import LambdaBuilder
from pulumi.asset import FileArchive
from pulumi.provider.provider import CheckFailure, InvokeResult
from aws_lambda_builders.exceptions import (
    LambdaBuilderError,
    UnsupportedArchitectureError,
)


class Architecture(Enum):
    ARM_64 = "arm64"
    X86_64 = "x86_64"


class BuildGoArgs:
    code: str
    """The path to the code to build"""

    architecture: Optional[str]
    """The Lambda architecture to build for"""

    @staticmethod
    def from_inputs(inputs: Any) -> "BuildGoArgs":
        return BuildGoArgs(code=inputs["code"], architecture=inputs["architecture"])

    def __init__(self, code: str, architecture: Optional[str] = None) -> None:
        self.code = code
        self.architecture = architecture


class BuildGoResult:
    asset: FileArchive
    """The built code asset"""

    def __init__(self, asset: FileArchive) -> None:
        self.asset = asset

    @staticmethod
    def to_invoke_results(result: "BuildGoResult") -> InvokeResult:
        return InvokeResult(
            outputs={
                "asset": result.asset,
            }
        )


def build_go(args: BuildGoArgs) -> InvokeResult:
    builder = LambdaBuilder("go", "modules", None)
    tmp_dir = tempfile.mkdtemp()
    arch = Architecture.X86_64.value
    if args.architecture != None:
        arch = args.architecture

    try:
        builder.build(
            source_dir=args.code,
            artifacts_dir=tmp_dir,
            scratch_dir=tempfile.gettempdir(),
            # manifest_path is a required argument, but is not
            # actually used by the go builder so it doesn't _really_ matter
            # what this value is.
            manifest_path=os.path.join(args.code, "go.mod"),
            runtime="provided",
            architecture=arch,
        )
    except UnsupportedArchitectureError as err:
        print(err)
        return InvokeResult(
            outputs={},
            failures=[CheckFailure(property="architecture", reason=err.__str__())],
        )
    # The only two input properties are code & architecture & lambda_builders only throws a specific
    # error for architecture. The rest of the errors we can return as generic errors
    except LambdaBuilderError as err:
        return InvokeResult(
            outputs={}, failures=[CheckFailure(property="", reason=err.__str__())]
        )

    return InvokeResult(outputs={"asset": FileArchive(tmp_dir)})
