package pgress

import (
	"context"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
		"github.com/influxdata/influxdb-client-go/v2/api"
	"github.com/vishal2911/algoTrading/angelOneGo/model"
	"github.com/vishal2911/algoTrading/angelOneGo/util"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgressStore struct {
	DB       *gorm.DB
	InfluxDB influxdb2.Client
	Bucket   string
	Org      string
}

func (store *PgressStore) NewStore() {
	// Initialize PostgreSQL connection
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore, "Setting DataBase", nil)

	db, err := gorm.Open(postgres.Open(model.DSN), &gorm.Config{})
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.NewStore,
			"error while connecting to database", err)
		panic(err)
	}

	err = db.AutoMigrate(
		model.User{},
	)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.NewStore,
			"error while automigrating database", err)
		panic(err)
	}
	store.DB = db

	// Initialize InfluxDB connection
	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore, "Setting InfluxDB", nil)

	influxDB := influxdb2.NewClient(model.InfluxDBURL, model.InfluxDBToken)
	health, err := influxDB.Health(context.Background())
	if err != nil || health.Status != "pass" {
		util.Log(model.LogLevelError, model.PgressPackageLevel, model.NewStore,
			"error while connecting to InfluxDB", err)
		panic(err)
	}
	store.InfluxDB = influxDB
	store.Bucket = model.InfluxDBBucket
	store.Org = model.InfluxDBOrg

	util.Log(model.LogLevelInfo, model.PgressPackageLevel, model.NewStore, "Setting DataBase completed .... ", nil)
}

// Example function to write data to InfluxDB
func (store *PgressStore) WriteToInflux(measurement string, tags map[string]string, fields map[string]interface{}, timestamp time.Time) {
	writeAPI := store.InfluxDB.WriteAPIBlocking(store.Org, store.Bucket)
	point := influxdb2.NewPoint(measurement, tags, fields, timestamp)
	err := writeAPI.WritePoint(context.Background(), point)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, "WriteToInflux",
			"error while writing to InfluxDB", err)
	}
}

// Example function to query data from InfluxDB
func (store *PgressStore) QueryInflux(query string) (*api.QueryTableResult, error) {
	queryAPI := store.InfluxDB.QueryAPI(store.Org)
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		util.Log(model.LogLevelError, model.PgressPackageLevel, "QueryInflux",
			"error while querying InfluxDB", err)
		return nil, err
	}
	return result, nil
}
