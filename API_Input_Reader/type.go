package api_input_reader

type SDC struct {
	ConnectionKey     string          `json:"connection_key"`
	Result            bool            `json:"result"`
	RedisKey          string          `json:"redis_key"`
	Filepath          string          `json:"filepath"`
	APIStatusCode     int             `json:"api_status_code"`
	RuntimeSessionID  string          `json:"runtime_session_id"`
	BusinessPartnerID *int            `json:"business_partner"`
	ServiceLabel      string          `json:"service_label"`
	InvoiceDocument   InvoiceDocument `json:"InvoiceDocument"`
	APISchema         string          `json:"api_schema"`
	Accepter          []string        `json:"accepter"`
	InvoiceDocumentID *int            `json:"invoice_document_id"`
	Deleted           bool            `json:"deleted"`
}

type InvoiceDocument struct {
	InvoiceDocument            *int                  `json:"InvoiceDocument"`
	CreationDate               *string               `json:"CreationDate"`
	LastChangeDate             *string               `json:"LastChangeDate"`
	BillToParty                *int                  `json:"BillToParty"`
	BillFromParty              *int                  `json:"BillFromParty"`
	InvoiceDocumentDate        *string               `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime        *string               `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate     *string               `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate       *string               `json:"InvoicePeriodEndDate"`
	AccountingPostingDate      *string               `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled *bool                 `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument   *int                  `json:"CancelledInvoiceDocument"`
	IsExportDelivery           *bool                 `json:"IsExportDelivery"`
	HeaderBillingConfStatus    *string               `json:"HeaderBillingConfStatus"`
	TotalNetAmount             *float32              `json:"TotalNetAmount"`
	TotalTaxAmount             *float32              `json:"TotalTaxAmount"`
	TotalGrossAmount           *float32              `json:"TotalGrossAmount"`
	TransactionCurrency        *string               `json:"TransactionCurrency"`
	Incoterms                  *string               `json:"Incoterms"`
	PaymentTerms               *string               `json:"PaymentTerms"`
	DueCalculationBaseDate     *string               `json:"DueCalculationBaseDate"`
	NetPaymentDays             *int                  `json:"NetPaymentDays"`
	PaymentMethod              *string               `json:"PaymentMethod"`
	HeaderPaymentBlockStatus   *bool                 `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument  *string               `json:"ExternalReferenceDocument"`
	DocumentHeaderText         *string               `json:"DocumentHeaderText"`
	HeaderPartner              []HeaderPartner       `json:"HeaderPartner"`
	HeaderPDF                  []HeaderPDF           `json:"HeaderPDF"`
	InvoiceDocumentItem        []InvoiceDocumentItem `json:"InvoiceDocumentItem"`
	Address                    Address               `json:"Address"`
}

