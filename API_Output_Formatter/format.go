package dpfm_api_output_formatter

import (
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
)

func ConvertToHeader(
	sdc *api_input_reader.SDC,
	orderID *[]api_processing_data_formatter.OrderID,
	headerOrdersHeader *[]api_processing_data_formatter.HeaderOrdersHeader,
	calculateDeliveryDocument *api_processing_data_formatter.CalculateDeliveryDocument,
) (*[]Header, error) {
	header := make([]Header, 0, len(*headerOrdersHeader))

	for _, v := range *headerOrdersHeader {
		orderId := v.OrderID
		for _, orderID := range *orderID {
			if *orderID.OrderID != *orderId {
				continue
			}

			header = append(header, Header{
				DeliveryDocument:              calculateDeliveryDocument.DeliveryDocumentLatestNumber,
				Buyer:                         v.Buyer,
				Seller:                        v.Seller,
				ReferenceDocument:             orderID.ReferenceDocument,
				ReferenceDocumentItem:         sdc.DeliveryDocument.ReferenceDocumentItem,
				OrderID:                       v.OrderID,
				OrderItem:                     sdc.DeliveryDocument.OrderItem,
				ContractType:                  v.ContractType,
				OrderValidityStartDate:        v.OrderValidityStartDate,
				OrderValidityEndDate:          v.OrderValidityEndDate,
				InvoiceScheduleStartDate:      v.InvoiceScheduleStartDate,
				InvoiceScheduleEndDate:        v.InvoiceScheduleEndDate,
				IssuingPlantTimeZone:          sdc.DeliveryDocument.IssuingPlantTimeZone,
				ReceivingPlantTimeZone:        sdc.DeliveryDocument.ReceivingPlantTimeZone,
				DocumentDate:                  sdc.DeliveryDocument.DocumentDate,
				PlannedGoodsIssueDate:         sdc.DeliveryDocument.PlannedGoodsIssueDate,
				PlannedGoodsIssueTime:         sdc.DeliveryDocument.PlannedGoodsIssueTime,
				PlannedGoodsReceiptDate:       sdc.DeliveryDocument.PlannedGoodsReceiptDate,
				PlannedGoodsReceiptTime:       sdc.DeliveryDocument.PlannedGoodsReceiptTime,
				BillingDocumentDate:           sdc.DeliveryDocument.BillingDocumentDate,
				CompleteDeliveryIsDefined:     sdc.DeliveryDocument.CompleteDeliveryIsDefined,
				OverallDeliveryStatus:         sdc.DeliveryDocument.OverallDeliveryStatus,
				CreationDate:                  sdc.DeliveryDocument.CreationDate,
				CreationTime:                  sdc.DeliveryDocument.CreationTime,
				IssuingBlockReason:            sdc.DeliveryDocument.IssuingBlockReason,
				ReceivingBlockReason:          sdc.DeliveryDocument.ReceivingBlockReason,
				GoodsIssueOrReceiptSlipNumber: sdc.DeliveryDocument.GoodsIssueOrReceiptSlipNumber,
				HeaderBillingStatus:           sdc.DeliveryDocument.HeaderBillingStatus,
				HeaderBillingConfStatus:       sdc.DeliveryDocument.HeaderBillingConfStatus,
				HeaderBillingBlockReason:      sdc.DeliveryDocument.HeaderBillingBlockReason,
				HeaderGrossWeight:             sdc.DeliveryDocument.HeaderGrossWeight,
				HeaderNetWeight:               sdc.DeliveryDocument.HeaderNetWeight,
				HeaderVolume:                  sdc.DeliveryDocument.HeaderVolume,
				HeaderVolumeUnit:              sdc.DeliveryDocument.HeaderVolumeUnit,
				HeaderWeightUnit:              sdc.DeliveryDocument.HeaderWeightUnit,
				Incoterms:                     v.Incoterms,
				IsExportImportDelivery:        v.IsExportImportDelivery,
				LastChangeDate:                sdc.DeliveryDocument.LastChangeDate,
				IssuingPlantBusinessPartner:   sdc.DeliveryDocument.IssuingPlantBusinessPartner,
				IssuingPlant:                  sdc.DeliveryDocument.IssuingPlant,
				ReceivingPlant:                sdc.DeliveryDocument.ReceivingPlant,
				ReceivingPlantBusinessPartner: sdc.DeliveryDocument.ReceivingPlantBusinessPartner,
				DeliverToParty:                sdc.DeliveryDocument.DeliverToParty,
				DeliverFromParty:              sdc.DeliveryDocument.DeliverFromParty,
				TransactionCurrency:           v.TransactionCurrency,
				OverallDelivReltdBillgStatus:  sdc.DeliveryDocument.OverallDelivReltdBillgStatus,
			})
		}
	}

	return &header, nil
}
