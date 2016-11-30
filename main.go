package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/oschwald/geoip2-golang"
)

var (
	schema         graphql.Schema
	queryLimit     = 65535
	queryLimitText = []byte{}
	geoDB          *geoip2.Reader
	geoip2file     = `/usr/share/GeoIP/GeoLite2-City.mmdb`
)

func main() {

	initVar()

	http.HandleFunc(`/api`, api)
	fmt.Println(`server started`)
	err := http.ListenAndServe(`:59999`, nil)
	if err != nil {
		panic("ListenAndServe fail: " + err.Error())
	}
}

func initVar() {

	var err error

	geoDB, err = geoip2.Open(geoip2file)
	if err != nil {
		log.Fatal(err)
	}

	schema, err = graphql.NewSchema(graphql.SchemaConfig{Query: geoipType})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	queryLimitText = []byte(fmt.Sprintf(`{"data":null,"errors":[{"message":"query too long (over limit %d)","locations":[]}]`, queryLimit))
}

func writeHttp(w http.ResponseWriter, s string) {
	w.Write([]byte(s))
}

func api(w http.ResponseWriter, r *http.Request) {

	buf := bytes.NewBuffer([]byte{})
	rlimit := queryLimit
	for {
		p := make([]byte, 1024)
		n, err := r.Body.Read(p)
		rlimit -= n
		if rlimit < 0 {
			w.Write(queryLimitText)
			return
		}
		if n > 0 {
			buf.Write(p[:n])
		}
		if err != nil {
			break
		}
	}

	query := buf.String()

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	rs := graphql.Do(params)

	rJSON, _ := json.Marshal(rs)

	w.Header().Set(`Content-Type`, `application/json; charset=utf-8`)

	w.Write([]byte(rJSON))
}

func getTime(p graphql.ResolveParams) (interface{}, error) {
	return fmt.Sprintf(`%v`, time.Now().Unix()), nil
}
