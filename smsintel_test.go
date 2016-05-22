package smsintel

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/kihamo/smsintel/procedure"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type ServerSuite struct {
	mux    *http.ServeMux
	server *httptest.Server
	client *SmsIntel
}

func (s *ServerSuite) SetUpTest(c *C) {
	s.mux = http.NewServeMux()
	s.server = httptest.NewUnstartedServer(s.mux)
	s.client = NewSmsIntel("12345", "54321")
}

func (s *ServerSuite) TearDownTest(c *C) {
	s.server.Close()
}

func (s *ServerSuite) ServerStart() {
	s.server.Start()
	s.client.SetOptions(map[string]string{
		"api_url": s.server.URL,
	})
}

func (s *ServerSuite) AddHandler(u string, file string) {
	s.mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.String() != u {
			http.NotFound(w, r)
			return
		}

		filename := filepath.Join(append([]string{"json"}, file)...)
		f, err := os.Open(filename)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		io.Copy(w, f)
	})
}

type ClientSuite struct {
	ServerSuite
}

var _ = Suite(&ClientSuite{})

func (s *ClientSuite) Test_Info_Success(c *C) {
	s.AddHandler("/info.php?login=12345&password=54321", "info_success.json")
	s.ServerStart()

	response, err := s.client.Info(&procedure.InfoInput{})

	c.Assert(err, IsNil)
	c.Assert(response.UserAccess.CanOrgListExport, Equals, true)
	c.Assert(response.UserAccess.CanDutyAndSkipped, Equals, false)
}
