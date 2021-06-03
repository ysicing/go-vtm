package vtm

import (
	"encoding/json"
)

type EventStatistics struct {
	Statistics struct {
		Matched *int `json:"matched"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetEventStatistics(name string) (*EventStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/events/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(EventStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
