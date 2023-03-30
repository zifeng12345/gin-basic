package waiting

import srv "nwd/service/waiting"

var wctl waitingCtl

type waitingCtl struct {
	service srv.IWaiting
}

func GetWaitingCtl() *waitingCtl {
	return &wctl
}

func Init(srv srv.IWaiting) {
	wctl.service = srv
}
