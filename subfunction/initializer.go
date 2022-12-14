package subfunction

import (
	"context"
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	"golang.org/x/xerrors"
)

type SubFunction struct {
	ctx context.Context
	db  *database.Mysql
	l   *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx: ctx,
		db:  db,
		l:   l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.MetaData, error) {
	var err error
	var metaData *api_processing_data_formatter.MetaData

	metaData, err = psdc.ConvertToMetaData(sdc)
	if err != nil {
		return nil, err
	}

	return metaData, nil
}

func (f *SubFunction) OrderItemByNumberSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItem, error) {
	var args []interface{}

	issuingPlantBusinessPartner := sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartner
	receivingPlantBusinessPartner := sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartner

	if len(*issuingPlantBusinessPartner) != len(*receivingPlantBusinessPartner) {
		return nil, nil
	}

	issuingPlant := sdc.DeliveryDocumentInputParameters.IssuingPlant
	receivingPlant := sdc.DeliveryDocumentInputParameters.ReceivingPlant

	if len(*issuingPlant) != len(*receivingPlant) {
		return nil, nil
	}

	dataKey, err := psdc.ConvertToOrderItemByNumberSpecificationKey(sdc, len(*issuingPlantBusinessPartner), len(*issuingPlant))
	if err != nil {
		return nil, err
	}

	for i := range *issuingPlantBusinessPartner {
		dataKey.IssuingPlantBusinessPartner[i] = (*issuingPlantBusinessPartner)[i]
		dataKey.ReceivingPlantBusinessPartner[i] = (*receivingPlantBusinessPartner)[i]
	}

	for i := range *issuingPlant {
		dataKey.IssuingPlant[i] = (*issuingPlant)[i]
		dataKey.ReceivingPlant[i] = (*receivingPlant)[i]
	}

	repeat1 := strings.Repeat("(?,?),", len(dataKey.IssuingPlantBusinessPartner)-1) + "(?,?)"
	for i := range dataKey.IssuingPlantBusinessPartner {
		args = append(args, dataKey.IssuingPlantBusinessPartner[i], dataKey.ReceivingPlantBusinessPartner[i])
	}

	repeat2 := strings.Repeat("(?,?),", len(dataKey.IssuingPlant)-1) + "(?,?)"
	for i := range dataKey.IssuingPlant {
		args = append(args, dataKey.IssuingPlant[i], dataKey.ReceivingPlant[i])
	}

	args = append(
		args,
		dataKey.IssuingPlantPartnerFunction,
		dataKey.ReceivingPlantPartnerFunction,
		dataKey.ItemCompleteDeliveryIsDefined,
		dataKey.ItemDeliveryBlockStatus,
		dataKey.ItemDeliveryStatus,
	)

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner) IN ( `+repeat1+` )
		AND (IssuingPlant, ReceivingPlant) IN ( `+repeat2+` )
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner, 
		IssuingPlant, ReceivingPlant, IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner) IN ( `+repeat1+` )
		AND (IssuingPlant, ReceivingPlant) IN ( `+repeat2+` )
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderItemByNumberSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemByRangeSpecification(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderItem, error) {
	dataKey, err := psdc.ConvertToOrderItemByRangeSpecificationKey(sdc)
	if err != nil {
		return nil, err
	}

	dataKey.IssuingPlantBusinessPartnerFrom = sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartnerFrom
	dataKey.IssuingPlantBusinessPartnerTo = sdc.DeliveryDocumentInputParameters.IssuingPlantBusinessPartnerTo
	dataKey.ReceivingPlantBusinessPartnerFrom = sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartnerFrom
	dataKey.ReceivingPlantBusinessPartnerTo = sdc.DeliveryDocumentInputParameters.ReceivingPlantBusinessPartnerTo
	dataKey.IssuingPlantFrom = sdc.DeliveryDocumentInputParameters.IssuingPlantFrom
	dataKey.IssuingPlantTo = sdc.DeliveryDocumentInputParameters.IssuingPlantTo
	dataKey.ReceivingPlantFrom = sdc.DeliveryDocumentInputParameters.ReceivingPlantFrom
	dataKey.ReceivingPlantTo = sdc.DeliveryDocumentInputParameters.ReceivingPlantTo

	var count *int
	err = f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner BETWEEN ? AND ?
		AND ReceivingPlantBusinessPartner BETWEEN ? AND ?
		AND IssuingPlant BETWEEN ? AND ?
		AND ReceivingPlant BETWEEN ? AND ?
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.IssuingPlantBusinessPartnerFrom, dataKey.IssuingPlantBusinessPartnerTo, dataKey.ReceivingPlantBusinessPartnerFrom, dataKey.ReceivingPlantBusinessPartnerTo, dataKey.IssuingPlantFrom, dataKey.IssuingPlantTo, dataKey.ReceivingPlantFrom, dataKey.ReceivingPlantTo, dataKey.IssuingPlantPartnerFunction, dataKey.ReceivingPlantPartnerFunction, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, IssuingPlantBusinessPartner, ReceivingPlantBusinessPartner, 
		IssuingPlant, ReceivingPlant, IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE IssuingPlantBusinessPartner BETWEEN ? AND ?
		AND ReceivingPlantBusinessPartner BETWEEN ? AND ?
		AND IssuingPlant BETWEEN ? AND ?
		AND ReceivingPlant BETWEEN ? AND ?
		AND (IssuingPlantPartnerFunction, ReceivingPlantPartnerFunction, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.IssuingPlantBusinessPartnerFrom, dataKey.IssuingPlantBusinessPartnerTo, dataKey.ReceivingPlantBusinessPartnerFrom, dataKey.ReceivingPlantBusinessPartnerTo, dataKey.IssuingPlantFrom, dataKey.IssuingPlantTo, dataKey.ReceivingPlantFrom, dataKey.ReceivingPlantTo, dataKey.IssuingPlantPartnerFunction, dataKey.ReceivingPlantPartnerFunction, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderIDByRangeSpecification(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderID(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.OrderID, error) {
	dataKey, err := psdc.ConvertToOrderIDKey(sdc)
	if err != nil {
		return nil, err
	}

	rows, err := f.db.Query(
		`SELECT OrderID, HeaderCompleteDeliveryIsDefined, OverallDeliveryStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_header_data
		WHERE (OrderID, HeaderCompleteDeliveryIsDefined) = (?, ?)
		AND OverallDeliveryStatus <> ?;`, dataKey.ReferenceDocument, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.OverallDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToOrderID(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(2)

	psdc.MetaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// // I-1. OrderItemの絞り込み
		// psdc.OrderItem, e = f.OrderItemByNumberSpecification(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// I-1. OrderItemの絞り込み
		psdc.OrderItem, e = f.OrderItemByRangeSpecification(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// I-2. ヘッダパートナプラントのデータ取得
		psdc.HeaderOrdersHeaderPartnerPlant, e = f.OrdersHeaderPartnerPlant(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// // II-0-1-1. OrderIDが未入出荷であり、かつ、OrderIDに入出荷伝票未登録残がある、明細の取得
		// psdc.OrderID, e = f.OrderID(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 1-1. オーダー参照レコード・値の取得（オーダーヘッダ）
		// psdc.HeaderOrdersHeader, e = f.OrdersHeader(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }

		// // 1-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）
		// psdc.HeaderOrdersHeaderPartner, e = f.OrdersHeaderPartner(sdc, psdc)
		// if e != nil {
		// 	err = e
		// 	return
		// }
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-7. DeliveryDocument
		psdc.CalculateDeliveryDocument, e = f.CalculateDeliveryDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	// osdc, err = f.SetValue(sdc, osdc, psdc)
	if err != nil {
		return err
	}

	return nil
}
