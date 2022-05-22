package account

import (
	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/utils"
)

var accountRepository AccountRepository

func injectionAccountRepository() {
	accountRepository = AccountRepository{
		Client: database.GetInstance().Account,
	}
}

func LoginService(id, password string) *ent.Account {
	injectionAccountRepository()
	return accountRepository.FindByIdAndPassword(id, utils.EncryptSHA256(password))
}

func LogoutService(id string) *ent.Account {
	injectionAccountRepository()
	return accountRepository.FindById(id)
}

func JoinService(params *JoinParams) {
	injectionAccountRepository()
	accountRepository.CreateAccount(params)
}
