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

// GetReleasesReader is a Reader for the GetReleases structure.
type GetReleasesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetReleasesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetReleasesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetReleasesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetReleasesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetReleasesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetReleasesOK creates a GetReleasesOK with default headers values
func NewGetReleasesOK() *GetReleasesOK {
	return &GetReleasesOK{}
}

/*GetReleasesOK handles this case with default header values.

Response of all namespaces and releases contained within
*/
type GetReleasesOK struct {
	Payload *GetReleasesOKBody
}

func (o *GetReleasesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/releases][%d] getReleasesOK  %+v", 200, o.Payload)
}

func (o *GetReleasesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetReleasesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetReleasesUnauthorized creates a GetReleasesUnauthorized with default headers values
func NewGetReleasesUnauthorized() *GetReleasesUnauthorized {
	return &GetReleasesUnauthorized{}
}

/*GetReleasesUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetReleasesUnauthorized struct {
}

func (o *GetReleasesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/releases][%d] getReleasesUnauthorized ", 401)
}

func (o *GetReleasesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleasesForbidden creates a GetReleasesForbidden with default headers values
func NewGetReleasesForbidden() *GetReleasesForbidden {
	return &GetReleasesForbidden{}
}

/*GetReleasesForbidden handles this case with default header values.

403 Forbidden
*/
type GetReleasesForbidden struct {
}

func (o *GetReleasesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/releases][%d] getReleasesForbidden ", 403)
}

func (o *GetReleasesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetReleasesInternalServerError creates a GetReleasesInternalServerError with default headers values
func NewGetReleasesInternalServerError() *GetReleasesInternalServerError {
	return &GetReleasesInternalServerError{}
}

/*GetReleasesInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetReleasesInternalServerError struct {
}

func (o *GetReleasesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/releases][%d] getReleasesInternalServerError ", 500)
}

func (o *GetReleasesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*GetReleasesOKBody get releases o k body
swagger:model GetReleasesOKBody
*/
type GetReleasesOKBody struct {
	models.Releases
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *GetReleasesOKBody) UnmarshalJSON(raw []byte) error {
	// GetReleasesOKBodyAO0
	var getReleasesOKBodyAO0 models.Releases
	if err := swag.ReadJSON(raw, &getReleasesOKBodyAO0); err != nil {
		return err
	}
	o.Releases = getReleasesOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o GetReleasesOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	getReleasesOKBodyAO0, err := swag.WriteJSON(o.Releases)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, getReleasesOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this get releases o k body
func (o *GetReleasesOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Releases
	if err := o.Releases.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *GetReleasesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetReleasesOKBody) UnmarshalBinary(b []byte) error {
	var res GetReleasesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
