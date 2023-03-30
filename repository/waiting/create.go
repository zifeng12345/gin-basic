package waiting

import (
	"nwd/model/request"
	"nwd/model/tables"
	"nwd/shared/database"
)

func (w *waitingRepo) Create(db database.IConnection, req request.Waiting) error {
	var waiting tables.Waiting = tables.Waiting{
		StoreId:     req.StoreId,
		UserName:    req.UserName,
		Desc:        req.Desc,
		PhoneNumber: req.PhoneNumber,
		Status:      req.Status,
		Number:      req.Number,
		Day:         req.Day,
		Pax:         req.Pax,
	}

	return db.Tables("waiting_lists").Create(&waiting).Error
}
