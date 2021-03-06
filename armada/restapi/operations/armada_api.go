// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	"github.com/go-openapi/runtime/yamlpc"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewArmadaAPI creates a new Armada instance
func NewArmadaAPI(spec *loads.Document) *ArmadaAPI {
	return &ArmadaAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		YamlConsumer:        yamlpc.YAMLConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		GetHealthHandler: GetHealthHandlerFunc(func(params GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHealth has not yet been implemented")
		}),
		GetReleasesHandler: GetReleasesHandlerFunc(func(params GetReleasesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetReleases has not yet been implemented")
		}),
		GetStatusHandler: GetStatusHandlerFunc(func(params GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation GetStatus has not yet been implemented")
		}),
		GetVersionsHandler: GetVersionsHandlerFunc(func(params GetVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetVersions has not yet been implemented")
		}),
		PostApplyManifestHandler: PostApplyManifestHandlerFunc(func(params PostApplyManifestParams) middleware.Responder {
			return middleware.NotImplemented("operation PostApplyManifest has not yet been implemented")
		}),
		PostRollbackReleaseNameHandler: PostRollbackReleaseNameHandlerFunc(func(params PostRollbackReleaseNameParams) middleware.Responder {
			return middleware.NotImplemented("operation PostRollbackReleaseName has not yet been implemented")
		}),
		PostTestReleaseNameHandler: PostTestReleaseNameHandlerFunc(func(params PostTestReleaseNameParams) middleware.Responder {
			return middleware.NotImplemented("operation PostTestReleaseName has not yet been implemented")
		}),
		PostTestsHandler: PostTestsHandlerFunc(func(params PostTestsParams) middleware.Responder {
			return middleware.NotImplemented("operation PostTests has not yet been implemented")
		}),
		PostValidateDesignHandler: PostValidateDesignHandlerFunc(func(params PostValidateDesignParams) middleware.Responder {
			return middleware.NotImplemented("operation PostValidateDesign has not yet been implemented")
		}),
	}
}

/*ArmadaAPI Armada provides operators a way to deploy or upgrade collection of helm
charts using a single command.
*/
type ArmadaAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer
	// YamlConsumer registers a consumer for a "application/x-yaml" mime type
	YamlConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// GetHealthHandler sets the operation handler for the get health operation
	GetHealthHandler GetHealthHandler
	// GetReleasesHandler sets the operation handler for the get releases operation
	GetReleasesHandler GetReleasesHandler
	// GetStatusHandler sets the operation handler for the get status operation
	GetStatusHandler GetStatusHandler
	// GetVersionsHandler sets the operation handler for the get versions operation
	GetVersionsHandler GetVersionsHandler
	// PostApplyManifestHandler sets the operation handler for the post apply manifest operation
	PostApplyManifestHandler PostApplyManifestHandler
	// PostRollbackReleaseNameHandler sets the operation handler for the post rollback release name operation
	PostRollbackReleaseNameHandler PostRollbackReleaseNameHandler
	// PostTestReleaseNameHandler sets the operation handler for the post test release name operation
	PostTestReleaseNameHandler PostTestReleaseNameHandler
	// PostTestsHandler sets the operation handler for the post tests operation
	PostTestsHandler PostTestsHandler
	// PostValidateDesignHandler sets the operation handler for the post validate design operation
	PostValidateDesignHandler PostValidateDesignHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *ArmadaAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ArmadaAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ArmadaAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ArmadaAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ArmadaAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ArmadaAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ArmadaAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ArmadaAPI
func (o *ArmadaAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.YamlConsumer == nil {
		unregistered = append(unregistered, "YamlConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.GetHealthHandler == nil {
		unregistered = append(unregistered, "GetHealthHandler")
	}

	if o.GetReleasesHandler == nil {
		unregistered = append(unregistered, "GetReleasesHandler")
	}

	if o.GetStatusHandler == nil {
		unregistered = append(unregistered, "GetStatusHandler")
	}

	if o.GetVersionsHandler == nil {
		unregistered = append(unregistered, "GetVersionsHandler")
	}

	if o.PostApplyManifestHandler == nil {
		unregistered = append(unregistered, "PostApplyManifestHandler")
	}

	if o.PostRollbackReleaseNameHandler == nil {
		unregistered = append(unregistered, "PostRollbackReleaseNameHandler")
	}

	if o.PostTestReleaseNameHandler == nil {
		unregistered = append(unregistered, "PostTestReleaseNameHandler")
	}

	if o.PostTestsHandler == nil {
		unregistered = append(unregistered, "PostTestsHandler")
	}

	if o.PostValidateDesignHandler == nil {
		unregistered = append(unregistered, "PostValidateDesignHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ArmadaAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ArmadaAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *ArmadaAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ArmadaAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		case "application/x-yaml":
			result["application/x-yaml"] = o.YamlConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *ArmadaAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ArmadaAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the armada API
func (o *ArmadaAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ArmadaAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/health"] = NewGetHealth(o.context, o.GetHealthHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/releases"] = NewGetReleases(o.context, o.GetReleasesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/status"] = NewGetStatus(o.context, o.GetStatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/versions"] = NewGetVersions(o.context, o.GetVersionsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/apply"] = NewPostApplyManifest(o.context, o.PostApplyManifestHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/rollback/{release_name}"] = NewPostRollbackReleaseName(o.context, o.PostRollbackReleaseNameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/test/{release_name}"] = NewPostTestReleaseName(o.context, o.PostTestReleaseNameHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/tests"] = NewPostTests(o.context, o.PostTestsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/validatedesign"] = NewPostValidateDesign(o.context, o.PostValidateDesignHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ArmadaAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ArmadaAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ArmadaAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ArmadaAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
