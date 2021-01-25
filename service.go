package cicil

type CicilService interface {
	GetToken() (digest string, dateNowReturn string)
	GetCheckoutURL(param CheckoutRequest) (resp CheckoutResponse, err error)
	SetCancelOrder(param CancelOrderRequest) (resp CancelOrderResponse, err error)
	UpdateStatus (param UpdateStatusRequest) (resp UpdateStatusResponse, err error)
}
