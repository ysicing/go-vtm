package vtm

import (
	"encoding/json"
)

type ExtrasUserCounters64Statistics struct {
	Statistics struct {
		Counter *int `json:"counter"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetExtrasUserCounters64Statistics() (*ExtrasUserCounters64Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/extras/user_counters_64")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ExtrasUserCounters64Statistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
