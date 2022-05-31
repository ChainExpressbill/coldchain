package account

import (
	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/utils"
	"github.com/gofiber/fiber/v2"
)

func LoginService(id, password string, c *fiber.Ctx) (*ent.Account, error) {
	account, err := FindByIdAndPassword(id, utils.EncryptSHA256(password))
	if err != nil {
		return nil, err
	}

	err2 := database.SetSession(account.ID, c)
	if err2 != nil {
		return nil, err
	}

	return account, nil
}

func LogoutService(id string) *ent.Account {
	return FindById(id)
}

func JoinService(params *JoinParams) {
	CreateAccount(params)
}
