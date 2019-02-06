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

// GetSiteStatusesReader is a Reader for the GetSiteStatuses structure.
type GetSiteStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSiteStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSiteStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetSiteStatusesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetSiteStatusesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetSiteStatusesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSiteStatusesOK creates a GetSiteStatusesOK with default headers values
func NewGetSiteStatusesOK() *GetSiteStatusesOK {
	return &GetSiteStatusesOK{}
}

/*GetSiteStatusesOK handles this case with default header values.

Generic String answer
*/
type GetSiteStatusesOK struct {
	Payload string
}

func (o *GetSiteStatusesOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/site_statuses][%d] getSiteStatusesOK  %+v", 200, o.Payload)
}

func (o *GetSiteStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSiteStatusesUnauthorized creates a GetSiteStatusesUnauthorized with default headers values
func NewGetSiteStatusesUnauthorized() *GetSiteStatusesUnauthorized {
	return &GetSiteStatusesUnauthorized{}
}

/*GetSiteStatusesUnauthorized handles this case with default header values.

401 Not authorized
*/
type GetSiteStatusesUnauthorized struct {
}

func (o *GetSiteStatusesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/site_statuses][%d] getSiteStatusesUnauthorized ", 401)
}

func (o *GetSiteStatusesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSiteStatusesForbidden creates a GetSiteStatusesForbidden with default headers values
func NewGetSiteStatusesForbidden() *GetSiteStatusesForbidden {
	return &GetSiteStatusesForbidden{}
}

/*GetSiteStatusesForbidden handles this case with default header values.

403 Forbidden
*/
type GetSiteStatusesForbidden struct {
}

func (o *GetSiteStatusesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/site_statuses][%d] getSiteStatusesForbidden ", 403)
}

func (o *GetSiteStatusesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSiteStatusesInternalServerError creates a GetSiteStatusesInternalServerError with default headers values
func NewGetSiteStatusesInternalServerError() *GetSiteStatusesInternalServerError {
	return &GetSiteStatusesInternalServerError{}
}

/*GetSiteStatusesInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type GetSiteStatusesInternalServerError struct {
}

func (o *GetSiteStatusesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/site_statuses][%d] getSiteStatusesInternalServerError ", 500)
}

func (o *GetSiteStatusesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}