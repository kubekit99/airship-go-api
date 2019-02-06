// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/runtime/yamlpc"

	"github.com/kubekit99/airship-go-api/armada/restapi/operations"
)

//go:generate swagger generate server --target ../../armada --name Armada --spec ../swagger/swagger.yml --exclude-main

func configureFlags(api *operations.ArmadaAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.ArmadaAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.YamlConsumer = yamlpc.YAMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetHealthHandler = operations.GetHealthHandlerFunc(func(params operations.GetHealthParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetHealth has not yet been implemented")
	})
	api.GetReleasesHandler = operations.GetReleasesHandlerFunc(func(params operations.GetReleasesParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetReleases has not yet been implemented")
	})
	api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetStatus has not yet been implemented")
	})
	api.GetVersionsHandler = operations.GetVersionsHandlerFunc(func(params operations.GetVersionsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetVersions has not yet been implemented")
	})
	api.PostApplyManifestHandler = operations.PostApplyManifestHandlerFunc(func(params operations.PostApplyManifestParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostApplyManifest has not yet been implemented")
	})
	api.PostRollbackReleaseNameHandler = operations.PostRollbackReleaseNameHandlerFunc(func(params operations.PostRollbackReleaseNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostRollbackReleaseName has not yet been implemented")
	})
	api.PostTestReleaseNameHandler = operations.PostTestReleaseNameHandlerFunc(func(params operations.PostTestReleaseNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostTestReleaseName has not yet been implemented")
	})
	api.PostTestsHandler = operations.PostTestsHandlerFunc(func(params operations.PostTestsParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostTests has not yet been implemented")
	})
	api.PostValidateDesignHandler = operations.PostValidateDesignHandlerFunc(func(params operations.PostValidateDesignParams) middleware.Responder {
		return middleware.NotImplemented("operation .PostValidateDesign has not yet been implemented")
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
