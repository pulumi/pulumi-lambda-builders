// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambdabuilders

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-lambda-builders/sdk/go/lambda-builders/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Builds a Golang Lambda Function into a Pulumi Asset that can be deployed.
//
// The below example uses a folder structure like this:
//
// The output of `buildGo` produces an asset that can be passed to the
// `aws.Lambda` `Code` property.
//
// ## Example Usage
//
// Basic usage:
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
//	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
//	lambdabuilders "github.com/pulumi/pulumi-lambda-builders/sdk/go/lambda-builders"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			builder, err := lambdabuilders.BuildGo(ctx, &lambdabuilders.BuildGoArgs{
//				Architecture: pulumi.StringRef("arm64"),
//				Code:         pulumi.StringRef("cmd/simple"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			lambdaRolePolicy, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
//				Statements: []iam.GetPolicyDocumentStatement{
//					{
//						Actions: []string{
//							"sts:AssumeRole",
//						},
//						Principals: []iam.GetPolicyDocumentStatementPrincipal{
//							{
//								Type: "Service",
//								Identifiers: []string{
//									"lambda.amazonaws.com",
//								},
//							},
//						},
//					},
//				},
//			}, nil)
//			if err != nil {
//				return err
//			}
//			role, err := iam.NewRole(ctx, "role", &iam.RoleArgs{
//				AssumeRolePolicy: pulumi.String(lambdaRolePolicy.Json),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = lambda.NewFunction(ctx, "function", &lambda.FunctionArgs{
//				Code: builder.Asset,
//				Architectures: pulumi.StringArray{
//			pulumi.String("arm64"),
//		},
//				Handler: pulumi.String("bootstrap"),
//				Role:    role.Arn,
//				Runtime: pulumi.String(lambda.RuntimeCustomAL2023),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func BuildGo(ctx *pulumi.Context, args *BuildGoArgs, opts ...pulumi.InvokeOption) (*BuildGoResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv BuildGoResult
	err := ctx.Invoke("lambda-builders:index:buildGo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type BuildGoArgs struct {
	// Lambda function architecture to build for. Valid values are `"x86_64"` and `"arm64"`. Default is `"x86_64"`.
	Architecture *string `pulumi:"architecture"`
	// The path to the go code to build
	Code *string `pulumi:"code"`
}

type BuildGoResult struct {
	// The archive that contains the golang binary that will be deployed to the Lambda Function.
	Asset pulumi.Archive `pulumi:"asset"`
}

func BuildGoOutput(ctx *pulumi.Context, args BuildGoOutputArgs, opts ...pulumi.InvokeOption) BuildGoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BuildGoResultOutput, error) {
			args := v.(BuildGoArgs)
			opts = internal.PkgInvokeDefaultOpts(opts)
			var rv BuildGoResult
			secret, err := ctx.InvokePackageRaw("lambda-builders:index:buildGo", args, &rv, "", opts...)
			if err != nil {
				return BuildGoResultOutput{}, err
			}

			output := pulumi.ToOutput(rv).(BuildGoResultOutput)
			if secret {
				return pulumi.ToSecret(output).(BuildGoResultOutput), nil
			}
			return output, nil
		}).(BuildGoResultOutput)
}

type BuildGoOutputArgs struct {
	// Lambda function architecture to build for. Valid values are `"x86_64"` and `"arm64"`. Default is `"x86_64"`.
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	// The path to the go code to build
	Code pulumi.StringPtrInput `pulumi:"code"`
}

func (BuildGoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildGoArgs)(nil)).Elem()
}

type BuildGoResultOutput struct{ *pulumi.OutputState }

func (BuildGoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BuildGoResult)(nil)).Elem()
}

func (o BuildGoResultOutput) ToBuildGoResultOutput() BuildGoResultOutput {
	return o
}

func (o BuildGoResultOutput) ToBuildGoResultOutputWithContext(ctx context.Context) BuildGoResultOutput {
	return o
}

// The archive that contains the golang binary that will be deployed to the Lambda Function.
func (o BuildGoResultOutput) Asset() pulumi.ArchiveOutput {
	return o.ApplyT(func(v BuildGoResult) pulumi.Archive { return v.Asset }).(pulumi.ArchiveOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildGoResultOutput{})
}
