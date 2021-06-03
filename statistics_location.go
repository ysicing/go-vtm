package vtm

import (
	"encoding/json"
)

type LocationStatistics struct {
	Statistics struct {
		Load      *int `json:"load"`
		Responses *int `json:"responses"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetLocationStatistics(name string) (*LocationStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/locations/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(LocationStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
