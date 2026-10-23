package platform

import (
	"context"
	"errors"

	"github.com/ubnetdef/ScoreTrak/pkg/platform/none"

	"github.com/ubnetdef/ScoreTrak/pkg/config"
	"github.com/ubnetdef/ScoreTrak/pkg/platform/docker"
	"github.com/ubnetdef/ScoreTrak/pkg/platform/kubernetes"
	"github.com/ubnetdef/ScoreTrak/pkg/platform/worker"
)

// Platform is an interface that allows ScoreTrak to deploy/remove the worker containers on a given environment like docker, docker swarm, or kubernetes.
type Platform interface {
	DeployWorkers(ctx context.Context, info worker.Info) error
	RemoveWorkers(ctx context.Context, info worker.Info) error
}

var ErrInvalidPlatform = errors.New("invalid platform specified")

const (
	Docker     = "docker"
	Swarm      = "swarm"
	Kubernetes = "kubernetes"
	None       = "none"
)

func NewPlatform(config config.StaticConfig) (Platform, error) {
	switch config.Platform.Use {
	case Docker, Swarm:
		return docker.NewDocker(config)
	case Kubernetes:
		return kubernetes.NewKubernetes(config)
	case None:
		return none.NewNonePlatform()
	default:
		return nil, ErrInvalidPlatform
	}
}
