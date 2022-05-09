// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/Omar-Belghaouti/microservices/client/models"
)

// AddProductReader is a Reader for the AddProduct structure.
type AddProductReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProductReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProductOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewAddProductUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 501:
		result := NewAddProductNotImplemented()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewAddProductOK creates a AddProductOK with default headers values
func NewAddProductOK() *AddProductOK {
	return &AddProductOK{}
}

/* AddProductOK describes a response with status code 200, with default header values.

Data structure representing a single product
*/
type AddProductOK struct {
	Payload *models.Product
}

func (o *AddProductOK) Error() string {
	return fmt.Sprintf("[POST /products/][%d] addProductOK  %+v", 200, o.Payload)
}
func (o *AddProductOK) GetPayload() *models.Product {
	return o.Payload
}

func (o *AddProductOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Product)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductUnprocessableEntity creates a AddProductUnprocessableEntity with default headers values
func NewAddProductUnprocessableEntity() *AddProductUnprocessableEntity {
	return &AddProductUnprocessableEntity{}
}

/* AddProductUnprocessableEntity describes a response with status code 422, with default header values.

Validation errors defined as an array of strings
*/
type AddProductUnprocessableEntity struct {
	Payload *models.ValidationError
}

func (o *AddProductUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /products/][%d] addProductUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *AddProductUnprocessableEntity) GetPayload() *models.ValidationError {
	return o.Payload
}

func (o *AddProductUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ValidationError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddProductNotImplemented creates a AddProductNotImplemented with default headers values
func NewAddProductNotImplemented() *AddProductNotImplemented {
	return &AddProductNotImplemented{}
}

/* AddProductNotImplemented describes a response with status code 501, with default header values.

Generic error message returned as a string
*/
type AddProductNotImplemented struct {
	Payload *models.GenericError
}

func (o *AddProductNotImplemented) Error() string {
	return fmt.Sprintf("[POST /products/][%d] addProductNotImplemented  %+v", 501, o.Payload)
}
func (o *AddProductNotImplemented) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *AddProductNotImplemented) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
