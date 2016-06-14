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
	"strings"
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

	var h string
	if strings.HasPrefix(r.Host, "www.") {
		h = r.Host[4:]
	}else{
		h = r.Host
	}

	if i, ok := v[h]; ok {
		r.Host = h + ":" + strconv.Itoa(i)
	}else{
		log.Println("unknown host '" + h + "'")
		r.Host = "localhost:8080"
	}

	r.URL.Host = r.Host
	r.URL.Scheme = "http"
}

func main(){

	go configUpdater()

	proxy := httputil.NewSingleHostReverseProxy(&url.URL{
		Scheme: "http",
	})
	proxy.Director = director
	http.ListenAndServe(":80", proxy)
}