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

// GetValidationByIDReader is a Reader for the GetValidationByID structure.
type GetValidationByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetValidationByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetValidationByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetValidationByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetValidationByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetValidationByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetValidationByIDOK creates a GetValidationByIDOK with default headers values
func NewGetValidationByIDOK() *GetValidationByIDOK {
	return &GetValidationByIDOK{}
}

/*GetValidationByIDOK handles this case with default header values.

Generic String answer
*/
type GetValidationByIDOK struct {
	Payload string
}

func (o *GetValidationByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}/validations/{validation-name}][%d] getValidationByIdOK  %+v", 200, o.Payload)
}

func (o *GetValidationByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetValidationByIDUnauthorized creates a GetValidationByIDUnauthorized with default headers values
func NewGetValidationByIDUnauthorized() *GetValidationByIDUnauthorized {
	return &GetValidationByIDUnauthorized{}
}

/*GetValidationByIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetValidationByIDUnauthorized struct {
}

func (o *GetValidationByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}/validations/{validation-name}][%d] getValidationByIdUnauthorized ", 401)
}

func (o *GetValidationByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationByIDForbidden creates a GetValidationByIDForbidden with default headers values
func NewGetValidationByIDForbidden() *GetValidationByIDForbidden {
	return &GetValidationByIDForbidden{}
}

/*GetValidationByIDForbidden handles this case with default header values.

403 Forbidden
*/
type GetValidationByIDForbidden struct {
}

func (o *GetValidationByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}/validations/{validation-name}][%d] getValidationByIdForbidden ", 403)
}

func (o *GetValidationByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetValidationByIDInternalServerError creates a GetValidationByIDInternalServerError with default headers values
func NewGetValidationByIDInternalServerError() *GetValidationByIDInternalServerError {
	return &GetValidationByIDInternalServerError{}
}

/*GetValidationByIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetValidationByIDInternalServerError struct {
}

func (o *GetValidationByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}/validations/{validation-name}][%d] getValidationByIdInternalServerError ", 500)
}

func (o *GetValidationByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}