package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"strings"
)

func (f *SubFunction) HeaderPartnerPlant(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.HeaderPartnerPlant, error) {
	var args []interface{}

	orderID := psdc.OrderID
	repeat := strings.Repeat("?,", len(*orderID)-1) + "?"
	for _, v := range *orderID {
		args = append(args, v.OrderID)
	}

	rows, err := f.db.Query(
		`SELECT OrderID, PartnerFunction, BusinessPartner, Plant
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_partner_plant_data
		WHERE OrderID IN ( `+repeat+` );`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToHeaderPartnerPlant(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}
