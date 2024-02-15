package config

import "os"

var Ip_server = os.Getenv("SERVER_ADDRESS")
var Port_server = os.Getenv("SERVER_PORT")
var Api_key = os.Getenv("API_KEY")
var Api_url = os.Getenv("API_URL")
