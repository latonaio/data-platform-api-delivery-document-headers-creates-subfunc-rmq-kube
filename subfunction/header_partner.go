package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) OrdersHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderOrdersHeaderPartner, error) {
	var args []interface{}

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, tag := range *orderID {
		args = append(args, tag.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, BusinessPartnerFullName, BusinessPartnerName, Organization, Country, Language, Currency, ExternalDocumentID, AddressID
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderOrdersHeaderPartner(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
