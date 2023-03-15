package repositoryimpl

import (
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain/repository"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/infrastructure/postgresql"
)

type softwarePkg struct {
	softwarePkgPR
}

func NewSoftwarePkg(cfg *Config) repository.PullRequest {
	return softwarePkg{
		softwarePkgPR{postgresql.NewDBTable(cfg.Table.SoftwarePkgPR)},
	}
}
