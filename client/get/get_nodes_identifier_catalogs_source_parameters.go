package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetNodesIdentifierCatalogsSourceParams creates a new GetNodesIdentifierCatalogsSourceParams object
// with the default values initialized.
func NewGetNodesIdentifierCatalogsSourceParams() *GetNodesIdentifierCatalogsSourceParams {
	return &GetNodesIdentifierCatalogsSourceParams{}
}

/*GetNodesIdentifierCatalogsSourceParams contains all the parameters to send to the API endpoint
for the get nodes identifier catalogs source operation typically these are written to a http.Request
*/
type GetNodesIdentifierCatalogsSourceParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
	/*Source
	  Source catalog name to fetch


	*/
	Source string
}

// WithIdentifier adds the identifier to the get nodes identifier catalogs source params
func (o *GetNodesIdentifierCatalogsSourceParams) WithIdentifier(identifier string) *GetNodesIdentifierCatalogsSourceParams {
	o.Identifier = identifier
	return o
}

// WithSource adds the source to the get nodes identifier catalogs source params
func (o *GetNodesIdentifierCatalogsSourceParams) WithSource(source string) *GetNodesIdentifierCatalogsSourceParams {
	o.Source = source
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierCatalogsSourceParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	// path param source
	if err := r.SetPathParam("source", o.Source); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}