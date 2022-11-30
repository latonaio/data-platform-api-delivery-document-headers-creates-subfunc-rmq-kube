package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	"data-platform-api-delivery-document-headers-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

func getBoolPtr(b bool) *bool {
	return &b
}

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
		ReferenceDocument:               sdc.DeliveryDocument.ReferenceDocument,
		HeaderCompleteDeliveryIsDefined: getBoolPtr(false),
		OverallDeliveryStatus:           "CL",
	}
	data := pm

	orderIDKey := OrderIDKey{
		ReferenceDocument:               data.ReferenceDocument,
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
	pm := &requests.OrderID{
		ReferenceDocument:               nil,
		OrderID:                         nil,
		HeaderCompleteDeliveryIsDefined: nil,
		OverallDeliveryStatus:           "",
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
			&pm.OrderID,
			&pm.HeaderCompleteDeliveryIsDefined,
			&pm.OverallDeliveryStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		pm.ReferenceDocument = pm.OrderID

		data := pm
		orderID = append(orderID, OrderID{
			ReferenceDocument:               data.ReferenceDocument,
			OrderID:                         data.OrderID,
			HeaderCompleteDeliveryIsDefined: data.HeaderCompleteDeliveryIsDefined,
			OverallDeliveryStatus:           data.OverallDeliveryStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentKey() (*CalculateDeliveryDocumentKey, error) {
	pm := &requests.CalculateDeliveryDocumentKey{
		ServiceLabel:             "",
		FieldNameWithNumberRange: "DeliveryDocument",
	}
	data := pm

	calculateDeliveryDocumentKey := CalculateDeliveryDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateDeliveryDocumentKey, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentQueryGets(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*CalculateDeliveryDocumentQueryGets, error) {
	pm := &requests.CalculateDeliveryDocumentQueryGets{
		ServiceLabel:                 "",
		FieldNameWithNumberRange:     "",
		DeliveryDocumentLatestNumber: nil,
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
			&pm.DeliveryDocumentLatestNumber,
		)
		if err != nil {
			return nil, err
		}
	}
	data := pm

	calculateDeliveryDocumentQueryGets := CalculateDeliveryDocumentQueryGets{
		ServiceLabel:                 data.ServiceLabel,
		FieldNameWithNumberRange:     data.FieldNameWithNumberRange,
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
	}

	return &calculateDeliveryDocumentQueryGets, nil
}

func (psdc *SDC) ConvertToCalculateDeliveryDocument(
	deliveryDocumentLatestNumber *int,
) (*CalculateDeliveryDocument, error) {
	pm := &requests.CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: nil,
		DeliveryDocument:             nil,
	}

	pm.DeliveryDocumentLatestNumber = deliveryDocumentLatestNumber
	data := pm

	calculateDeliveryDocument := CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
		DeliveryDocument:             data.DeliveryDocument,
	}

	return &calculateDeliveryDocument, nil
}

// Header
func (psdc *SDC) ConvertToHeaderOrdersHeader(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeader, error) {
	var headerOrdersHeader []HeaderOrdersHeader
	pm := &requests.HeaderOrdersHeader{
		DeliveryDocument:         nil,
		OrderID:                  nil,
		OrderType:                "",
		Buyer:                    nil,
		Seller:                   nil,
		ContractType:             nil,
		OrderValidityStartDate:   nil,
		OrderValidityEndDate:     nil,
		InvoiceScheduleStartDate: nil,
		InvoiceScheduleEndDate:   nil,
		TransactionCurrency:      nil,
		Incoterms:                nil,
		IsExportImportDelivery:   nil,
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
			&pm.OrderID,
			&pm.OrderType,
			&pm.Buyer,
			&pm.Seller,
			&pm.ContractType,
			&pm.OrderValidityStartDate,
			&pm.OrderValidityEndDate,
			&pm.InvoiceScheduleStartDate,
			&pm.InvoiceScheduleEndDate,
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.IsExportImportDelivery,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeader = append(headerOrdersHeader, HeaderOrdersHeader{
			DeliveryDocument:         data.DeliveryDocument,
			OrderID:                  data.OrderID,
			OrderType:                data.OrderType,
			Buyer:                    data.Buyer,
			Seller:                   data.Seller,
			ContractType:             data.ContractType,
			OrderValidityStartDate:   data.OrderValidityStartDate,
			OrderValidityEndDate:     data.OrderValidityEndDate,
			InvoiceScheduleStartDate: data.InvoiceScheduleStartDate,
			InvoiceScheduleEndDate:   data.InvoiceScheduleEndDate,
			TransactionCurrency:      data.TransactionCurrency,
			Incoterms:                data.Incoterms,
			IsExportImportDelivery:   data.IsExportImportDelivery,
		})
	}

	return &headerOrdersHeader, nil
}

func (psdc *SDC) ConvertToHeaderOrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeaderPartner, error) {
	var headerOrdersHeaderPartner []HeaderOrdersHeaderPartner

	pm := &requests.HeaderOrdersHeaderPartner{
		DeliveryDocument:        nil,
		OrderID:                 nil,
		PartnerFunction:         "",
		BusinessPartner:         nil,
		BusinessPartnerFullName: nil,
		BusinessPartnerName:     nil,
		Organization:            nil,
		Country:                 nil,
		Language:                nil,
		Currency:                nil,
		ExternalDocumentID:      nil,
		AddressID:               nil,
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
			DeliveryDocument:        data.DeliveryDocument,
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

func (psdc *SDC) ConvertToHeaderOrdersHeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]HeaderOrdersHeaderPartnerPlant, error) {
	var headerOrdersHeaderPartnerPlant []HeaderOrdersHeaderPartnerPlant

	pm := &requests.HeaderOrdersHeaderPartnerPlant{
		DeliveryDocument: nil,
		OrderID:          nil,
		PartnerFunction:  "",
		BusinessPartner:  nil,
		Plant:            "",
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
			&pm.OrderID,
			&pm.PartnerFunction,
			&pm.BusinessPartner,
			&pm.Plant,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		headerOrdersHeaderPartnerPlant = append(headerOrdersHeaderPartnerPlant, HeaderOrdersHeaderPartnerPlant{
			DeliveryDocument: data.DeliveryDocument,
			OrderID:          data.OrderID,
			PartnerFunction:  data.PartnerFunction,
			BusinessPartner:  data.BusinessPartner,
			Plant:            data.Plant,
		})
	}

	return &headerOrdersHeaderPartnerPlant, nil
}
