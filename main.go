package main

import (
	"GoBlogLambda/dynamo"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	BlogPostId string
}

type MyResponse struct {
	Message string
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	var cucker = dynamo.GetBlogLocation()
	return MyResponse{Message: cucker}, nil
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
