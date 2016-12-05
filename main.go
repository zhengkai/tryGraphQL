package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/language/ast"
	"github.com/oschwald/geoip2-golang"
)

var (
	schema         graphql.Schema
	queryLimit     = 65535
	queryLimitText = []byte{}
	geoDB          *geoip2.Reader
	geoIspDB       *geoip2.Reader
	geoIp2file     = `/usr/share/GeoIP/GeoIP2-City.mmdb`
	geoLite2File   = `/usr/share/GeoIP/GeoLite2-City.mmdb`
	geoIp2IspFile  = `/usr/share/GeoIP/GeoIP2-ISP.mmdb`

	statusRequest      uint64 = 0
	statusRequestError uint64 = 0

	statusTimeStart = time.Now()
)

type queryJson struct {
	Query         string `json:"query"`
	OperationName string `json:"operationName"`
	Variables     string `json:"variables"`
}

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

	dbFile := geoIp2file
	if _, err = os.Stat(dbFile); err == nil {
		dbFile = geoLite2File
	}

	geoDB, err = geoip2.Open(dbFile)
	if err != nil {
		log.Fatal(err)
	}

	if _, err = os.Stat(geoIp2IspFile); err == nil {
		geoIspDB, err = geoip2.Open(geoIp2IspFile)
		if err != nil {
			log.Fatal(err)
		}
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

	var err error

	if r.Method == `GET` {
		buf.WriteString(r.FormValue(`query`))
	} else {
		rlimit := queryLimit
		for {
			p := make([]byte, 1024)
			n, err := r.Body.Read(p)
			rlimit -= n
			if rlimit < 0 {
				statusRequestError++
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
	}

	query := buf.String()

	var qJson queryJson
	err = json.Unmarshal([]byte(query), &qJson)
	if err == nil {
		query = qJson.Query
	}

	// fmt.Println(r.Method, `query =`, query)

	params := graphql.Params{
		Schema:        schema,
		RequestString: query,
	}
	rs := graphql.Do(params)
	if rs.Errors != nil {
		statusRequestError++
	}

	rJSON, _ := json.Marshal(rs)

	w.Header().Set(`Content-Type`, `application/json; charset=utf-8`)
	w.Header().Set(`Access-Control-Allow-Origin`, `*`)
	w.Write([]byte(rJSON))
}

func getTime(p graphql.ResolveParams) (interface{}, error) {
	return fmt.Sprintf(`%v`, time.Now().Unix()), nil
}

func getStatus(p graphql.ResolveParams) (interface{}, error) {

	timeDiff := time.Since(statusTimeStart) / time.Second

	return Status{
		Uptime:       int64(timeDiff),
		UptimeText:   (timeDiff * time.Second).String(),
		Request:      statusRequest,
		RequestError: statusRequestError,
	}, nil
}

func selectedFields(p graphql.ResolveParams) []string {

	set := p.Info.FieldASTs[0].SelectionSet

	fields := make([]string, len(set.Selections))
	for i, f := range set.Selections {
		switch f.(type) {
		case *ast.Field:
			fields[i] = f.(*ast.Field).Name.Value
		case *ast.FragmentSpread:
			fields[i] = f.(*ast.FragmentSpread).Name.Value
		}
	}
	return fields
}

func isContain(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
