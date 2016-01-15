package post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewPostLookupsParams creates a new PostLookupsParams object
// with the default values initialized.
func NewPostLookupsParams() *PostLookupsParams {
	return &PostLookupsParams{}
}

/*PostLookupsParams contains all the parameters to send to the API endpoint
for the post lookups operation typically these are written to a http.Request
*/
type PostLookupsParams struct {

	/*Content
	  foo

	*/
	Content interface{}
}

// WithContent adds the content to the post lookups params
func (o *PostLookupsParams) WithContent(content interface{}) *PostLookupsParams {
	o.Content = content
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostLookupsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if err := r.SetBodyParam(o.Content); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}