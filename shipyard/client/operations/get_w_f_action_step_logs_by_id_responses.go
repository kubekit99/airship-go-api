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

// GetWFActionStepLogsByIDReader is a Reader for the GetWFActionStepLogsByID structure.
type GetWFActionStepLogsByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWFActionStepLogsByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWFActionStepLogsByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWFActionStepLogsByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetWFActionStepLogsByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWFActionStepLogsByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWFActionStepLogsByIDOK creates a GetWFActionStepLogsByIDOK with default headers values
func NewGetWFActionStepLogsByIDOK() *GetWFActionStepLogsByIDOK {
	return &GetWFActionStepLogsByIDOK{}
}

/*GetWFActionStepLogsByIDOK handles this case with default header values.

Generic String answer
*/
type GetWFActionStepLogsByIDOK struct {
	Payload string
}

func (o *GetWFActionStepLogsByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/steps/{step-id}/logs][%d] getWFActionStepLogsByIdOK  %+v", 200, o.Payload)
}

func (o *GetWFActionStepLogsByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWFActionStepLogsByIDUnauthorized creates a GetWFActionStepLogsByIDUnauthorized with default headers values
func NewGetWFActionStepLogsByIDUnauthorized() *GetWFActionStepLogsByIDUnauthorized {
	return &GetWFActionStepLogsByIDUnauthorized{}
}

/*GetWFActionStepLogsByIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetWFActionStepLogsByIDUnauthorized struct {
}

func (o *GetWFActionStepLogsByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/steps/{step-id}/logs][%d] getWFActionStepLogsByIdUnauthorized ", 401)
}

func (o *GetWFActionStepLogsByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWFActionStepLogsByIDForbidden creates a GetWFActionStepLogsByIDForbidden with default headers values
func NewGetWFActionStepLogsByIDForbidden() *GetWFActionStepLogsByIDForbidden {
	return &GetWFActionStepLogsByIDForbidden{}
}

/*GetWFActionStepLogsByIDForbidden handles this case with default header values.

403 Forbidden
*/
type GetWFActionStepLogsByIDForbidden struct {
}

func (o *GetWFActionStepLogsByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/steps/{step-id}/logs][%d] getWFActionStepLogsByIdForbidden ", 403)
}

func (o *GetWFActionStepLogsByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWFActionStepLogsByIDInternalServerError creates a GetWFActionStepLogsByIDInternalServerError with default headers values
func NewGetWFActionStepLogsByIDInternalServerError() *GetWFActionStepLogsByIDInternalServerError {
	return &GetWFActionStepLogsByIDInternalServerError{}
}

/*GetWFActionStepLogsByIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetWFActionStepLogsByIDInternalServerError struct {
}

func (o *GetWFActionStepLogsByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/steps/{step-id}/logs][%d] getWFActionStepLogsByIdInternalServerError ", 500)
}

func (o *GetWFActionStepLogsByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
