package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetCatalogsIdentifierParams creates a new GetCatalogsIdentifierParams object
// with the default values initialized.
func NewGetCatalogsIdentifierParams() *GetCatalogsIdentifierParams {
	return &GetCatalogsIdentifierParams{}
}

/*GetCatalogsIdentifierParams contains all the parameters to send to the API endpoint
for the get catalogs identifier operation typically these are written to a http.Request
*/
type GetCatalogsIdentifierParams struct {

	/*Identifier
	  identifier of a catalog

	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get catalogs identifier params
func (o *GetCatalogsIdentifierParams) WithIdentifier(identifier string) *GetCatalogsIdentifierParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetCatalogsIdentifierParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}