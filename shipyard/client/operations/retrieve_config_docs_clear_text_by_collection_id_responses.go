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

// RetrieveConfigDocsClearTextByCollectionIDReader is a Reader for the RetrieveConfigDocsClearTextByCollectionID structure.
type RetrieveConfigDocsClearTextByCollectionIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RetrieveConfigDocsClearTextByCollectionIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRetrieveConfigDocsClearTextByCollectionIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewRetrieveConfigDocsClearTextByCollectionIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewRetrieveConfigDocsClearTextByCollectionIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewRetrieveConfigDocsClearTextByCollectionIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRetrieveConfigDocsClearTextByCollectionIDOK creates a RetrieveConfigDocsClearTextByCollectionIDOK with default headers values
func NewRetrieveConfigDocsClearTextByCollectionIDOK() *RetrieveConfigDocsClearTextByCollectionIDOK {
	return &RetrieveConfigDocsClearTextByCollectionIDOK{}
}

/*RetrieveConfigDocsClearTextByCollectionIDOK handles this case with default header values.

Generic String answer
*/
type RetrieveConfigDocsClearTextByCollectionIDOK struct {
	Payload string
}

func (o *RetrieveConfigDocsClearTextByCollectionIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/configdocs/{collection-id}][%d] retrieveConfigDocsClearTextByCollectionIdOK  %+v", 200, o.Payload)
}

func (o *RetrieveConfigDocsClearTextByCollectionIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRetrieveConfigDocsClearTextByCollectionIDUnauthorized creates a RetrieveConfigDocsClearTextByCollectionIDUnauthorized with default headers values
func NewRetrieveConfigDocsClearTextByCollectionIDUnauthorized() *RetrieveConfigDocsClearTextByCollectionIDUnauthorized {
	return &RetrieveConfigDocsClearTextByCollectionIDUnauthorized{}
}

/*RetrieveConfigDocsClearTextByCollectionIDUnauthorized handles this case with default header values.

401 Not authorized
*/
type RetrieveConfigDocsClearTextByCollectionIDUnauthorized struct {
}

func (o *RetrieveConfigDocsClearTextByCollectionIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/configdocs/{collection-id}][%d] retrieveConfigDocsClearTextByCollectionIdUnauthorized ", 401)
}

func (o *RetrieveConfigDocsClearTextByCollectionIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveConfigDocsClearTextByCollectionIDForbidden creates a RetrieveConfigDocsClearTextByCollectionIDForbidden with default headers values
func NewRetrieveConfigDocsClearTextByCollectionIDForbidden() *RetrieveConfigDocsClearTextByCollectionIDForbidden {
	return &RetrieveConfigDocsClearTextByCollectionIDForbidden{}
}

/*RetrieveConfigDocsClearTextByCollectionIDForbidden handles this case with default header values.

403 Forbidden
*/
type RetrieveConfigDocsClearTextByCollectionIDForbidden struct {
}

func (o *RetrieveConfigDocsClearTextByCollectionIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/configdocs/{collection-id}][%d] retrieveConfigDocsClearTextByCollectionIdForbidden ", 403)
}

func (o *RetrieveConfigDocsClearTextByCollectionIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewRetrieveConfigDocsClearTextByCollectionIDInternalServerError creates a RetrieveConfigDocsClearTextByCollectionIDInternalServerError with default headers values
func NewRetrieveConfigDocsClearTextByCollectionIDInternalServerError() *RetrieveConfigDocsClearTextByCollectionIDInternalServerError {
	return &RetrieveConfigDocsClearTextByCollectionIDInternalServerError{}
}

/*RetrieveConfigDocsClearTextByCollectionIDInternalServerError handles this case with default header values.

500 Internal Server Error
*/
type RetrieveConfigDocsClearTextByCollectionIDInternalServerError struct {
}

func (o *RetrieveConfigDocsClearTextByCollectionIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/configdocs/{collection-id}][%d] retrieveConfigDocsClearTextByCollectionIdInternalServerError ", 500)
}

func (o *RetrieveConfigDocsClearTextByCollectionIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
