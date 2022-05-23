package orders

type OrderSearchParams struct {
	Orderer   string `json:"orderer,omitempty"`
	Receiver  string `json:"receiver,omitempty"`
	StartDate string `json:"startDate,omitempty"`
	EndDate   string `json:"endDate,omitempty"`
	Page      string `json:"page,omitempty"`
	Size      string `json:"size,omitempty"`
}

type OrderRequestBody struct {
	// jwt 적용 전 임시로 id를 받아서 처리
	AccountId        string `json:"accountId,omitempty"`
	Oid              string `json:"oid,omitempty"`
	Orderer          string `json:"orderer,omitempty"`
	Receiver         string `json:"receiver,omitempty"`
	DrugName         string `json:"drugName,omitempty"`
	DrugStandard     string `json:"drugStandard,omitempty"`
	Quantity         int    `json:"quantity,omitempty"`
	RegisterName     string `json:"registerName,omitempty"`
	StorageCondition string `json:"storageCondition,omitempty"`
}
