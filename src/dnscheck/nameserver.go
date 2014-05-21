package main

import (
	"encoding/json"
	"io/ioutil"
	// "os"
	"strings"
	"time"
)

type DnsServer interface {
	Address() string
}

type NameServer struct {
	CheckedAt      time.Time `json:"checked_at"`
	City           string    `json:"city"`
	CountryID      string    `json:"country_id"`
	CreatedAt      time.Time `json:"created_at"`
	Error          string    `json:"error"`
	Id             int       `json:"id"`
	Ip             string    `json:"ip"`
	Name           string    `json:"name"`
	State          string    `json:"state"`
	StateChangedAt time.Time `json:"state_changed_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	Version        string    `json:"version"`
}

func (n *NameServer) Address() string {
	addr := n.Ip
	if strings.Contains(addr, ":") {
		addr = n.Name
	}
	return addr + ":53"
}

func (n *NameServer) Addr() string {
	addr := n.Ip
	if strings.Contains(addr, ":") {
		addr = n.Name
	}
	return addr
}

func loadNameServers(path string) (ns []NameServer, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &ns)
	return
}
