package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdklambdagoalpha/v2"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type AppServerlessCdkGoStackProps struct {
	awscdk.StackProps
}

func NewAppServerlessCdkGoStack(scope constructs.Construct, id string, props *AppServerlessCdkGoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	// create Lambda function
	getHandler := awscdklambdagoalpha.NewGoFunction(stack, jsii.String("CheckoutDocker"), &awscdklambdagoalpha.GoFunctionProps{
		Runtime: awslambda.Runtime_GO_1_X(),
		Entry:   jsii.String("../Golang/cmd/checkout/main.go"),
		Bundling: &awscdklambdagoalpha.BundlingOptions{
			DockerImage: awscdk.DockerImage_FromBuild(jsii.String("./"), &awscdk.DockerBuildOptions{}),
		
		},
	})

	// create API Gateway
	restApi := awsapigateway.NewRestApi(stack, jsii.String("goApi"), &awsapigateway.RestApiProps{
		RestApiName:    jsii.String("Checkout"),
		CloudWatchRole: jsii.Bool(false),
	})

	// create API Gateway resource
	restApi.Root().AddResource(jsii.String("checkout"), &awsapigateway.ResourceOptions{}).AddMethod(
		jsii.String("POST"),
		awsapigateway.NewLambdaIntegration(getHandler, &awsapigateway.LambdaIntegrationOptions{}),
		restApi.Root().DefaultMethodOptions(),
	)

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewAppServerlessCdkGoStack(app, "goStack-", &AppServerlessCdkGoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return nil
}