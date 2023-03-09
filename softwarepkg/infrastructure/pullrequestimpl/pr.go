package pullrequestimpl

import (
	sdk "github.com/opensourceways/go-gitee/gitee"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
)

type pullrequestImpl struct {
	cli iClient
	cfg Config
}

type iClient interface {
	GetBot() (sdk.User, error)
	CreatePullRequest(org, repo, title, body, head, base string, canModify bool) (sdk.PullRequest, error)
}

func (impl *pullrequestImpl) Create(pkg *domain.SoftwarePkg) (domain.PullRequest, error) {
	return domain.PullRequest{}, nil
}

func (impl *pullrequestImpl) Merge(*domain.PullRequest) error {
	return nil
}

func (impl *pullrequestImpl) Close(*domain.PullRequest) error {
	return nil
}
