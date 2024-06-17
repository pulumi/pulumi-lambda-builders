PROVIDER_VERSION ?= 1.0.0-alpha.0+dev
VERSION_GENERIC = $(shell pulumictl convert-version --language generic --version "$(PROVIDER_VERSION)")

PACK            := lambda-builders
PACK_           := lambda_builders
PROJECT         := github.com/pulumi/pulumi-${PACK}

PROVIDER        := pulumi-resource-${PACK}
CODEGEN         := pulumi-gen-${PACK}
VERSION_PATH    := provider/pkg/version.Version

WORKING_DIR     := $(shell pwd)
SCHEMA_PATH     := ${WORKING_DIR}/schema.json

SRC             := provider/cmd/pulumi-resource-${PACK}

# Need to pick up locally pinned pulumi-language-* plugins
export PULUMI_IGNORE_AMBIENT_PLUGINS = true
export GOPATH := $(shell go env GOPATH)

# Ensure the codegen file is present so that the hard-coded "Tar provider binaries" step doesn't fail
codegen: .pulumi/bin/pulumi # Required by CI
	mkdir -p bin && touch bin/pulumi-gen-lambda-builders

provider: build_provider # Required by CI
test_provider: # Required by CI
generate_schema: # Required by CI
local_generate: generate # Required by CI

generate:: gen_go_sdk gen_dotnet_sdk gen_nodejs_sdk gen_python_sdk generate_java

build:: build_provider build_dotnet_sdk build_nodejs_sdk build_python_sdk

install:: install_dotnet_sdk install_nodejs_sdk

ensure:: tidy

tidy: tidy_provider tidy_examples
	cd sdk && go mod tidy

tidy_examples:
	cd examples && go mod tidy

tidy_provider:
	cd provider && go mod tidy


# Provider

PROVIDER_FILES =  bin/PulumiPlugin.yaml bin/requirements.txt bin/run-provider.py
PROVIDER_FILES += bin/pulumi-resource-${PACK}.cmd bin/pulumi-resource-${PACK}

build_provider::	bin/venv bin/${PACK}-provider ${PROVIDER_FILES}

bin/venv:		${SRC}/requirements.txt
	rm -rf $@
	python3 -m venv $@
	./bin/venv/bin/python -m pip install -r $<

bin/${PACK}-provider:	${SRC}/	${SRC}/${PACK_}_provider/VERSION
	rm -rf $@
	cp ${WORKING_DIR}/schema.json ${SRC}/${PACK_}_provider/schema.json
	./bin/venv/bin/python -m pip install --no-deps provider/cmd/pulumi-resource-${PACK}/ -t bin/ --upgrade

bin/PulumiPlugin.yaml:			${SRC}/PulumiPlugin.yaml
bin/requirements.txt:			${SRC}/requirements.txt
bin/pulumi-resource-${PACK}.cmd:	${SRC}/pulumi-resource-${PACK}.cmd
bin/pulumi-resource-${PACK}:		${SRC}/pulumi-resource-${PACK}
bin/run-provider.py:			${SRC}/run-provider.py

bin/%:
	cp -f $< $@

${SRC}/${PACK_}_provider/VERSION:
	echo "${VERSION_GENERIC}" > ${SRC}/${PACK_}_provider/VERSION

# Go SDK

gen_go_sdk: .pulumi/bin/pulumi
	rm -rf sdk/go
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language go --version ${VERSION_GENERIC}
build_go_sdk::
generate_go: gen_go_sdk # Required by CI
build_go: # Required by CI
install_go_sdk:: # Required by CI


# .NET SDK

gen_dotnet_sdk: .pulumi/bin/pulumi
	rm -rf sdk/dotnet
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language dotnet --version ${VERSION_GENERIC}

build_dotnet_sdk:: DOTNET_VERSION := ${VERSION}
build_dotnet_sdk:: gen_dotnet_sdk
	cd sdk/dotnet/ && \
		echo "${DOTNET_VERSION}" >version.txt && \
		dotnet build

install_dotnet_sdk::
	rm -rf ${WORKING_DIR}/nuget
	mkdir -p ${WORKING_DIR}/nuget
	find . -name '*.nupkg' -print -exec cp -p {} ${WORKING_DIR}/nuget \;

generate_dotnet: gen_dotnet_sdk # Required by CI
build_dotnet: build_dotnet_sdk # Required by CI


# Node.js SDK

gen_nodejs_sdk: .pulumi/bin/pulumi
	rm -rf sdk/nodejs
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language nodejs --version ${VERSION_GENERIC}

build_nodejs_sdk:: gen_nodejs_sdk
	cd sdk/nodejs/ && \
		yarn install && \
		yarn run tsc --version && \
		yarn run tsc && \
		cp ../../README.md ../../LICENSE package.json yarn.lock ./bin/

generate_nodejs: gen_nodejs_sdk # Required by CI
build_nodejs: build_nodejs_sdk # Required by CI
install_nodejs_sdk:: build_nodejs_sdk
	yarn unlink ${PACK} || true
	yarn link --cwd ${WORKING_DIR}/sdk/nodejs/bin


# Python SDK

gen_python_sdk: .pulumi/bin/pulumi
	rm -rf sdk/python
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language python --version ${VERSION_GENERIC}
	cp ${WORKING_DIR}/README.md sdk/python

build_python_sdk:: PYPI_VERSION := ${VERSION}
build_python_sdk:: gen_python_sdk
	cd sdk/python/ && \
		rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
		python3 -m venv venv && \
		./venv/bin/python -m pip install build && \
		cd ./bin && ../venv/bin/python -m build .

generate_python: gen_python_sdk # Required by CI
build_python: build_python_sdk # Required by CI
install_python_sdk:: # Required by CI

# Java SDK

generate_java: .pulumi/bin/pulumi # Required by CI
	.pulumi/bin/pulumi package gen-sdk ${SCHEMA_PATH} --language java --version ${VERSION_GENERIC}
	cp ${WORKING_DIR}/README.md sdk/java
build_java: # Required by CI
	cd sdk/java && gradle --console=plain build

install_java_sdk: # Required by CI

# Output tarballs for plugin distribution. Example use:
#
# pulumi plugin install resource lambda-builders 0.0.1 --file pulumi-resource-lambda-builders-v0.0.1-linux-amd64.tar.gz

dist::	build_provider
	rm -rf dist
	mkdir -p dist
	(cd bin && tar --gzip --exclude venv --exclude pulumi-resource-${PACK}.cmd -cf ../dist/pulumi-resource-${PACK}-v${VERSION}-linux-amd64.tar.gz .)
	cp dist/pulumi-resource-${PACK}-v${VERSION}-linux-amd64.tar.gz dist/pulumi-resource-${PACK}-v${VERSION}-darwin-amd64.tar.gz
	cp dist/pulumi-resource-${PACK}-v${VERSION}-linux-amd64.tar.gz dist/pulumi-resource-${PACK}-v${VERSION}-darwin-arm64.tar.gz
	(cd bin && tar --gzip --exclude venv --exclude pulumi-resource-${PACK} -cf ../dist/pulumi-resource-${PACK}-v${VERSION}-windows-amd64.tar.gz .)

# Pulumi for codegen
.pulumi/bin/pulumi: PULUMI_VERSION := $(shell cat .pulumi.version)
.pulumi/bin/pulumi: HOME := $(WORKING_DIR)
.pulumi/bin/pulumi: .pulumi.version
	curl -fsSL https://get.pulumi.com | sh -s -- -version "$(PULUMI_VERSION)"
