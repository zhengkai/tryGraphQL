package main

import "github.com/graphql-go/graphql"

type Geo struct {
	Ip        string      `json:"ip"`
	City      string      `json:"city"`
	Country   GeoCountry  `json:"country"`
	Location  GeoLocation `json:"location"`
	Continent string      `json:"continent"`
	Postal    string      `json:"postal_code"`
}

type GeoCountry struct {
	IsoCode string `json:"iso_code"`
	Name    string `json:"name"`
}

type GeoLocation struct {
	AccuracyRadius uint16  `json:"accuracy_radius"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	MetroCode      uint    `json:"metro_code"`
	TimeZone       string  `json:"time_zone"`
}

var geoCountry = graphql.NewObject(graphql.ObjectConfig{
	Name: "Country",
	Fields: graphql.Fields{
		"iso_code": &graphql.Field{
			Type: graphql.String,
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var geoLang = graphql.NewEnum(graphql.EnumConfig{
	Name: "language",
	Values: graphql.EnumValueConfigMap{
		`ru`: &graphql.EnumValueConfig{},
		`zh`: &graphql.EnumValueConfig{},
		`de`: &graphql.EnumValueConfig{},
		`en`: &graphql.EnumValueConfig{},
		`es`: &graphql.EnumValueConfig{},
		`fr`: &graphql.EnumValueConfig{},
		`ja`: &graphql.EnumValueConfig{},
		`pt`: &graphql.EnumValueConfig{},
	},
})

var geoLocation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Fields: graphql.Fields{
		"time_zone": &graphql.Field{
			Type: graphql.String,
		},
		"latitude": &graphql.Field{
			Type: graphql.Float,
		},
		"longitude": &graphql.Field{
			Type: graphql.Float,
		},
		"metro_code": &graphql.Field{
			Type: graphql.Int,
		},
		"accuracy_radius": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var geoType = graphql.NewObject(graphql.ObjectConfig{
	Name: "GeoInfo",
	Fields: graphql.Fields{
		"ip": &graphql.Field{
			Type: graphql.String,
		},
		"city": &graphql.Field{
			Type: graphql.String,
		},
		"postal_code": &graphql.Field{
			Type: graphql.String,
		},
		"continent": &graphql.Field{
			Type: graphql.String,
		},
		"country": &graphql.Field{
			Type: geoCountry,
		},
		"location": &graphql.Field{
			Type: geoLocation,
		},
	},
})

var geoipType = graphql.NewObject(graphql.ObjectConfig{
	Name:        "GeoIP",
	Description: "GeoIP Lookup",
	Fields: graphql.Fields{
		"geo": &graphql.Field{
			Type:        geoType,
			Description: `查询 IP 归属地`,
			Args: graphql.FieldConfigArgument{
				"ip": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"lang": &graphql.ArgumentConfig{
					Type:         geoLang,
					DefaultValue: `en`,
				},
			},
			Resolve: geolookup,
		},
		"time": &graphql.Field{
			Type:        graphql.String,
			Description: `返回服务器时间戳`,
			Resolve:     getTime,
		},
	},
})
