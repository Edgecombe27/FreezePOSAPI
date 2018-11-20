package common

import (
	"os"
	"github.com/edgecombe27/freeze_pos_api/menu"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

var DB *pg.DB
func Init() *pg.DB {
	DB = pg.Connect(&pg.Options{
		User:     os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Addr:     os.Getenv("POSTGRES_ADDRESS"),
	})
	err := CreateSchema()
	if err != nil {
		panic(err)
	}
	return DB
}

func CreateSchema() error {
	for _, model := range []interface{}{(*menu.Menu)(nil)} {
		err := DB.CreateTable(model, &orm.CreateTableOptions{IfNotExists: true})
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDB() *pg.DB {
	return DB
}
