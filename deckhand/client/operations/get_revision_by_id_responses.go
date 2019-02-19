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

// GetRevisionByIDReader is a Reader for the GetRevisionByID structure.
type GetRevisionByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRevisionByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRevisionByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetRevisionByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRevisionByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRevisionByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRevisionByIDOK creates a GetRevisionByIDOK with default headers values
func NewGetRevisionByIDOK() *GetRevisionByIDOK {
	return &GetRevisionByIDOK{}
}

/*GetRevisionByIDOK handles this case with default header values.

Generic String answer
*/
type GetRevisionByIDOK struct {
	Payload string
}

func (o *GetRevisionByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}][%d] getRevisionByIdOK  %+v", 200, o.Payload)
}

func (o *GetRevisionByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRevisionByIDUnauthorized creates a GetRevisionByIDUnauthorized with default headers values
func NewGetRevisionByIDUnauthorized() *GetRevisionByIDUnauthorized {
	return &GetRevisionByIDUnauthorized{}
}

/*GetRevisionByIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetRevisionByIDUnauthorized struct {
}

func (o *GetRevisionByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}][%d] getRevisionByIdUnauthorized ", 401)
}

func (o *GetRevisionByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevisionByIDForbidden creates a GetRevisionByIDForbidden with default headers values
func NewGetRevisionByIDForbidden() *GetRevisionByIDForbidden {
	return &GetRevisionByIDForbidden{}
}

/*GetRevisionByIDForbidden handles this case with default header values.

403 Forbidden
*/
type GetRevisionByIDForbidden struct {
}

func (o *GetRevisionByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}][%d] getRevisionByIdForbidden ", 403)
}

func (o *GetRevisionByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRevisionByIDInternalServerError creates a GetRevisionByIDInternalServerError with default headers values
func NewGetRevisionByIDInternalServerError() *GetRevisionByIDInternalServerError {
	return &GetRevisionByIDInternalServerError{}
}

/*GetRevisionByIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetRevisionByIDInternalServerError struct {
}

func (o *GetRevisionByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/revisions/{revision-id}][%d] getRevisionByIdInternalServerError ", 500)
}

func (o *GetRevisionByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
