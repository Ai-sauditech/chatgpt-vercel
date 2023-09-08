package main

import (
	"os"

	"github.com/InfluxCommunity/influxdb3-go/influx"
)

func main() {
	// Create client
	url := "https://us-central1-1.gcp.cloud2.influxdata.com"
	token := os.Getenv("INFLUXDB_TOKEN")

	// Create a new client using an InfluxDB server base URL and an authentication token
	client, err := influx.New(influx.Configs{
		HostURL:   url,
		AuthToken: token,
	})

	if err != nil {
		panic(err)
	}
	// Close client at the end and escalate error if present
	defer func(client *influx.Client) {
		err := client.Close()
		if err != nil {
			panic(err)
		}
	}(client)

	database := "influxtelegraf"
}
