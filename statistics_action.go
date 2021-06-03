package vtm

import (
	"encoding/json"
)

type ActionStatistics struct {
	Statistics struct {
		Processed *int `json:"processed"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetActionStatistics(name string) (*ActionStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/actions/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ActionStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
