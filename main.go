package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync"
)

type consumer struct {
	lock sync.Mutex
	server *http.Server
	log *log.Logger
}

func (c *consumer) start(addr string) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		c.lock.Lock()
		defer c.lock.Unlock()
		log.Printf("Received a %q request:\n", req.Method)
		switch req.Method {
		case http.MethodPost:
			log.Printf("Headers: %s", req.Header)
			data, err := ioutil.ReadAll(req.Body)
			bodyStr := string(data)
			log.Printf("Body: %q, Error: %s", bodyStr, err)
		}
		log.Printf("----------------------\n")
	})
	c.server = &http.Server{
		Handler: mux,
		Addr: addr,
	}
	if err := c.server.ListenAndServe(); err != nil {
		c.log.Fatal(err)
	}
}

func main() {
	c := &consumer{
		log: log.New(os.Stdout, "simple-sink", log.Lshortfile),
	}
	addr := ":8080"
	c.log.Printf("starting sink on %q...\n", addr)
	c.start(addr)
}