# Pulumi Lambda Builders

## Background

TODO

## Prerequisites

- Pulumi CLI
- Python 3.6+
- Node.js
- Yarn
- Go 1.17
- Node.js (to build the Node SDK)
- .NET Code SDK (to build the .NET SDK)

## Build and Test

### Regenerate SDKs
make generate

### Build and install the provider and SDKs
make build
make install

### Ensure the pulumi-provider-xyz script is on PATH (for testing)
$ export PATH=$PATH:$PWD/bin

### Test Node.js SDK
$ cd examples/simple
$ yarn install
$ yarn link @pulumi/xyz
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up

