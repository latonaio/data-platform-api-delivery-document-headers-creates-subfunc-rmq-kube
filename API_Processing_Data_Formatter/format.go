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

func (psdc *SDC) ConvertToOrderItemByNumberSpecificationKey(sdc *api_input_reader.SDC, businessPartnerLength int, plantLength int) (*OrderItemKey, error) {
	pm := &requests.OrderItemKey{
		IssuingPlantPartnerFunction:   "DELIVERFROM",
		ReceivingPlantPartnerFunction: "DELIVERTO",
		ItemCompleteDeliveryIsDefined: getBoolPtr(false),
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       getBoolPtr(false),
	}

	for i := 0; i < businessPartnerLength; i++ {
		pm.IssuingPlantBusinessPartner = append(pm.IssuingPlantBusinessPartner, nil)
		pm.ReceivingPlantBusinessPartner = append(pm.ReceivingPlantBusinessPartner, nil)
	}

	for i := 0; i < plantLength; i++ {
		pm.IssuingPlant = append(pm.IssuingPlant, nil)
		pm.ReceivingPlant = append(pm.ReceivingPlant, nil)
	}

	data := pm
	orderItemKey := OrderItemKey{
		OrderID:                           data.OrderID,
		OrderItem:                         data.OrderItem,
		IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
		IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
		IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
		ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
		ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
		ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
		IssuingPlant:                      data.IssuingPlant,
		IssuingPlantFrom:                  data.IssuingPlantFrom,
		IssuingPlantTo:                    data.IssuingPlantTo,
		ReceivingPlant:                    data.ReceivingPlant,
		ReceivingPlantFrom:                data.ReceivingPlantFrom,
		ReceivingPlantTo:                  data.ReceivingPlantTo,
		ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
	}

	return &orderItemKey, nil
}

