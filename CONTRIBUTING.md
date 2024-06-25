# Contributing to the Pulumi ecosystem

Do you want to contribute to Pulumi? Awesome! We are so happy to have you. We
have a few tips and housekeeping items to help you get up and running.

## Code of Conduct

Please make sure to read and observe our [Code of
Conduct](./CODE-OF-CONDUCT.md)

## Community Expectations

Please read about our [contribution guidelines
here.](https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#communications)

## Setting up your development environment

### Pulumi prerequisites

Please refer to the [main Pulumi repo](https://github.com/pulumi/pulumi/)'s
[CONTRIBUTING.md file](
<https://github.com/pulumi/pulumi/blob/master/CONTRIBUTING.md#developing>) for
details on how to get set up with Pulumi.

### AWS Lambda Builders Prerequisites

- Pulumi CLI
- Python 3.6+
- Node.js
- Yarn
- Go 1.17
- Node.js (to build the Node SDK)
- .NET Code SDK (to build the .NET SDK)

### Build and Test

#### Build and install the provider and SDKs
make build
make install

#### Regenerate SDKs
make generate

### Ensure the pulumi-provider-lambda-builders script is on PATH (for testing)
$ export PATH=$PATH:$PWD/bin

### Test Go SDK
$ cd examples/simple-go
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up

## Committing Generated Code

You must generate and check in the SDKs on each pull request containing a code
change, e.g. adding a new resource to `schema.json`.

1. Run `make build` to re-generate the provider binary.
1. Run `make generate` to generate the SDKs
1. Open a pull request containing all changes
1. *Note:* If a large number of seemingly-unrelated diffs are produced by `make
   build` (for example, lots of changes to comments unrelated to the
   change you are making), ensure that the latest dependencies for the provider
   are installed by running `go mod tidy` in the `provider/` directory of this
   repository.

## Running Integration Tests

The examples and integration tests in this repository will create and destroy
real AWS cloud resources while running. Before running these tests, make sure
that you have [configured Pulumi with AWS](https://pulumi.io/install/aws.html)
successfully once before.

The only additional step you need to take to run tests in this repo is to set
the `AWS_REGION` environment variable to the region you'd like to create test
resources in. The integration tests do try to clean up after themselves by
deleting everything that was created, but in the event of bugs or test failures
you may need to go into the AWS Console and delete resources yourself.

Once you have set `AWS_REGION` and configured your AWS credentials, `make test`
will run your integration tests.

