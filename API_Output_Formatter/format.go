package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Header, error) {
	calculateInvoiceDocument := psdc.CalculateInvoiceDocument
	headerOrdersHeader := psdc.HeaderOrdersHeader
	header := make([]Header, 0, len(*headerOrdersHeader))

	for _, v := range *headerOrdersHeader {
		header = append(header, Header{
			InvoiceDocument:            calculateInvoiceDocument.InvoiceDocumentLatestNumber,
			CreationDate:               sdc.InvoiceDocument.CreationDate,
			LastChangeDate:             sdc.InvoiceDocument.LastChangeDate,
			BillToParty:                v.BillToParty,
			BillFromParty:              v.BillFromParty,
			BillToCountry:              v.BillToCountry,
			BillFromCountry:            v.BillFromCountry,
			InvoiceDocumentDate:        sdc.InvoiceDocument.InvoiceDocumentDate,
			InvoiceDocumentTime:        sdc.InvoiceDocument.InvoiceDocumentTime,
			InvoicePeriodStartDate:     sdc.InvoiceDocument.InvoicePeriodStartDate,
			InvoicePeriodEndDate:       sdc.InvoiceDocument.InvoicePeriodEndDate,
			AccountingPostingDate:      sdc.InvoiceDocument.AccountingPostingDate,
			InvoiceDocumentIsCancelled: sdc.InvoiceDocument.InvoiceDocumentIsCancelled,
			CancelledInvoiceDocument:   sdc.InvoiceDocument.CancelledInvoiceDocument,
			IsExportImportDelivery:     sdc.InvoiceDocument.IsExportDelivery,
			// HeaderBillingIsConfirmed:   sdc.InvoiceDocument.HeaderBillingIsConfirmed,
			HeaderBillingConfStatus:   sdc.InvoiceDocument.HeaderBillingConfStatus,
			TotalNetAmount:            v.TotalNetAmount,
			TotalTaxAmount:            v.TotalTaxAmount,
			TotalGrossAmount:          v.TotalGrossAmount,
			TransactionCurrency:       v.TransactionCurrency,
			Incoterms:                 v.Incoterms,
			PaymentTerms:              v.PaymentTerms,
			DueCalculationBaseDate:    v.DueCalculationBaseDate,
			NetPaymentDays:            sdc.InvoiceDocument.NetPaymentDays,
			PaymentMethod:             v.PaymentMethod,
			HeaderPaymentBlockStatus:  sdc.InvoiceDocument.HeaderPaymentBlockStatus,
			ExternalReferenceDocument: sdc.InvoiceDocument.ExternalReferenceDocument,
			DocumentHeaderText:        sdc.InvoiceDocument.DocumentHeaderText,
		})
	}

	return &header, nil
}

func ConvertToHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartner, error) {
	calculateInvoiceDocument := psdc.CalculateInvoiceDocument
	headerOrdersHeaderPartner := psdc.HeaderOrdersHeaderPartner
	headerPartner := make([]HeaderPartner, 0, len(*headerOrdersHeaderPartner))

	for _, v := range *headerOrdersHeaderPartner {
		headerPartner = append(headerPartner, HeaderPartner{
			InvoiceDocument:         calculateInvoiceDocument.InvoiceDocumentLatestNumber,
			PartnerFunction:         v.PartnerFunction,
			BusinessPartner:         v.BusinessPartner,
			BusinessPartnerFullName: v.BusinessPartnerFullName,
			BusinessPartnerName:     v.BusinessPartnerName,
			Organization:            v.Organization,
			Country:                 v.Country,
			Language:                v.Language,
			Currency:                v.Currency,
			ExternalDocumentID:      v.ExternalDocumentID,
			AddressID:               v.AddressID,
		})
	}

	return &headerPartner, nil
}
