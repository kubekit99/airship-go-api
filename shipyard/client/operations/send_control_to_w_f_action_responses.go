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

// SendControlToWFActionReader is a Reader for the SendControlToWFAction structure.
type SendControlToWFActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SendControlToWFActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSendControlToWFActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewSendControlToWFActionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewSendControlToWFActionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewSendControlToWFActionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSendControlToWFActionOK creates a SendControlToWFActionOK with default headers values
func NewSendControlToWFActionOK() *SendControlToWFActionOK {
	return &SendControlToWFActionOK{}
}

/*SendControlToWFActionOK handles this case with default header values.

Generic String answer
*/
type SendControlToWFActionOK struct {
	Payload string
}

func (o *SendControlToWFActionOK) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/actions/{action-id}/control/{control-verb}][%d] sendControlToWFActionOK  %+v", 200, o.Payload)
}

func (o *SendControlToWFActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSendControlToWFActionUnauthorized creates a SendControlToWFActionUnauthorized with default headers values
func NewSendControlToWFActionUnauthorized() *SendControlToWFActionUnauthorized {
	return &SendControlToWFActionUnauthorized{}
}

/*SendControlToWFActionUnauthorized handles this case with default header values.

401 Not authorized
*/
type SendControlToWFActionUnauthorized struct {
}

func (o *SendControlToWFActionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/actions/{action-id}/control/{control-verb}][%d] sendControlToWFActionUnauthorized ", 401)
}

func (o *SendControlToWFActionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendControlToWFActionForbidden creates a SendControlToWFActionForbidden with default headers values
func NewSendControlToWFActionForbidden() *SendControlToWFActionForbidden {
	return &SendControlToWFActionForbidden{}
}

/*SendControlToWFActionForbidden handles this case with default header values.

403 Forbidden
*/
type SendControlToWFActionForbidden struct {
}

func (o *SendControlToWFActionForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/actions/{action-id}/control/{control-verb}][%d] sendControlToWFActionForbidden ", 403)
}

func (o *SendControlToWFActionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSendControlToWFActionInternalServerError creates a SendControlToWFActionInternalServerError with default headers values
func NewSendControlToWFActionInternalServerError() *SendControlToWFActionInternalServerError {
	return &SendControlToWFActionInternalServerError{}
}

/*SendControlToWFActionInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type SendControlToWFActionInternalServerError struct {
}

func (o *SendControlToWFActionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/actions/{action-id}/control/{control-verb}][%d] sendControlToWFActionInternalServerError ", 500)
}

func (o *SendControlToWFActionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
