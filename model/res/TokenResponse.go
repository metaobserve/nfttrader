package res

import "github.com/dometa/model"

type TokenResponse struct {
	Token       string           `json:token`
	Status      model.StatusType `json:status`
	Description string           `json:description`
}
