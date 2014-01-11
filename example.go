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

  switch(query) {
  case "info":
    info, _ := b.Info(opts)
    fmt.Printf("----------\n%#v\n----------\n", info)

  case "plan":
    plan, _ := b.Plan(opts)
    fmt.Printf("----------\n%#v\n----------\n", plan)

  case "project":
    project, _ := b.Project(opts)
    fmt.Printf("----------\n%#v\n----------\n", project)

  case "queue":
    queue, _ := b.Queue(opts)
    fmt.Printf("----------\n%#v\n----------\n", queue)

  case "result":
    result, _ := b.Result(opts)
    fmt.Printf("----------\n%#v\n----------\n", result)

  }
}
