package res

import "github.com/dometa/model"

type MintResponse struct {
	Status      model.StatusType `json:status`
	Description string           `json:description`
}
