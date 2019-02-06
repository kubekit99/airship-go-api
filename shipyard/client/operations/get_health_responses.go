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

// GetHealthReader is a Reader for the GetHealth structure.
type GetHealthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetHealthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetHealthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetHealthUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetHealthForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetHealthInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHealthOK creates a GetHealthOK with default headers values
func NewGetHealthOK() *GetHealthOK {
	return &GetHealthOK{}
}

/*GetHealthOK handles this case with default header values.

Generic String answer
*/
type GetHealthOK struct {
	Payload string
}

func (o *GetHealthOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthOK  %+v", 200, o.Payload)
}

func (o *GetHealthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetHealthUnauthorized creates a GetHealthUnauthorized with default headers values
func NewGetHealthUnauthorized() *GetHealthUnauthorized {
	return &GetHealthUnauthorized{}
}

/*GetHealthUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetHealthUnauthorized struct {
}

func (o *GetHealthUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthUnauthorized ", 401)
}

func (o *GetHealthUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthForbidden creates a GetHealthForbidden with default headers values
func NewGetHealthForbidden() *GetHealthForbidden {
	return &GetHealthForbidden{}
}

/*GetHealthForbidden handles this case with default header values.

403 Forbidden
*/
type GetHealthForbidden struct {
}

func (o *GetHealthForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthForbidden ", 403)
}

func (o *GetHealthForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthInternalServerError creates a GetHealthInternalServerError with default headers values
func NewGetHealthInternalServerError() *GetHealthInternalServerError {
	return &GetHealthInternalServerError{}
}

/*GetHealthInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetHealthInternalServerError struct {
}

func (o *GetHealthInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthInternalServerError ", 500)
}

func (o *GetHealthInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}