package requests

type OrdersHeader struct {
	OrderID                int     `json:"OrderID"`
	OrderType              string  `json:"OrderType"`
	Buyer                  int     `json:"Buyer"`
	Seller                 int     `json:"Seller"`
	ContractType           *string `json:"ContractType"`
	OrderValidityStartDate *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate   *string `json:"OrderValidityEndDate"`
	TransactionCurrency    string  `json:"TransactionCurrency"`
	Incoterms              *string `json:"Incoterms"`
	BillFromParty          *int    `json:"BillFromParty"`
	BillToParty            *int    `json:"BillToParty"`
	BillFromCountry        *string `json:"BillFromCountry"`
	BillToCountry          *string `json:"BillToCountry"`
	Payer                  *int    `json:"Payer"`
	Payee                  *int    `json:"Payee"`
	IsExportImportDelivery *bool   `json:"IsExportImportDelivery"`
}
