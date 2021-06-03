package vtm

import (
	"encoding/json"
)

type NetworkInterfaceStatistics struct {
	Statistics struct {
		Collisions *int `json:"collisions"`
		RxBytes    *int `json:"rx_bytes"`
		RxErrors   *int `json:"rx_errors"`
		RxPackets  *int `json:"rx_packets"`
		TxBytes    *int `json:"tx_bytes"`
		TxErrors   *int `json:"tx_errors"`
		TxPackets  *int `json:"tx_packets"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetNetworkInterfaceStatistics(name string) (*NetworkInterfaceStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/network_interface/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(NetworkInterfaceStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
