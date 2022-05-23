package account

import (
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/utils"
)

func LoginService(id, password string) *ent.Account {
	return FindByIdAndPassword(id, utils.EncryptSHA256(password))
}

func LogoutService(id string) *ent.Account {
	return FindById(id)
}

func JoinService(params *JoinParams) {
	CreateAccount(params)
}
