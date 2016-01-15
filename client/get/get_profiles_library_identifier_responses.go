package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type GetProfilesLibraryIdentifierReader struct {
	formats strfmt.Registry
}

func (o *GetProfilesLibraryIdentifierReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetProfilesLibraryIdentifierOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetProfilesLibraryIdentifierNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetProfilesLibraryIdentifierDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetProfilesLibraryIdentifierOK creates a GetProfilesLibraryIdentifierOK with default headers values
func NewGetProfilesLibraryIdentifierOK() *GetProfilesLibraryIdentifierOK {
	return &GetProfilesLibraryIdentifierOK{}
}

/*GetProfilesLibraryIdentifierOK

return profile

*/
type GetProfilesLibraryIdentifierOK struct {
	Payload GetProfilesLibraryIdentifierOKBodyBody
}

func (o *GetProfilesLibraryIdentifierOK) Error() string {
	return fmt.Sprintf("[GET /profiles/library/{identifier}][%d] getProfilesLibraryIdentifierOK  %+v", 200, o.Payload)
}

func (o *GetProfilesLibraryIdentifierOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetProfilesLibraryIdentifierNotFound creates a GetProfilesLibraryIdentifierNotFound with default headers values
func NewGetProfilesLibraryIdentifierNotFound() *GetProfilesLibraryIdentifierNotFound {
	return &GetProfilesLibraryIdentifierNotFound{}
}

/*GetProfilesLibraryIdentifierNotFound

There is no profile in the library with identifier.

*/
type GetProfilesLibraryIdentifierNotFound struct {
	Payload *models.Error
}

func (o *GetProfilesLibraryIdentifierNotFound) Error() string {
	return fmt.Sprintf("[GET /profiles/library/{identifier}][%d] getProfilesLibraryIdentifierNotFound  %+v", 404, o.Payload)
}

func (o *GetProfilesLibraryIdentifierNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetProfilesLibraryIdentifierDefault creates a GetProfilesLibraryIdentifierDefault with default headers values
func NewGetProfilesLibraryIdentifierDefault(code int) *GetProfilesLibraryIdentifierDefault {
	return &GetProfilesLibraryIdentifierDefault{
		_statusCode: code,
	}
}

/*GetProfilesLibraryIdentifierDefault

Unexpected error
*/
type GetProfilesLibraryIdentifierDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get profiles library identifier default response
func (o *GetProfilesLibraryIdentifierDefault) Code() int {
	return o._statusCode
}

func (o *GetProfilesLibraryIdentifierDefault) Error() string {
	return fmt.Sprintf("[GET /profiles/library/{identifier}][%d] GetProfilesLibraryIdentifier default  %+v", o._statusCode, o.Payload)
}

func (o *GetProfilesLibraryIdentifierDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

/*GetProfilesLibraryIdentifierOKBodyBody GetProfilesLibraryIdentifierOKBodyBody get profiles library identifier o k body body

swagger:model GetProfilesLibraryIdentifierOKBodyBody
*/
type GetProfilesLibraryIdentifierOKBodyBody interface{}