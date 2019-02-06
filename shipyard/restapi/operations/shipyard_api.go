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

// NewShipyardAPI creates a new Shipyard instance
func NewShipyardAPI(spec *loads.Document) *ShipyardAPI {
	return &ShipyardAPI{
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
		TxtProducer:         runtime.TextProducer(),
		GetActionByActionIDHandler: GetActionByActionIDHandlerFunc(func(params GetActionByActionIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActionByActionID has not yet been implemented")
		}),
		GetActionControlHandler: GetActionControlHandlerFunc(func(params GetActionControlParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActionControl has not yet been implemented")
		}),
		GetActionStepByIDHandler: GetActionStepByIDHandlerFunc(func(params GetActionStepByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActionStepByID has not yet been implemented")
		}),
		GetActionStepLogsHandler: GetActionStepLogsHandlerFunc(func(params GetActionStepLogsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActionStepLogs has not yet been implemented")
		}),
		GetActionValidationByIDHandler: GetActionValidationByIDHandlerFunc(func(params GetActionValidationByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActionValidationByID has not yet been implemented")
		}),
		GetActionsHandler: GetActionsHandlerFunc(func(params GetActionsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetActions has not yet been implemented")
		}),
		GetCommitConfigDocsHandler: GetCommitConfigDocsHandlerFunc(func(params GetCommitConfigDocsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetCommitConfigDocs has not yet been implemented")
		}),
		GetConfigHandler: GetConfigHandlerFunc(func(params GetConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation GetConfig has not yet been implemented")
		}),
		GetConfigDocByIDHandler: GetConfigDocByIDHandlerFunc(func(params GetConfigDocByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetConfigDocByID has not yet been implemented")
		}),
		GetConfigDocsHandler: GetConfigDocsHandlerFunc(func(params GetConfigDocsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetConfigDocs has not yet been implemented")
		}),
		GetHealthHandler: GetHealthHandlerFunc(func(params GetHealthParams) middleware.Responder {
			return middleware.NotImplemented("operation GetHealth has not yet been implemented")
		}),
		GetNoteDetailsByIDHandler: GetNoteDetailsByIDHandlerFunc(func(params GetNoteDetailsByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetNoteDetailsByID has not yet been implemented")
		}),
		GetRenderedConfigDocsHandler: GetRenderedConfigDocsHandlerFunc(func(params GetRenderedConfigDocsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetRenderedConfigDocs has not yet been implemented")
		}),
		GetSiteStatusesHandler: GetSiteStatusesHandlerFunc(func(params GetSiteStatusesParams) middleware.Responder {
			return middleware.NotImplemented("operation GetSiteStatuses has not yet been implemented")
		}),
		GetStatusHandler: GetStatusHandlerFunc(func(params GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation GetStatus has not yet been implemented")
		}),
		GetVersionsHandler: GetVersionsHandlerFunc(func(params GetVersionsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetVersions has not yet been implemented")
		}),
		GetWorkflowByIDHandler: GetWorkflowByIDHandlerFunc(func(params GetWorkflowByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation GetWorkflowByID has not yet been implemented")
		}),
		GetWorkflowsHandler: GetWorkflowsHandlerFunc(func(params GetWorkflowsParams) middleware.Responder {
			return middleware.NotImplemented("operation GetWorkflows has not yet been implemented")
		}),
		PostApplyManifestHandler: PostApplyManifestHandlerFunc(func(params PostApplyManifestParams) middleware.Responder {
			return middleware.NotImplemented("operation PostApplyManifest has not yet been implemented")
		}),
		ProbeLivenessHandler: ProbeLivenessHandlerFunc(func(params ProbeLivenessParams) middleware.Responder {
			return middleware.NotImplemented("operation ProbeLiveness has not yet been implemented")
		}),
		ProbeReadinessHandler: ProbeReadinessHandlerFunc(func(params ProbeReadinessParams) middleware.Responder {
			return middleware.NotImplemented("operation ProbeReadiness has not yet been implemented")
		}),
	}
}

