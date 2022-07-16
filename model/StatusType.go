package model

type StatusType int

const (
	StatusType_SUCCESS      StatusType = 200
	StatusType_NOTFUND      StatusType = 300
	StatusType_NOPERMISSION StatusType = 400
	StatusType_EXCEPTION    StatusType = 500
)

func (status StatusType) String() string {
	switch status {
	case StatusType_SUCCESS:
		return "SUCCESS"
	case StatusType_NOTFUND:
		return "NOTFUND"
	case StatusType_NOPERMISSION:
		return "NOPERMISSION"
	case StatusType_EXCEPTION:
		return "EXCEPTION"
	default:
		return "EXCEPTION"
	}
}
