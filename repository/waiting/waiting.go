package waiting

import (
	"nwd/model/request"
	"nwd/shared/database"
)

type waitingRepo struct {
}

type IWaitingRepo interface {
	Create(db database.IConnection, req request.Waiting) error
}

var w = waitingRepo{}

func GetWaitingRepo() IWaitingRepo {
	return &w
}
