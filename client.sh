#!/bin/bash -x

# curl -s -d '{geo(ip:"soulogic.com")  {ip city} time}' 127.0.0.1:59999/api | jq
# curl -s -d '{geo(ip:"127.0.0.1")     {ip city} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"127.0.0.1") {ip city country{iso_code name}} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"202.130.251.3") {ip city country{iso_code name}} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"202.130.251.3") {ip city country{iso_code name} location{ accuracy_radius latitude longitude metro_code time_zone }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"202.130.251.3") {ip city country{iso_code name} location{ metro_code time_zone }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"8.8.8.8") {ip city country{iso_code name} postal_code continent location{ metro_code time_zone latitude longitude accuracy_radius }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"202.130.251.3") {ip city country{iso_code name} postal_code continent location{ metro_code time_zone latitude longitude accuracy_radius }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"106.187.48.203") {ip city country{iso_code name} postal_code continent location{ metro_code time_zone latitude longitude accuracy_radius }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"2400:8900::f03c:91ff:fedf:dbf2",lang:ja) {ip city country{iso_code name} postal_code continent location{ metro_code time_zone latitude longitude accuracy_radius }} time}' 127.0.0.1:59999/api | jq
curl -s -d '{a:geo(ip:"106.187.48.203") {city} b:geo(ip:"202.130.251.3") {city}}' 127.0.0.1:59999/api | jq
curl -s -d '{geo(ip:"202.130.251.3",lang:pt) {ip city country{iso_code name} postal_code continent location{ metro_code time_zone latitude longitude accuracy_radius }} time}' 127.0.0.1:59999/api | jq
#curl -s -d '{__schema { types { name fields { name description}  } } }' 127.0.0.1:59999/api | jq

# curl -s -d '{ __schema { queryType { name, fields { name, description } } } }' 127.0.0.1:59999/api | jq
# curl -v -s -d '{ __schema { } }' 127.0.0.1:59999/api | jq
