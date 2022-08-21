package util

import (
	"context"
	"github.com/google/go-github/v43/github"
)

type Connector struct {
	Client  *github.Client
	Context context.Context
}

func (connector Connector) GetDatabase(name string) Database {
	repos, _, err := connector.Client.Repositories.List(connector.Context, "", nil)
	if err != nil {
		panic(err)
	}
	for _, repo := range repos {
		if repo.GetName() == name {
			return Database{repo}
		}
	}
	repo := &github.Repository{
		Name:    github.String(name),
		Private: github.Bool(true),
	}
	_, _, err = connector.Client.Repositories.Create(connector.Context, "", repo)
	if err != nil {
		panic(err)
	}
	return Database{repo}
}
