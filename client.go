package gitbase

import (
	"context"

	"github.com/google/go-github/v43/github"
)

func InitDatabase(ctx context.Context, token string) *github.Client {
	auth := initAuthClient(ctx, token)
	return github.NewClient(auth)
}
