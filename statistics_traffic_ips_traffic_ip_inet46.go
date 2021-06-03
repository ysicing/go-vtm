package vtm

import (
	"encoding/json"
)

type TrafficIpsTrafficIpInet46Statistics struct {
	Statistics struct {
		State *string `json:"state"`
		Time  *int    `json:"time"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetTrafficIpsTrafficIpInet46Statistics(name string) (*TrafficIpsTrafficIpInet46Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/traffic_ips/traffic_ip_inet46/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(TrafficIpsTrafficIpInet46Statistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
