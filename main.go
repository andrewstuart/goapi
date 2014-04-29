package main

import (
  "github.com/go-martini/martini"
  "labix.org/v2/mgo"
  "fmt"
  "net/http"
  "encoding/json"
)

type jsn map[string]interface{}

func main() {
  m := martini.Classic()

  sess, err := mgo.Dial("mongodb://172.17.0.2")
  if(err != nil) {
    fmt.Println(err)
  }

  db := sess.DB("test")

  m.Map(db) //Let martini inject the db instance

  m.Get("/", func (d *mgo.Database, r http.ResponseWriter) {
    db.C("hello").Insert(jsn{"foo": "Hello"})

    var res []interface{}
    db.C("hello").Find(nil).All(&res)

    enc := json.NewEncoder(r);

    enc.Encode(res);
  })

  m.Run()
}
