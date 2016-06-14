package main

import (
	"net/http"
	"net/http/httputil"
	"time"
	"net/url"
	"sync/atomic"
	"io/ioutil"
	"log"
	"encoding/json"
	"strconv"
	"bytes"
	"path"
)

var redirectToApache atomic.Value

func configUpdater(){
	lastValue := []byte{}
	for {
		b, err := ioutil.ReadFile(path.Join("./config.json"))
		if err != nil {
			log.Println(err)
		} else {
			value := make(map[string]int)

			if !bytes.Equal(lastValue, b) {
				lastValue = b
				log.Println("Loaded config")
			}

			err := json.Unmarshal(b, &value)
			if err != nil {
				log.Println(err)
			}else{
				redirectToApache.Store(value)
			}
		}

		time.Sleep(30 * time.Second)
	}
}

func director(r *http.Request) {
	v := redirectToApache.Load().(map[string]int)
	if i, ok := v[r.Host]; ok {
		r.Host = r.Host + ":" + strconv.Itoa(i)
	}else{
		log.Println("unknown host '" + r.Host + "'")
		r.Host = "localhost:8080"
	}

	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main(){

	go configUpdater()

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
		Host:   "localhost:80",
	})
	proxy.Director = director
	http.ListenAndServe(":80", proxy)
}