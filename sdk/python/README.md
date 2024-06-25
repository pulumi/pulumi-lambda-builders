# Pulumi Lambda Builders

---
> [!NOTE]
> Pulumi Lambda Builders is currently experimental

---

## Background

Pulumi Lambda Builders is a library that provides utilities for easily
building/bundling Lambda Function code. The library currently supports building
`go` Lambdas, but will eventually support the below languages/build tools.

- Java with Gradle
- Java with Maven
- Dotnet with amazon.lambda.tools
- Python with Pip
- Javascript with Npm
- Typescript with esbuild
- Ruby with Bundler
- Go with Mod
- Rust with Cargo

This library integrates with the
[aws-lambda-builders](https://github.com/aws/aws-lambda-builders) library which
provides the building utilities.

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @pulumi/lambda-builders
```

or `yarn`:

```bash
yarn add @pulumi/lambda-builders
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-lambda-builders
```

### Go

To use from Go, use `go get` to grab the latest version of the library

```bash
go get github.com/pulumi/pulumi-lambda-builders/sdk
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Pulumi.LambdaBuilders
```

## References

* TODO: [Tutorial]()
* [API Reference Documentation](https://www.pulumi.com/registry/packages/lambda-builders/api-docs/)
* [Examples](./examples)
