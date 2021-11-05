package fhemgo

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"time"
)

const fhemCsrfKey = "X-Fhem-Csrftoken"

type FhemClient struct {
	client    *http.Client
	server    string
	port      int    `default:"1773"`
	useSsl    string `default:"False"`
	protocol  string `default:"https"`
	username  string
	password  string
	csrf      bool `default:"True"`
	csrfToken string
	catfile   string
	baseUrl   string
}

// NewClient creates a new FHEM client with basic functionality
func NewClient(server string, port int, https bool, local bool) *FhemClient {
	client := &http.Client{Timeout: time.Minute}
	return &FhemClient{
		server:    server,
		port:      port,
		client:    client,
		baseUrl:   getBaseUrl(server, port, https, local),
		csrfToken: fetchToken(client, getBaseUrl(server, port, https, local)),
	}
}

// fetchToken fetches the current CSRF token which is needed to send commands
func fetchToken(client *http.Client, baseurl string) string {
	reqUrl := fmt.Sprintf("%s?room=foo", baseurl)
	res, err := client.Get(reqUrl)
	if err != nil || res.StatusCode != http.StatusOK {
		log.WithError(err).Errorf("Could not fetch csrf token from baseurl %s", baseurl)
	}
	return res.Header.Get(fhemCsrfKey)
}

func (f *FhemClient) Get(devType string) {
	rurl := buildUrl(f.baseUrl, f.csrfToken, "list", devType)
	res, err := f.client.Get(rurl)
	if err != nil || res.StatusCode != http.StatusOK {
		log.WithError(err)
	}

}

func buildUrl(base string, token string, cmd string, devType string) string {
	baseUrl, err := url.Parse(base)
	if err != nil {
		log.WithError(err).Errorf("Malformed URL: %s", base)
	}

	params := url.Values{}
	params.Add("fwcsrf", token)
	params.Add("cmd", cmd)
	params.Add("TYPE", devType)

	baseUrl.RawQuery = params.Encode() // Escape Query Parameters

	return baseUrl.String()
}

func getBaseUrl(server string, port int, https bool, local bool) string {
	if local {
		return fmt.Sprintf("localhost:%d/fhem", port)
	}
	if https {
		return fmt.Sprintf("https://%s:%d/fhem", server, port)
	}
	return fmt.Sprintf("http://%s:%d/fhem", server, port)
}

func readBody(response http.Response) (*[]byte, error) {
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.WithError(err).Errorf("Could not read response body %s", response.Body)
		return nil, err
	}
	return &body, nil
}
