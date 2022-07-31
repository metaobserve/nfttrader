package res

import "github.com/dometa/model"

type WalletLoginResponse struct {
	Status      model.StatusType `json:status`
	Description string           `json:description`
}
