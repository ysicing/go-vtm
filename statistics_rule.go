package vtm

import (
	"encoding/json"
)

type RuleStatistics struct {
	Statistics struct {
		Aborts                *int `json:"aborts"`
		Discards              *int `json:"discards"`
		ExecutionTimeWarnings *int `json:"execution_time_warnings"`
		Executions            *int `json:"executions"`
		PoolSelect            *int `json:"pool_select"`
		Responds              *int `json:"responds"`
		Retries               *int `json:"retries"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetRuleStatistics(name string) (*RuleStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/rules/" + name)
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(RuleStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
