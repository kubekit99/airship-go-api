// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"fmt"

	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	swag "github.com/go-openapi/swag"

	"github.com/kubekit99/airship-go-api/promenade/restapi/operations"
)

//go:generate swagger generate server --target ../../promenade --name Promenade --spec ../swagger/swagger.yml

func configureFlags(api *operations.PromenadeAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.PromenadeAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.TxtProducer = runtime.TextProducer()

	api.GetConfigHandler = operations.GetConfigHandlerFunc(
		func(params operations.GetConfigParams) middleware.Responder {
			name := swag.StringValue(params.Name)
			if name == "" {
				name = "all"
			}

			configfile := fmt.Sprintf("Config for %s!", name)
			return operations.NewGetConfigOK().WithPayload(configfile)
		})
	api.ProbeLivenessHandler = operations.ProbeLivenessHandlerFunc(
		func(params operations.ProbeLivenessParams) middleware.Responder {
			return operations.NewProbeLivenessOK().WithPayload("promenade alive")
		})
	api.ProbeReadinessHandler = operations.ProbeReadinessHandlerFunc(
		func(params operations.ProbeReadinessParams) middleware.Responder {
			return operations.NewProbeReadinessOK().WithPayload("promenade ready")
		})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
