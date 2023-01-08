package services

import (
	"errors"

	"github.com/yosa12978/MyShirts/internal/helpers"
	"github.com/yosa12978/MyShirts/internal/models"
	"github.com/yosa12978/MyShirts/internal/repos"
)

type AccountService interface {
	GetAccount(id string) (models.Account, error)
	VerifyToken(token string) bool
	Delete(id string, token string) error
	CreateAccount(account models.Account) error
}

type accountService struct {
	accountRepository repos.AccountRepo
}

func NewAccountService() AccountService {
	return &accountService{
		accountRepository: repos.NewAccountRepoMongo(),
	}
}

func (as *accountService) GetAccount(id string) (models.Account, error) {
	return as.accountRepository.GetByID(id)
}

func (as *accountService) VerifyToken(token string) bool {
	_, err := as.accountRepository.GetByToken(token)
	return err == nil
}

func (as *accountService) Delete(id string, token string) error {
	acc, err := as.accountRepository.GetByToken(token)
	if err != nil {
		return err
	}
	if acc.Role != "ADMIN" && acc.Token != token {
		return errors.New("you have no permission to do this")
	}
	return as.accountRepository.Delete(id)
}

func (as *accountService) CreateAccount(account models.Account) error {
	(&account).Salt = helpers.NewSalt()
	(&account).Password = helpers.NewMD5(account.Password + (&account).Salt)
	(&account).Token = helpers.NewUserToken()
	(&account).Role = models.ROLE_USER
	return as.accountRepository.Create(account)
}
