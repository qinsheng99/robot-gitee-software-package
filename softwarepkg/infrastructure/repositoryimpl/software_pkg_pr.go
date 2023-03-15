package repositoryimpl

import (
	"github.com/google/uuid"

	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/domain"
	"github.com/opensourceways/robot-gitee-software-package/softwarepkg/infrastructure/postgresql"
)

type softwarePkgPR struct {
	cli dbClient
}

func (s softwarePkgPR) Add(p *domain.PullRequest) error {
	u, err := uuid.Parse(p.Pkg.Id)
	if err != nil {
		return err
	}

	var do SoftwarePkgPRDO
	s.toSoftwarePkgPRDO(p, u, &do)

	filter := SoftwarePkgPRDO{PkgId: u}

	return s.cli.Insert(&filter, &do)
}

func (s softwarePkgPR) Save(p *domain.PullRequest) error {
	u, err := uuid.Parse(p.Pkg.Id)
	if err != nil {
		return err
	}
	filter := SoftwarePkgPRDO{PkgId: u}

	var do SoftwarePkgPRDO
	s.toSoftwarePkgPRDO(p, u, &do)

	return s.cli.UpdateRecord(&filter, &do)
}

func (s softwarePkgPR) Find(num int) (domain.PullRequest, error) {
	filter := SoftwarePkgPRDO{Num: num}

	var res SoftwarePkgPRDO
	if err := s.cli.GetRecord(&filter, &res); err != nil {
		return domain.PullRequest{}, err
	}

	return res.toDomainPullRequest(), nil
}

func (s softwarePkgPR) FindAll(isMerged bool) ([]domain.PullRequest, error) {
	filter := SoftwarePkgPRDO{}
	if isMerged {
		filter.Merge = mergeStatus
	} else {
		filter.Merge = unMergeStatus
	}
	
	var res []SoftwarePkgPRDO

	if err := s.cli.GetRecords(
		&filter,
		&res,
		postgresql.Pagination{},
		nil,
	); err != nil {
		return nil, err
	}

	var p = make([]domain.PullRequest, len(res))

	for i := range res {
		p[i] = res[i].toDomainPullRequest()
	}

	return p, nil
}

func (s softwarePkgPR) Remove(num int) error {
	filter := SoftwarePkgPRDO{Num: num}

	return s.cli.DeleteRecord(&filter)
}
