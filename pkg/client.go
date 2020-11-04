package fhem_go

type Client struct{
	server string
	port int `default:"1773"`
	useSsl string `default:"False"`
	protocol string `default:"https"`
	username string
	password string
	csrf  bool `default:"True"`
	catfile string
}