// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Builds a Golang Lambda Function into a Pulumi Asset that can be deployed.
 *
 * The below example uses a folder structure like this:
 *
 * The output of `buildGo` produces an asset that can be passed to the
 * `aws.Lambda` `Code` property.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as lambda_builders from "@pulumi/lambda-builders";
 *
 * const builder = lambda_builders.buildGo({
 *     architecture: "arm64",
 *     code: "cmd/simple",
 * });
 * const lambdaRolePolicy = aws.iam.getPolicyDocumentOutput({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["lambda.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const role = new aws.iam.Role("role", {
 *    assumeRolePolicy: lambdaRolePolicy.apply(lambdaRolePolicy => lambdaRolePolicy.json),
 * });
 * new aws.lambda.Function("function", {
 *     code: builder.asset,
 *     architectures: ["arm64"],
 *     handler: "bootstrap",
 *     role: role.arn,
 *     runtime: aws.lambda.Runtime.CustomAL2023,
 * });
 * ```
 */
export function buildGo(args?: BuildGoArgs, opts?: pulumi.InvokeOptions): Promise<BuildGoResult> {
    args = args || {};

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("lambda-builders:index:buildGo", {
        "architecture": args.architecture,
        "code": args.code,
    }, opts);
}

export interface BuildGoArgs {
    /**
     * Lambda function architecture to build for. Valid values are `"x86_64"` and `"arm64"`. Default is `"x86_64"`.
     */
    architecture?: string;
    /**
     * The path to the go code to build
     */
    code?: string;
}

export interface BuildGoResult {
    /**
     * The archive that contains the golang binary that will be deployed to the Lambda Function.
     */
    readonly asset?: pulumi.asset.Archive;
}
/**
 * Builds a Golang Lambda Function into a Pulumi Asset that can be deployed.
 *
 * The below example uses a folder structure like this:
 *
 * The output of `buildGo` produces an asset that can be passed to the
 * `aws.Lambda` `Code` property.
 *
 * ## Example Usage
 *
 * Basic usage:
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 * import * as lambda_builders from "@pulumi/lambda-builders";
 *
 * const builder = lambda_builders.buildGo({
 *     architecture: "arm64",
 *     code: "cmd/simple",
 * });
 * const lambdaRolePolicy = aws.iam.getPolicyDocumentOutput({
 *     statements: [{
 *         actions: ["sts:AssumeRole"],
 *         principals: [{
 *             type: "Service",
 *             identifiers: ["lambda.amazonaws.com"],
 *         }],
 *     }],
 * });
 * const role = new aws.iam.Role("role", {
 *    assumeRolePolicy: lambdaRolePolicy.apply(lambdaRolePolicy => lambdaRolePolicy.json),
 * });
 * new aws.lambda.Function("function", {
 *     code: builder.asset,
 *     architectures: ["arm64"],
 *     handler: "bootstrap",
 *     role: role.arn,
 *     runtime: aws.lambda.Runtime.CustomAL2023,
 * });
 * ```
 */
export function buildGoOutput(args?: BuildGoOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<BuildGoResult> {
    return pulumi.output(args).apply((a: any) => buildGo(a, opts))
}

export interface BuildGoOutputArgs {
    /**
     * Lambda function architecture to build for. Valid values are `"x86_64"` and `"arm64"`. Default is `"x86_64"`.
     */
    architecture?: pulumi.Input<string>;
    /**
     * The path to the go code to build
     */
    code?: pulumi.Input<string>;
}
