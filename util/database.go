package util

import (
	"fmt"
	"github.com/google/go-github/v43/github"
)

type Database struct {
	repo *github.Repository
}

func (database Database) GetTable(name string) {
	fmt.Println(database.repo.GetBranchesURL())
}
