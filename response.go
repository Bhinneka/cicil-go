package cicil

type CheckoutResponse struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	PoNumber string `json:"po_number"`
	Url      string `json:"url"`
}

type CancelOrderResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type UpdateStatusResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}