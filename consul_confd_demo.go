package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/benschw/consul-clb-go/clb"
	"io/ioutil"
	"launchpad.net/goyaml"
	"log"
	"net/http"
	"os"
)

func status(w http.ResponseWriter, req *http.Request) {
	health := "OK"

	if _, err := os.Stat("/tmp/fail-healthcheck"); err == nil {
		//file exists
		health = "FAIL"
	}
	fmt.Fprintf(w, "%+v", health)
}

// [{"Node":"agent-one","Address":"172.20.20.11","ServiceID":"web","ServiceName":"web","ServiceTags":["rails"],"ServicePort":80}]

func getAddress(svcName string) (string, error) {
	c := clb.NewClb("127.0.0.1", "8600", clb.Random)

	srvRecord := svcName + ".service.consul"
	address, err := c.GetAddress(srvRecord)
	if err != nil {
		return "", err
	}

	return address.String(), nil
}

type Config struct {
	Foo  string
	Foo2 string
}

func getFoo() (string, error) {
	fileData, err := ioutil.ReadFile("/opt/my-config.yaml")
	if err != nil {
		return "", err
	}

	data := Config{}

	err = goyaml.Unmarshal(fileData, &data)
	if err != nil {
		return "", err
	}
	log.Printf("%+v", data)
	return data.Foo, nil
}

func foo(w http.ResponseWriter, req *http.Request) {
	name, _ := os.Hostname()
	foo, err := getFoo()
	if err != nil {
		log.Print(err)
	}

	fmt.Fprintf(w, "{\"Name\": \"%s\", \"Foo\": \"%s\"}", name, foo)

}

type FooName struct {
	Name string
	Foo  string
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

	fmt.Fprintf(w, "Discovered Address: '%s'\nService Response: Host Name: '%s'\nService Response: Foo: '%s'", addStr, data.Name, data.Foo)

}

var addr = flag.String("addr", ":8080", "http service address")                  // Q=17, R=18
var consulAddr = *flag.String("consul-addr", "localhost:8500", "consul address") // Q=17, R=18

func main() {
	flag.Parse()
	http.Handle("/status", http.HandlerFunc(status))
	http.Handle("/foo", http.HandlerFunc(foo))
	http.Handle("/demo", http.HandlerFunc(demo))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
