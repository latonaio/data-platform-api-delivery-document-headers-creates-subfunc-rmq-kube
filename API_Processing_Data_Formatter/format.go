package api_processing_data_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	"data-platform-api-delivery-document-headers-creates-subfunc/DPFM_API_Caller/requests"
	"database/sql"
	"fmt"
)

// Initializer
func (psdc *SDC) ConvertToMetaData(sdc *api_input_reader.SDC) *MetaData {
	pm := &requests.MetaData{
		BusinessPartnerID: sdc.BusinessPartnerID,
		ServiceLabel:      sdc.ServiceLabel,
	}
	data := pm

	metaData := MetaData{
		BusinessPartnerID: data.BusinessPartnerID,
		ServiceLabel:      data.ServiceLabel,
	}

	return &metaData
}

func (psdc *SDC) ConvertToProcessType() *ProcessType {
	pm := &requests.ProcessType{}
	data := pm

	processType := ProcessType{
		BulkProcess:       data.BulkProcess,
		IndividualProcess: data.IndividualProcess,
	}

	return &processType
}

func (psdc *SDC) ConvertToOrderIDKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		IssuingPlantPartnerFunction:   "DELIVERFROM",
		ReceivingPlantPartnerFunction: "DELIVERTO",
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	orderIDKey := OrderIDKey{
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

	return &orderIDKey
}

func (psdc *SDC) ConvertToOrderIDByArraySpec(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
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
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDByRangeSpec(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
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
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcessKey() *OrderIDKey {
	pm := &requests.OrderIDKey{
		ItemCompleteDeliveryIsDefined: false,
		ItemDeliveryStatus:            "CL",
		ItemDeliveryBlockStatus:       false,
	}

	data := pm
	orderIDKey := OrderIDKey{
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

	return &orderIDKey
}

func (psdc *SDC) ConvertToOrderIDInIndividualProcess(rows *sql.Rows) (*[]OrderID, error) {
	var orderID []OrderID

	for i := 0; true; i++ {
		pm := &requests.OrderID{}

		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_item_data'テーブルに対象のレコードが存在しません。")
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
		orderID = append(orderID, OrderID{
			OrderID:                       data.OrderID,
			OrderItem:                     data.OrderItem,
			IssuingPlantPartnerFunction:   data.IssuingPlantPartnerFunction,
			IssuingPlantBusinessPartner:   data.IssuingPlantBusinessPartner,
			ReceivingPlantPartnerFunction: data.ReceivingPlantPartnerFunction,
			ReceivingPlantBusinessPartner: data.ReceivingPlantBusinessPartner,
			IssuingPlant:                  data.IssuingPlant,
			ReceivingPlant:                data.ReceivingPlant,
			ItemCompleteDeliveryIsDefined: data.ItemCompleteDeliveryIsDefined,
			ItemDeliveryStatus:            data.ItemDeliveryStatus,
			ItemDeliveryBlockStatus:       data.ItemDeliveryBlockStatus,
		})
	}

	return &orderID, nil
}

func (psdc *SDC) ConvertToHeaderPartnerPlant(rows *sql.Rows) (*[]HeaderPartnerPlant, error) {
	var headerPartnerPlant []HeaderPartnerPlant

	for i := 0; true; i++ {
		pm := &requests.HeaderPartnerPlant{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_partner_plant_data'テーブルに対象のレコードが存在しません。")
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
		headerPartnerPlant = append(headerPartnerPlant, HeaderPartnerPlant{
			OrderID:         data.OrderID,
			PartnerFunction: data.PartnerFunction,
			BusinessPartner: data.BusinessPartner,
			Plant:           data.Plant,
		})
	}

	return &headerPartnerPlant, nil
}

// Header
func (psdc *SDC) ConvertToCalculateDeliveryDocumentKey() *CalculateDeliveryDocumentKey {
	pm := &requests.CalculateDeliveryDocumentKey{
		FieldNameWithNumberRange: "DeliveryDocument",
	}

	data := pm
	calculateDeliveryDocumentKey := CalculateDeliveryDocumentKey{
		ServiceLabel:             data.ServiceLabel,
		FieldNameWithNumberRange: data.FieldNameWithNumberRange,
	}

	return &calculateDeliveryDocumentKey
}

func (psdc *SDC) ConvertToCalculateDeliveryDocumentQueryGets(rows *sql.Rows) (*CalculateDeliveryDocumentQueryGets, error) {
	pm := &requests.CalculateDeliveryDocumentQueryGets{}

	for i := 0; true; i++ {
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_number_range_latest_number_data'テーブルに対象のレコードが存在しません。")
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

func (psdc *SDC) ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber *int, deliveryDocument int) *CalculateDeliveryDocument {
	pm := &requests.CalculateDeliveryDocument{}

	pm.DeliveryDocumentLatestNumber = deliveryDocumentLatestNumber
	pm.DeliveryDocument = deliveryDocument

	data := pm
	calculateDeliveryDocument := CalculateDeliveryDocument{
		DeliveryDocumentLatestNumber: data.DeliveryDocumentLatestNumber,
		DeliveryDocument:             data.DeliveryDocument,
	}

	return &calculateDeliveryDocument
}

func (psdc *SDC) ConvertToOrdersHeader(rows *sql.Rows) (*[]OrdersHeader, error) {
	var ordersHeader []OrdersHeader

	for i := 0; true; i++ {
		pm := &requests.OrdersHeader{}
		if !rows.Next() {
			if i == 0 {
				return nil, fmt.Errorf("'data_platform_orders_header_data'テーブルに対象のレコードが存在しません。")
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
			&pm.TransactionCurrency,
			&pm.Incoterms,
			&pm.BillFromParty,
			&pm.BillToParty,
			&pm.BillFromCountry,
			&pm.BillToCountry,
			&pm.Payer,
			&pm.Payee,
			&pm.IsExportImportDelivery,
		)
		if err != nil {
			return nil, err
		}

		data := pm
		ordersHeader = append(ordersHeader, OrdersHeader{
			OrderID:                data.OrderID,
			OrderType:              data.OrderType,
			Buyer:                  data.Buyer,
			Seller:                 data.Seller,
			ContractType:           data.ContractType,
			OrderValidityStartDate: data.OrderValidityStartDate,
			OrderValidityEndDate:   data.OrderValidityEndDate,
			TransactionCurrency:    data.TransactionCurrency,
			Incoterms:              data.Incoterms,
			BillFromParty:          data.BillFromParty,
			BillToParty:            data.BillToParty,
			BillFromCountry:        data.BillFromCountry,
			BillToCountry:          data.BillToCountry,
			Payer:                  data.Payer,
			Payee:                  data.Payee,
			IsExportImportDelivery: data.IsExportImportDelivery,
		})
	}

	return &ordersHeader, nil
}

func (psdc *SDC) ConvertToOrdersHeaderPartner(rows *sql.Rows) (*[]OrdersHeaderPartner, error) {
	var ordersHeaderPartner []OrdersHeaderPartner

	pm := &requests.OrdersHeaderPartner{}

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
		ordersHeaderPartner = append(ordersHeaderPartner, OrdersHeaderPartner{
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

	return &ordersHeaderPartner, nil
}
