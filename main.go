//Goapi is an api build on top of mgo and martini intended to be RESTful and work out of the box.
package main

import (
  "github.com/go-martini/martini"
  "labix.org/v2/mgo"
  "log"
  "net/http"
  "encoding/json"
)

type jsn map[string]interface{}

func main() {
  m := martini.Classic()

  sess, err := mgo.Dial("mongodb://172.18.0.3")
  if(err != nil) {
    log.Fatal(err)
  }

  db := sess.DB("test")

  m.Map(db) //Let martini inject the db instance

  m.Post("/:coll", func (d *mgo.Database, req *http.Request, p martini.Params) (int, string) {
    var j jsn

    de := json.NewDecoder(req.Body)

    err1 := de.Decode(&j)
    if(err1 != nil) {
      return 400, "Bad JSON\n"
    }

    err2 := db.C(p["coll"]).Insert(j)
    if(err2 != nil) {
      return 500, "Server error"
    } else {
      return 200, "OKAY\n"
    }
  })

  m.Get("/:coll", func (d *mgo.Database, r http.ResponseWriter, p martini.Params) {

    var res []interface{}
    db.C(p["coll"]).Find(nil).All(&res)

    enc := json.NewEncoder(r);

    enc.Encode(res);
  })

  m.Run()
}
