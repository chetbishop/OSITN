package main

import (
	//"fmt"
	"github.com/chetbishop/OSINT/database"
	//"github.com/chetbishop/OSINT/url"
	"github.com/chetbishop/OSINT/domains"
)

func main() {
	//url.UrlGet()
	database.CreateDatabase("default")
	domains.DomainsAddDB("default", "bcbs.com")
}
