// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostChartrepoRepoChartsNameVersionLabelsReader is a Reader for the PostChartrepoRepoChartsNameVersionLabels structure.
type PostChartrepoRepoChartsNameVersionLabelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChartrepoRepoChartsNameVersionLabelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostChartrepoRepoChartsNameVersionLabelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostChartrepoRepoChartsNameVersionLabelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostChartrepoRepoChartsNameVersionLabelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostChartrepoRepoChartsNameVersionLabelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostChartrepoRepoChartsNameVersionLabelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostChartrepoRepoChartsNameVersionLabelsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostChartrepoRepoChartsNameVersionLabelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostChartrepoRepoChartsNameVersionLabelsOK creates a PostChartrepoRepoChartsNameVersionLabelsOK with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsOK() *PostChartrepoRepoChartsNameVersionLabelsOK {
	return &PostChartrepoRepoChartsNameVersionLabelsOK{}
}

/*PostChartrepoRepoChartsNameVersionLabelsOK handles this case with default header values.

The label is successfully marked to the chart version.
*/
type PostChartrepoRepoChartsNameVersionLabelsOK struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsOK) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsOK ", 200)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsBadRequest creates a PostChartrepoRepoChartsNameVersionLabelsBadRequest with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsBadRequest() *PostChartrepoRepoChartsNameVersionLabelsBadRequest {
	return &PostChartrepoRepoChartsNameVersionLabelsBadRequest{}
}

/*PostChartrepoRepoChartsNameVersionLabelsBadRequest handles this case with default header values.

Bad request
*/
type PostChartrepoRepoChartsNameVersionLabelsBadRequest struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsBadRequest) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsBadRequest ", 400)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsUnauthorized creates a PostChartrepoRepoChartsNameVersionLabelsUnauthorized with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsUnauthorized() *PostChartrepoRepoChartsNameVersionLabelsUnauthorized {
	return &PostChartrepoRepoChartsNameVersionLabelsUnauthorized{}
}

/*PostChartrepoRepoChartsNameVersionLabelsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostChartrepoRepoChartsNameVersionLabelsUnauthorized struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsUnauthorized ", 401)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsForbidden creates a PostChartrepoRepoChartsNameVersionLabelsForbidden with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsForbidden() *PostChartrepoRepoChartsNameVersionLabelsForbidden {
	return &PostChartrepoRepoChartsNameVersionLabelsForbidden{}
}

/*PostChartrepoRepoChartsNameVersionLabelsForbidden handles this case with default header values.

Operation is forbidden or quota exceeded
*/
type PostChartrepoRepoChartsNameVersionLabelsForbidden struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsForbidden) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsForbidden ", 403)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsNotFound creates a PostChartrepoRepoChartsNameVersionLabelsNotFound with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsNotFound() *PostChartrepoRepoChartsNameVersionLabelsNotFound {
	return &PostChartrepoRepoChartsNameVersionLabelsNotFound{}
}

/*PostChartrepoRepoChartsNameVersionLabelsNotFound handles this case with default header values.

Not found
*/
type PostChartrepoRepoChartsNameVersionLabelsNotFound struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsNotFound) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsNotFound ", 404)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsConflict creates a PostChartrepoRepoChartsNameVersionLabelsConflict with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsConflict() *PostChartrepoRepoChartsNameVersionLabelsConflict {
	return &PostChartrepoRepoChartsNameVersionLabelsConflict{}
}

/*PostChartrepoRepoChartsNameVersionLabelsConflict handles this case with default header values.

Conflicts
*/
type PostChartrepoRepoChartsNameVersionLabelsConflict struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsConflict) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsConflict ", 409)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsNameVersionLabelsInternalServerError creates a PostChartrepoRepoChartsNameVersionLabelsInternalServerError with default headers values
func NewPostChartrepoRepoChartsNameVersionLabelsInternalServerError() *PostChartrepoRepoChartsNameVersionLabelsInternalServerError {
	return &PostChartrepoRepoChartsNameVersionLabelsInternalServerError{}
}

/*PostChartrepoRepoChartsNameVersionLabelsInternalServerError handles this case with default header values.

Internal server error occurred
*/
type PostChartrepoRepoChartsNameVersionLabelsInternalServerError struct {
}

func (o *PostChartrepoRepoChartsNameVersionLabelsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts/{name}/{version}/labels][%d] postChartrepoRepoChartsNameVersionLabelsInternalServerError ", 500)
}

func (o *PostChartrepoRepoChartsNameVersionLabelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}