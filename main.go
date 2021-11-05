package main

import (
	"github.com/oppermax/fhem-go"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"time"
)

type FhemClient struct {
	url string
	port int
	socket string
	httpClient *http.Client
}

func main(){
	fhem := fhemgo.NewClient("192.168.178.22", 8083)

	//res, err := fhem.httpClient.Get("http://192.168.178.22:8083/fhem?cmd=list%20TYPE=dummy&fwcsrf=csrf_55444420021900")
	//if err != nil {
	//	log.WithError(err)
	//}
	//b, _ := io.ReadAll(res.Body)
	//log.Info(string(b))
}