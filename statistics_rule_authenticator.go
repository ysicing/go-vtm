package vtm

import (
	"encoding/json"
)

type RuleAuthenticatorStatistics struct {
	Statistics struct {
		Errors   *int `json:"errors"`
		Fails    *int `json:"fails"`
		Passes   *int `json:"passes"`
		Requests *int `json:"requests"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetRuleAuthenticatorStatistics(name string) (*RuleAuthenticatorStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/rule_authenticators/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(RuleAuthenticatorStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
