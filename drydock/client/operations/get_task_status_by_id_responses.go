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

// GetTaskStatusByIDReader is a Reader for the GetTaskStatusByID structure.
type GetTaskStatusByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTaskStatusByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetTaskStatusByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetTaskStatusByIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetTaskStatusByIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetTaskStatusByIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTaskStatusByIDOK creates a GetTaskStatusByIDOK with default headers values
func NewGetTaskStatusByIDOK() *GetTaskStatusByIDOK {
	return &GetTaskStatusByIDOK{}
}

/*GetTaskStatusByIDOK handles this case with default header values.

Generic String answer
*/
type GetTaskStatusByIDOK struct {
	Payload string
}

func (o *GetTaskStatusByIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks/{task-id}][%d] getTaskStatusByIdOK  %+v", 200, o.Payload)
}

func (o *GetTaskStatusByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTaskStatusByIDUnauthorized creates a GetTaskStatusByIDUnauthorized with default headers values
func NewGetTaskStatusByIDUnauthorized() *GetTaskStatusByIDUnauthorized {
	return &GetTaskStatusByIDUnauthorized{}
}

/*GetTaskStatusByIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetTaskStatusByIDUnauthorized struct {
}

func (o *GetTaskStatusByIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks/{task-id}][%d] getTaskStatusByIdUnauthorized ", 401)
}

func (o *GetTaskStatusByIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskStatusByIDForbidden creates a GetTaskStatusByIDForbidden with default headers values
func NewGetTaskStatusByIDForbidden() *GetTaskStatusByIDForbidden {
	return &GetTaskStatusByIDForbidden{}
}

/*GetTaskStatusByIDForbidden handles this case with default header values.

403 Forbidden
*/
type GetTaskStatusByIDForbidden struct {
}

func (o *GetTaskStatusByIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks/{task-id}][%d] getTaskStatusByIdForbidden ", 403)
}

func (o *GetTaskStatusByIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetTaskStatusByIDInternalServerError creates a GetTaskStatusByIDInternalServerError with default headers values
func NewGetTaskStatusByIDInternalServerError() *GetTaskStatusByIDInternalServerError {
	return &GetTaskStatusByIDInternalServerError{}
}

/*GetTaskStatusByIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetTaskStatusByIDInternalServerError struct {
}

func (o *GetTaskStatusByIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/tasks/{task-id}][%d] getTaskStatusByIdInternalServerError ", 500)
}

func (o *GetTaskStatusByIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