/*ShipyardAPI Shipyard provides operators a way to deploy or upgrade collection of helm
charts using a single command.
*/
type ShipyardAPI struct {
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

	// TxtProducer registers a producer for a "text/plain" mime type
	TxtProducer runtime.Producer

	// GetActionByActionIDHandler sets the operation handler for the get action by action Id operation
	GetActionByActionIDHandler GetActionByActionIDHandler
	// GetActionControlHandler sets the operation handler for the get action control operation
	GetActionControlHandler GetActionControlHandler
	// GetActionStepByIDHandler sets the operation handler for the get action step by Id operation
	GetActionStepByIDHandler GetActionStepByIDHandler
	// GetActionStepLogsHandler sets the operation handler for the get action step logs operation
	GetActionStepLogsHandler GetActionStepLogsHandler
	// GetActionValidationByIDHandler sets the operation handler for the get action validation by Id operation
	GetActionValidationByIDHandler GetActionValidationByIDHandler
	// GetActionsHandler sets the operation handler for the get actions operation
	GetActionsHandler GetActionsHandler
	// GetCommitConfigDocsHandler sets the operation handler for the get commit config docs operation
	GetCommitConfigDocsHandler GetCommitConfigDocsHandler
	// GetConfigHandler sets the operation handler for the get config operation
	GetConfigHandler GetConfigHandler
	// GetConfigDocByIDHandler sets the operation handler for the get config doc by Id operation
	GetConfigDocByIDHandler GetConfigDocByIDHandler
	// GetConfigDocsHandler sets the operation handler for the get config docs operation
	GetConfigDocsHandler GetConfigDocsHandler
	// GetHealthHandler sets the operation handler for the get health operation
	GetHealthHandler GetHealthHandler
	// GetNoteDetailsByIDHandler sets the operation handler for the get note details by Id operation
	GetNoteDetailsByIDHandler GetNoteDetailsByIDHandler
	// GetRenderedConfigDocsHandler sets the operation handler for the get rendered config docs operation
	GetRenderedConfigDocsHandler GetRenderedConfigDocsHandler
	// GetSiteStatusesHandler sets the operation handler for the get site statuses operation
	GetSiteStatusesHandler GetSiteStatusesHandler
	// GetStatusHandler sets the operation handler for the get status operation
	GetStatusHandler GetStatusHandler
	// GetVersionsHandler sets the operation handler for the get versions operation
	GetVersionsHandler GetVersionsHandler
	// GetWorkflowByIDHandler sets the operation handler for the get workflow by Id operation
	GetWorkflowByIDHandler GetWorkflowByIDHandler
	// GetWorkflowsHandler sets the operation handler for the get workflows operation
	GetWorkflowsHandler GetWorkflowsHandler
	// PostApplyManifestHandler sets the operation handler for the post apply manifest operation
	PostApplyManifestHandler PostApplyManifestHandler
	// ProbeLivenessHandler sets the operation handler for the probe liveness operation
	ProbeLivenessHandler ProbeLivenessHandler
	// ProbeReadinessHandler sets the operation handler for the probe readiness operation
	ProbeReadinessHandler ProbeReadinessHandler

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
func (o *ShipyardAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *ShipyardAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *ShipyardAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *ShipyardAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *ShipyardAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *ShipyardAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *ShipyardAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the ShipyardAPI
func (o *ShipyardAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.YamlConsumer == nil {
		unregistered = append(unregistered, "YamlConsumer")
	}

	if o.TxtProducer == nil {
		unregistered = append(unregistered, "TxtProducer")
	}

	if o.GetActionByActionIDHandler == nil {
		unregistered = append(unregistered, "GetActionByActionIDHandler")
	}

	if o.GetActionControlHandler == nil {
		unregistered = append(unregistered, "GetActionControlHandler")
	}

	if o.GetActionStepByIDHandler == nil {
		unregistered = append(unregistered, "GetActionStepByIDHandler")
	}

	if o.GetActionStepLogsHandler == nil {
		unregistered = append(unregistered, "GetActionStepLogsHandler")
	}

	if o.GetActionValidationByIDHandler == nil {
		unregistered = append(unregistered, "GetActionValidationByIDHandler")
	}

	if o.GetActionsHandler == nil {
		unregistered = append(unregistered, "GetActionsHandler")
	}

	if o.GetCommitConfigDocsHandler == nil {
		unregistered = append(unregistered, "GetCommitConfigDocsHandler")
	}

	if o.GetConfigHandler == nil {
		unregistered = append(unregistered, "GetConfigHandler")
	}

	if o.GetConfigDocByIDHandler == nil {
		unregistered = append(unregistered, "GetConfigDocByIDHandler")
	}

	if o.GetConfigDocsHandler == nil {
		unregistered = append(unregistered, "GetConfigDocsHandler")
	}

	if o.GetHealthHandler == nil {
		unregistered = append(unregistered, "GetHealthHandler")
	}

	if o.GetNoteDetailsByIDHandler == nil {
		unregistered = append(unregistered, "GetNoteDetailsByIDHandler")
	}

	if o.GetRenderedConfigDocsHandler == nil {
		unregistered = append(unregistered, "GetRenderedConfigDocsHandler")
	}

	if o.GetSiteStatusesHandler == nil {
		unregistered = append(unregistered, "GetSiteStatusesHandler")
	}

	if o.GetStatusHandler == nil {
		unregistered = append(unregistered, "GetStatusHandler")
	}

	if o.GetVersionsHandler == nil {
		unregistered = append(unregistered, "GetVersionsHandler")
	}

	if o.GetWorkflowByIDHandler == nil {
		unregistered = append(unregistered, "GetWorkflowByIDHandler")
	}

	if o.GetWorkflowsHandler == nil {
		unregistered = append(unregistered, "GetWorkflowsHandler")
	}

	if o.PostApplyManifestHandler == nil {
		unregistered = append(unregistered, "PostApplyManifestHandler")
	}

	if o.ProbeLivenessHandler == nil {
		unregistered = append(unregistered, "ProbeLivenessHandler")
	}

	if o.ProbeReadinessHandler == nil {
		unregistered = append(unregistered, "ProbeReadinessHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *ShipyardAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *ShipyardAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *ShipyardAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *ShipyardAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

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
func (o *ShipyardAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "text/plain":
			result["text/plain"] = o.TxtProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *ShipyardAPI) HandlerFor(method, path string) (http.Handler, bool) {
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

// Context returns the middleware context for the shipyard API
func (o *ShipyardAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *ShipyardAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions/{action-id}"] = NewGetActionByActionID(o.context, o.GetActionByActionIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions/{action-id}/control/{control-verb}"] = NewGetActionControl(o.context, o.GetActionControlHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions/{action-id}/steps/{step-id}"] = NewGetActionStepByID(o.context, o.GetActionStepByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions/{action-id}/steps/{step-id}/logs"] = NewGetActionStepLogs(o.context, o.GetActionStepLogsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions/{action-id}/validations/{validation-id}"] = NewGetActionValidationByID(o.context, o.GetActionValidationByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/actions"] = NewGetActions(o.context, o.GetActionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/commitconfigdocs"] = NewGetCommitConfigDocs(o.context, o.GetCommitConfigDocsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/config"] = NewGetConfig(o.context, o.GetConfigHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/configdocs/{collection-id}"] = NewGetConfigDocByID(o.context, o.GetConfigDocByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/configdocs"] = NewGetConfigDocs(o.context, o.GetConfigDocsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/health"] = NewGetHealth(o.context, o.GetHealthHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/notedetails/{note-id}"] = NewGetNoteDetailsByID(o.context, o.GetNoteDetailsByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/renderedconfigdocs"] = NewGetRenderedConfigDocs(o.context, o.GetRenderedConfigDocsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/site_statuses"] = NewGetSiteStatuses(o.context, o.GetSiteStatusesHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/status"] = NewGetStatus(o.context, o.GetStatusHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/versions"] = NewGetVersions(o.context, o.GetVersionsHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/workflows/{workflow-id}"] = NewGetWorkflowByID(o.context, o.GetWorkflowByIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/api/v1.0/workflows"] = NewGetWorkflows(o.context, o.GetWorkflowsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/api/v1.0/apply"] = NewPostApplyManifest(o.context, o.PostApplyManifestHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/liveness"] = NewProbeLiveness(o.context, o.ProbeLivenessHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/readiness"] = NewProbeReadiness(o.context, o.ProbeReadinessHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *ShipyardAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *ShipyardAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *ShipyardAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *ShipyardAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
