package vtm

import (
	"encoding/json"
)

type ExtrasUserCounters32Statistics struct {
	Statistics struct {
		Counter *int `json:"counter"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetExtrasUserCounters32Statistics() (*ExtrasUserCounters32Statistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/extras/user_counters_32")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ExtrasUserCounters32Statistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
