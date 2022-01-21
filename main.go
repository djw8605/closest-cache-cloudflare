package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)



type Caches struct {
	Cache string
	Distance float64
}

type Result struct {
	Caches []Caches `json:"caches"`
}

func (tp *Caches) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Printf("Error whilde decoding %v\n", err)
		return err
	}
	tp.Cache = v[0].(string)
	tp.Distance = v[1].(float64)
	return nil
}

func main() {
	url := "https://cache-location.osgstorage.org/_caches"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: tr,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Got error %s", err.Error())
		return
	}
	req.Header.Set("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36")
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("Got error %s", err.Error())
		return
	}
	defer response.Body.Close()

	read, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	var result Result
	err = json.Unmarshal(read, &result)
	if err != nil {
		fmt.Println("Failed to parse:")
		fmt.Println(string(read))
		panic(err)
	}


	fmt.Println(result.Caches[0].Cache)


}

