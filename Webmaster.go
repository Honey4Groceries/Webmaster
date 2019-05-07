package main

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"net/http"
	// "net/url"
	"io/ioutil"
)

func main() {
	c, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic("redis is dead")
		return
	}
	defer c.Close()

	worker("https://products.wholefoodsmarket.com/api/data/stores")
}

// A worker makes a get request to a passed in URL, and gets a json back in response to dump into channel
// These get requests are to be made to a worker lambda function and so the custom response that you will
// get will be of a defined type. that way we really should make a custom struct with the predefined
// json schema. Once we have a custom struct, then we may parse the response json to data, which is of type
// of the defined json. Then we just queue data back into the channel

// What do we need to keep track of in a webscraper?
// status code: keep track of 404s
// we need an enum keeping track of what grocery store it has been scraped from.
// data: json containing data scraped
// urls: urls that was scraped

func worker(url string) {
	var data []map[string]interface{}
	// so let's first make a http request!
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		panic(err)
	}

	fmt.Println(data)
}
