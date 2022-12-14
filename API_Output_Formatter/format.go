package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"encoding/json"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]Header, error) {
	orderID := psdc.OrderID
	headerOrdersHeader := psdc.HeaderOrdersHeader
	calculateDeliveryDocument := psdc.CalculateDeliveryDocument
	headers := make([]Header, 0, len(*headerOrdersHeader))

	for _, v := range *headerOrdersHeader {
		orderId := v.OrderID
		for _, orderID := range *orderID {
			if *orderID.OrderID != *orderId {
				continue
			}
			header := Header{}
			inputHeader := sdc.DeliveryDocument
			inputData, err := json.Marshal(inputHeader)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(inputData, &header)
			if err != nil {
				return nil, err
			}

			data, err := json.Marshal(v)
			if err != nil {
				return nil, err
			}
			err = json.Unmarshal(data, &header)
			if err != nil {
				return nil, err
			}

			header.DeliveryDocument = calculateDeliveryDocument.DeliveryDocumentLatestNumber
			header.ReferenceDocument = orderID.ReferenceDocument
			headers = append(headers, header)
		}
	}

	return &headers, nil
}

func ConvertToHeaderPartner(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]HeaderPartner, error) {
	headerOrdersHeaderPartner := psdc.HeaderOrdersHeaderPartner
	calculateDeliveryDocument := psdc.CalculateDeliveryDocument
	headerPartners := make([]HeaderPartner, 0, len(*headerOrdersHeaderPartner))

	for _, v := range *headerOrdersHeaderPartner {
		headerPartner := HeaderPartner{}
		inputHeaderPartner := sdc.DeliveryDocument.HeaderPartner[0]
		inputData, err := json.Marshal(inputHeaderPartner)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(inputData, &headerPartner)
		if err != nil {
			return nil, err
		}

		data, err := json.Marshal(v)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(data, &headerPartner)
		if err != nil {
			return nil, err
		}

		headerPartner.DeliveryDocument = calculateDeliveryDocument.DeliveryDocumentLatestNumber
		headerPartners = append(headerPartners, headerPartner)
	}

	return &headerPartners, nil
}

// func ConvertToHeaderPartnerPlant(
// 	sdc *api_input_reader.SDC,
// 	psdc *api_processing_data_formatter.SDC,
// ) (*[]HeaderPartnerPlant, error) {
// 	headerOrdersHeaderPartnerPlant := psdc.HeaderOrdersHeaderPartnerPlant
// 	calculateDeliveryDocument := psdc.CalculateDeliveryDocument
// 	headerPartnerPlants := make([]HeaderPartnerPlant, 0, len(*headerOrdersHeaderPartnerPlant))

// 	for _, v := range *headerOrdersHeaderPartnerPlant {
// 		headerPartnerPlant := HeaderPartnerPlant{}
// 		inputHeaderPartnerPlant := sdc.DeliveryDocument.HeaderPartner[0].HeaderPartnerPlant[0]
// 		inputData, err := json.Marshal(inputHeaderPartnerPlant)
// 		if err != nil {
// 			return nil, err
// 		}
// 		err = json.Unmarshal(inputData, &headerPartnerPlant)
// 		if err != nil {
// 			return nil, err
// 		}

// 		data, err := json.Marshal(v)
// 		if err != nil {
// 			return nil, err
// 		}
// 		err = json.Unmarshal(data, &headerPartnerPlant)
// 		if err != nil {
// 			return nil, err
// 		}

// 		headerPartnerPlant.DeliveryDocument = calculateDeliveryDocument.DeliveryDocumentLatestNumber
// 		headerPartnerPlants = append(headerPartnerPlants, headerPartnerPlant)
// 	}

// 	return &headerPartnerPlants, nil
// }
