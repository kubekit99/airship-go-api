// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/kubekit99/airship-go-api/armada/models"
)

// GetStatusReader is a Reader for the GetStatus structure.
type GetStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStatusOK creates a GetStatusOK with default headers values
func NewGetStatusOK() *GetStatusOK {
	return &GetStatusOK{}
}

/*GetStatusOK handles this case with default header values.

Response of Tiller statuses
*/
type GetStatusOK struct {
	Payload *GetStatusOKBody
}

func (o *GetStatusOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/status][%d] getStatusOK  %+v", 200, o.Payload)
}

func (o *GetStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetStatusOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetStatusUnauthorized creates a GetStatusUnauthorized with default headers values
func NewGetStatusUnauthorized() *GetStatusUnauthorized {
	return &GetStatusUnauthorized{}
}

/*GetStatusUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetStatusUnauthorized struct {
}

func (o *GetStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/status][%d] getStatusUnauthorized ", 401)
}

func (o *GetStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatusForbidden creates a GetStatusForbidden with default headers values
func NewGetStatusForbidden() *GetStatusForbidden {
	return &GetStatusForbidden{}
}

/*GetStatusForbidden handles this case with default header values.

403 Forbidden
*/
type GetStatusForbidden struct {
}

func (o *GetStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/status][%d] getStatusForbidden ", 403)
}

func (o *GetStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetStatusInternalServerError creates a GetStatusInternalServerError with default headers values
func NewGetStatusInternalServerError() *GetStatusInternalServerError {
	return &GetStatusInternalServerError{}
}

/*GetStatusInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetStatusInternalServerError struct {
}

func (o *GetStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/status][%d] getStatusInternalServerError ", 500)
}

func (o *GetStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetStatusOKBody get status o k body
swagger:model GetStatusOKBody
*/
type GetStatusOKBody struct {
	models.Status
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetStatusOKBody) UnmarshalJSON(raw []byte) error {
	// GetStatusOKBodyAO0
	var getStatusOKBodyAO0 models.Status
	if err := swag.ReadJSON(raw, &getStatusOKBodyAO0); err != nil {
		return err
	}
	o.Status = getStatusOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetStatusOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getStatusOKBodyAO0, err := swag.WriteJSON(o.Status)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getStatusOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get status o k body
func (o *GetStatusOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Status
	if err := o.Status.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetStatusOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetStatusOKBody) UnmarshalBinary(b []byte) error {
	var res GetStatusOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
