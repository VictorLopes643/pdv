package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
    // "encoding/json"
	"github.com/aws/aws-lambda-go/events"
    "context"
    // "fmt"
    "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/appsync"
)




type InputEvent  struct {
    Title *string `json:"title"`
    Price *float64 `json:"price"`
}
type outputEvent struct {
	Message string `json:"message"`
}

type Query struct {
    Query string                 `json:"query"`
    Variables map[string]interface{} `json:"variables,omitempty"`
}


func main() {
	lambda.Start(Handler)
}



func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

   // Cria uma sessão AWS
   sess := session.Must(session.NewSessionWithOptions(session.Options{
    Config: aws.Config{
        Region: aws.String("us-west-1"),
    },
    Profile: "soohc",
}))

// Define a consulta
query := `query MyQuery {
    listProdutos {
      items {
        name
        price
      }
    }
  }
}`

// Cria um cliente AppSync
client := appsync.New(sess)



// Executa a consulta
resp, err := client.ListGraphqlApis(client, query)
if err != nil {
    log.Fatal(err)
}

// Imprime a resposta
fmt.Println(resp)

    // fmt.Println(request.Body)
    
    // Imprime a resposta
    // fmt.Println(resp)
    // var event InputEvent
    // err := json.Unmarshal([]byte(request.Body), &event)
    // if err != nil {
    //     fmt.Println(err)
    //     return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 400}, nil
    // }

    // fmt.Println("Hello World!")
	// fmt.Println(*event.Title)
    // fmt.Println(*event.Price)
    return events.APIGatewayProxyResponse{Body: "It worked!", StatusCode: 200}, nil
}