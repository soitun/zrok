// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti-test-kitchen/zrok/rest_model_zrok"
)

// UnshareReader is a Reader for the Unshare structure.
type UnshareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnshareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUnshareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUnshareUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUnshareNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUnshareInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUnshareOK creates a UnshareOK with default headers values
func NewUnshareOK() *UnshareOK {
	return &UnshareOK{}
}

/*
UnshareOK describes a response with status code 200, with default header values.

share removed
*/
type UnshareOK struct {
}

// IsSuccess returns true when this unshare o k response has a 2xx status code
func (o *UnshareOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unshare o k response has a 3xx status code
func (o *UnshareOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unshare o k response has a 4xx status code
func (o *UnshareOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this unshare o k response has a 5xx status code
func (o *UnshareOK) IsServerError() bool {
	return false
}

// IsCode returns true when this unshare o k response a status code equal to that given
func (o *UnshareOK) IsCode(code int) bool {
	return code == 200
}

func (o *UnshareOK) Error() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareOK ", 200)
}

func (o *UnshareOK) String() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareOK ", 200)
}

func (o *UnshareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnshareUnauthorized creates a UnshareUnauthorized with default headers values
func NewUnshareUnauthorized() *UnshareUnauthorized {
	return &UnshareUnauthorized{}
}

/*
UnshareUnauthorized describes a response with status code 401, with default header values.

unauthorized
*/
type UnshareUnauthorized struct {
}

// IsSuccess returns true when this unshare unauthorized response has a 2xx status code
func (o *UnshareUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unshare unauthorized response has a 3xx status code
func (o *UnshareUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unshare unauthorized response has a 4xx status code
func (o *UnshareUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this unshare unauthorized response has a 5xx status code
func (o *UnshareUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this unshare unauthorized response a status code equal to that given
func (o *UnshareUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UnshareUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareUnauthorized ", 401)
}

func (o *UnshareUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareUnauthorized ", 401)
}

func (o *UnshareUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnshareNotFound creates a UnshareNotFound with default headers values
func NewUnshareNotFound() *UnshareNotFound {
	return &UnshareNotFound{}
}

/*
UnshareNotFound describes a response with status code 404, with default header values.

not found
*/
type UnshareNotFound struct {
}

// IsSuccess returns true when this unshare not found response has a 2xx status code
func (o *UnshareNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unshare not found response has a 3xx status code
func (o *UnshareNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unshare not found response has a 4xx status code
func (o *UnshareNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this unshare not found response has a 5xx status code
func (o *UnshareNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this unshare not found response a status code equal to that given
func (o *UnshareNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UnshareNotFound) Error() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareNotFound ", 404)
}

func (o *UnshareNotFound) String() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareNotFound ", 404)
}

func (o *UnshareNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnshareInternalServerError creates a UnshareInternalServerError with default headers values
func NewUnshareInternalServerError() *UnshareInternalServerError {
	return &UnshareInternalServerError{}
}

/*
UnshareInternalServerError describes a response with status code 500, with default header values.

internal server error
*/
type UnshareInternalServerError struct {
	Payload rest_model_zrok.ErrorMessage
}

// IsSuccess returns true when this unshare internal server error response has a 2xx status code
func (o *UnshareInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this unshare internal server error response has a 3xx status code
func (o *UnshareInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unshare internal server error response has a 4xx status code
func (o *UnshareInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this unshare internal server error response has a 5xx status code
func (o *UnshareInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this unshare internal server error response a status code equal to that given
func (o *UnshareInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UnshareInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareInternalServerError  %+v", 500, o.Payload)
}

func (o *UnshareInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /unshare][%d] unshareInternalServerError  %+v", 500, o.Payload)
}

func (o *UnshareInternalServerError) GetPayload() rest_model_zrok.ErrorMessage {
	return o.Payload
}

func (o *UnshareInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
