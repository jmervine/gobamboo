package bamboo

import (
  "fmt"
  "testing"
  "github.com/stretchr/testify/assert"
  "net/http"
  "net/http/httptest"
  "io/ioutil"
)

// Setup
func PrefabBamboo() (b *Bamboo) {
  b          = new(Bamboo)
  b.Domain   = "http://bamboo.example.com"
  b.Username = "USERNAME"
  b.Password = "PASSWORD"
  return
}

// Mocks
func MockRequest(api string) (*httptest.Server) {
  body, err := ioutil.ReadFile("./mocks/"+api+".json")
  if err != nil {
    panic(err)
  }
  return httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(res, string(body))
  }))
}

// Tests
func TestBambooVariables(t *testing.T) {
  b := PrefabBamboo()

  assert.IsType(t, b, new(Bamboo), "b should be *Bamboo")
  assert.Equal(t, b.Domain, "http://bamboo.example.com", "Domain should be set")
  assert.Equal(t, b.Username, "USERNAME", "Username should be set")
  assert.Equal(t, b.Password, "PASSWORD", "Password should be set")
}

func TestGenerateUrl(t *testing.T) {
  b := PrefabBamboo()

  b.GenerateUrl("TEST_API_KEY", make(map[string]string))
  assert.Contains(t, b.Url, "bamboo.example.com", "Url should contain Domain")
  assert.Contains(t, b.Url, "USERNAME", "Url should contain Username")
  assert.Contains(t, b.Url, "PASSWORD", "Url should contain Password")
  assert.Contains(t, b.Url, "TEST_API_KEY", "Url should contain API Key")

  opts := make(map[string]string)
  opts["TESTKEY"] = "TESTVAL"
  b.GenerateUrl("TEST_API_KEY", opts)

  assert.Contains(t, b.Url, "TESTKEY=TESTVAL", "Url should contain passed argument")
}

func TestInfo(t *testing.T) {
  ts := MockRequest("info")
  defer ts.Close()

  b := PrefabBamboo()
  b.Url = ts.URL
  info, err := b.Info(make(map[string]string))

  assert.NoError(t, err, "Info shouldn't error.")
  assert.Equal(t, info.State, "RUNNING", "info.State should be RUNNING")
}

func TestPlan(t *testing.T) {
  ts := MockRequest("plan")
  defer ts.Close()

  b := PrefabBamboo()
  b.Url = ts.URL
  plan, err := b.Plan(make(map[string]string))

  assert.NoError(t, err, "Plan shouldn't error.")
  assert.Equal(t, plan.Link.Href, "http://bamboo.example.com/rest/api/latest/plan", "should be equal")
  assert.Equal(t, plan.Plans.Plan[0].Link.Href, "http://bamboo.example.com/rest/api/latest/plan/EX-EXPL", "should be equal")
  assert.Equal(t, plan.Plans.Plan[0].PlanKey.Key, "EX-EXPL", "should be equal")
  assert.Equal(t, plan.Plans.Plan[0].Name, "Example Plan", "should be equal")
}

func TestProject(t *testing.T) {
  ts := MockRequest("project")
  defer ts.Close()

  b := PrefabBamboo()
  b.Url = ts.URL
  project, err := b.Project(make(map[string]string))

  assert.NoError(t, err, "Project shouldn't error.")
  assert.Equal(t, project.Link.Href, "http://bamboo.example.com/rest/api/latest/project", "should be equal")
  assert.Equal(t, project.Projects.Project[0].Link.Href, "http://bamboo.example.com/rest/api/latest/project/EXPROJ", "should be equal")
  assert.Equal(t, project.Projects.Project[0].Name, "Example Project", "should be equal")
}

func TestQueue(t *testing.T) {
  ts := MockRequest("queue")
  defer ts.Close()

  b := PrefabBamboo()
  b.Url = ts.URL
  queue, err := b.Queue(make(map[string]string))

  assert.NoError(t, err, "Queue shouldn't error.")
  assert.Equal(t, queue.Link.Href, "http://bamboo.example.com/rest/api/latest/queue", "should be equal")
  assert.Equal(t, queue.QueuedBuilds.Size, 0, "should be equal")
  assert.IsType(t, queue.QueuedBuilds.Queue, *(new([]QueueType)), "should be empty")
  assert.Empty(t, queue.QueuedBuilds.Queue, "should be same type")
}

func TestResult(t *testing.T) {
  ts := MockRequest("result")
  defer ts.Close()

  b := PrefabBamboo()
  b.Url = ts.URL
  result, err := b.Result(make(map[string]string))

  assert.NoError(t, err, "Result shouldn't error.")
  assert.Equal(t, result.Results.Expand, "result", "should be equal")
  assert.Equal(t, result.Results.Result[0].Link.Href, "http://bamboo.example.com/rest/api/latest/result/EX-EXPL-5", "should be equal")
  assert.Equal(t, result.Results.Result[0].Plan.PlanKey.Key, "EX-EXPL", "should be equal")
  assert.Equal(t, result.Results.Result[0].PlanResultKey.ResultNumber, 5, "should be equal")
  assert.Equal(t, result.Results.Result[0].PlanResultKey.EntityKey.Key, "EX-EXPL", "should be equal")
}
