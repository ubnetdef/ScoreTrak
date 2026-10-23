package none

import (
	"context"

	"github.com/ubnetdef/ScoreTrak/pkg/platform/util"
	"github.com/ubnetdef/ScoreTrak/pkg/platform/worker"
)

type None struct{}

func NewNonePlatform() (d *None, err error) {
	return &None{}, nil
}

func (d *None) DeployWorkers(ctx context.Context, info worker.Info) (err error) {
	return util.ErrSkippedOperation
}

func (d *None) RemoveWorkers(ctx context.Context, info worker.Info) error {
	return util.ErrSkippedOperation
}
