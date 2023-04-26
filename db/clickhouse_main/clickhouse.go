package main

import (
	"database/sql"
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	click "gorm.io/driver/clickhouse"
	"gorm.io/gorm"
	"simple-video-server/config"
	"time"
	//"github.com/ClickHouse/clickhouse-go/v2"
	//"github.com/ClickHouse/clickhouse-go/v2"
)

type Request struct {
	TraceID string `json:"trace_id"`
	UID     uint   `json:"uid"`
}

func (r *Request) TableName() string {
	return "request"
}

func main() {
	//connect()
	//conn, err := connect2()
	//if err != nil {
	//	panic(err)
	//}
	//
	//ctx := context.Background()
	//rows, err := conn.Query(ctx, "SELECT trace_id, api_url as uuid_str FROM simple_video_server.request LIMIT 5")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for rows.Next() {
	//	var (
	//		trace_id, api_url string
	//	)
	//	if err := rows.Scan(&api_url, &trace_id); err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	log.Printf("trace_id: %s, api_url: %s\n", trace_id, api_url)
	//}
	sqlDb := connect3()
	db, err := gorm.Open(click.New(click.Config{
		Conn: sqlDb,
	}))
	if err != nil {
		panic(err)
	}
	var log Request
	db.Find(&log)

	fmt.Println("log ", log)
}

//func connect3() (*gorm.DB, error) {
//	dsn := "clickhouse://:@localhost:9000/simple_video_server?dial_timeout=10s&read_timeout=20s"
//	//gorm.Open(clickhouseopen)
//	db, err := gorm.Open(clickhouse.New(clickhouse.Config{
//		DSN:                          dsn,
//		Conn:                         conn,     // initialize with existing database conn
//		DisableDatetimePrecision:     true,     // disable datetime64 precision, not supported before clickhouse 20.4
//		DontSupportRenameColumn:      true,     // rename column not supported before clickhouse 20.4
//		DontSupportEmptyDefaultValue: false,    // do not consider empty strings as valid default values
//		SkipInitializeWithVersion:    false,    // smart configure based on used version
//		DefaultGranularity:           3,        // 1 granule = 8192 rows
//		DefaultCompression:           "LZ4",    // default compression algorithm. LZ4 is lossless
//		DefaultIndexType:             "minmax", // index stores extremes of the expression
//		DefaultTableEngineOpts:       "ENGINE=MergeTree() ORDER BY tuple()",
//	}), &gorm.Config{})
//
//	return db, err
//}

//func connect() {
//	// Open database connection
//	db, err := sql.Open("clickhouse", "tcp://192.168.10.100:9000?username=&compress=1")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer db.Close()
//
//	// Execute query
//	rows, err := db.Query("SELECT trace_id, api_url FROM simple_video_server.request")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer rows.Close()
//
//	// Iterate over rows
//	for rows.Next() {
//		var trace_id string
//		var api_url string
//		// expected 13 destination arguments in Scan, not 2
//		if err := rows.Scan(&trace_id, &api_url); err != nil {
//			fmt.Println(err)
//			return
//		}
//		fmt.Printf("%s %s\n", trace_id, api_url)
//	}
//}

func connect3() *sql.DB {
	db := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{"192.168.10.182:9000"},
		Auth: clickhouse.Auth{
			Database: "simple_video_server",
		},
		Settings: clickhouse.Settings{
			"max_execution_time": 60,
		},
		DialTimeout: 5 * time.Second,
		//Compression: &clickhouse.Compression{
		//	clickhouse.CompressionLZ4,
		//},
		Debug: config.Env.Debug,
	})

	return db
}

//func connect2() (driver.Conn, error) {
//	ctx := context.Background()
//	//conn, err := clickhouse.Open(&clickhouse.Options{
//		Addr: []string{"192.168.10.100:9000"},
//		Auth: clickhouse.Auth{
//			Database: "simple_video_server",
//			//Username: "default",
//			//Password: "default",
//		},
//		ClientInfo: clickhouse.ClientInfo{
//			Products: []struct {
//				Name    string
//				Version string
//			}{
//				{Name: "an-example-go-client", Version: "0.1"},
//			},
//		},
//		Debugf: func(format string, v ...interface{}) {
//			fmt.Printf(format, v)
//		},
//		// 需要注释TLS不然会报: tls: first record does not look like a TLS handshake
//		//TLS: &tls.Config{
//		//	InsecureSkipVerify: true,
//		//},
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	if err := conn.Ping(ctx); err != nil {
//		if exception, ok := err.(*clickhouse.Exception); ok {
//			fmt.Printf("Exception [%d] %s \n%s\n", exception.Code, exception.Message, exception.StackTrace)
//		}
//		return nil, err
//	}
//
//	return conn, nil
//
//}
