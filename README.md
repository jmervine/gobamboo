# gobamboo

Go Lib for making HTTP GET requests to a Bamboo server and working with the
results as structs.

``` golang
package main

// example and test

import (
  "fmt"
  "github.com/jmervine/gobamboo" // package is bamboo
)

func main() {
  b    := new(bamboo.Bamboo)
  opts := make(map[string]string)

  b.Domain   = "http://bamboo.example.com"
  b.Username = "USERNAME"
  b.Password = "PASSWORD"

  // info
  info, _ := b.Info(opts) // returns *bamboo.Info
  fmt.Printf("----------\n%#v\n----------\n", info)

  // b.Url must now be reset or regenerated
  //
  //  Note: I had to build it this way for testing. I need to revist this and
  //        see if I can find an effective way around this, that still allows
  //        for unit testing.
  //
  // Reset example:
  b.Url = ""

  // plan
  plan, _ := b.Plan(opts) // returns *bamboo.Plan
  fmt.Printf("----------\n%#v\n----------\n", plan)

  // b.Url must now be reset or regenerated
  //
  // Regenerate example:
  b.GenerateUrl("project")

  // project
  project, _ := b.Project(opts) // returns *bamboo.Project
  fmt.Printf("----------\n%#v\n----------\n", project)

  // queue
  b.GenerateUrl("queue")
  queue, _ := b.Queue(opts) // returns *bamboo.Queue
  fmt.Printf("----------\n%#v\n----------\n", queue)ambo

  // result
  b.GenerateUrl("result")
  result, _ := b.Result(opts) // returns *bamboo.Result
  fmt.Printf("----------\n%#v\n----------\n", result)
}
```

See [example.go](example.go) for a more useful full example.
