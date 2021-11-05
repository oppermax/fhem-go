package main

import (
	"github.com/oppermax/fhem-go/pkg"
)

func main(){
	fhem := fhemgo.NewClient("", 8083, false, true)

	fhem.Get("dummy")

	//res, err := fhem.httpClient.Get("http://192.168.178.22:8083/fhem?cmd=list%20TYPE=dummy&fwcsrf=csrf_55444420021900")
	//if err != nil {
	//	log.WithError(err)
	//}
	//b, _ := io.ReadAll(res.Body)
	//log.Info(string(b))
}