type HeaderPartner struct {
	PartnerFunction         *string              `json:"PartnerFunction"`
	BusinessPartner         *int                 `json:"BusinessPartner"`
	BusinessPartnerFullName *string              `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string              `json:"BusinessPartnerName"`
	Organization            *string              `json:"Organization"`
	Country                 *string              `json:"Country"`
	Language                *string              `json:"Language"`
	Currency                *string              `json:"Currency"`
	ExternalDocumentID      *string              `json:"ExternalDocumentID"`
	AddressID               *int                 `json:"AddressID"`
	HeaderPartnerContact    HeaderPartnerContact `json:"HeaderPartnerContact"`
}

type HeaderPartnerContact struct {
	PartnerFunction   *string `json:"PartnerFunction"`
	ContactID         *int    `json:"ContactID"`
	ContactPersonName *string `json:"ContactPersonName"`
	EmailAddress      *string `json:"EmailAddress"`
	PhoneNumber       *string `json:"PhoneNumber"`
	MobilePhoneNumber *string `json:"MobilePhoneNumber"`
	FaxNumber         *string `json:"FaxNumber"`
	ContactTag1       *string `json:"ContactTag1"`
	ContactTag2       *string `json:"ContactTag2"`
	ContactTag3       *string `json:"ContactTag3"`
	ContactTag4       *string `json:"ContactTag4"`
}

type HeaderPDF struct {
	DocType                  *string `json:"DocType"`
	DocVersionID             *int    `json:"DocVersionID"`
	DocID                    *string `json:"DocID"`
	DocIssuerBusinessPartner *int    `json:"DocIssuerBusinessPartner"`
	FileName                 *string `json:"FileName"`
}

type InvoiceDocumentItem struct {
	InvoiceDocumentItem            *int               `json:"InvoiceDocumentItem"`
	DocumentItemCategory           *string            `json:"DocumentItemCategory"`
	InvoiceDocumentItemText        *string            `json:"InvoiceDocumentItemText"`
	CreationDate                   *string            `json:"CreationDate"`
	CreationTime                   *string            `json:"CreationTime"`
	ItemBillingConfStatus          *string            `json:"ItemBillingConfStatus"`
	BusinessArea                   *string            `json:"BusinessArea"`
	DeliverToParty                 *int               `json:"DeliverToParty"`
	DeliverFromParty               *int               `json:"DeliverFromParty"`
	ProductStandardID              *string            `json:"ProductStandardID"`
	Batch                          *string            `json:"Batch"`
	Product                        *string            `json:"Product"`
	ProductGroup                   *string            `json:"ProductGroup"`
	ProductionPlantPartnerFunction *string            `json:"ProductionPlantPartnerFunction"`
	ProductionPlantBusinessPartner *string            `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                *string            `json:"ProductionPlant"`
	ProductionPlantStorageLocation *string            `json:"ProductionPlantStorageLocation"`
	IssuingPlantPartnerFunction    *string            `json:"IssuingPlantPartnerFunction"`
	IssuingPlantBusinessPartner    *string            `json:"IssuingPlantBusinessPartner"`
	IssuingPlant                   *string            `json:"IssuingPlant"`
	IssuingPlantStorageLocation    *string            `json:"IssuingPlantStorageLocation"`
	ReceivingPlantPartnerFunction  *string            `json:"ReceivingPlantPartnerFunction"`
	ReceivingPlantBusinessPartner  *string            `json:"ReceivingPlantBusinessPartner"`
	ReceivingPlant                 *string            `json:"ReceivingPlant"`
	ReceivingPlantStorageLocation  *string            `json:"ReceivingPlantStorageLocation"`
	ServicesRenderedDate           *string            `json:"ServicesRenderedDate"`
	InvoiceQuantity                *float32           `json:"InvoiceQuantity"`
	InvoiceQuantityUnit            *string            `json:"InvoiceQuantityUnit"`
	InvoiceQuantityInBaseUnit      *float32           `json:"InvoiceQuantityInBaseUnit"`
	BaseUnit                       *string            `json:"BaseUnit"`
	ActualGoodsIssueDate           *string            `json:"ActualGoodsIssueDate"`
	ActualGoodsIssueTime           *string            `json:"ActualGoodsIssueTime"`
	ActualGoodsReceiptDate         *string            `json:"ActualGoodsReceiptDate"`
	ActualGoodsReceiptTime         *string            `json:"ActualGoodsReceiptTime"`
	ItemGrossWeight                *float32           `json:"ItemGrossWeight"`
	ItemNetWeight                  *float32           `json:"ItemNetWeight"`
	ItemWeightUnit                 *string            `json:"ItemWeightUnit"`
	ItemVolume                     *float32           `json:"ItemVolume"`
	ItemVolumeUnit                 *string            `json:"ItemVolumeUnit"`
	NetAmount                      *float32           `json:"NetAmount"`
	TaxAmount                      *float32           `json:"TaxAmount"`
	GrossAmount                    *float32           `json:"GrossAmount"`
	GoodsIssueOrReceiptSlipNumber  *string            `json:"GoodsIssueOrReceiptSlipNumber"`
	TransactionCurrency            *string            `json:"TransactionCurrency"`
	PricingDate                    *string            `json:"PricingDate"`
	ProductTaxClassification       *string            `json:"ProductTaxClassification"`
	Project                        *string            `json:"Project"`
	OrderID                        *int               `json:"OrderID"`
	OrderItem                      *int               `json:"OrderItem"`
	OrderType                      *string            `json:"OrderType"`
	ContractType                   *string            `json:"ContractType"`
	OrderVaridityStartDate         *string            `json:"OrderVaridityStartDate"`
	OrderValidityEndDate           *string            `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate       *string            `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate         *string            `json:"InvoiceScheduleEndDate"`
	DeliveryDocument               *int               `json:"DeliveryDocument"`
	DeliveryDocumentItem           *int               `json:"DeliveryDocumentItem"`
	OriginDocument                 *int               `json:"OriginDocument"`
	OriginDocumentItem             *int               `json:"OriginDocumentItem"`
	ReferenceDocument              *int               `json:"ReferenceDocument"`
	ReferenceDocumentItem          *int               `json:"ReferenceDocumentItem"`
	ReferenceDocumentType          *string            `json:"ReferenceDocumentType"`
	ExternalReferenceDocument      *string            `json:"ExternalReferenceDocument"`
	ExternalReferenceDocumentItem  *string            `json:"ExternalReferenceDocumentItem"`
	OrderDistributionChannel       *string            `json:"OrderDistributionChannel"`
	TaxCode                        *string            `json:"TaxCode"`
	TaxRate                        *float32           `json:"TaxRate"`
	CountryOfOrigin                *string            `json:"CountryOfOrigin"`
	ItemPartner                    ItemPartner        `json:"ItemPartner"`
	ItemPricingElement             ItemPricingElement `json:"ItemPricingElement"`
}

type ItemPartner struct {
	PartnerFunction *string `json:"PartnerFunction"`
	BusinessPartner *int    `json:"BusinessPartner"`
}

type ItemPricingElement struct {
	PricingProcedureStep       *int     `json:"PricingProcedureStep"`
	PricingProcedureCounter    *int     `json:"PricingProcedureCounter"`
	ConditionType              *string  `json:"ConditionType"`
	PricingDate                *string  `json:"PricingDate"`
	ConditionRateValue         *float32 `json:"ConditionRateValue"`
	ConditionCurrency          *string  `json:"ConditionCurrency"`
	ConditionQuantity          *float32 `json:"ConditionQuantity"`
	ConditionQuantityUnit      *string  `json:"ConditionQuantityUnit"`
	ConditionRecord            *int     `json:"ConditionRecord"`
	ConditionSequentialNumber  *int     `json:"ConditionSequentialNumber"`
	TaxCode                    *string  `json:"TaxCode"`
	ConditionAmount            *float32 `json:"ConditionAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	ConditionIsManuallyChanged *bool    `json:"ConditionIsManuallyChanged"`
}

type Address struct {
	AddressID   *int    `json:"AddressID"`
	PostalCode  *string `json:"PostalCode"`
	LocalRegion *string `json:"LocalRegion"`
	Country     *string `json:"Country"`
	District    *string `json:"District"`
	StreetName  *string `json:"StreetName"`
	CityName    *string `json:"CityName"`
	Building    *string `json:"Building"`
	Floor       *int    `json:"Floor"`
	Room        *int    `json:"Room"`
}
