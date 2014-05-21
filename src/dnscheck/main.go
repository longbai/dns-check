package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/slene/iploc"
)

type Result struct {
	NameServer        string
	NameServerCountry string
	DomainIp          string
	DomainCountry     string
}

func main() {
	dns := flag.String("ns", "", "name servers")
	_domain := flag.String("domain", "", "domain")
	output := flag.String("o", "", "output json file")

	flag.Parse()
	domain := *_domain
	if *dns == "" || domain == "" {
		fmt.Println("invalid args")
		flag.PrintDefaults()
		os.Exit(1)
	}

	ns, err := loadNameServers(*dns)
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		os.Exit(1)
	}

	if !strings.HasSuffix(domain, ".") {
		domain += "."
	}

	// replace iplocFilePath to your iploc.dat path
	iplocFilePath, _ := filepath.Abs("bin/iploc.dat")

	// simple set a true param can preload all ipinfo
	// need allocate more memory > 30M
	// and speed can grow up about 40 percent than not preload
	iploc.IpLocInit(iplocFilePath, true)
	var rets []Result
	for _, v := range ns {
		ip, err := Query(&v, domain)
		if err != nil {
			continue
		}
		var r Result
		r.NameServer = v.Addr()
		r.NameServerCountry = iploc.COUNTRIES_ZH[v.CountryID]
		r.DomainIp = ip.String()
		info, err := iploc.GetIpInfo(ip.String())
		if err == nil {
			r.DomainCountry = info.Country
		}
		rets = append(rets, r)
		fmt.Printf("%#v\n", r)
	}
	if *output != "" {
		data, err := json.MarshalIndent(rets, "", "")
		if err != nil {
			fmt.Println(err)
			return
		}
		err = ioutil.WriteFile(*output, data, 0666)
		if err != nil {
			fmt.Println(err)
		}
	}
}
