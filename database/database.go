package database

/*
	If you want to use database, just uncomment the code below.
*/

// import (
// 	"database/sql"
// 	"errors"
// 	"fmt"

// 	"goink/config"

// 	_ "github.com/lib/pq"
// )

// var db *sql.DB

// func init() {
// 	psqlInfo := fmt.Sprintf(
// 		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
// 		config.Database.Host,
// 		config.Database.Port,
// 		config.Database.User,
// 		config.Database.Password,
// 		config.Database.Name,
// 	)

// 	preDB, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	err = preDB.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	db = preDB
// }
