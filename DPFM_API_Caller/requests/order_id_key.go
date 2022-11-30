package requests

type OrderIDKey struct {
	HeaderCompleteDeliveryIsDefined *bool  `json:"HeaderCompleteDeliveryIsDefined"`
	OverallDeliveryStatus           string `json:"OverallDeliveryStatus"`
}
