// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new operations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for operations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetConfig get config API
*/
func (a *Client) GetConfig(params *GetConfigParams) (*GetConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getConfig",
		Method:             "GET",
		PathPattern:        "/config",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetConfigReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetConfigOK), nil

}

/*
GetHealth Returns the '/health'
*/
func (a *Client) GetHealth(params *GetHealthParams) (*GetHealthOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetHealthParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getHealth",
		Method:             "GET",
		PathPattern:        "/api/v1.0/health",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetHealthReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetHealthOK), nil

}

/*
GetJoinScripts Get join script for node
*/
func (a *Client) GetJoinScripts(params *GetJoinScriptsParams) (*GetJoinScriptsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetJoinScriptsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getJoinScripts",
		Method:             "GET",
		PathPattern:        "/api/v1.0/join-scripts",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetJoinScriptsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetJoinScriptsOK), nil

}

/*
GetVersions Returns list of all supported versions of Promenade. Currently this returns a static value.
*/
func (a *Client) GetVersions(params *GetVersionsParams) (*GetVersionsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetVersionsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getVersions",
		Method:             "GET",
		PathPattern:        "/versions",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetVersionsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetVersionsOK), nil

}

/*
ProbeLiveness probe liveness API
*/
func (a *Client) ProbeLiveness(params *ProbeLivenessParams) (*ProbeLivenessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProbeLivenessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "probeLiveness",
		Method:             "GET",
		PathPattern:        "/liveness",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProbeLivenessReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProbeLivenessOK), nil

}

/*
ProbeReadiness probe readiness API
*/
func (a *Client) ProbeReadiness(params *ProbeReadinessParams) (*ProbeReadinessOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProbeReadinessParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "probeReadiness",
		Method:             "GET",
		PathPattern:        "/readiness",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProbeReadinessReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProbeReadinessOK), nil

}

/*
UpdateNodeLabels Update Node Labels
*/
func (a *Client) UpdateNodeLabels(params *UpdateNodeLabelsParams) (*UpdateNodeLabelsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateNodeLabelsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateNodeLabels",
		Method:             "PUT",
		PathPattern:        "/api/v1.0/node-labels/{node-name}",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateNodeLabelsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateNodeLabelsOK), nil

}

/*
ValidateDesign Validate documents
*/
func (a *Client) ValidateDesign(params *ValidateDesignParams) (*ValidateDesignOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewValidateDesignParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "validateDesign",
		Method:             "POST",
		PathPattern:        "/api/v1.0/validatedesign",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ValidateDesignReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ValidateDesignOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
