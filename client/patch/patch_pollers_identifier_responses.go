package patch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type PatchPollersIdentifierReader struct {
	formats strfmt.Registry
}

func (o *PatchPollersIdentifierReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchPollersIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewPatchPollersIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewPatchPollersIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewPatchPollersIdentifierOK creates a PatchPollersIdentifierOK with default headers values
func NewPatchPollersIdentifierOK() *PatchPollersIdentifierOK {
	return &PatchPollersIdentifierOK{}
}

/*PatchPollersIdentifierOK

Specifics of the patched poller

*/
type PatchPollersIdentifierOK struct {
	Payload PatchPollersIdentifierOKBodyBody
}

func (o *PatchPollersIdentifierOK) Error() string {
	return fmt.Sprintf("[PATCH /pollers/{identifier}][%d] patchPollersIdentifierOK  %+v", 200, o.Payload)
}

func (o *PatchPollersIdentifierOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPatchPollersIdentifierNotFound creates a PatchPollersIdentifierNotFound with default headers values
func NewPatchPollersIdentifierNotFound() *PatchPollersIdentifierNotFound {
	return &PatchPollersIdentifierNotFound{}
}

/*PatchPollersIdentifierNotFound

There is no poller with specified identifier.

*/
type PatchPollersIdentifierNotFound struct {
	Payload *models.Error
}

func (o *PatchPollersIdentifierNotFound) Error() string {
	return fmt.Sprintf("[PATCH /pollers/{identifier}][%d] patchPollersIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *PatchPollersIdentifierNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPatchPollersIdentifierDefault creates a PatchPollersIdentifierDefault with default headers values
func NewPatchPollersIdentifierDefault(code int) *PatchPollersIdentifierDefault {
	return &PatchPollersIdentifierDefault{
		_statusCode: code,
	}
}

/*PatchPollersIdentifierDefault

Unexpected error
*/
type PatchPollersIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the patch pollers identifier default response
func (o *PatchPollersIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *PatchPollersIdentifierDefault) Error() string {
	return fmt.Sprintf("[PATCH /pollers/{identifier}][%d] PatchPollersIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *PatchPollersIdentifierDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*PatchPollersIdentifierOKBodyBody PatchPollersIdentifierOKBodyBody patch pollers identifier o k body body

swagger:model PatchPollersIdentifierOKBodyBody
*/
type PatchPollersIdentifierOKBodyBody interface{}