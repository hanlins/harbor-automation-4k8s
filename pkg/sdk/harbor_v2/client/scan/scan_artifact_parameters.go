// Code generated by go-swagger; DO NOT EDIT.

package scan

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
)

// NewScanArtifactParams creates a new ScanArtifactParams object
// with the default values initialized.
func NewScanArtifactParams() *ScanArtifactParams {
	var ()
	return &ScanArtifactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewScanArtifactParamsWithTimeout creates a new ScanArtifactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewScanArtifactParamsWithTimeout(timeout time.Duration) *ScanArtifactParams {
	var ()
	return &ScanArtifactParams{

		timeout: timeout,
	}
}

// NewScanArtifactParamsWithContext creates a new ScanArtifactParams object
// with the default values initialized, and the ability to set a context for a request
func NewScanArtifactParamsWithContext(ctx context.Context) *ScanArtifactParams {
	var ()
	return &ScanArtifactParams{

		Context: ctx,
	}
}

// NewScanArtifactParamsWithHTTPClient creates a new ScanArtifactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewScanArtifactParamsWithHTTPClient(client *http.Client) *ScanArtifactParams {
	var ()
	return &ScanArtifactParams{
		HTTPClient: client,
	}
}

/*ScanArtifactParams contains all the parameters to send to the API endpoint
for the scan artifact operation typically these are written to a http.Request
*/
type ScanArtifactParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectName
	  The name of the project

	*/
	ProjectName string
	/*Reference
	  The reference of the artifact, can be digest or tag

	*/
	Reference string
	/*RepositoryName
	  The name of the repository. If it contains slash, encode it with URL encoding. e.g. a/b -> a%252Fb

	*/
	RepositoryName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the scan artifact params
func (o *ScanArtifactParams) WithTimeout(timeout time.Duration) *ScanArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the scan artifact params
func (o *ScanArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the scan artifact params
func (o *ScanArtifactParams) WithContext(ctx context.Context) *ScanArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the scan artifact params
func (o *ScanArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the scan artifact params
func (o *ScanArtifactParams) WithHTTPClient(client *http.Client) *ScanArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the scan artifact params
func (o *ScanArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the scan artifact params
func (o *ScanArtifactParams) WithXRequestID(xRequestID *string) *ScanArtifactParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the scan artifact params
func (o *ScanArtifactParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectName adds the projectName to the scan artifact params
func (o *ScanArtifactParams) WithProjectName(projectName string) *ScanArtifactParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the scan artifact params
func (o *ScanArtifactParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithReference adds the reference to the scan artifact params
func (o *ScanArtifactParams) WithReference(reference string) *ScanArtifactParams {
	o.SetReference(reference)
	return o
}

// SetReference adds the reference to the scan artifact params
func (o *ScanArtifactParams) SetReference(reference string) {
	o.Reference = reference
}

// WithRepositoryName adds the repositoryName to the scan artifact params
func (o *ScanArtifactParams) WithRepositoryName(repositoryName string) *ScanArtifactParams {
	o.SetRepositoryName(repositoryName)
	return o
}

// SetRepositoryName adds the repositoryName to the scan artifact params
func (o *ScanArtifactParams) SetRepositoryName(repositoryName string) {
	o.RepositoryName = repositoryName
}

// WriteToRequest writes these params to a swagger request
func (o *ScanArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param project_name
	if err := r.SetPathParam("project_name", o.ProjectName); err != nil {
		return err
	}

	// path param reference
	if err := r.SetPathParam("reference", o.Reference); err != nil {
		return err
	}

	// path param repository_name
	if err := r.SetPathParam("repository_name", o.RepositoryName); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}