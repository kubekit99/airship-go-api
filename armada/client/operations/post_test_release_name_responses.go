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

// PostTestReleaseNameReader is a Reader for the PostTestReleaseName structure.
type PostTestReleaseNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTestReleaseNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostTestReleaseNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewPostTestReleaseNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostTestReleaseNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostTestReleaseNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostTestReleaseNameOK creates a PostTestReleaseNameOK with default headers values
func NewPostTestReleaseNameOK() *PostTestReleaseNameOK {
	return &PostTestReleaseNameOK{}
}

/*PostTestReleaseNameOK handles this case with default header values.

Response of a test of a specified release name
*/
type PostTestReleaseNameOK struct {
	Payload *PostTestReleaseNameOKBody
}

func (o *PostTestReleaseNameOK) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/test/{release_name}][%d] postTestReleaseNameOK  %+v", 200, o.Payload)
}

func (o *PostTestReleaseNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PostTestReleaseNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTestReleaseNameUnauthorized creates a PostTestReleaseNameUnauthorized with default headers values
func NewPostTestReleaseNameUnauthorized() *PostTestReleaseNameUnauthorized {
	return &PostTestReleaseNameUnauthorized{}
}

/*PostTestReleaseNameUnauthorized handles this case with default header values.

401 Not authorized
*/
type PostTestReleaseNameUnauthorized struct {
}

func (o *PostTestReleaseNameUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/test/{release_name}][%d] postTestReleaseNameUnauthorized ", 401)
}

func (o *PostTestReleaseNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTestReleaseNameForbidden creates a PostTestReleaseNameForbidden with default headers values
func NewPostTestReleaseNameForbidden() *PostTestReleaseNameForbidden {
	return &PostTestReleaseNameForbidden{}
}

/*PostTestReleaseNameForbidden handles this case with default header values.

403 Forbidden
*/
type PostTestReleaseNameForbidden struct {
}

func (o *PostTestReleaseNameForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/test/{release_name}][%d] postTestReleaseNameForbidden ", 403)
}

func (o *PostTestReleaseNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostTestReleaseNameInternalServerError creates a PostTestReleaseNameInternalServerError with default headers values
func NewPostTestReleaseNameInternalServerError() *PostTestReleaseNameInternalServerError {
	return &PostTestReleaseNameInternalServerError{}
}

/*PostTestReleaseNameInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type PostTestReleaseNameInternalServerError struct {
}

func (o *PostTestReleaseNameInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v1.0/test/{release_name}][%d] postTestReleaseNameInternalServerError ", 500)
}

func (o *PostTestReleaseNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*PostTestReleaseNameOKBody post test release name o k body
swagger:model PostTestReleaseNameOKBody
*/
type PostTestReleaseNameOKBody struct {
	models.Testresult
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostTestReleaseNameOKBody) UnmarshalJSON(raw []byte) error {
	// PostTestReleaseNameOKBodyAO0
	var postTestReleaseNameOKBodyAO0 models.Testresult
	if err := swag.ReadJSON(raw, &postTestReleaseNameOKBodyAO0); err != nil {
		return err
	}
	o.Testresult = postTestReleaseNameOKBodyAO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostTestReleaseNameOKBody) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	postTestReleaseNameOKBodyAO0, err := swag.WriteJSON(o.Testresult)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postTestReleaseNameOKBodyAO0)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post test release name o k body
func (o *PostTestReleaseNameOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with models.Testresult
	if err := o.Testresult.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (o *PostTestReleaseNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostTestReleaseNameOKBody) UnmarshalBinary(b []byte) error {
	var res PostTestReleaseNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}