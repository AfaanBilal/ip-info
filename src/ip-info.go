package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	var ip string
	fmt.Println("IP Info")
	fmt.Print("Please enter the IP: ")
	fmt.Scanln(&ip)

	res, err := http.Get("https://api.iplegit.com/full?ip=" + ip)
	if err != nil {
		log.Fatal(err)
	}

	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	type IpInfo struct {
		Ip          string
		Type        string `json:"type"`
		CountryCode string
		CountryName string
		Latitude    float32
		Longitude   float32
		City        string
		Region      string
		ISP         string
		ASN         int
	}

	var ipInfo IpInfo
	json.Unmarshal(data, &ipInfo)

	fmt.Printf("IP:           %s\n", ipInfo.Ip)
	fmt.Printf("Type:         %s\n", ipInfo.Type)
	fmt.Printf("Country Code: %s\n", ipInfo.CountryCode)
	fmt.Printf("Country Name: %s\n", ipInfo.CountryName)
	fmt.Printf("Latitude:     %.6f\n", ipInfo.Latitude)
	fmt.Printf("Longitude:    %.6f\n", ipInfo.Longitude)
	fmt.Printf("City:         %s\n", ipInfo.City)
	fmt.Printf("Region:       %s\n", ipInfo.Region)
	fmt.Printf("ISP:          %s\n", ipInfo.ISP)
	fmt.Printf("ASN:          %d\n", ipInfo.ASN)
	fmt.Print("\nData provided by IPLegit.com")
}
