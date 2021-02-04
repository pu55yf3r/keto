// Code generated by go-swagger; DO NOT EDIT.

package engines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/ory/keto/internal/httpclient/models"
)

// NewUpsertOryAccessControlPolicyParams creates a new UpsertOryAccessControlPolicyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpsertOryAccessControlPolicyParams() *UpsertOryAccessControlPolicyParams {
	return &UpsertOryAccessControlPolicyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpsertOryAccessControlPolicyParamsWithTimeout creates a new UpsertOryAccessControlPolicyParams object
// with the ability to set a timeout on a request.
func NewUpsertOryAccessControlPolicyParamsWithTimeout(timeout time.Duration) *UpsertOryAccessControlPolicyParams {
	return &UpsertOryAccessControlPolicyParams{
		timeout: timeout,
	}
}

// NewUpsertOryAccessControlPolicyParamsWithContext creates a new UpsertOryAccessControlPolicyParams object
// with the ability to set a context for a request.
func NewUpsertOryAccessControlPolicyParamsWithContext(ctx context.Context) *UpsertOryAccessControlPolicyParams {
	return &UpsertOryAccessControlPolicyParams{
		Context: ctx,
	}
}

// NewUpsertOryAccessControlPolicyParamsWithHTTPClient creates a new UpsertOryAccessControlPolicyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpsertOryAccessControlPolicyParamsWithHTTPClient(client *http.Client) *UpsertOryAccessControlPolicyParams {
	return &UpsertOryAccessControlPolicyParams{
		HTTPClient: client,
	}
}

/* UpsertOryAccessControlPolicyParams contains all the parameters to send to the API endpoint
   for the upsert ory access control policy operation.

   Typically these are written to a http.Request.
*/
type UpsertOryAccessControlPolicyParams struct {

	// Body.
	Body *models.OryAccessControlPolicy

	/* Flavor.

	   The ORY Access Control Policy flavor. Can be "regex", "glob", and "exact".
	*/
	Flavor string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the upsert ory access control policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertOryAccessControlPolicyParams) WithDefaults() *UpsertOryAccessControlPolicyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the upsert ory access control policy params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpsertOryAccessControlPolicyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) WithTimeout(timeout time.Duration) *UpsertOryAccessControlPolicyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) WithContext(ctx context.Context) *UpsertOryAccessControlPolicyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) WithHTTPClient(client *http.Client) *UpsertOryAccessControlPolicyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) WithBody(body *models.OryAccessControlPolicy) *UpsertOryAccessControlPolicyParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) SetBody(body *models.OryAccessControlPolicy) {
	o.Body = body
}

// WithFlavor adds the flavor to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) WithFlavor(flavor string) *UpsertOryAccessControlPolicyParams {
	o.SetFlavor(flavor)
	return o
}

// SetFlavor adds the flavor to the upsert ory access control policy params
func (o *UpsertOryAccessControlPolicyParams) SetFlavor(flavor string) {
	o.Flavor = flavor
}

// WriteToRequest writes these params to a swagger request
func (o *UpsertOryAccessControlPolicyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param flavor
	if err := r.SetPathParam("flavor", o.Flavor); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
