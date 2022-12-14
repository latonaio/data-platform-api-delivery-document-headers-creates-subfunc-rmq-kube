package requests

type HeaderOrdersHeader struct {
	DeliveryDocument         *int    `json:"DeliveryDocument"`
	OrderID                  *int    `json:"OrderID"`
	OrderType                string  `json:"OrderType"`
	Buyer                    *int    `json:"Buyer"`
	Seller                   *int    `json:"Seller"`
	ContractType             *string `json:"ContractType"`
	OrderValidityStartDate   *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate     *string `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate *string `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate   *string `json:"InvoiceScheduleEndDate"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	IsExportImportDelivery   *bool   `json:"IsExportImportDelivery"`
}
