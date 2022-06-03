package orders

import (
	"time"

	"github.com/ChainExpressbill/coldchain/ent"
)

type OrderSearchParams struct {
	Orderer   string    `json:"orderer,omitempty"`
	Receiver  string    `json:"receiver,omitempty"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Page      int       `json:"page"`
	Size      int       `json:"size"`
}

type GetOrdersResponse struct {
	OrderList  []*ent.Order `json:"orderList"`
	Page       int          `json:"page"`
	Size       int          `json:"size"`
	TotalPage  int          `json:"totalSize"`
	TotalCount int          `json:"totalCount"`
}

type OrderRequestBody struct {
	// jwt 적용 전 임시로 id를 받아서 처리
	AccountId        string `json:"accountId,omitempty"`
	Id               int    `json:"id,omitempty"`
	Oid              string `json:"oid,omitempty"`
	Orderer          string `json:"orderer,omitempty"`
	Receiver         string `json:"receiver,omitempty"`
	DrugName         string `json:"drugName,omitempty"`
	DrugStandard     string `json:"drugStandard,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
	RegisterName     string `json:"registerName,omitempty"`
	StorageCondition string `json:"storageCondition,omitempty"`
}
