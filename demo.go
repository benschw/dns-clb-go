package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/benschw/dns-clb-go/clb"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getAddress(svcName string) (string, error) {
	c := clb.NewClb("127.0.0.1", "8600", clb.Random)

	srvRecord := svcName + ".service.consul"
	address, err := c.GetAddress(srvRecord)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

func foo(w http.ResponseWriter, req *http.Request) {
	name, _ := os.Hostname()

	fmt.Fprintf(w, "{\"Name\": \"%s\"}", name)

}

type FooName struct {
	Name string
}

func demo(w http.ResponseWriter, req *http.Request) {
	addStr, err := getAddress("my-svc")
	if err != nil {
		log.Fatal(err)
	}

	url := "http://" + addStr + "/foo"

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data FooName
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Discovered Address: '%s'\nService Response: Host Name: '%s'", addStr, data.Name)

}

var addr = flag.String("addr", ":8080", "http service address")                  // Q=17, R=18
var consulAddr = *flag.String("consul-addr", "localhost:8500", "consul address") // Q=17, R=18

func main() {
	flag.Parse()
	http.Handle("/foo", http.HandlerFunc(foo))
	http.Handle("/demo", http.HandlerFunc(demo))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
