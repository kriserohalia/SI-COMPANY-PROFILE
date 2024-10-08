package manager

import "github.com/kriserohalia/SI-COMPANY-PROFILE/server/repository"

type RepoManager interface {
	AuthRepo() repository.AuthRepo
	UserRepo() repository.UserRepo
	CategoryRepo() repository.CategoryRepo
	ProductRepo() repository.ProdukRepo
	DetailRepo() repository.ProductDetailRepo
	CustomRepo() repository.CustomRepo
	TransactionRepo() repository.TransactionRepo
	CompanyRepo() repository.CompanyRepo
	ServiceRepo() repository.ServiceRepo
}

type repoManager struct {
	infra Inframanager
}

// ServiceRepo implements RepoManager.
func (repo *repoManager) ServiceRepo() repository.ServiceRepo {
	return repository.NewServiceRepo(repo.infra.Connection())
}

// CompanyRepo implements RepoManager.
func (repo *repoManager) CompanyRepo() repository.CompanyRepo {
	return repository.NewRepoCompany(repo.infra.Connection())
}

// TransactionRepo implements RepoManager.
func (repo *repoManager) TransactionRepo() repository.TransactionRepo {
	return repository.NewTransactionRepo(repo.infra.Connection())
}

// CustomRepo implements RepoManager.
func (repo *repoManager) CustomRepo() repository.CustomRepo {
	return repository.NewCustomRepo(repo.infra.Connection())
}

// DetailRepo implements RepoManager.
func (repo *repoManager) DetailRepo() repository.ProductDetailRepo {
	return repository.NewProductDetail(repo.infra.Connection())
}

// ProductRepo implements RepoManager.
func (repo *repoManager) ProductRepo() repository.ProdukRepo {
	return repository.NewProductRepo(repo.infra.Connection())

}

// CategoryRepo implements RepoManager.
func (repo *repoManager) CategoryRepo() repository.CategoryRepo {
	return repository.NewCategoryRepo(repo.infra.Connection())
}

// AuthRepo implements RepoManager.
func (repo *repoManager) AuthRepo() repository.AuthRepo {
	return repository.NewAuthRepo(repo.infra.Connection())
}

func (repo *repoManager) UserRepo() repository.UserRepo {
	return repository.NewUserRepo(repo.infra.Connection())
}

func NewRepoManager(infra Inframanager) RepoManager {
	return &repoManager{infra}

}
