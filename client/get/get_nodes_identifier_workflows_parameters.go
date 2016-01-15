package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetNodesIdentifierWorkflowsParams creates a new GetNodesIdentifierWorkflowsParams object
// with the default values initialized.
func NewGetNodesIdentifierWorkflowsParams() *GetNodesIdentifierWorkflowsParams {
	return &GetNodesIdentifierWorkflowsParams{}
}

/*GetNodesIdentifierWorkflowsParams contains all the parameters to send to the API endpoint
for the get nodes identifier workflows operation typically these are written to a http.Request
*/
type GetNodesIdentifierWorkflowsParams struct {

	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
}

// WithIdentifier adds the identifier to the get nodes identifier workflows params
func (o *GetNodesIdentifierWorkflowsParams) WithIdentifier(identifier string) *GetNodesIdentifierWorkflowsParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodesIdentifierWorkflowsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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