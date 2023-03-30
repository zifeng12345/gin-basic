package waiting

import (
	"nwd/model/request"
	"nwd/shared/database"
)

func (w *waiting) Create(req request.Waiting) error {
	w.repo.Create(database.GetConnection(), req)
	return nil
}
