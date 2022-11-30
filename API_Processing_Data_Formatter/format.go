package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	"data-platform-api-invoice-document-headers-creates-subfunc-rmq/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) (*MetaData, error) {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData, nil
}

func (psdc *SDC) ConvertToOrderIDKey(sdc *api_input_reader.SDC) (*OrderIDKey, error) {
	pm := &requests.OrderIDKey{
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:           "CL",
	}
	data := pm

	orderIDKey := OrderIDKey{
		HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
		OverallDeliveryStatus:           data.OverallDeliveryStatus,
	}

	return &orderIDKey, nil
}

func (psdc *SDC) ConvertToOrderID(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderID, error) {
	var orderID []OrderID
	pm := &requests.OrderID{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderID = append(orderID, OrderID{
			OrderID:                         data.OrderID,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			OverallDeliveryStatus:           data.OverallDeliveryStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToDeliveryDocumentKey(sdc *api_input_reader.SDC) (*DeliveryDocumentKey, error) {
	pm := &requests.DeliveryDocumentKey{
		CompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:     "CL",
	}
	data := pm

	deliveryDocumentKey := DeliveryDocumentKey{
		CompleteDeliveryIsDefined: data.CompleteDeliveryIsDefined,
		OverallDeliveryStatus:     data.OverallDeliveryStatus,
	}

	return &deliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToDeliveryDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]DeliveryDocument, error) {
	var deliveryDocument []DeliveryDocument
	pm := &requests.DeliveryDocument{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.CompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		deliveryDocument = append(deliveryDocument, DeliveryDocument{
			DeliveryDocument:          data.DeliveryDocument,
			CompleteDeliveryIsDefined: data.CompleteDeliveryIsDefined,
			OverallDeliveryStatus:     data.OverallDeliveryStatus,
		})
	}

	return &deliveryDocument, nil
}

func getBoolPtr(b bool) *bool {
	return &b
}

// Header
func (psdc *SDC) ConvertToHeaderOrdersHeader(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeader, error) {
	var headerOrdersHeader []HeaderOrdersHeader
	pm := &requests.HeaderOrdersHeader{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.TotalNetAmount,
			&pm.TransactionCurrency,
			&pm.TotalTaxAmount,
			&pm.TotalGrossAmount,
			&pm.Incoterms,
			&pm.PaymentTerms,
			&pm.PaymentMethod,
			&pm.BillToCountry,
			&pm.BillFromCountry,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeader = append(headerOrdersHeader, HeaderOrdersHeader{
			InvoiceDocument:         data.InvoiceDocument,
			OrderID:                 data.OrderID,
			InvoiceDocumentType:     data.InvoiceDocumentType,
			BillToParty:             data.BillToParty,
			BillFromParty:           data.BillFromParty,
			BillToPartyLanguage:     data.BillToPartyLanguage,
			BillFromPartyLanguage:   data.BillFromPartyLanguage,
			TotalNetAmount:          data.TotalNetAmount,
			TransactionCurrency:     data.TransactionCurrency,
			BusinessPartnerCurrency: data.BusinessPartnerCurrency,
			TotalTaxAmount:          data.TotalTaxAmount,
			TotalGrossAmount:        data.TotalGrossAmount,
			Incoterms:               data.Incoterms,
			PaymentTerms:            data.PaymentTerms,
			DueCalculationBaseDate:  data.DueCalculationBaseDate,
			PaymentMethod:           data.PaymentMethod,
			BillToAddressID:         data.BillToAddressID,
			BillFromAddressID:       data.BillFromAddressID,
			BillToCountry:           data.BillToCountry,
			BillToLocalRegion:       data.BillToLocalRegion,
			BillFromCountry:         data.BillFromCountry,
			BillFromLocalRegion:     data.BillFromLocalRegion,
		})
	}

	return &headerOrdersHeader, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentKey() (*CalculateInvoiceDocumentKey, error) {
	pm := &requests.CalculateInvoiceDocumentKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "InvoiceDocument",
	}
	data := pm

	calculateInvoiceDocumentKey := CalculateInvoiceDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateInvoiceDocumentKey, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocumentQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateInvoiceDocumentQueryGets, error) {
	pm := &requests.CalculateInvoiceDocumentQueryGets{
		ServiceLabel:                "",
		FieldNameWithNumberRange:    "",
		InvoiceDocumentLatestNumber: nil,
	}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.ServiceLabel,
			&pm.FieldNameWithNumberRange,
			&pm.InvoiceDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	calculateInvoiceDocumentQueryGets := CalculateInvoiceDocumentQueryGets{
		ServiceLabel:                data.ServiceLabel,
		FieldNameWithNumberRange:    data.FieldNameWithNumberRange,
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
	}

	return &calculateInvoiceDocumentQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateInvoiceDocument(
	invoiceDocumentLatestNumber *int,
) (*CalculateInvoiceDocument, error) {
	pm := &requests.CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: nil,
		InvoiceDocument:             nil,
	}

	pm.InvoiceDocumentLatestNumber = invoiceDocumentLatestNumber
	data := pm

	calculateInvoiceDocument := CalculateInvoiceDocument{
		InvoiceDocumentLatestNumber: data.InvoiceDocumentLatestNumber,
		InvoiceDocument:             data.InvoiceDocument,
	}

	return &calculateInvoiceDocument, nil
}

func (psdc *SDC) ConvertToHeaderOrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeaderPartner, error) {
	var headerOrdersHeaderPartner []HeaderOrdersHeaderPartner
	pm := &requests.HeaderOrdersHeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeaderPartner = append(headerOrdersHeaderPartner, HeaderOrdersHeaderPartner{
			InvoiceDocument:         data.InvoiceDocument,
			OrderID:                 data.OrderID,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &headerOrdersHeaderPartner, nil
}

func (psdc *SDC) ConvertToHeaderDeliveryDocumentHeader(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderDeliveryDocumentHeader, error) {
	var headerdeliveryDocumentHeader []HeaderDeliveryDocumentHeader
	pm := &requests.HeaderDeliveryDocumentHeader{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.Buyer,
			&pm.Seller,
			&pm.ReferenceDocument,
			&pm.ReferenceDocumentItem,
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			// &pm.IssuingPlantTimeZone,
			// &pm.ReceivingPlantTimeZone,
			&pm.DocumentDate,
			&pm.PlannedGoodsIssueDate,
			&pm.PlannedGoodsIssueTime,
			&pm.PlannedGoodsReceiptDate,
			&pm.PlannedGoodsReceiptTime,
			&pm.BillingDocumentDate,
			&pm.CompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.IssuingBlockReason,
			&pm.ReceivingBlockReason,
			&pm.GoodsIssueOrReceiptSlipNumber,
			&pm.HeaderBillingStatus,
			&pm.HeaderBillingConfStatus,
			&pm.HeaderBillingBlockReason,
			&pm.HeaderGrossWeight,
			&pm.HeaderNetWeight,
			&pm.HeaderWeightUnit,
			&pm.Incoterms,
			&pm.BillToCountry,
			&pm.BillFromCountry,
			&pm.IsExportImportDelivery,
			&pm.LastChangeDate,
			&pm.IssuingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlantBusinessPartner,
			&pm.ReceivingPlant,
			&pm.DeliverToParty,
			&pm.DeliverFromParty,
			&pm.TransactionCurrency,
			&pm.OverallDelivReltdBillgStatus,
			&pm.StockIsFullyConfirmed,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerdeliveryDocumentHeader = append(headerdeliveryDocumentHeader, HeaderDeliveryDocumentHeader{
			InvoiceDocument:          data.InvoiceDocument,
			DeliveryDocument:         data.DeliveryDocument,
			Buyer:                    data.Buyer,
			Seller:                   data.Seller,
			ReferenceDocument:        data.ReferenceDocument,
			ReferenceDocumentItem:    data.ReferenceDocumentItem,
			OrderID:                  data.OrderID,
			OrderItem:                data.OrderItem,
			ContractType:             data.ContractType,
			OrderValidityStartDate:   data.OrderValidityStartDate,
			OrderValidityEndDate:     data.OrderValidityEndDate,
			InvoiceScheduleStartDate: data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:   data.InvoiceScheduleEndDate,
			// IssuingPlantTimeZone:          data.IssuingPlantTimeZone,
			// ReceivingPlantTimeZone:        data.ReceivingPlantTimeZone,
			DocumentDate:                  data.DocumentDate,
			PlannedGoodsIssueDate:         data.PlannedGoodsIssueDate,
			PlannedGoodsIssueTime:         data.PlannedGoodsIssueTime,
			PlannedGoodsReceiptDate:       data.PlannedGoodsReceiptDate,
			PlannedGoodsReceiptTime:       data.PlannedGoodsReceiptTime,
			BillingDocumentDate:           data.BillingDocumentDate,
			CompleteDeliveryIsDefined:     data.CompleteDeliveryIsDefined,
			OverallDeliveryStatus:         data.OverallDeliveryStatus,
			CreationDate:                  data.CreationDate,
			CreationTime:                  data.CreationTime,
			IssuingBlockReason:            data.IssuingBlockReason,
			ReceivingBlockReason:          data.ReceivingBlockReason,
			GoodsIssueOrReceiptSlipNumber: data.GoodsIssueOrReceiptSlipNumber,
			HeaderBillingStatus:           data.HeaderBillingStatus,
			HeaderBillingConfStatus:       data.HeaderBillingConfStatus,
			HeaderBillingBlockReason:      data.HeaderBillingBlockReason,
			HeaderGrossWeight:             data.HeaderGrossWeight,
			HeaderNetWeight:               data.HeaderNetWeight,
			HeaderWeightUnit:              data.HeaderWeightUnit,
			Incoterms:                     data.Incoterms,
			BillToCountry:                 data.BillToCountry,
			BillFromCountry:               data.BillFromCountry,
			IsExportImportDelivery:        data.IsExportImportDelivery,
			LastChangeDate:                data.LastChangeDate,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			ReceivingPlant:                data.ReceivingPlant,
			DeliverToParty:                data.DeliverToParty,
			DeliverFromParty:              data.DeliverFromParty,
			TransactionCurrency:           data.TransactionCurrency,
			OverallDelivReltdBillgStatus:  data.OverallDelivReltdBillgStatus,
			StockIsFullyConfirmed:         data.StockIsFullyConfirmed,
		})
	}

	return &headerdeliveryDocumentHeader, nil
}

func (psdc *SDC) ConvertToHeaderDeliveryDocumentHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderDeliveryDocumentHeaderPartner, error) {
	var headerDeliveryDocumentHeaderPartner []HeaderDeliveryDocumentHeaderPartner
	pm := &requests.HeaderDeliveryDocumentHeaderPartner{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.DeliveryDocument,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.BusinessPartnerFullName,
			&pm.BusinessPartnerName,
			&pm.Organization,
			&pm.Country,
			&pm.Language,
			&pm.Currency,
			&pm.ExternalDocumentID,
			&pm.AddressID,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerDeliveryDocumentHeaderPartner = append(headerDeliveryDocumentHeaderPartner, HeaderDeliveryDocumentHeaderPartner{
			InvoiceDocument:         data.InvoiceDocument,
			DeliveryDocument:        data.DeliveryDocument,
			PartnerFunction:         data.PartnerFunction,
			BusinessPartner:         data.BusinessPartner,
			BusinessPartnerFullName: data.BusinessPartnerFullName,
			BusinessPartnerName:     data.BusinessPartnerName,
			Organization:            data.Organization,
			Country:                 data.Country,
			Language:                data.Language,
			Currency:                data.Currency,
			ExternalDocumentID:      data.ExternalDocumentID,
			AddressID:               data.AddressID,
		})
	}

	return &headerDeliveryDocumentHeaderPartner, nil
}
