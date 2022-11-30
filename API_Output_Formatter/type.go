package dpfm_api_output_formatter

type SDC struct {
	ConnectionKey       string   `json:"connection_key"`
	Result              bool     `json:"result"`
	RedisKey            string   `json:"redis_key"`
	Filepath            string   `json:"filepath"`
	APIStatusCode       int      `json:"api_status_code"`
	RuntimeSessionID    string   `json:"runtime_session_id"`
	BusinessPartnerID   *int     `json:"business_partner"`
	ServiceLabel        string   `json:"service_label"`
	Message             Message  `json:"message"`
	APISchema           string   `json:"api_schema"`
	Accepter            []string `json:"accepter"`
	Deleted             bool     `json:"deleted"`
	SQLUpdateResult     *bool    `json:"sql_update_result"`
	SQLUpdateError      string   `json:"sql_update_error"`
	SubfuncResult       *bool    `json:"subfunc_result"`
	SubfuncError        string   `json:"subfunc_error"`
	ExconfResult        *bool    `json:"exconf_result"`
	ExconfError         string   `json:"exconf_error"`
	APIProcessingResult *bool    `json:"api_processing_result"`
	APIProcessingError  string   `json:"api_processing_error"`
}

type Message struct {
	Header        []Header        `json:"Header"`
	HeaderPartner []HeaderPartner `json:"HeaderPartner"`
}

type Header struct {
	InvoiceDocument            *int     `json:"InvoiceDocument"`
	CreationDate               *string  `json:"CreationDate"`
	LastChangeDate             *string  `json:"LastChangeDate"`
	BillToParty                *int     `json:"BillToParty"`
	BillFromParty              *int     `json:"BillFromParty"`
	BillToCountry              *string  `json:"BillToCountry"`
	BillFromCountry            *string  `json:"BillFromCountry"`
	InvoiceDocumentDate        *string  `json:"InvoiceDocumentDate"`
	InvoiceDocumentTime        *string  `json:"InvoiceDocumentTime"`
	InvoicePeriodStartDate     *string  `json:"InvoicePeriodStartDate"`
	InvoicePeriodEndDate       *string  `json:"InvoicePeriodEndDate"`
	AccountingPostingDate      *string  `json:"AccountingPostingDate"`
	InvoiceDocumentIsCancelled *bool    `json:"InvoiceDocumentIsCancelled"`
	CancelledInvoiceDocument   *int     `json:"CancelledInvoiceDocument"`
	IsExportImportDelivery     *bool    `json:"IsExportImportDelivery"`
	HeaderBillingIsConfirmed   *bool    `json:"HeaderBillingIsConfirmed"`
	HeaderBillingConfStatus    *string  `json:"HeaderBillingConfStatus"`
	TotalNetAmount             *float32 `json:"TotalNetAmount"`
	TotalTaxAmount             *float32 `json:"TotalTaxAmount"`
	TotalGrossAmount           *float32 `json:"TotalGrossAmount"`
	TransactionCurrency        *string  `json:"TransactionCurrency"`
	Incoterms                  *string  `json:"Incoterms"`
	PaymentTerms               *string  `json:"PaymentTerms"`
	DueCalculationBaseDate     *string  `json:"DueCalculationBaseDate"`
	NetPaymentDays             *int     `json:"NetPaymentDays"`
	PaymentMethod              *string  `json:"PaymentMethod"`
	HeaderPaymentBlockStatus   *bool    `json:"HeaderPaymentBlockStatus"`
	ExternalReferenceDocument  *string  `json:"ExternalReferenceDocument"`
	DocumentHeaderText         *string  `json:"DocumentHeaderText"`
}

type HeaderPartner struct {
	InvoiceDocument         *int    `json:"InvoiceDocument`
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
