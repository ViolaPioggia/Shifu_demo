package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	targetUrl := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	for {
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		average := calculateAverage(body)
		log.Println("Average:", average)
		time.Sleep(2 * time.Second)
	}
}

func calculateAverage(data []byte) float64 {
	sum := 0
	count := 0
	for _, value := range data {
		sum += int(value)
		count++
	}
	if count > 0 {
		return float64(sum) / float64(count)
	}
	return 0
}
