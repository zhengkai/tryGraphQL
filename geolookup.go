package main

import (
	"errors"
	"fmt"
	"net"
	"strings"

	"github.com/graphql-go/graphql"
)

var langAlias = map[string]string{
	`zh`: `zh-CN`,
	`pt`: `pt-BR`,
}

func geolookup(p graphql.ResolveParams) (r interface{}, err error) {

	ip, _ := p.Args["ip"].(string)
	ip = strings.TrimSpace(ip)

	sLang, _ := p.Args["lang"].(string)

	statusRequest++

	if tmpLang, ok := langAlias[sLang]; ok {
		sLang = tmpLang
	}

	pip := net.ParseIP(ip)
	if !pip.IsGlobalUnicast() {
		err = errors.New(fmt.Sprintf(`ip "%s" is not a global unicast address`, ip))
		return
	}

	record, err := geoDB.City(pip)
	if err != nil {
		return
	}
	if record.City.GeoNameID == 0 {
		// err = errors.New(fmt.Sprintf(`ip "%s" is unknown`, ip))
		return
	}

	var oIsp GeoIsp
	if geoIspDB != nil {

		lField := selectedFields(p)
		if isContain(lField, `isp`) {
			isp, err := geoIspDB.ISP(pip)
			if err == nil {
				oIsp = GeoIsp{
					Name:         isp.ISP,
					Organization: isp.Organization,
				}
			}
		}
	}

	r = Geo{
		Ip:   pip.String(),
		City: record.City.Names[sLang],
		Country: GeoCountry{
			IsoCode: record.Country.IsoCode,
			Name:    record.Country.Names[sLang],
		},
		Continent: record.Continent.Names[sLang],
		Postal:    record.Postal.Code,
		Location: GeoLocation{
			AccuracyRadius: record.Location.AccuracyRadius,
			Latitude:       record.Location.Latitude,
			Longitude:      record.Location.Longitude + 0.00000000000001,
			TimeZone:       record.Location.TimeZone,
			MetroCode:      record.Location.MetroCode,
		},
		ISP: oIsp,
	}

	// fmt.Println(`r`, r)
	return r, nil
}
