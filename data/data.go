package data

import (
  "database/sql"
)

var Db *sql.DB

func init() {
  var err error
  Db, err = sql.Open("postgres", "dbname=myapp sslmode=disable")
  if err != nil {
    log.Fatal(err)
  }
  return
}
