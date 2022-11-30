package requests

type OrderID struct {
	InvoiceDocument                 *int   `json:"InvoiceDocument"`
	OrderID                         *int   `json:"OrderID`
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}
