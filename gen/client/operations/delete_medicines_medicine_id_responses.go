// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteMedicinesMedicineIDReader is a Reader for the DeleteMedicinesMedicineID structure.
type DeleteMedicinesMedicineIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMedicinesMedicineIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteMedicinesMedicineIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteMedicinesMedicineIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteMedicinesMedicineIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /medicines/{medicine_id}] DeleteMedicinesMedicineID", response, response.Code())
	}
}

// NewDeleteMedicinesMedicineIDOK creates a DeleteMedicinesMedicineIDOK with default headers values
func NewDeleteMedicinesMedicineIDOK() *DeleteMedicinesMedicineIDOK {
	return &DeleteMedicinesMedicineIDOK{}
}

/*
DeleteMedicinesMedicineIDOK describes a response with status code 200, with default header values.

Medicine deleted successfully
*/
type DeleteMedicinesMedicineIDOK struct {
	Payload string
}

// IsSuccess returns true when this delete medicines medicine Id o k response has a 2xx status code
func (o *DeleteMedicinesMedicineIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete medicines medicine Id o k response has a 3xx status code
func (o *DeleteMedicinesMedicineIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete medicines medicine Id o k response has a 4xx status code
func (o *DeleteMedicinesMedicineIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete medicines medicine Id o k response has a 5xx status code
func (o *DeleteMedicinesMedicineIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete medicines medicine Id o k response a status code equal to that given
func (o *DeleteMedicinesMedicineIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete medicines medicine Id o k response
func (o *DeleteMedicinesMedicineIDOK) Code() int {
	return 200
}

func (o *DeleteMedicinesMedicineIDOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdOK %s", 200, payload)
}

func (o *DeleteMedicinesMedicineIDOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdOK %s", 200, payload)
}

func (o *DeleteMedicinesMedicineIDOK) GetPayload() string {
	return o.Payload
}

func (o *DeleteMedicinesMedicineIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMedicinesMedicineIDBadRequest creates a DeleteMedicinesMedicineIDBadRequest with default headers values
func NewDeleteMedicinesMedicineIDBadRequest() *DeleteMedicinesMedicineIDBadRequest {
	return &DeleteMedicinesMedicineIDBadRequest{}
}

/*
DeleteMedicinesMedicineIDBadRequest describes a response with status code 400, with default header values.

bad request
*/
type DeleteMedicinesMedicineIDBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this delete medicines medicine Id bad request response has a 2xx status code
func (o *DeleteMedicinesMedicineIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete medicines medicine Id bad request response has a 3xx status code
func (o *DeleteMedicinesMedicineIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete medicines medicine Id bad request response has a 4xx status code
func (o *DeleteMedicinesMedicineIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete medicines medicine Id bad request response has a 5xx status code
func (o *DeleteMedicinesMedicineIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this delete medicines medicine Id bad request response a status code equal to that given
func (o *DeleteMedicinesMedicineIDBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the delete medicines medicine Id bad request response
func (o *DeleteMedicinesMedicineIDBadRequest) Code() int {
	return 400
}

func (o *DeleteMedicinesMedicineIDBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdBadRequest %s", 400, payload)
}

func (o *DeleteMedicinesMedicineIDBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdBadRequest %s", 400, payload)
}

func (o *DeleteMedicinesMedicineIDBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteMedicinesMedicineIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteMedicinesMedicineIDUnauthorized creates a DeleteMedicinesMedicineIDUnauthorized with default headers values
func NewDeleteMedicinesMedicineIDUnauthorized() *DeleteMedicinesMedicineIDUnauthorized {
	return &DeleteMedicinesMedicineIDUnauthorized{}
}

/*
DeleteMedicinesMedicineIDUnauthorized describes a response with status code 401, with default header values.

internal server error
*/
type DeleteMedicinesMedicineIDUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this delete medicines medicine Id unauthorized response has a 2xx status code
func (o *DeleteMedicinesMedicineIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete medicines medicine Id unauthorized response has a 3xx status code
func (o *DeleteMedicinesMedicineIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete medicines medicine Id unauthorized response has a 4xx status code
func (o *DeleteMedicinesMedicineIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete medicines medicine Id unauthorized response has a 5xx status code
func (o *DeleteMedicinesMedicineIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete medicines medicine Id unauthorized response a status code equal to that given
func (o *DeleteMedicinesMedicineIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete medicines medicine Id unauthorized response
func (o *DeleteMedicinesMedicineIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteMedicinesMedicineIDUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdUnauthorized %s", 401, payload)
}

func (o *DeleteMedicinesMedicineIDUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /medicines/{medicine_id}][%d] deleteMedicinesMedicineIdUnauthorized %s", 401, payload)
}

func (o *DeleteMedicinesMedicineIDUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *DeleteMedicinesMedicineIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
