package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	adapter "github.com/awslabs/aws-lambda-go-api-proxy/chi"
	"github.com/go-chi/chi/v5"
)

var chiLambda *adapter.ChiLambda
var router chi.Router

func init() {
	router = chi.NewRouter()
	
	// Add your routes here
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from Lambda!"))
	})

	chiLambda = adapter.New(router.(*chi.Mux))
}

func main() {
	// Check if running in Lambda environment
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") != "" {
		lambda.Start(chiLambda.ProxyWithContext)
	} else {
		// Local development server
		log.Println("Starting server on :8080...")
		http.ListenAndServe(":8080", router)
	}
} 