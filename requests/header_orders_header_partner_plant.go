package requests

type HeaderOrdersHeaderPartnerPlant struct {
	DeliveryDocument *int   `json:"DeliveryDocument`
	OrderID          *int   `json:"OrderID"`
	PartnerFunction  string `json:"PartnerFunction"`
	BusinessPartner  *int   `json:"BusinessPartner"`
	Plant            string `json:"Plant"`
}
