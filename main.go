package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

var (
	fields         graphql.Fields
	rootQuery      graphql.ObjectConfig
	schemaConfig   graphql.SchemaConfig
	schema         graphql.Schema
	queryLimit     = 65535
	queryLimitText = []byte{}
)

func main() {

	fields = graphql.Fields{
		"hello": &graphql.Field{
			Type:    graphql.String,
			Resolve: helloworld,
		},
		"time": &graphql.Field{
			Type:    graphql.String,
			Resolve: getTime,
		},
	}
	rootQuery = graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig = graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}

	var err error
	schema, err = graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}

	queryLimitText = []byte(fmt.Sprintf(`{"data":null,"errors":[{"message":"query too long (over limit %d)","locations":[{"line":0,"column":0}]}]`, queryLimit))

	http.HandleFunc(`/api`, api)
	fmt.Println(`server started`)
	err = http.ListenAndServe(`:59999`, nil)
	if err != nil {
		panic("ListenAndServe fail: " + err.Error())
	}
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

	w.Write([]byte(rJSON))
}

func helloworld(p graphql.ResolveParams) (interface{}, error) {
	return "world", nil
}

func getTime(p graphql.ResolveParams) (interface{}, error) {
	return fmt.Sprintf(`%v`, time.Now().Unix()), nil
}
