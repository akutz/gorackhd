package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewPostNodesIdentifierObmParams creates a new PostNodesIdentifierObmParams object
// with the default values initialized.
func NewPostNodesIdentifierObmParams() *PostNodesIdentifierObmParams {
	return &PostNodesIdentifierObmParams{}
}

/*PostNodesIdentifierObmParams contains all the parameters to send to the API endpoint
for the post nodes identifier obm operation typically these are written to a http.Request
*/
type PostNodesIdentifierObmParams struct {

	/*Body
	  obm settings to apply.


	*/
	Body interface{}
	/*Identifier
	  Mac addresses and unique aliases to identify the node by, |
	expect a string or an array of strings.


	*/
	Identifier string
}

// WithBody adds the body to the post nodes identifier obm params
func (o *PostNodesIdentifierObmParams) WithBody(body interface{}) *PostNodesIdentifierObmParams {
	o.Body = body
	return o
}

// WithIdentifier adds the identifier to the post nodes identifier obm params
func (o *PostNodesIdentifierObmParams) WithIdentifier(identifier string) *PostNodesIdentifierObmParams {
	o.Identifier = identifier
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostNodesIdentifierObmParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param identifier
	if err := r.SetPathParam("identifier", o.Identifier); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}