package subfunction

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"fmt"
)

func (f *SubFunction) SetValue(
	sdc *api_input_reader.SDC,
	osdc *dpfm_api_output_formatter.SDC,
	orderID *[]api_processing_data_formatter.OrderID,
	headerOrdersHeader *[]api_processing_data_formatter.HeaderOrdersHeader,
	headerOrdersHeaderPartner *[]api_processing_data_formatter.HeaderOrdersHeaderPartner,
	headerOrdersHeaderPartnerPlant *[]api_processing_data_formatter.HeaderOrdersHeaderPartnerPlant,
	calculateDeliveryDocument *api_processing_data_formatter.CalculateDeliveryDocument,
) (*dpfm_api_output_formatter.SDC, error) {
	var outHeader *[]dpfm_api_output_formatter.Header
	var err error

	outHeader, err = dpfm_api_output_formatter.ConvertToHeader(sdc, orderID, headerOrdersHeader, calculateDeliveryDocument)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	osdc.Message = dpfm_api_output_formatter.Message{
		Header: *outHeader,
	}

	return osdc, nil
}
