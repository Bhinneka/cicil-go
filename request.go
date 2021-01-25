package cicil

type CheckoutRequest struct {
	Transaction Transaction  `json:"transaction"`
	PushURL     string       `json:"push_url"`
	RedirectURL string       `json:"redirect_url"`
	Buyer       Buyer        `json:"buyer"`
	Shipment    Shipment     `json:"shipment"`
	SellerList  []SellerList `json:"seller_list,omitempty"`
	BackURL     string       `json:"back_url"`
}

type Transaction struct {
	TotalAmount   int        `json:"total_amount"`
	TransactionID string     `json:"transaction_id"`
	ItemLists     []ItemList `json:"item_list"`
}

type ItemList struct {
	ItemID   string `json:"item_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
	Category string `json:"category,omitempty"`
	Url      string `json:"url,omitempty"`
	SellerID string `json:"seller_id,omitempty"`
}

type Buyer struct {
	Fullname   string `json:"fullname"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Address    string `json:"address"`
	City       string `json:"city"`
	District   string `json:"district"`
	PostalCode string `json:"postal_code"`
	Company    string `json:"company"`
	Country    string `json:"country"`
}

type Shipment struct {
	ShipmentProvider string `json:"shipment_provider"`
	ShipmentPrice    int    `json:"shipment_price"`
	ShipmentTax      int    `json:"shipment_tax"`
	Address          string `json:"address"`
	City             string `json:"city"`
	District         string `json:"district"`
	PostalCode       string `json:"postal_code"`
	Phone            string `json:"phone"`
	Company          string `json:"company"`
	Name             string `json:"name"`
	Country          string `json:"country"`
}

type SellerList struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Url   string `json:"url"`
}

type CancelOrderRequest struct {
	PONumber      string `json:"po_number"`
	Reason        string `json:"reason"`
	CancelledBy   string `json:"cancelled_by"`
	TotalAmount   int    `json:"total_amount"`
	TransactionID string `json:"transaction_id"`
}

type UpdateStatusRequest struct {
	PONumber      string   `json:"po_number"`
	POStatus      string   `json:"po_status"`
	TransactionID string   `json:"transaction_id"`
	Shipment      Shipment `json:"shipment,omitempty"`
	Reason        string   `json:"reason,omitempty"`
}
