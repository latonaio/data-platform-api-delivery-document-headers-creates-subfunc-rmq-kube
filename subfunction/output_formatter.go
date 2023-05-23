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
	psdc *api_processing_data_formatter.SDC,
) (*dpfm_api_output_formatter.SDC, error) {
	var header *[]dpfm_api_output_formatter.Header
	var headerPartner *[]dpfm_api_output_formatter.HeaderPartner
	var headerPartnerPlant *[]dpfm_api_output_formatter.HeaderPartnerPlant
	var err error

	header, err = dpfm_api_output_formatter.ConvertToHeader(sdc, psdc)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	headerPartner, err = dpfm_api_output_formatter.ConvertToHeaderPartner(sdc, psdc)
	if err != nil {
		fmt.Printf("err = %+v \n", err)
		return nil, err
	}

	// headerPartnerPlant, err = dpfm_api_output_formatter.ConvertToHeaderPartnerPlant(sdc, psdc)
	// if err != nil {
	// 	fmt.Printf("err = %+v \n", err)
	// 	return nil, err
	// }

	osdc.Message = dpfm_api_output_formatter.Message{
		Header:             header,
		HeaderPartner:      headerPartner,
		HeaderPartnerPlant: headerPartnerPlant,
	}

	return osdc, nil
}
