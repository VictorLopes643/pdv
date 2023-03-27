package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
    "encoding/json"
	"github.com/aws/aws-lambda-go/events"

)




type InputEvent  struct {
    Title *string `json:"title"`
    Price *float64 `json:"price"`
}
type outputEvent struct {
	Message string `json:"message"`
}

func main() {
	lambda.Start(Handler)
}


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
    fmt.Println(request.Body)
    
    var event InputEvent
    err := json.Unmarshal([]byte(request.Body), &event)
    if err != nil {
        fmt.Println(err)
        return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
    }

    fmt.Println("Hello World!")
	fmt.Println(*event.Title)
    fmt.Println(*event.Price)
    return events.APIGatewayProxyResponse{Body: "It worked!", StatusCode: 200}, nil
}