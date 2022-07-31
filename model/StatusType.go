package model

type StatusType int

const (
	StatusType_SUCCESS     StatusType = 200
	StatusType_SOURCEMOVED StatusType = 300
	StatusType_CUSTOMERROR StatusType = 400
	StatusType_FAILURE     StatusType = 500
)

func (status StatusType) String() string {
	switch status {
	case StatusType_SUCCESS:
		return "SUCCESS"
	case StatusType_SOURCEMOVED:
		return "SOURCEMOVED"
	case StatusType_CUSTOMERROR:
		return "CUSTOMERROR"
	case StatusType_FAILURE:
		return "FAILURE"
	default:
		return "SERVEREXCEPTION"
	}
}
