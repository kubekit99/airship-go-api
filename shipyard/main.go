package main

import (
	"log"

	loads "github.com/go-openapi/loads"
	"github.com/kubekit99/airship-go-api/shipyard/restapi"
	"github.com/kubekit99/airship-go-api/shipyard/restapi/operations"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewShipyardAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.ConfigureFlags()
	server.Port = 8081

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
