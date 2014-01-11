package bamboo

import (
  "fmt"
  "errors"
  "strings"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

const PATH     = string("/rest/api/latest")
const EXT      = string("json")

/***
 * JSON Objects SHARED
 ***/
type LinkType struct {
  Href string `json:"href"`
  Rel  string `json:"rel"`
}

type PlanType struct {
  Name        string   `json:"name"`
  ShortName   string   `json:"shortName"`
  Key         string   `json:"key"`
  ShortKey    string   `json:"shortKey"`
  Enabled     bool     `json:"enabled"`
  Type        string   `json:"type"`
  Link        LinkType
  PlanKey     PlanKeyType
}

type PlanKeyType struct {
  Key string `json:"key"`
}

/***
 * JSON Objects for /info
 ***/
type Info struct {
  Version       string   `json:"version"`
  Edition       string   `json:"edition"`
  BuildDate     string   `json:"buildDate"`
  BuildNumber   string   `json:"buildNumber"`
  State         string   `json:"state"`
}

/***
 * JSON Objects for /result
 ***/
type Result struct {
  Expand   string `json:"expand"`
  Link     LinkType
  Results  ResultsType
}

type ResultsType struct {
  Size       int    `json:"size"`
  Expand     string `json:"expand"`
  StartIndex int    `json:"start-index"`
  MaxResults int    `json:"max-result"`
  Result     []ResultType
}

type ResultType struct {
  Id             int      `json:"id"`
  Key            string   `json:"key"`
  State          string   `json:"state"`
  LifeCycleState string   `json:"lifeCycleState"`
  Link           LinkType
  Plan           PlanType
  PlanResultKey  PlanResultKeyType
}

type PlanResultKeyType struct {
  Key            string `json:"key"`
  ResultNumber   int    `json:"resultNumber"`
  EntityKey      EntityKeyType
}

type EntityKeyType struct {
  Key string `json:"key"`
}

/***
 * JSON Objects for /queue
 ***/
type Queue struct {
  Expand   string `json:"expand"`
  Link     LinkType
  QueuedBuilds QueuedBuildsType
}

type QueuedBuildsType struct {
  Size       int    `json:"size"`
  StartIndex int    `json:"start-index"`
  MaxResults int    `json:"max-result"`
  Queue      []QueueType
}

type QueueType struct {
  Name        string   `json:"name"`
  Key         string   `json:"key"`
  Link        LinkType
}

/***
 * JSON Objects for /project
 ***/
type Project struct {
  Expand   string `json:"expand"`
  Link     LinkType
  Projects ProjectsType
}

type ProjectsType struct {
  Size       int    `json:"size"`
  Expand     string `json:"expand"`
  StartIndex int    `json:"start-index"`
  MaxResults int    `json:"max-result"`
  Project    []ProjectType
}

type ProjectType struct {
  Name        string   `json:"name"`
  Key         string   `json:"key"`
  Link        LinkType
}

/***
 * JSON Objects for /plan
 ***/
type Plan struct {
  Expand string `json:"expand"`
  Link   LinkType
  Plans  PlansType
}

type PlansType struct {
  Size       int    `json:"size"`
  Expand     string `json:"expand"`
  StartIndex int    `json:"start-index"`
  MaxResults int    `json:"max-result"`
  Plan       []PlanType
}

/***
 * end JSON Objects
 ***/

func parse(j []byte, s interface{}) (interface{}, error) {
  err := json.Unmarshal(j, &s)
  return s, err
}

func Request(bamboo *Bamboo) (contents []byte, err error) {
  response, err := http.Get(bamboo.Url)
  if err != nil {
    return
  }

  defer response.Body.Close()

  if response.StatusCode != 200 {
    err = errors.New(string(response.Status))
    return
  }

  contents, err = ioutil.ReadAll(response.Body)
  return
}

type Bamboo struct {
  Domain   string
  Username string
  Password string
  Url      string
}

func (bamboo *Bamboo) GenerateUrl(api string, opts map[string]string) (err error, url string) {
  if bamboo.Domain == "" {
    err = errors.New("Bamboo domain is required.")
    return
  }
  domain := bamboo.Domain
  if !strings.HasPrefix(domain, "https://") && !strings.HasPrefix(domain, "http://") {
    domain = "http://"+domain
  }

  if bamboo.Username == "" || bamboo.Password == "" {
    err = errors.New("Bamboo Username and Password are required.")
    return
  }

  if !strings.HasPrefix(api, "/") {
    api = "/"+api
  }

  ext := EXT
  if ext != "" && !strings.HasPrefix(ext, ".") {
    ext = "."+ext
  }

  url = domain + PATH + api + ext + "?os_authType=basic&os_username=" + bamboo.Username + "&os_password=" + bamboo.Password

  // flatten opts
  for k,v := range(opts) {
    url = url + "&"+k+"="+v
  }

  bamboo.Url = url
  return
}

func checkError(err error) {
  if err != nil {
    fmt.Printf("Error: ", err)
  }
}

func (bamboo *Bamboo) Info(opts map[string]string) (*Info, error) {
  if bamboo.Url == "" {
    err, _ := bamboo.GenerateUrl("info", opts)
    if err != nil { return nil, err }
  }

  contents, err := Request(bamboo)
  if err != nil { return nil, err }

  var info Info
  res, err := parse(contents, &info)
  return res.(*Info), err
}

func (bamboo *Bamboo) Plan(opts map[string]string) (*Plan, error) {
  if bamboo.Url == "" {
    err, _ := bamboo.GenerateUrl("plan", opts)
    if err != nil { return nil, err }
  }

  contents, err := Request(bamboo)
  if err != nil { return nil, err }

  var plan Plan
  res, err := parse(contents, &plan)
  return res.(*Plan), err
}

func (bamboo *Bamboo) Project(opts map[string]string) (*Project, error) {
  if bamboo.Url == "" {
    err, _ := bamboo.GenerateUrl("project", opts)
    if err != nil { return nil, err }
  }

  contents, err := Request(bamboo)
  if err != nil { return nil, err }

  var project Project
  res, err := parse(contents, &project)
  return res.(*Project), err
}

func (bamboo *Bamboo) Queue(opts map[string]string) (*Queue, error) {
  if bamboo.Url == "" {
    err, _ := bamboo.GenerateUrl("queue", opts)
    if err != nil { return nil, err }
  }

  contents, err := Request(bamboo)
  if err != nil { return nil, err }

  var queue Queue
  res, err := parse(contents, &queue)
  return res.(*Queue), err
}

func (bamboo *Bamboo) Result(opts map[string]string) (*Result, error) {
  if bamboo.Url == "" {
    err, _ := bamboo.GenerateUrl("result", opts)
    if err != nil { return nil, err }
  }

  contents, err := Request(bamboo)
  if err != nil { return nil, err }

  var result Result
  res, err := parse(contents, &result)
  return res.(*Result), err
}

