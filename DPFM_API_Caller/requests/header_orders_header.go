package requests

type HeaderOrdersHeader struct {
	InvoiceDocument         *int     `json:"InvoiceDocument"`
	OrderID                 *int     `json:"OrderID`
	InvoiceDocumentType     *string  `json:"InvoiceDocumentType"`
	BillToParty             *int     `json:"BillToParty"`
	BillFromParty           *int     `json:"BillFromParty"`
	BillToPartyLanguage     *string  `json:"BillToPartyLanguage"`
	BillFromPartyLanguage   *string  `json:"BillFromPartyLanguage"`
	TotalNetAmount          *float32 `json:"TotalNetAmount"`
	TransactionCurrency     *string  `json:"TransactionCurrency"`
	BusinessPartnerCurrency *string  `json:"BusinessPartnerCurrency"`
	TotalTaxAmount          *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount        *float32 `json:"TotalGrossAmount"`
	Incoterms               *string  `json:"Incoterms"`
	PaymentTerms            *string  `json:"PaymentTerms"`
	DueCalculationBaseDate  *string  `json:"DueCalculationBaseDate"`
	PaymentMethod           *string  `json:"PaymentMethod"`
	BillToAddressID         *int     `json:"BillToAddressID"`
	BillFromAddressID       *int     `json:"BillFromAddressID"`
	BillToCountry           *string  `json:"BillToCountry"`
	BillToLocalRegion       *string  `json:"BillToLocalRegion"`
	BillFromCountry         *string  `json:"BillFromCountry"`
	BillFromLocalRegion     *string  `json:"BillFromLocalRegion"`
}
