package clustermap

import "github.com/kris-nova/kubicorn/apis/cluster"

func NewSimpleAzureCluster(name string) *cluster.Cluster {
	return &cluster.Cluster{
		Name: name,
		ServerPools: []*cluster.ServerPool{
			{
				PoolType: cluster.ServerPoolType_Master,
				Cloud:    cluster.ServerPoolCloud_Azure,
				Name:     "azure-master",
				Count:    1,
				Networks: []*cluster.Network{
					{
						NetworkType: cluster.NetworkType_Public,
					},
				},
			},
			{
				PoolType: cluster.ServerPoolType_Node,
				Cloud:    cluster.ServerPoolCloud_Azure,
				Name:     "azure-node",
				Count:    3,
				Networks: []*cluster.Network{
					{
						NetworkType: cluster.NetworkType_Public,
					},
				},
			},
		},
	}
}