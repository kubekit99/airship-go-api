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

// ProbeReadinessReader is a Reader for the ProbeReadiness structure.
type ProbeReadinessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProbeReadinessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewProbeReadinessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewProbeReadinessOK creates a ProbeReadinessOK with default headers values
func NewProbeReadinessOK() *ProbeReadinessOK {
	return &ProbeReadinessOK{}
}

/*ProbeReadinessOK handles this case with default header values.

Generic String answer
*/
type ProbeReadinessOK struct {
	Payload string
}

func (o *ProbeReadinessOK) Error() string {
	return fmt.Sprintf("[GET /readiness][%d] probeReadinessOK  %+v", 200, o.Payload)
}

func (o *ProbeReadinessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
