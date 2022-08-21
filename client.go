package GitBase

import (
	"context"
	"github.com/google/go-github/v43/github"
	"github.com/ignema/GitBase/lib"
	"github.com/ignema/GitBase/util"
)

func Connect(ctx context.Context, token string) util.Connector {
	auth := lib.InitAuthClient(ctx, token)
	return util.Connector{Client: github.NewClient(auth), Context: ctx}
}
