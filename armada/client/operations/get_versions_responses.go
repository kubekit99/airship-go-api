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

// GetVersionsReader is a Reader for the GetVersions structure.
type GetVersionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetVersionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetVersionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetVersionsOK creates a GetVersionsOK with default headers values
func NewGetVersionsOK() *GetVersionsOK {
	return &GetVersionsOK{}
}

/*GetVersionsOK handles this case with default header values.

Response of getting Armada versions
*/
type GetVersionsOK struct {
	Payload *GetVersionsOKBody
}

func (o *GetVersionsOK) Error() string {
	return fmt.Sprintf("[GET /versions][%d] getVersionsOK  %+v", 200, o.Payload)
}

func (o *GetVersionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetVersionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetVersionsOKBody get versions o k body
swagger:model GetVersionsOKBody
*/
type GetVersionsOKBody struct {
	models.Versions
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetVersionsOKBody) UnmarshalJSON(raw []byte) error {
	// GetVersionsOKBodyAO0
	var getVersionsOKBodyAO0 models.Versions
	if err := swag.ReadJSON(raw, &getVersionsOKBodyAO0); err != nil {
		return err
	}
	o.Versions = getVersionsOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetVersionsOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getVersionsOKBodyAO0, err := swag.WriteJSON(o.Versions)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getVersionsOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get versions o k body
func (o *GetVersionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Versions
	if err := o.Versions.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetVersionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetVersionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetVersionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}