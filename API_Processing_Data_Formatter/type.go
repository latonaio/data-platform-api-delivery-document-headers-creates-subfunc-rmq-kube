package api_processing_data_formatter

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Plant         string   `json:"plant/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	MetaData                       *MetaData                         `json:"MetaData"`
	OrderID                        *[]OrderID                        `json:"OrderID`
	CalculateDeliveryDocument      *CalculateDeliveryDocument        `json:"CalculateDeliveryDocument`
	HeaderOrdersHeader             *[]HeaderOrdersHeader             `json:"HeaderOrdersHeader"`
	HeaderOrdersHeaderPartner      *[]HeaderOrdersHeaderPartner      `json"HeaderOrdersHeaderPartner"`
	HeaderOrdersHeaderPartnerPlant *[]HeaderOrdersHeaderPartnerPlant `json"HeaderOrdersHeaderPartnerPlant"`
}

type MetaData struct {
	BusinessPartnerID *int   `json:"business_partner"`
	ServiceLabel      string `json:"service_label"`
}

type OrderIDKey struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type OrderID struct {
	ReferenceDocument               *int   `json:"ReferenceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}

type CalculateDeliveryDocumentKey struct {
	ServiceLabel             string `json:"service_label"`
	FieldNameWithNumberRange string
}

type CalculateDeliveryDocumentQueryGets struct {
	ServiceLabel                 string `json:"service_label"`
	FieldNameWithNumberRange     string
	DeliveryDocumentLatestNumber *int
}

type CalculateDeliveryDocument struct {
	DeliveryDocumentLatestNumber *int
	DeliveryDocument             *int `json:"DeliveryDocument"`
}

type HeaderOrdersHeader struct {
	DeliveryDocument         *int    `json:"DeliveryDocument`
	OrderID                  *int    `json:"OrderID"`
	OrderType                string  `json:"OrderType"`
	Buyer                    *int    `json:"Buyer"`
	Seller                   *int    `json:"Seller"`
	ContractType             *string `json:"ContractType"`
	OrderValidityStartDate   *string `json:"OrderValidityStartDate"`
	OrderValidityEndDate     *string `json:"OrderValidityEndDate"`
	InvoiceScheduleStartDate *string `json:"InvoiceScheduleStartDate"`
	InvoiceScheduleEndDate   *string `json:"InvoiceScheduleEndDate"`
	TransactionCurrency      *string `json:"TransactionCurrency"`
	Incoterms                *string `json:"Incoterms"`
	IsExportImportDelivery   *bool   `json:"IsExportImportDelivery"`
}

type HeaderOrdersHeaderPartner struct {
	DeliveryDocument        *int    `json:"DeliveryDocument`
	OrderID                 *int    `json:"OrderID"`
	PartnerFunction         string  `json:"PartnerFunction"`
	BusinessPartner         *int    `json:"BusinessPartner"`
	BusinessPartnerFullName *string `json:"BusinessPartnerFullName"`
	BusinessPartnerName     *string `json:"BusinessPartnerName"`
	Organization            *string `json:"Organization"`
	Country                 *string `json:"Country"`
	Language                *string `json:"Language"`
	Currency                *string `json:"Currency"`
	ExternalDocumentID      *string `json:"ExternalDocumentID"`
	AddressID               *int    `json:"AddressID"`
}

type HeaderOrdersHeaderPartnerPlant struct {
	DeliveryDocument *int   `json:"DeliveryDocument`
	OrderID          *int   `json:"OrderID"`
	PartnerFunction  string `json:"PartnerFunction"`
	BusinessPartner  *int   `json:"BusinessPartner"`
	Plant            string `json:"Plant"`
}
