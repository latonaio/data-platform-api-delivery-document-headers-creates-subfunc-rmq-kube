package requests

type DeliveryDocumentKey struct {
	CompleteDeliveryIsDefined *bool  `json:"CompleteDeliveryIsDefined"`
	OverallDeliveryStatus     string `json:"OverallDeliveryStatus"`
}
