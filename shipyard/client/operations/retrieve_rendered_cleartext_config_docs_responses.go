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

// RetrieveRenderedCleartextConfigDocsReader is a Reader for the RetrieveRenderedCleartextConfigDocs structure.
type RetrieveRenderedCleartextConfigDocsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveRenderedCleartextConfigDocsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveRenderedCleartextConfigDocsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRetrieveRenderedCleartextConfigDocsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRetrieveRenderedCleartextConfigDocsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRetrieveRenderedCleartextConfigDocsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveRenderedCleartextConfigDocsOK creates a RetrieveRenderedCleartextConfigDocsOK with default headers values
func NewRetrieveRenderedCleartextConfigDocsOK() *RetrieveRenderedCleartextConfigDocsOK {
	return &RetrieveRenderedCleartextConfigDocsOK{}
}

/*RetrieveRenderedCleartextConfigDocsOK handles this case with default header values.

Generic String answer
*/
type RetrieveRenderedCleartextConfigDocsOK struct {
	Payload string
}

func (o *RetrieveRenderedCleartextConfigDocsOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/renderedconfigdocs][%d] retrieveRenderedCleartextConfigDocsOK  %+v", 200, o.Payload)
}

func (o *RetrieveRenderedCleartextConfigDocsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveRenderedCleartextConfigDocsUnauthorized creates a RetrieveRenderedCleartextConfigDocsUnauthorized with default headers values
func NewRetrieveRenderedCleartextConfigDocsUnauthorized() *RetrieveRenderedCleartextConfigDocsUnauthorized {
	return &RetrieveRenderedCleartextConfigDocsUnauthorized{}
}

/*RetrieveRenderedCleartextConfigDocsUnauthorized handles this case with default header values.

401 Not authorized
*/
type RetrieveRenderedCleartextConfigDocsUnauthorized struct {
}

func (o *RetrieveRenderedCleartextConfigDocsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/renderedconfigdocs][%d] retrieveRenderedCleartextConfigDocsUnauthorized ", 401)
}

func (o *RetrieveRenderedCleartextConfigDocsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveRenderedCleartextConfigDocsForbidden creates a RetrieveRenderedCleartextConfigDocsForbidden with default headers values
func NewRetrieveRenderedCleartextConfigDocsForbidden() *RetrieveRenderedCleartextConfigDocsForbidden {
	return &RetrieveRenderedCleartextConfigDocsForbidden{}
}

/*RetrieveRenderedCleartextConfigDocsForbidden handles this case with default header values.

403 Forbidden
*/
type RetrieveRenderedCleartextConfigDocsForbidden struct {
}

func (o *RetrieveRenderedCleartextConfigDocsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/renderedconfigdocs][%d] retrieveRenderedCleartextConfigDocsForbidden ", 403)
}

func (o *RetrieveRenderedCleartextConfigDocsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveRenderedCleartextConfigDocsInternalServerError creates a RetrieveRenderedCleartextConfigDocsInternalServerError with default headers values
func NewRetrieveRenderedCleartextConfigDocsInternalServerError() *RetrieveRenderedCleartextConfigDocsInternalServerError {
	return &RetrieveRenderedCleartextConfigDocsInternalServerError{}
}

/*RetrieveRenderedCleartextConfigDocsInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type RetrieveRenderedCleartextConfigDocsInternalServerError struct {
}

func (o *RetrieveRenderedCleartextConfigDocsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/renderedconfigdocs][%d] retrieveRenderedCleartextConfigDocsInternalServerError ", 500)
}

func (o *RetrieveRenderedCleartextConfigDocsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
