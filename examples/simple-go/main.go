package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/iam"
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/lambda"
	lambdabuilders "github.com/pulumi/pulumi-lambda-builders/sdk/go/lambda-builders"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		wd, _ := os.Getwd()
		path := filepath.Join(wd, "cmd/simple")
		log.New(os.Stderr, "", 0).Println(path)
		build, err := lambdabuilders.BuildGo(ctx, &lambdabuilders.BuildGoArgs{
			Architecture: pulumi.StringRef("arm64"),
			Code:         pulumi.StringRef(path),
		})
		if err != nil {
			return err
		}
		doc, err := iam.GetPolicyDocument(ctx, &iam.GetPolicyDocumentArgs{
			Statements: []iam.GetPolicyDocumentStatement{
				{
					Actions: []string{"sts:AssumeRole"},
					Effect:  pulumi.StringRef("Allow"),
					Principals: []iam.GetPolicyDocumentStatementPrincipal{
						{
							Type:        "Service",
							Identifiers: []string{"lambda.amazonaws.com"},
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		role, err := iam.NewRole(ctx, "example-role", &iam.RoleArgs{
			AssumeRolePolicy: pulumi.String(doc.Json),
		})
		if err != nil {
			return err
		}
		f, err := lambda.NewFunction(ctx, "example-function", &lambda.FunctionArgs{
			Architectures: pulumi.ToStringArray([]string{"arm64"}),
			Code:          build.Asset,
			Handler:       pulumi.String("bootstrap"),
			Role:          role.Arn,
			Runtime:       pulumi.String(lambda.RuntimeCustomAL2023),
		})
		if err != nil {
			return err
		}

		// invoke the function to ensure that bundling worked correctly
		res, err := lambda.NewInvocation(ctx, "invoke-function", &lambda.InvocationArgs{
			FunctionName: f.Name,
			Input:        pulumi.String(""),
		})
		if err != nil {
			return err
		}
		ctx.Export("result", res.Result)
		return nil
	})
}
