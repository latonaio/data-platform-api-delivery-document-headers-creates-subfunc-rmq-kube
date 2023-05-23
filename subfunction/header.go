package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"strings"

	"golang.org/x/xerrors"
)

func (f *SubFunction) CalculateDeliveryDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateDeliveryDocument, error) {
	metaData := psdc.MetaData
	dataKey := psdc.ConvertToCalculateDeliveryDocumentKey()

	dataKey.ServiceLabel = metaData.ServiceLabel

	rows, err := f.db.Query(
		`SELECT ServiceLabel, FieldNameWithNumberRange, LatestNumber
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_number_range_latest_number_data
		WHERE (ServiceLabel, FieldNameWithNumberRange) = (?, ?);`, dataKey.ServiceLabel, dataKey.FieldNameWithNumberRange,
	)
	if err != nil {
		return nil, err
	}

	dataQueryGets, err := psdc.ConvertToCalculateDeliveryDocumentQueryGets(rows)
	if err != nil {
		return nil, err
	}

	if dataQueryGets.DeliveryDocumentLatestNumber == nil {
		return nil, xerrors.Errorf("'data_platform_number_range_latest_number_data'テーブルのLatestNumberがNULLです。")
	}

	deliveryDocumentLatestNumber := dataQueryGets.DeliveryDocumentLatestNumber
	deliveryDocument := *dataQueryGets.DeliveryDocumentLatestNumber + 1

	data := psdc.ConvertToCalculateDeliveryDocument(deliveryDocumentLatestNumber, deliveryDocument)

	return data, err
}

func (f *SubFunction) OrdersHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrdersHeader, error) {
	var args []interface{}

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, v := range *orderID {
		args = append(args, v.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderType, Buyer, Seller, ContractType, ValidityStartDate, ValidityEndDate, 
		TransactionCurrency, Incoterms, BillFromParty, BillToParty, BillFromCountry, BillToCountry, 
		Payer, Payee, IsExportImportDelivery
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrdersHeader(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
