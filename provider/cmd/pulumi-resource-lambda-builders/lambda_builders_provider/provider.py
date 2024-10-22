#  Copyright 2016-2024, Pulumi Corporation.
#
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

from typing import Mapping, Any

from pulumi.provider.provider import InvokeResult
import pulumi.provider as provider

import lambda_builders_provider
from lambda_builders_provider.build_go import BuildGoArgs, build_go


class Provider(provider.Provider):

    def __init__(self) -> None:
        super().__init__(
            lambda_builders_provider.__version__, lambda_builders_provider.__schema__
        )

    def invoke(self, token: str, args: Mapping[str, Any]) -> InvokeResult:
        if token == "lambda-builders:index:buildGo":
            return build_go(BuildGoArgs.from_inputs(args))

        raise Exception(f"Unknown function {token}")
