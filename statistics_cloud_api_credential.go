package vtm

import (
	"encoding/json"
)

type CloudApiCredentialStatistics struct {
	Statistics struct {
		NodeCreations  *int `json:"node_creations"`
		NodeDeletions  *int `json:"node_deletions"`
		StatusRequests *int `json:"status_requests"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCloudApiCredentialStatistics(name string) (*CloudApiCredentialStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/cloud_api_credentials/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CloudApiCredentialStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
