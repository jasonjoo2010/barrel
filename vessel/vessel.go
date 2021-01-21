package vessel

import (
	"context"

	dockerClient "github.com/docker/docker/client"
	"github.com/projectcalico/libcalico-go/lib/clientv3"
	"github.com/projecteru2/barrel/store"
	"github.com/projecteru2/barrel/types"
)

// ContainerVessel .
type ContainerVessel interface {
	ListContainers() ([]types.ContainerInfo, error)
	// GetContainerByID(containerID string) (types.ContainerInfo, error)
	UpdateContainer(context.Context, types.ContainerInfo) error
	DeleteContainer(context.Context, types.ContainerInfo) error
}

// Vessel .
type Vessel interface {
	NodeName() string
	ContainerVessel() ContainerVessel
	CalicoIPAllocator() CalicoIPAllocator
	FixedIPAllocator() FixedIPAllocator
}

type vesselImpl struct {
	nodeName         string
	containerVessel  ContainerVessel
	fixedIPAllocator FixedIPAllocator
}

// NewVessel .
func NewVessel(nodeName string, cliv3 clientv3.Interface, dockerCli *dockerClient.Client, driverName string, stor store.Store) Vessel {
	return vesselImpl{
		nodeName:         nodeName,
		fixedIPAllocator: NewFixedIPAllocator(NewIPPoolManager(cliv3, dockerCli, driverName), stor),
	}
}

func (impl vesselImpl) NodeName() string {
	return impl.nodeName
}

func (impl vesselImpl) ContainerVessel() ContainerVessel {
	return impl.containerVessel
}

func (impl vesselImpl) CalicoIPAllocator() CalicoIPAllocator {
	return impl.fixedIPAllocator
}

func (impl vesselImpl) FixedIPAllocator() FixedIPAllocator {
	return impl.fixedIPAllocator
}