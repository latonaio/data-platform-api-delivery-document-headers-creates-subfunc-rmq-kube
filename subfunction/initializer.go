package subfunction

import (
	"context"
	api_input_reader "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-invoice-document-headers-creates-subfunc-rmq/API_Processing_Data_Formatter"

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
		WHERE HeaderCompleteDeliveryIsDefined = ?
		AND OverallDeliveryStatus <> ?;`, dataKey.HeaderCompleteDeliveryIsDefined, dataKey.OverallDeliveryStatus,
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

func (f *SubFunction) DeliveryDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*[]api_processing_data_formatter.DeliveryDocument, error) {
	dataKey, err := psdc.ConvertToDeliveryDocumentKey(sdc)
	if err != nil {
		return nil, err
	}

	rows, err := f.db.Query(
		`SELECT DeliveryDocument, CompleteDeliveryIsDefined, OverallDeliveryStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_delivery_document_header_data
		WHERE CompleteDeliveryIsDefined = ?
		AND OverallDeliveryStatus <> ?;`, dataKey.CompleteDeliveryIsDefined, dataKey.OverallDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}

	data, err := psdc.ConvertToDeliveryDocument(sdc, rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CalculateInvoiceDocument(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) (*api_processing_data_formatter.CalculateInvoiceDocument, error) {
	metaData := psdc.MetaData
	dataKey, err := psdc.ConvertToCalculateInvoiceDocumentKey()
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

	dataQueryGets, err := psdc.ConvertToCalculateInvoiceDocumentQueryGets(sdc, rows)
	if err != nil {
		return nil, err
	}

	calculateInvoiceDocument := CalculateInvoiceDocument(*dataQueryGets.InvoiceDocumentLatestNumber)

	data, err := psdc.ConvertToCalculateInvoiceDocument(calculateInvoiceDocument)
	if err != nil {
		return nil, err
	}

	return data, err
}

func CalculateInvoiceDocument(latestNumber int) *int {
	res := latestNumber + 1
	return &res
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}
	wg.Add(3)

	psdc.MetaData, err = f.MetaData(sdc, psdc)
	if err != nil {
		return err
	}

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-3. [OrderID]が未請求であり、かつ、[OrderID]に入出荷伝票未登録残がある、明細の取得
		psdc.OrderID, e = f.OrderID(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// I-1-1. オーダ参照レコード・値の取得（オーダーヘッダ）
		psdc.HeaderOrdersHeader, e = f.OrdersHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// I-1-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）
		psdc.HeaderOrdersHeaderPartner, e = f.OrdersHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 2-3. [Delivery Document]が未請求であり、かつ、[Delivery Document]に入出荷伝票未登録残がある、明細の取得
		psdc.DeliveryDocument, e = f.DeliveryDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		//I-2-1. 入出荷伝票参照レコード・値の取得（入出荷伝票ヘッダ）
		psdc.HeaderDeliveryDocumentHeader, e = f.DeliveryDocumentHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// I-1-2. オーダー参照レコード・値の取得（オーダーヘッダパートナ）
		psdc.HeaderDeliveryDocumentHeaderPartner, e = f.DeliveryDocumentHeaderPartner(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-1 InvoiceDocument
		psdc.CalculateInvoiceDocument, e = f.CalculateInvoiceDocument(sdc, psdc)
		if e != nil {
			err = e
			return
		}
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	osdc, err = f.SetValue(sdc, psdc, osdc)
	if err != nil {
		return err
	}

	return nil
}
