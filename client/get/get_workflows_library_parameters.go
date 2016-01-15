package get

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
)

// NewGetWorkflowsLibraryParams creates a new GetWorkflowsLibraryParams object
// with the default values initialized.
func NewGetWorkflowsLibraryParams() *GetWorkflowsLibraryParams {
	return &GetWorkflowsLibraryParams{}
}

/*GetWorkflowsLibraryParams contains all the parameters to send to the API endpoint
for the get workflows library operation typically these are written to a http.Request
*/
type GetWorkflowsLibraryParams struct {
}

// WriteToRequest writes these params to a swagger request
func (o *GetWorkflowsLibraryParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}