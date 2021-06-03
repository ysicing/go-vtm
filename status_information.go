package vtm

import (
	"encoding/json"
)

type SystemInformation struct {
	Information struct {
		Platform  *string `json:"platform"`
		TmVersion *string `json:"tm_version"`
		Uuid      *string `json:"uuid"`
	} `json:"information"`
}

func (vtm VirtualTrafficManager) GetSystemInformation() (*SystemInformation, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/information")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(SystemInformation)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
