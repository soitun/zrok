// Code generated by go-swagger; DO NOT EDIT.

package share

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// UnaccessOKCode is the HTTP code returned for type UnaccessOK
const UnaccessOKCode int = 200

/*
UnaccessOK access removed

swagger:response unaccessOK
*/
type UnaccessOK struct {
}

// NewUnaccessOK creates UnaccessOK with default headers values
func NewUnaccessOK() *UnaccessOK {

	return &UnaccessOK{}
}

// WriteResponse to the client
func (o *UnaccessOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}

// UnaccessUnauthorizedCode is the HTTP code returned for type UnaccessUnauthorized
const UnaccessUnauthorizedCode int = 401

/*
UnaccessUnauthorized unauthorized

swagger:response unaccessUnauthorized
*/
type UnaccessUnauthorized struct {
}

// NewUnaccessUnauthorized creates UnaccessUnauthorized with default headers values
func NewUnaccessUnauthorized() *UnaccessUnauthorized {

	return &UnaccessUnauthorized{}
}

// WriteResponse to the client
func (o *UnaccessUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(401)
}

// UnaccessNotFoundCode is the HTTP code returned for type UnaccessNotFound
const UnaccessNotFoundCode int = 404

/*
UnaccessNotFound not found

swagger:response unaccessNotFound
*/
type UnaccessNotFound struct {
}

// NewUnaccessNotFound creates UnaccessNotFound with default headers values
func NewUnaccessNotFound() *UnaccessNotFound {

	return &UnaccessNotFound{}
}

// WriteResponse to the client
func (o *UnaccessNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// UnaccessInternalServerErrorCode is the HTTP code returned for type UnaccessInternalServerError
const UnaccessInternalServerErrorCode int = 500

/*
UnaccessInternalServerError internal server error

swagger:response unaccessInternalServerError
*/
type UnaccessInternalServerError struct {
}

// NewUnaccessInternalServerError creates UnaccessInternalServerError with default headers values
func NewUnaccessInternalServerError() *UnaccessInternalServerError {

	return &UnaccessInternalServerError{}
}

// WriteResponse to the client
func (o *UnaccessInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}