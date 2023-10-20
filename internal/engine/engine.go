package engine

import (
	"github.com/alwindoss/dux/internal/feedback"
	"github.com/alwindoss/dux/internal/user"
	"go.etcd.io/bbolt"
)

func Run(dbLoc string) error {
	db, err := bbolt.Open(dbLoc, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()

	feedMgr := feedback.New(db)
	userMgr := user.New(db)
	runApp(feedMgr, userMgr)
	return nil
}
