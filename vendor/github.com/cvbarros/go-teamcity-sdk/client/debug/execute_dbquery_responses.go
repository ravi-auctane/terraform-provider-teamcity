// Code generated by go-swagger; DO NOT EDIT.

package debug

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// ExecuteDBQueryReader is a Reader for the ExecuteDBQuery structure.
type ExecuteDBQueryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteDBQueryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExecuteDBQueryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExecuteDBQueryOK creates a ExecuteDBQueryOK with default headers values
func NewExecuteDBQueryOK() *ExecuteDBQueryOK {
	return &ExecuteDBQueryOK{}
}

/*ExecuteDBQueryOK handles this case with default header values.

successful operation
*/
type ExecuteDBQueryOK struct {
	Payload string
}

func (o *ExecuteDBQueryOK) Error() string {
	return fmt.Sprintf("[GET /app/rest/debug/database/query/{query}][%d] executeDBQueryOK  %+v", 200, o.Payload)
}

func (o *ExecuteDBQueryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}