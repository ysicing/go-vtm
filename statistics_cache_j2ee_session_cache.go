package vtm

import (
	"encoding/json"
)

type CacheJ2EeSessionCacheStatistics struct {
	Statistics struct {
		Entries    *int `json:"entries"`
		EntriesMax *int `json:"entries_max"`
		HitRate    *int `json:"hit_rate"`
		Hits       *int `json:"hits"`
		Lookups    *int `json:"lookups"`
		Misses     *int `json:"misses"`
		Oldest     *int `json:"oldest"`
	} `json:"statistics"`
}

func (vtm VirtualTrafficManager) GetCacheJ2EeSessionCacheStatistics() (*CacheJ2EeSessionCacheStatistics, *vtmErrorResponse) {
	conn := vtm.connector.getChildConnector("/tm/6.2/status/local_tm/statistics/cache/j2ee_session_cache")
	data, ok := conn.get()
	if ok != true {
		object := new(vtmErrorResponse)
		json.NewDecoder(data).Decode(object)
		return nil, object
	}
	object := new(CacheJ2EeSessionCacheStatistics)
	if err := json.NewDecoder(data).Decode(object); err != nil {
		panic(err)
	}
	return object, nil
}
