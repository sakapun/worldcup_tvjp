package main

import (
  //"github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "log"
)

//func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
//  log.Print("call")
//  return &events.APIGatewayProxyResponse{
//    StatusCode: 200,
//    Body:       "Hello, World",
//  }, nil
//}
func handler(input string) (string, error) {
  log.Print("call")
  return input, nil
}

func main() {
  // Make the handler available for Remote Procedure Call by AWS Lambda
  log.Print("start")
  lambda.Start(handler)
}
