package vtm

import (
	"encoding/json"
)

type TrafficIpsTrafficIpStatistics struct {
	Statistics struct {
		State *string `json:"state"`
		Time  *int    `json:"time"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetTrafficIpsTrafficIpStatistics(name string) (*TrafficIpsTrafficIpStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/traffic_ips/traffic_ip/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(TrafficIpsTrafficIpStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
