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

// GetWFActionValidationByIDReader is a Reader for the GetWFActionValidationByID structure.
type GetWFActionValidationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWFActionValidationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWFActionValidationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetWFActionValidationByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetWFActionValidationByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetWFActionValidationByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWFActionValidationByIDOK creates a GetWFActionValidationByIDOK with default headers values
func NewGetWFActionValidationByIDOK() *GetWFActionValidationByIDOK {
	return &GetWFActionValidationByIDOK{}
}

/*GetWFActionValidationByIDOK handles this case with default header values.

Generic String answer
*/
type GetWFActionValidationByIDOK struct {
	Payload string
}

func (o *GetWFActionValidationByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/validations/{validation-id}][%d] getWFActionValidationByIdOK  %+v", 200, o.Payload)
}

func (o *GetWFActionValidationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWFActionValidationByIDUnauthorized creates a GetWFActionValidationByIDUnauthorized with default headers values
func NewGetWFActionValidationByIDUnauthorized() *GetWFActionValidationByIDUnauthorized {
	return &GetWFActionValidationByIDUnauthorized{}
}

/*GetWFActionValidationByIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetWFActionValidationByIDUnauthorized struct {
}

func (o *GetWFActionValidationByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/validations/{validation-id}][%d] getWFActionValidationByIdUnauthorized ", 401)
}

func (o *GetWFActionValidationByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWFActionValidationByIDForbidden creates a GetWFActionValidationByIDForbidden with default headers values
func NewGetWFActionValidationByIDForbidden() *GetWFActionValidationByIDForbidden {
	return &GetWFActionValidationByIDForbidden{}
}

/*GetWFActionValidationByIDForbidden handles this case with default header values.

403 Forbidden
*/
type GetWFActionValidationByIDForbidden struct {
}

func (o *GetWFActionValidationByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/validations/{validation-id}][%d] getWFActionValidationByIdForbidden ", 403)
}

func (o *GetWFActionValidationByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWFActionValidationByIDInternalServerError creates a GetWFActionValidationByIDInternalServerError with default headers values
func NewGetWFActionValidationByIDInternalServerError() *GetWFActionValidationByIDInternalServerError {
	return &GetWFActionValidationByIDInternalServerError{}
}

/*GetWFActionValidationByIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetWFActionValidationByIDInternalServerError struct {
}

func (o *GetWFActionValidationByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/actions/{action-id}/validations/{validation-id}][%d] getWFActionValidationByIdInternalServerError ", 500)
}

func (o *GetWFActionValidationByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
