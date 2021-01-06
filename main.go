package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/epsagon/epsagon-go/epsagon"
)

func main() {
	log.Println("enter main")
	token := os.Getenv("EPSAGON_TOKEN")
	lambda.Start(epsagon.WrapLambdaHandler(
	  epsagon.NewTracerConfig("Epsagon Application", token),
	  myHandler))

	http.HandleFunc("/", helloServer)
	log.Println("Server started on 8080.")
	http.ListenAndServe(":8080", nil)	
}

func myHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("In myHandler, received body: ", request.Body)
	return events.APIGatewayProxyResponse{Body: request.Body, StatusCode: 200}, nil
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, fmt.Sprintf("Hello, %s!", name))
	log.Println(fmt.Sprintf("Hello, %s!", name))
}
