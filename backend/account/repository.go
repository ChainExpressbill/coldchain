package account

import (
	"context"

	"github.com/ChainExpressbill/coldchain/database"
	"github.com/ChainExpressbill/coldchain/ent"
	"github.com/ChainExpressbill/coldchain/ent/account"
	"github.com/ChainExpressbill/coldchain/utils"
)

type AccountRepository struct {
	Client *ent.AccountClient
}

var ctx = context.TODO()

func FindById(id string) *ent.Account {
	result := database.GetInstance().Account.
		Query().
		Where(account.ID(id)).
		OnlyX(ctx)

	return result
}

func FindByIdAndPassword(id, password string) *ent.Account {
	result := database.GetInstance().Account.
		Query().
		Where(account.ID(id)).
		Where(account.Password(password)).
		Select(
			account.FieldID,
			account.FieldName,
			account.FieldEmailAddress,
			account.FieldCreated,
			account.FieldUpdated,
		).
		OnlyX(ctx)

	return result
}

func CreateAccount(params *JoinParams) {
	database.GetInstance().Account.
		Create().
		SetID(params.Id).
		SetPassword(utils.EncryptSHA256(params.Password)).
		SetName(params.Name).
		SetEmailAddress(params.EmailAddress).
		SaveX(ctx)
}
