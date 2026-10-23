package configrepo

import (
	"context"

	"github.com/ubnetdef/ScoreTrak/pkg/config"
)

type Repo interface {
	Get(ctx context.Context) (*config.DynamicConfig, error)
	Update(context.Context, *config.DynamicConfig) error
}
