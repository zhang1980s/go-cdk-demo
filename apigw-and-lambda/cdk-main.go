package main

import (
	"os"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsapigateway"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type DemoStackProps struct {
	awscdk.StackProps
}

func NewDemoStack(scope constructs.Construct, id string, props *DemoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps

	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	DemoHandler := awslambda.NewFunction(stack, jsii.String("DemoFunction"), &awslambda.FunctionProps{
		Runtime:    awslambda.Runtime_GO_1_X(),
		Code:       awslambda.Code_FromAsset(jsii.String("app/bin/."), nil),
		Handler:    jsii.String("main"),
		MemorySize: jsii.Number(128),
	})

	awsapigateway.NewLambdaRestApi(stack, jsii.String("endpoint"), &awsapigateway.LambdaRestApiProps{
		Handler: DemoHandler,
	})

	return stack
}

func main() {
	app := awscdk.NewApp(nil)

	NewDemoStack(app, "DemoStack", &DemoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String(os.Getenv("CDK_DEFAULT_ACCOUNT")),
		Region:  jsii.String(os.Getenv("CDK_DEFAULT_REGION")),
	}
}
