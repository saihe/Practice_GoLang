// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// GetEmployeeIDOKCode is the HTTP code returned for type GetEmployeeIDOK
const GetEmployeeIDOKCode int = 200

/*GetEmployeeIDOK OK

swagger:response getEmployeeIdOK
*/
type GetEmployeeIDOK struct {

	/*
	  In: Body
	*/
	Payload *GetEmployeeIDOKBody `json:"body,omitempty"`
}

// NewGetEmployeeIDOK creates GetEmployeeIDOK with default headers values
func NewGetEmployeeIDOK() *GetEmployeeIDOK {

	return &GetEmployeeIDOK{}
}

// WithPayload adds the payload to the get employee Id o k response
func (o *GetEmployeeIDOK) WithPayload(payload *GetEmployeeIDOKBody) *GetEmployeeIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get employee Id o k response
func (o *GetEmployeeIDOK) SetPayload(payload *GetEmployeeIDOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetEmployeeIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
