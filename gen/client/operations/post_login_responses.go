// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/hamza-sharif/homeopathic-doctor-assistant/gen/models"
)

// PostLoginReader is a Reader for the PostLogin structure.
type PostLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostLoginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /login] PostLogin", response, response.Code())
	}
}

// NewPostLoginOK creates a PostLoginOK with default headers values
func NewPostLoginOK() *PostLoginOK {
	return &PostLoginOK{}
}

/*
PostLoginOK describes a response with status code 200, with default header values.

List of patients
*/
type PostLoginOK struct {
	Payload *models.User
}

// IsSuccess returns true when this post login o k response has a 2xx status code
func (o *PostLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post login o k response has a 3xx status code
func (o *PostLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login o k response has a 4xx status code
func (o *PostLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post login o k response has a 5xx status code
func (o *PostLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post login o k response a status code equal to that given
func (o *PostLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the post login o k response
func (o *PostLoginOK) Code() int {
	return 200
}

func (o *PostLoginOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginOK %s", 200, payload)
}

func (o *PostLoginOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginOK %s", 200, payload)
}

func (o *PostLoginOK) GetPayload() *models.User {
	return o.Payload
}

func (o *PostLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLoginBadRequest creates a PostLoginBadRequest with default headers values
func NewPostLoginBadRequest() *PostLoginBadRequest {
	return &PostLoginBadRequest{}
}

/*
PostLoginBadRequest describes a response with status code 400, with default header values.

bad request
*/
type PostLoginBadRequest struct {
	Payload interface{}
}

// IsSuccess returns true when this post login bad request response has a 2xx status code
func (o *PostLoginBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post login bad request response has a 3xx status code
func (o *PostLoginBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login bad request response has a 4xx status code
func (o *PostLoginBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post login bad request response has a 5xx status code
func (o *PostLoginBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post login bad request response a status code equal to that given
func (o *PostLoginBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the post login bad request response
func (o *PostLoginBadRequest) Code() int {
	return 400
}

func (o *PostLoginBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginBadRequest %s", 400, payload)
}

func (o *PostLoginBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginBadRequest %s", 400, payload)
}

func (o *PostLoginBadRequest) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLoginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLoginUnauthorized creates a PostLoginUnauthorized with default headers values
func NewPostLoginUnauthorized() *PostLoginUnauthorized {
	return &PostLoginUnauthorized{}
}

/*
PostLoginUnauthorized describes a response with status code 401, with default header values.

internal server error
*/
type PostLoginUnauthorized struct {
	Payload interface{}
}

// IsSuccess returns true when this post login unauthorized response has a 2xx status code
func (o *PostLoginUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post login unauthorized response has a 3xx status code
func (o *PostLoginUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login unauthorized response has a 4xx status code
func (o *PostLoginUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post login unauthorized response has a 5xx status code
func (o *PostLoginUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post login unauthorized response a status code equal to that given
func (o *PostLoginUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the post login unauthorized response
func (o *PostLoginUnauthorized) Code() int {
	return 401
}

func (o *PostLoginUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginUnauthorized %s", 401, payload)
}

func (o *PostLoginUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginUnauthorized %s", 401, payload)
}

func (o *PostLoginUnauthorized) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLoginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostLoginNotFound creates a PostLoginNotFound with default headers values
func NewPostLoginNotFound() *PostLoginNotFound {
	return &PostLoginNotFound{}
}

/*
PostLoginNotFound describes a response with status code 404, with default header values.

internal server error
*/
type PostLoginNotFound struct {
	Payload interface{}
}

// IsSuccess returns true when this post login not found response has a 2xx status code
func (o *PostLoginNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post login not found response has a 3xx status code
func (o *PostLoginNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post login not found response has a 4xx status code
func (o *PostLoginNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post login not found response has a 5xx status code
func (o *PostLoginNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post login not found response a status code equal to that given
func (o *PostLoginNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the post login not found response
func (o *PostLoginNotFound) Code() int {
	return 404
}

func (o *PostLoginNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginNotFound %s", 404, payload)
}

func (o *PostLoginNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /login][%d] postLoginNotFound %s", 404, payload)
}

func (o *PostLoginNotFound) GetPayload() interface{} {
	return o.Payload
}

func (o *PostLoginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
PostLoginBody post login body
swagger:model PostLoginBody
*/
type PostLoginBody struct {

	// password
	Password string `json:"password,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this post login body
func (o *PostLoginBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this post login body based on context it is used
func (o *PostLoginBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostLoginBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostLoginBody) UnmarshalBinary(b []byte) error {
	var res PostLoginBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}