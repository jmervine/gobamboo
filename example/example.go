package main

// example and test

import (
  "github.com/jmervine/gobamboo"
  "fmt"
  "flag"
  "os"
)

func main() {

  var domain, username, password, query string
  flag.StringVar(&domain,   "d", "", "bamboo domain name")
  flag.StringVar(&username, "u", "", "bamboo username")
  flag.StringVar(&password, "p", "", "bamboo password")
  flag.StringVar(&query,    "q", "", "bamboo api to query")
  flag.Parse()

  flag.Usage = func() {
    fmt.Fprintf(os.Stderr, "Usage of 'example':\n")
    flag.PrintDefaults()
  }

  if domain == "" || username == "" || password == "" || query == "" {
    flag.Usage()
    os.Exit(0)
  }

  b    := new(bamboo.Bamboo)
  opts := make(map[string]string)

  b.Domain   = domain
  b.Username = username
  b.Password = password

  fn := func (q string, o map[string]string) (interface{}, error) {
    switch(q) {
    case "plan":
      return b.Plan(o)
    case "project":
      return b.Project(o)
    case "queue":
      return b.Queue(o)
    case "result":
      return b.Result(o)
    default:
      return b.Info(o)
    }
  }

  result, _ := fn(query)
  fmt.Printf("----------\n%#v\n----------\n", result)
}
