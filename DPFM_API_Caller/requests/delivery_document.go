package requests

type DeliveryDocument struct {
	InvoiceDocument           *int   `json:"InvoiceDocument"`
	DeliveryDocument          *int   `json:"DeliveryDocument`
	CompleteDeliveryIsDefined *bool  `json:"ompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}
