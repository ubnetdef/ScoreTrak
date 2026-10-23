package policyrepo

import (
	"context"

	"github.com/ubnetdef/ScoreTrak/pkg/policy"
)

type Repo interface {
	Get(ctx context.Context) (*policy.Policy, error)
	Update(ctx context.Context, u *policy.Policy) error
}
