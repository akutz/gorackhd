package dhcp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/emccode/gorackhd/models"
)

type GetDhcpLeaseMacReader struct {
	formats strfmt.Registry
}

func (o *GetDhcpLeaseMacReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDhcpLeaseMacOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDhcpLeaseMacDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewGetDhcpLeaseMacOK creates a GetDhcpLeaseMacOK with default headers values
func NewGetDhcpLeaseMacOK() *GetDhcpLeaseMacOK {
	return &GetDhcpLeaseMacOK{}
}

/*GetDhcpLeaseMacOK

A single lease
*/
type GetDhcpLeaseMacOK struct {
	Payload []*models.Lease
}

func (o *GetDhcpLeaseMacOK) Error() string {
	return fmt.Sprintf("[GET /dhcp/lease/{mac}][%d] getDhcpLeaseMacOK  %+v", 200, o.Payload)
}

func (o *GetDhcpLeaseMacOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetDhcpLeaseMacDefault creates a GetDhcpLeaseMacDefault with default headers values
func NewGetDhcpLeaseMacDefault(code int) *GetDhcpLeaseMacDefault {
	return &GetDhcpLeaseMacDefault{
		_statusCode: code,
	}
}

/*GetDhcpLeaseMacDefault

NotFound error
*/
type GetDhcpLeaseMacDefault struct {
	_statusCode int
}

// Code gets the status code for the get dhcp lease mac default response
func (o *GetDhcpLeaseMacDefault) Code() int {
	return o._statusCode
}

func (o *GetDhcpLeaseMacDefault) Error() string {
	return fmt.Sprintf("[GET /dhcp/lease/{mac}][%d] GetDhcpLeaseMac default ", o._statusCode)
}

func (o *GetDhcpLeaseMacDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}