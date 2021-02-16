package main

import (
	"flag"
	"fmt"
	"os"

	influxdb "github.com/influxdata/influxdb/client/v2"
)

func main() {
	influxUrl := flag.String("influxUrl", "http://localhost:8086", "influx endpoint")
	influxUser := flag.String("user", "root", "influx user")
	influxPassword := flag.String("password", "root", "influx user's password")
	databaseName := flag.String("database", "", "Database name to create")
	retention := flag.String("retention", "30d", "Retention duration")
	replication := flag.String("replication", "1", "Replication count")
	flag.Parse()

	if len(*databaseName) == 0 {
		fmt.Println("Database name is required")
		flag.Usage()
		os.Exit(1)
	}

	client, err := influxdb.NewHTTPClient(influxdb.HTTPConfig{
		Addr: *influxUrl,
		Username: *influxUser,
		Password: *influxPassword,
		InsecureSkipVerify: true,
	})

	if err != nil {
		panic(fmt.Sprintf("Error connecting to Influx DB: %s", err))
	}

	_, err = client.Query(influxdb.Query{Command: fmt.Sprintf("create database %s", *databaseName) })
	if err != nil {
		panic(fmt.Sprintf("Error creating DB '%s': %s", *databaseName, err))
	}

	retentionCommand := fmt.Sprintf("create RETENTION POLICY \"default\" ON %s DURATION %s REPLICATION %s DEFAULT", *databaseName, *retention, *replication)
	_, err = client.Query(influxdb.Query{Command: retentionCommand})
	if err != nil {
		panic(fmt.Sprintf("Error creating default retention policy: %s", err))
	}
}
