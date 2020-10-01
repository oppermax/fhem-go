package fhem


type Client struct {
	server string
	port int
	useSsl bool
	protocol string
	username string
	password string
	csrf bool
	cafile string
	loglevel int
}