func (psdc *SDC) ConvertToOrderItemByNumberSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderItem, error) {
	var orderItem []OrderItem

	for i := 0; true; i++ {
		pm := &requests.OrderItem{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.IssuingPlantBusinessPartner,
			&pm.ReceivingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlant,
			&pm.IssuingPlantPartnerFunction,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderItem = append(orderItem, OrderItem{
			OrderID:                           data.OrderID,
			OrderItem:                         data.OrderItem,
			IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
			IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
			IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
			ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
			ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
			ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
			IssuingPlant:                      data.IssuingPlant,
			IssuingPlantFrom:                  data.IssuingPlantFrom,
			IssuingPlantTo:                    data.IssuingPlantTo,
			ReceivingPlant:                    data.ReceivingPlant,
			ReceivingPlantFrom:                data.ReceivingPlantFrom,
			ReceivingPlantTo:                  data.ReceivingPlantTo,
			ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
		})
	}

	return &orderItem, nil
}

func (psdc *SDC) ConvertToOrderItemByRangeSpecificationKey(sdc *api_input_reader.SDC) (*OrderItemKey, error) {
	pm := &requests.OrderItemKey{
		IssuingPlantPartnerFunction:   "DELIVERFROM",
		ReceivingPlantPartnerFunction: "DELIVERTO",
		ItemCompleteDeliveryIsDefined: getBoolPtr(false),
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       getBoolPtr(false),
	}

	data := pm
	orderItemKey := OrderItemKey{
		OrderID:                           data.OrderID,
		OrderItem:                         data.OrderItem,
		IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
		IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
		IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
		ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
		ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
		ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
		IssuingPlant:                      data.IssuingPlant,
		IssuingPlantFrom:                  data.IssuingPlantFrom,
		IssuingPlantTo:                    data.IssuingPlantTo,
		ReceivingPlant:                    data.ReceivingPlant,
		ReceivingPlantFrom:                data.ReceivingPlantFrom,
		ReceivingPlantTo:                  data.ReceivingPlantTo,
		ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
	}

	return &orderItemKey, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpecification(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderItem, error) {
	var orderItem []OrderItem

	for i := 0; true; i++ {
		pm := &requests.OrderItem{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.IssuingPlantBusinessPartner,
			&pm.ReceivingPlantBusinessPartner,
			&pm.IssuingPlant,
			&pm.ReceivingPlant,
			&pm.IssuingPlantPartnerFunction,
			&pm.ReceivingPlantPartnerFunction,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderItem = append(orderItem, OrderItem{
			OrderID:                           data.OrderID,
			OrderItem:                         data.OrderItem,
			IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
			IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
			IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
			ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
			ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
			ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
			IssuingPlant:                      data.IssuingPlant,
			IssuingPlantFrom:                  data.IssuingPlantFrom,
			IssuingPlantTo:                    data.IssuingPlantTo,
			ReceivingPlant:                    data.ReceivingPlant,
			ReceivingPlantFrom:                data.ReceivingPlantFrom,
			ReceivingPlantTo:                  data.ReceivingPlantTo,
			ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
		})
	}

	return &orderItem, nil
}

func (psdc *SDC) ConvertToOrderItemByReferenceDocumentKey(sdc *api_input_reader.SDC) (*OrderItemKey, error) {
	pm := &requests.OrderItemKey{
		ItemCompleteDeliveryIsDefined: getBoolPtr(false),
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       getBoolPtr(false),
	}

	data := pm
	orderItemKey := OrderItemKey{
		OrderID:                           data.OrderID,
		OrderItem:                         data.OrderItem,
		IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
		IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
		IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
		IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
		ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
		ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
		ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
		ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
		IssuingPlant:                      data.IssuingPlant,
		IssuingPlantFrom:                  data.IssuingPlantFrom,
		IssuingPlantTo:                    data.IssuingPlantTo,
		ReceivingPlant:                    data.ReceivingPlant,
		ReceivingPlantFrom:                data.ReceivingPlantFrom,
		ReceivingPlantTo:                  data.ReceivingPlantTo,
		ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
		ItemDeliveryStatus:                data.ItemDeliveryStatus,
		ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
	}

	return &orderItemKey, nil
}

func (psdc *SDC) ConvertToOrderItemByReferenceDocument(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrderItem, error) {
	var orderItem []OrderItem

	for i := 0; true; i++ {
		pm := &requests.OrderItem{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("DBに対象のレコードが存在しません。")
			} else {
				break
			}
		}
		err := rows.Scan(
			&pm.OrderID,
			&pm.OrderItem,
			&pm.ItemCompleteDeliveryIsDefined,
			&pm.ItemDeliveryStatus,
			&pm.ItemDeliveryBlockStatus,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return nil, err
		}

		data := pm
		orderItem = append(orderItem, OrderItem{
			OrderID:                           data.OrderID,
			OrderItem:                         data.OrderItem,
			IssuingPlantPartnerFunction:       data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:       data.IssuingPlantBusinessPartner,
			IssuingPlantBusinessPartnerFrom:   data.IssuingPlantBusinessPartnerFrom,
			IssuingPlantBusinessPartnerTo:     data.IssuingPlantBusinessPartnerTo,
			ReceivingPlantPartnerFunction:     data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner:     data.ReceivingPlantBusinessPartner,
			ReceivingPlantBusinessPartnerFrom: data.ReceivingPlantBusinessPartnerFrom,
			ReceivingPlantBusinessPartnerTo:   data.ReceivingPlantBusinessPartnerTo,
			IssuingPlant:                      data.IssuingPlant,
			IssuingPlantFrom:                  data.IssuingPlantFrom,
			IssuingPlantTo:                    data.IssuingPlantTo,
			ReceivingPlant:                    data.ReceivingPlant,
			ReceivingPlantFrom:                data.ReceivingPlantFrom,
			ReceivingPlantTo:                  data.ReceivingPlantTo,
			ItemCompleteDeliveryIsDefined:     data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:                data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:           data.ItemDeliveryBlockStatus,
		})
	}

	return &orderItem, nil
}

func (psdc *SDC) ConvertToOrdersHeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	rows *sql.Rows,
) (*[]OrdersHeaderPartnerPlant, error) {
	var ordersHeaderPartnerPlant []OrdersHeaderPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.OrdersHeaderPartnerPlant{}
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
		ordersHeaderPartnerPlant = append(ordersHeaderPartnerPlant, OrdersHeaderPartnerPlant{
			DeliveryDocument: data.DeliveryDocument,
			OrderID:          data.OrderID,
			PartnerFunction:  data.PartnerFunction,
			BusinessPartner:  data.BusinessPartner,
			Plant:            data.Plant,
		})
	}

	return &ordersHeaderPartnerPlant, nil
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
