// Code generated by go-swagger; DO NOT EDIT.

package terminal_audit_log

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/cgascoig/intersight-go/cli/models"
)

// GetTerminalAuditLogsReader is a Reader for the GetTerminalAuditLogs structure.
type GetTerminalAuditLogsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTerminalAuditLogsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTerminalAuditLogsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewGetTerminalAuditLogsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetTerminalAuditLogsOK creates a GetTerminalAuditLogsOK with default headers values
func NewGetTerminalAuditLogsOK() *GetTerminalAuditLogsOK {
	return &GetTerminalAuditLogsOK{}
}

/*GetTerminalAuditLogsOK handles this case with default header values.

List of terminalAuditLogs for the given filter criteria
*/
type GetTerminalAuditLogsOK struct {
	Payload *models.TerminalAuditLogList
}

func (o *GetTerminalAuditLogsOK) Error() string {
	return fmt.Sprintf("[GET /terminal/AuditLogs][%d] getTerminalAuditLogsOK  %+v", 200, o.Payload)
}

func (o *GetTerminalAuditLogsOK) GetPayload() *models.TerminalAuditLogList {
	return o.Payload
}

func (o *GetTerminalAuditLogsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TerminalAuditLogList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTerminalAuditLogsDefault creates a GetTerminalAuditLogsDefault with default headers values
func NewGetTerminalAuditLogsDefault(code int) *GetTerminalAuditLogsDefault {
	return &GetTerminalAuditLogsDefault{
		_statusCode: code,
	}
}

/*GetTerminalAuditLogsDefault handles this case with default header values.

Unexpected error
*/
type GetTerminalAuditLogsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the get terminal audit logs default response
func (o *GetTerminalAuditLogsDefault) Code() int {
	return o._statusCode
}

func (o *GetTerminalAuditLogsDefault) Error() string {
	return fmt.Sprintf("[GET /terminal/AuditLogs][%d] GetTerminalAuditLogs default  %+v", o._statusCode, o.Payload)
}

func (o *GetTerminalAuditLogsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *GetTerminalAuditLogsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}