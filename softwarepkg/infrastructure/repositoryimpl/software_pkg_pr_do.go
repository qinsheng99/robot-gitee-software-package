package repositoryimpl

import (
	"time"

	"github.com/google/uuid"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
)

const (
	mergeStatus   = 1
	unMergeStatus = 2
)

type SoftwarePkgPRDO struct {
	// must set "uuid" as the name of column
	PkgId     uuid.UUID `gorm:"column:uuid;type:uuid"`
	Num       int       `gorm:"column:num"`
	Link      string    `gorm:"column:link"`
	Merge     int       `gorm:"column:merge"`
	PkgName   string    `gorm:"column:pkg_name"`
	CreatedAt int64     `gorm:"column:created_at"`
	UpdatedAt int64     `gorm:"column:updated_at"`
}

func (s softwarePkgPR) toSoftwarePkgPRDO(p *domain.PullRequest, id uuid.UUID, do *SoftwarePkgPRDO) {
	*do = SoftwarePkgPRDO{
		PkgId:     id,
		Num:       p.Num,
		Link:      p.Link,
		PkgName:   p.Pkg.Name,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}

	if p.IsMerged() {
		do.Merge = mergeStatus
	} else {
		do.Merge = unMergeStatus
	}
}

func (do *SoftwarePkgPRDO) toDomainPullRequest() (pr domain.PullRequest) {
	pr.Link = do.Link
	pr.Num = do.Num

	if do.Merge == mergeStatus {
		pr.SetMerged()
	}

	pr.Pkg.Name = do.PkgName
	pr.Pkg.Id = do.PkgId.String()

	return
}
