package requests

type CalculateInvoiceDocumentQueryGets struct {
	ServiceLabel                string `json:"service_label"`
	FieldNameWithNumberRange    string `json:"FieldNameWithNumberRange"`
	InvoiceDocumentLatestNumber *int   `json:"InvoiceDocumentLatestNumber"`
}
