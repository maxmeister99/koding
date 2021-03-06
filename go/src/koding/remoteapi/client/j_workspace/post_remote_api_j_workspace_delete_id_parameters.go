package j_workspace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJWorkspaceDeleteIDParams creates a new PostRemoteAPIJWorkspaceDeleteIDParams object
// with the default values initialized.
func NewPostRemoteAPIJWorkspaceDeleteIDParams() *PostRemoteAPIJWorkspaceDeleteIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJWorkspaceDeleteIDParamsWithTimeout creates a new PostRemoteAPIJWorkspaceDeleteIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJWorkspaceDeleteIDParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJWorkspaceDeleteIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteIDParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJWorkspaceDeleteIDParamsWithContext creates a new PostRemoteAPIJWorkspaceDeleteIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJWorkspaceDeleteIDParamsWithContext(ctx context.Context) *PostRemoteAPIJWorkspaceDeleteIDParams {
	var ()
	return &PostRemoteAPIJWorkspaceDeleteIDParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJWorkspaceDeleteIDParams contains all the parameters to send to the API endpoint
for the post remote API j workspace delete ID operation typically these are written to a http.Request
*/
type PostRemoteAPIJWorkspaceDeleteIDParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJWorkspaceDeleteIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) WithContext(ctx context.Context) *PostRemoteAPIJWorkspaceDeleteIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) WithBody(body models.DefaultSelector) *PostRemoteAPIJWorkspaceDeleteIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) WithID(id string) *PostRemoteAPIJWorkspaceDeleteIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the post remote API j workspace delete ID params
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJWorkspaceDeleteIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
