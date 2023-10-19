package engine

import (
	"fmt"

	"github.com/alwindoss/dux/internal/feedback"
	"github.com/alwindoss/dux/internal/user"
	"github.com/dgraph-io/badger/v4"
)

func Run(dbLoc string) error {
	opt := badger.DefaultOptions(dbLoc)
	db, err := badger.Open(opt)
	if err != nil {
		err = fmt.Errorf("unable to open badger DB: %w", err)
		return err
	}
	defer db.Close()
	feedMgr := feedback.New(db)
	userMgr := user.New(db)
	runApp(feedMgr, userMgr)
	return nil
}
