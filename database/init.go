package database

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	database, err := gorm.Open("sqlite3", "./gors.db")
	if err != nil {
		fmt.Printf("opdate sqlite error : %s", err.Error())
		os.Exit(1)
	}
	db = database

	// Migrate the schema
	db.AutoMigrate(&Record{})
}
