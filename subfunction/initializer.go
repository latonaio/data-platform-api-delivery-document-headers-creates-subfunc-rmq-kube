package subfunction

import (
	"context"
	api_input_reader "data-platform-api-delivery-document-headers-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-delivery-document-headers-creates-subfunc/API_Processing_Data_Formatter"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
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
	var metaData *api_processing_data_formatter.MetaData
	var orderID *[]api_processing_data_formatter.OrderID
	var headerOrdersHeader *[]api_processing_data_formatter.HeaderOrdersHeader
	var headerOrdersHeaderPartner *[]api_processing_data_formatter.HeaderOrdersHeaderPartner
	var headerOrdersHeaderPartnerPlant *[]api_processing_data_formatter.HeaderOrdersHeaderPartnerPlant
	var calculateDeliveryDocument *api_processing_data_formatter.CalculateDeliveryDocument
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(2)

	metaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// II-0-1-1. OrderIDが未入出荷であり、かつ、OrderIDに入出荷伝票未登録残がある、明細の取得
		orderID, e = f.OrderID(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-1. オーダー参照レコード・値の取得（オーダーヘッダ）
		headerOrdersHeader, e = f.OrdersHeader(orderID, sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）
		headerOrdersHeaderPartner, e = f.OrdersHeaderPartner(orderID, sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-4. オーダー参照レコード・値の取得（オーダーヘッダパートナプラント）
		headerOrdersHeaderPartnerPlant, e = f.OrdersHeaderPartnerPlant(orderID, sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-7. DeliveryDocument
		calculateDeliveryDocument, e = f.CalculateDeliveryDocument(metaData, sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	osdc, err = f.SetValue(sdc, osdc, orderID, headerOrdersHeader, headerOrdersHeaderPartner, headerOrdersHeaderPartnerPlant, calculateDeliveryDocument)
	if err != nil {
		return err
	}

	return nil
}