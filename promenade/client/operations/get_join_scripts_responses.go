// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// GetJoinScriptsReader is a Reader for the GetJoinScripts structure.
type GetJoinScriptsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJoinScriptsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetJoinScriptsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetJoinScriptsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetJoinScriptsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetJoinScriptsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetJoinScriptsOK creates a GetJoinScriptsOK with default headers values
func NewGetJoinScriptsOK() *GetJoinScriptsOK {
	return &GetJoinScriptsOK{}
}

/*GetJoinScriptsOK handles this case with default header values.

Generic String answer
*/
type GetJoinScriptsOK struct {
	Payload string
}

func (o *GetJoinScriptsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/join-scripts][%d] getJoinScriptsOK  %+v", 200, o.Payload)
}

func (o *GetJoinScriptsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJoinScriptsUnauthorized creates a GetJoinScriptsUnauthorized with default headers values
func NewGetJoinScriptsUnauthorized() *GetJoinScriptsUnauthorized {
	return &GetJoinScriptsUnauthorized{}
}

/*GetJoinScriptsUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetJoinScriptsUnauthorized struct {
}

func (o *GetJoinScriptsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/join-scripts][%d] getJoinScriptsUnauthorized ", 401)
}

func (o *GetJoinScriptsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJoinScriptsForbidden creates a GetJoinScriptsForbidden with default headers values
func NewGetJoinScriptsForbidden() *GetJoinScriptsForbidden {
	return &GetJoinScriptsForbidden{}
}

/*GetJoinScriptsForbidden handles this case with default header values.

403 Forbidden
*/
type GetJoinScriptsForbidden struct {
}

func (o *GetJoinScriptsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/join-scripts][%d] getJoinScriptsForbidden ", 403)
}

func (o *GetJoinScriptsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetJoinScriptsInternalServerError creates a GetJoinScriptsInternalServerError with default headers values
func NewGetJoinScriptsInternalServerError() *GetJoinScriptsInternalServerError {
	return &GetJoinScriptsInternalServerError{}
}

/*GetJoinScriptsInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetJoinScriptsInternalServerError struct {
}

func (o *GetJoinScriptsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/join-scripts][%d] getJoinScriptsInternalServerError ", 500)
}

func (o *GetJoinScriptsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
