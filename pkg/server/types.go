package server

import (
	"gitlab.alipay-inc.com/afe/mosn/pkg/types"
	"gitlab.alipay-inc.com/afe/mosn/pkg/api/v2"
	"gitlab.alipay-inc.com/afe/mosn/pkg/log"
)

type Config struct {
	LogPath    string
	LogLevel   log.LogLevel
}

type Server interface {
	AddListener(lc v2.ListenerConfig)

	Start()

	Restart()

	Close()
}

type NetworkFilterFactoryCb func(manager types.FilterManager)

type NetworkFilterConfigFactory interface {
	CreateFilterFactory(clusterManager types.ClusterManager) NetworkFilterFactoryCb
}

type ClusterConfigFactoryCb interface {
	UpdateClusterConfig(configs []v2.Cluster) error
}

type ClusterHostFactoryCb interface {
	UpdateClusterHost(cluster string, priority uint32, hosts []v2.Host) error
}

type ClusterManagerFilter interface {
	OnCreated(cccb ClusterConfigFactoryCb, chcb ClusterHostFactoryCb)
}
