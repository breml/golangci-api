package request

import (
	"fmt"
	"strings"

	"github.com/golangci/golangci-api/pkg/logutil"
)

type Repo struct {
	Provider string `request:",url"`
	Owner    string `request:",url"`
	Name     string `request:",url"`
}

func (r Repo) FullName() string {
	return strings.ToLower(fmt.Sprintf("%s/%s", r.Owner, r.Name))
}

func (r Repo) String() string {
	return fmt.Sprintf("%s/%s", r.Provider, r.FullName())
}

func (r Repo) FillLogContext(lctx logutil.Context) {
	lctx["repo"] = r.String()
}
