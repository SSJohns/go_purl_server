package main

import (
	"fmt"
	"io/ioutil"
  "encoding/json"
	"log"
	"net/http"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

type PURL struct {
  purl_id int
  note string
  file_name string
  last_accessed
  local_url string
  redirect_uri string
}

func (p *URL) save() error {
  fileurl := p.local_url
}

func main() {
  db, err := sql.Open("mysql",
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  if db.Ping() {
    fmt.Println("good")
  } else{
    log.Fatal(err)
  }
}
