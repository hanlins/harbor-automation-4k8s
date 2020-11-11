// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/szlabs/harbor-automation-4k8s/pkg/sdk/harbor/models"
)

// PostProjectsProjectIDRobotsReader is a Reader for the PostProjectsProjectIDRobots structure.
type PostProjectsProjectIDRobotsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostProjectsProjectIDRobotsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostProjectsProjectIDRobotsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostProjectsProjectIDRobotsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostProjectsProjectIDRobotsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostProjectsProjectIDRobotsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostProjectsProjectIDRobotsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostProjectsProjectIDRobotsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostProjectsProjectIDRobotsCreated creates a PostProjectsProjectIDRobotsCreated with default headers values
func NewPostProjectsProjectIDRobotsCreated() *PostProjectsProjectIDRobotsCreated {
	return &PostProjectsProjectIDRobotsCreated{}
}

/*PostProjectsProjectIDRobotsCreated handles this case with default header values.

Project member created successfully.
*/
type PostProjectsProjectIDRobotsCreated struct {
	/*The URL of the created resource
	 */
	Location string

	Payload *models.RobotAccountPostRep
}

func (o *PostProjectsProjectIDRobotsCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsCreated  %+v", 201, o.Payload)
}

func (o *PostProjectsProjectIDRobotsCreated) GetPayload() *models.RobotAccountPostRep {
	return o.Payload
}

func (o *PostProjectsProjectIDRobotsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	o.Payload = new(models.RobotAccountPostRep)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostProjectsProjectIDRobotsBadRequest creates a PostProjectsProjectIDRobotsBadRequest with default headers values
func NewPostProjectsProjectIDRobotsBadRequest() *PostProjectsProjectIDRobotsBadRequest {
	return &PostProjectsProjectIDRobotsBadRequest{}
}

/*PostProjectsProjectIDRobotsBadRequest handles this case with default header values.

Project id is not valid.
*/
type PostProjectsProjectIDRobotsBadRequest struct {
}

func (o *PostProjectsProjectIDRobotsBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsBadRequest ", 400)
}

func (o *PostProjectsProjectIDRobotsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDRobotsUnauthorized creates a PostProjectsProjectIDRobotsUnauthorized with default headers values
func NewPostProjectsProjectIDRobotsUnauthorized() *PostProjectsProjectIDRobotsUnauthorized {
	return &PostProjectsProjectIDRobotsUnauthorized{}
}

/*PostProjectsProjectIDRobotsUnauthorized handles this case with default header values.

User need to log in first.
*/
type PostProjectsProjectIDRobotsUnauthorized struct {
}

func (o *PostProjectsProjectIDRobotsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsUnauthorized ", 401)
}

func (o *PostProjectsProjectIDRobotsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDRobotsForbidden creates a PostProjectsProjectIDRobotsForbidden with default headers values
func NewPostProjectsProjectIDRobotsForbidden() *PostProjectsProjectIDRobotsForbidden {
	return &PostProjectsProjectIDRobotsForbidden{}
}

/*PostProjectsProjectIDRobotsForbidden handles this case with default header values.

User in session does not have permission to the project.
*/
type PostProjectsProjectIDRobotsForbidden struct {
}

func (o *PostProjectsProjectIDRobotsForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsForbidden ", 403)
}

func (o *PostProjectsProjectIDRobotsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDRobotsConflict creates a PostProjectsProjectIDRobotsConflict with default headers values
func NewPostProjectsProjectIDRobotsConflict() *PostProjectsProjectIDRobotsConflict {
	return &PostProjectsProjectIDRobotsConflict{}
}

/*PostProjectsProjectIDRobotsConflict handles this case with default header values.

An robot account with same name already exist in the project.
*/
type PostProjectsProjectIDRobotsConflict struct {
}

func (o *PostProjectsProjectIDRobotsConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsConflict ", 409)
}

func (o *PostProjectsProjectIDRobotsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostProjectsProjectIDRobotsInternalServerError creates a PostProjectsProjectIDRobotsInternalServerError with default headers values
func NewPostProjectsProjectIDRobotsInternalServerError() *PostProjectsProjectIDRobotsInternalServerError {
	return &PostProjectsProjectIDRobotsInternalServerError{}
}

/*PostProjectsProjectIDRobotsInternalServerError handles this case with default header values.

Unexpected internal errors.
*/
type PostProjectsProjectIDRobotsInternalServerError struct {
}

func (o *PostProjectsProjectIDRobotsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_id}/robots][%d] postProjectsProjectIdRobotsInternalServerError ", 500)
}

func (o *PostProjectsProjectIDRobotsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}