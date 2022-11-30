package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderOrdersHeader, error) {
	var args []interface{}

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, tag := range *orderID {
		args = append(args, tag.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderType, Buyer, Seller, ContractType, VaridityStartDate, VaridityEndDate, InvoiceScheduleStartDate, InvoiceScheduleEndDate, TransactionCurrency, Incoterms, IsExportImportDelivery
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderOrdersHeader(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateDeliveryDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateDeliveryDocument, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateDeliveryDocumentKey()
	if err != nil {
		return nil, err
	}

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateDeliveryDocumentQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateDeliveryDocument := CalculateDeliveryDocument(*dataQueryGets.DeliveryDocumentLatestNumber)

	data, err := psdc.ConvertToCalculateDeliveryDocument(calculateDeliveryDocument)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateDeliveryDocument(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}
