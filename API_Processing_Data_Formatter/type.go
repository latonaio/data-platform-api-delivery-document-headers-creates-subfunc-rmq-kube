package api_processing_data_formatter

type SDC struct {
	MetaData                            *MetaData                              `json:"MetaData"`
	OrderID                             *[]OrderID                             `json:"OrderID"`
	DeliveryDocument                    *[]DeliveryDocument                    `json:"DeliveryDocument"`
	CalculateInvoiceDocument            *CalculateInvoiceDocument              `json:"CalculateInvoiceDocument"`
	HeaderOrdersHeader                  *[]HeaderOrdersHeader                  `json:"HeaderOrdersHeader"`
	HeaderDeliveryDocumentHeader        *[]HeaderDeliveryDocumentHeader        `json:"HeaderDeliveryDocumentHeader"`
	HeaderOrdersHeaderPartner           *[]HeaderOrdersHeaderPartner           `json:"HeaderOrdersHeaderPartner"`
	HeaderDeliveryDocumentHeaderPartner *[]HeaderDeliveryDocumentHeaderPartner `json:"HeaderDeliveryDocumentHeaderPartner"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type OrderID struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type DeliveryDocumentKey struct {
	CompleteDeliveryIsDefined *bool  `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}

type DeliveryDocument struct {
	InvoiceDocument           *int   `json:"InvoiceDocument"`
	DeliveryDocument          *int   `json:"DeliveryDocument`
	CompleteDeliveryIsDefined *bool  `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}

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

type HeaderDeliveryDocumentHeader struct {
	InvoiceDocument               *int     `json:"InvoiceDocument"`
	DeliveryDocument              *int     `json:"DeliveryDocument"`
	Buyer                         *int     `json:"Buyer`
	Seller                        *int     `json:"Seller"`
	ReferenceDocument             *int     `json:"ReferenceDocument"`
	ReferenceDocumentItem         *int     `json:"ReferenceDocumentItem"`
	OrderID                       *string  `json:"OrderID"`
	OrderItem                     *string  `json:"OrderItem"`
	ContractType                  *string  `json:"ContractType"`
	OrderValidityStartDate        *string  `json:"OrderValidityStartDate"`
	OrderValidityEndDate          *string  `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate      *string  `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate        *string  `json:"InvoiceScheduleEndDate"`
	IssuingPlantTimeZone          *string  `json:"IssuingPlantTimeZone"`
	ReceivingPlantTimeZone        *string  `json:"ReceivingPlantTimeZone"`
	DocumentDate                  *string  `json:"DocumentDate"`
	PlannedGoodsIssueDate         *string  `json:"PlannedGoodsIssueDate"`
	PlannedGoodsIssueTime         *string  `json:"PlannedGoodsIssueTime"`
	PlannedGoodsReceiptDate       *string  `json:"PlannedGoodsReceiptDate"`
	PlannedGoodsReceiptTime       *string  `json:"PlannedGoodsReceiptTime"`
	BillingDocumentDate           *string  `json:"BillingDocumentDate"`
	CompleteDeliveryIsDefined     *bool    `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus         *string  `json:"OverallDeliveryStatus"`
	CreationDate                  *string  `json:"CreationDate`
	CreationTime                  *string  `json:"CreationTime"`
	IssuingBlockReason            *bool    `json:"IssuingBlockReason"`
	ReceivingBlockReason          *bool    `json:"ReceivingBlockReason"`
	GoodsIssueOrReceiptSlipNumber *string  `json:"GoodsIssueOrReceiptSlipNumber"`
	HeaderBillingStatus           *string  `json:"HeaderBillingStatus"`
	HeaderBillingConfStatus       *string  `json:"HeaderBillingConfStatus"`
	HeaderBillingBlockReason      *bool    `json:"HeaderBillingBlockReason"`
	HeaderGrossWeight             *float32 `json:"HeaderGrossWeight"`
	HeaderNetWeight               *float32 `json:"HeaderNetWeight"`
	HeaderWeightUnit              *string  `json:"HeaderWeightUnit"`
	Incoterms                     *string  `json:"Incoterms"`
	BillToCountry                 *string  `json:"BillToCountry"`
	BillFromCountry               *string  `json:"BillFromCountry"`
	IsExportImportDelivery        *bool    `json:"IsExportImportDelivery"`
	LastChangeDate                *string  `json:"LastChangeDate"`
	IssuingPlantBusinessPartner   *int     `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                  *string  `json:"IssuingPlant"`
	ReceivingPlantBusinessPartner *int     `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                *string  `json:"ReceivingPlant"`
	DeliverToParty                *int     `json:"DeliverToParty"`
	DeliverFromParty              *int     `json:"DeliverFromParty"`
	TransactionCurrency           *string  `json:"TransactionCurrency"`
	OverallDelivReltdBillgStatus  *string  `json:"OverallDelivReltdBillgStatus"`
	StockIsFullyConfirmed         *bool    `json:"StockIsFullyConfirmed"`
}

type CalculateInvoiceDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string `json:"FieldNameWithNumberRange"`
}

type CalculateInvoiceDocumentQueryGets struct {
	ServiceLabel                string `json:"service_label"`
	FieldNameWithNumberRange    string `json:"FieldNameWithNumberRange"`
	InvoiceDocumentLatestNumber *int   `json:"InvoiceDocumentLatestNumber"`
}

type CalculateInvoiceDocument struct {
	InvoiceDocumentLatestNumber *int `json:"InvoiceDocumentLatestNumber"`
	InvoiceDocument             *int `json:"InvoiceDocument"`
}

type HeaderOrdersHeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument"`
	OrderID                 *int    `json:"OrderID`
	PartnerFunction         *string `json"PartnerFunction"`
	BusinessPartner         *int    `json"BusinessPartner"`
	BusinessPartnerFullName *string `json"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json"BusinessPartnerName"`
	Organization            *string `json"Organization"`
	Country                 *string `json"Country"`
	Language                *string `json"Language"`
	Currency                *string `json"Currency"`
	ExternalDocumentID      *string `json"ExternalDocumentID"`
	AddressID               *int    `json"AddressID"`
}

type HeaderDeliveryDocumentHeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument"`
	DeliveryDocument        *int    `json:"DeliveryDocument"`
	PartnerFunction         *string `json"PartnerFunction"`
	BusinessPartner         *int    `json"BusinessPartner"`
	BusinessPartnerFullName *string `json"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json"BusinessPartnerName"`
	Organization            *string `json"Organization"`
	Country                 *string `json"Country"`
	Language                *string `json"Language"`
	Currency                *string `json"Currency"`
	ExternalDocumentID      *string `json"ExternalDocumentID"`
	AddressID               *int    `json"AddressID"`
}
