package godbtt

import (
	"fmt"

	"github.com/tidwall/buntdb"
)

type DBTT struct {
	option Option
	db     *buntdb.DB
}

func NewDBTT(option Option) *DBTT {
	return &DBTT{
		option: option,
	}
}

func (dbtt *DBTT) Start() error {
	fmt.Println("DBTT Start")

	db, err := buntdb.Open(":memory:")
	if err != nil {
		fmt.Println("DBTT Start err", err)
		return err
	}
	dbtt.db = db
	return nil
}

func (dbtt *DBTT) Stop() {
	fmt.Println("DBTT Stop")
	if dbtt.db != nil {
		dbtt.db.Close()
		dbtt.db = nil
	}
}

func Test() {
	fmt.Println("111111111111111")
}
