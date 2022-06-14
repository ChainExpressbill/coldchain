package orders

import (
	"time"

	"github.com/ChainExpressbill/coldchain/ent"
)

type OrderSearchParams struct {
	Orderer   string    `json:"orderer,omitempty"`
	Receiver  string    `json:"receiver,omitempty"`
	Oid       string    `json:"oid,omitempty"`
	StartDate time.Time `json:"startDate"`
	EndDate   time.Time `json:"endDate"`
	Page      int       `json:"page"`
	Size      int       `json:"size"`
}

type GetOrdersResponse struct {
	OrderList  []*ent.Order `json:"orderList"`
	Page       int          `json:"page"`
	Size       int          `json:"size"`
	TotalPage  int          `json:"totalPage"`
	TotalCount int          `json:"totalCount"`
}

type OrderRequestBody struct {
	// jwt 적용 전 임시로 id를 받아서 처리
	AccountId           string `json:"accountId,omitempty"`
	Id                  int    `json:"id,omitempty"`
	Oid                 string `json:"oid,omitempty"`
	Orderer             string `json:"orderer"`
	Receiver            string `json:"receiver"`
	DrugName            string `json:"drugName"`
	DrugStandard        string `json:"drugStandard"`
	Quantity            int    `json:"quantity"`
	RegisterName        string `json:"registerName,omitempty"`
	StorageCondition    string `json:"storageCondition"`
	Memo                string `json:"memo,omitempty"`
	DeliveryDriverName  string `json:"deliveryDriverName"`
	DeliveryDriverTelNo string `json:"deliveryDriverTelNo"`
}
