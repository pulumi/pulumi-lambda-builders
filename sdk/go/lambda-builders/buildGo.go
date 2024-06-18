// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package lambdabuilders

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-lambda-builders/sdk/go/lambda-builders/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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
	Architecture *string `pulumi:"architecture"`
	Code         *string `pulumi:"code"`
}

type BuildGoResult struct {
	Asset pulumi.Archive `pulumi:"asset"`
}

func BuildGoOutput(ctx *pulumi.Context, args BuildGoOutputArgs, opts ...pulumi.InvokeOption) BuildGoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (BuildGoResult, error) {
			args := v.(BuildGoArgs)
			r, err := BuildGo(ctx, &args, opts...)
			var s BuildGoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(BuildGoResultOutput)
}

type BuildGoOutputArgs struct {
	Architecture pulumi.StringPtrInput `pulumi:"architecture"`
	Code         pulumi.StringPtrInput `pulumi:"code"`
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

func (o BuildGoResultOutput) Asset() pulumi.ArchiveOutput {
	return o.ApplyT(func(v BuildGoResult) pulumi.Archive { return v.Asset }).(pulumi.ArchiveOutput)
}

func init() {
	pulumi.RegisterOutputType(BuildGoResultOutput{})
}
