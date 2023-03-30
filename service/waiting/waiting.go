package waiting

import (
	"nwd/model/request"
	repo "nwd/repository/waiting"
)

type waiting struct {
	repo repo.IWaitingRepo
}

type IWaiting interface {
	Create(req request.Waiting) error
}

var w = waiting{}

func GetWaitingSrv() IWaiting {
	return &w
}

func Init(repo repo.IWaitingRepo) {
	w.repo = repo
}
