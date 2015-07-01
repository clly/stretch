// Author: Connor Kelly

package stretch

import "encoding/json"

type ClusterHealth struct {
	ClusterName         string `json:"cluster_name,omitempty"`
	Status              string `json:"status,omitempty"`
	TimedOut            bool   `json:"timed_out,omitempty"`
	NumberOfNodes       int    `json:"number_of_nodes,omitempty"`
	NumberOfDataNodes   int    `json:"number_of_data_nodes,omitempty"`
	ActivePrimaryShards int    `json:"active_primary_shards,omitempty"`
	ActiveShards        int    `json:"active_shards,omitempty"`
	RelocatingShards    int    `json:"relocating_shards,omitempty"`
	InitializingShards  int    `json:"initializing_shards,omitempty"`
	UnassignedShards    int    `json:"unassigned_shards,omitempty"`
}

func GetShortClusterHealth() string {
	endpoint := "/_cluster/health"
	jsonHealth := makeRequest("localhost:9200", endpoint)
	var health ClusterHealth
	json.Unmarshal(jsonHealth, &health)
	status := health.Status
	return status
}

func GetClusterHealth() ClusterHealth {
	endpoint := "/_cluster/health"
	jsonHealth := makeRequest("localhost:9200", endpoint)
	var health ClusterHealth
	json.Unmarshal(jsonHealth, &health)
	return health
}
