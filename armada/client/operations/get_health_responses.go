// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

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

	case 204:
		result := NewGetHealthNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 503:
		result := NewGetHealthServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetHealthNoContent creates a GetHealthNoContent with default headers values
func NewGetHealthNoContent() *GetHealthNoContent {
	return &GetHealthNoContent{}
}

/*GetHealthNoContent handles this case with default header values.

Indicates the system is healthy. This is currently the default return.
*/
type GetHealthNoContent struct {
}

func (o *GetHealthNoContent) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthNoContent ", 204)
}

func (o *GetHealthNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetHealthServiceUnavailable creates a GetHealthServiceUnavailable with default headers values
func NewGetHealthServiceUnavailable() *GetHealthServiceUnavailable {
	return &GetHealthServiceUnavailable{}
}

/*GetHealthServiceUnavailable handles this case with default header values.

Indicates the system is not healthy.  This is not yet implemented.
*/
type GetHealthServiceUnavailable struct {
}

func (o *GetHealthServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v1.0/health][%d] getHealthServiceUnavailable ", 503)
}

func (o *GetHealthServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
