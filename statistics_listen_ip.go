package vtm

import (
	"encoding/json"
)

type ListenIpStatistics struct {
	Statistics struct {
		BytesIn       *int `json:"bytes_in"`
		BytesOut      *int `json:"bytes_out"`
		CurrentConn   *int `json:"current_conn"`
		MaxConn       *int `json:"max_conn"`
		TotalRequests *int `json:"total_requests"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetListenIpStatistics(name string) (*ListenIpStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/listen_ips/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(ListenIpStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
