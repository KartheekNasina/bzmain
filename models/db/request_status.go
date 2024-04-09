// requestStatus.go
package models

type RequestStatus string

const (
	Requested RequestStatus = "requested"
	Approved  RequestStatus = "approved"
	Declined  RequestStatus = "declined"
	Confirmed RequestStatus = "confirmed"
	Canceled  RequestStatus = "canceled"
)
