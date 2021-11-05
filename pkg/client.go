package fhemgo

import (
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const csrfPrefix = "&fwcsrf="

type FhemClient struct {
	client   *http.Client
	server   string
	port     int    `default:"1773"`
	useSsl   string `default:"False"`
	protocol string `default:"https"`
	username string
	password string
	csrf     bool `default:"True"`
	catfile  string
}

func NewClient(url string, port int) *FhemClient {
	return &FhemClient{
		server: url,
		port:   port,
		client: &http.Client{Timeout: time.Minute},
	}
}

// FetchToken fetches the current CSRF token which is needed to send commands
// http://192.168.178.22:8083/fhem?cmd=list%20TYPE=dummy&fwcsrf=csrf_55444420021900
func (f *FhemClient) FetchToken() {
	res, err := f.client.Get("http://192.168.178.22:8083/fhem?cmd=tokenconnect")
	if err != nil {
		log.WithError(err)
	}

	res.Header

}
