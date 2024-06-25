from unittest import TestCase
from typing import cast
import os

from build_go import BuildGoArgs, build_go
from pulumi.asset import FileArchive


class TestLambdaBuilder(TestCase):
    TEST_DATA_FOLDER = os.path.join(
        os.path.dirname(__file__), "../../../../examples/simple-go/cmd/simple"
    )

    def test_success(self):
        res = build_go(BuildGoArgs(code=self.TEST_DATA_FOLDER, architecture="arm64"))
        asset = cast(FileArchive, res.outputs["asset"])
        files = os.listdir(asset.path)
        self.assertEqual(files[0], "bootstrap")

    def test_file_does_not_exist_error(self):
        res = build_go(BuildGoArgs(code="does-not-exist"))
        self.assertEqual(res.outputs, {})
        self.assertIsNotNone(res.failures)
        if res.failures is not None:
            self.assertEqual(
                res.failures[0].reason,
                "GoModulesBuilder:Build - [Errno 2] No such file or directory: 'does-not-exist'",
            )

    def test_invalid_architecture(self):
        res = build_go(BuildGoArgs(code=self.TEST_DATA_FOLDER, architecture="abc"))
        self.assertEqual(res.outputs, {})
        self.assertIsNotNone(res.failures)
        if res.failures is not None:
            self.assertEqual(
                res.failures[0].reason,
                "GoModulesBuilder:Validation - Architecture abc is not supported for runtime provided",
            )
