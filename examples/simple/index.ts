import * as xyz from "@pulumi/lambda-builders";
import * as path from 'path';
import * as aws from '@pulumi/aws';

const builder = xyz.buildGo({
    code: path.join(__dirname, 'cmd'),
    architecture: 'x86_64',
});

const role = new aws.iam.Role('example-role', {
  assumeRolePolicy: aws.iam.assumeRolePolicyForPrincipal(
    aws.iam.Principals.LambdaPrincipal,
  ),
});

new aws.lambda.Function('handler', {
    code: builder.then(val => val.asset),
    handler: 'bootstrap',
    role: role.arn,
    runtime: 'provided.al2023',
});
