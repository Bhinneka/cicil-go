package cicil

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

const (
	BaseURL        = "https://sandbox-api.cicil.dev/v1"
	MerchantID     = "YOUR_MERCHANT_ID"
	MerchantSecret = "YOUR_MERCHANT_SERCRET"
	APIKey         = "YOUR_API_KEY"
	DefaultTimeout = 20 * time.Second
)

func TestCicilCreation(t *testing.T){
	t.Run("Test Cicil Creation", func(t *testing.T) {
		//construct cicil
		//and add required parameters
		cic := New(BaseURL,APIKey, MerchantID, MerchantSecret,DefaultTimeout)

		if cic == nil {
			t.Error("Cannot Call Cicil Constructor")
		}
	})
}

func TestCicil_GetCheckoutURL(t *testing.T) {
	orderData := []byte(`{"transaction": {"total_amount": 13119000,"transaction_id": "ORD10111808","item_list": [{"item_id":"SKU101112","type": "product","name":"Notebook CICIL C12","price":12999000,"category":"laptop","url":"https://www.tokocicil.com/product/sku101112","quantity":1,"seller_id":"tokocicil-official"},{"item_id":"SKU131415","type": "product","name":"Sticker Aja","price":60000,"category":"accessories","url":"https://www.tokocicil.com/product/sku131415","quantity":2},{"item_id":"insurance","type": "fee","name":"Insurance Fee","price":5000,"quantity":1},{"item_id":"shipment_cost","type": "shipment_cost","name":"Shipping Fee","price":45000,"quantity":1},{"item_id":"CICIL1212","type": "discount","name":"Promo Harbolnas CICIL","price":-50000,"quantity":1}]},"buyer": {"fullname": "John Doe","email": "john.doe@mail.com","phone": "085322984060","address": "Jl. Sd Inpres RT.003/RW.006 No.174A 13950 Cakung, Pulogebang","city": "Jakarta Timur","district": "JK","postal_code": "11630","company": "Cicil","country": "ID"},"shipment": {"shipment_provider": "Flat rate","shipping_price": 40000,"shipping_tax": 0,"name": "John Doe","address": "Jl. Sd Inpres RT.003/RW.006 No.174A 13950 Cakung, Pulogebang","city": "Jakarta Timur","district": "Jakarta Timur","postal_code": "11630","phone": "085322984060","company": "Cicil","country": "ID"},"push_url": "https://api.tokocicil.com/update","redirect_url": "https://toko.cicil.dev","back_url": "https://toko.cicil.dev"}`)

	t.Run("Test Cicil Get Order URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{ "message": "", "po_number": "PO210122-143655", "status": "success", "url": "https://sandbox-staging.cicil.dev/po/UE8yMTAxMjItMTQzNjU1" }`))
		}))
		//close server
		defer ts.Close()

		//construct cicil
		amm := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var order CheckoutRequest
		err := json.Unmarshal(orderData, &order)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errGetOrderURL := amm.GetCheckoutURL(order)
		if errGetOrderURL != nil {
			t.Errorf("GetOrderURL() returned an error: %s", errGetOrderURL.Error())
		}
	})
	t.Run("Test Negative Case Cicil Get Order URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"invalid parameter on order request"}`))

		}))
		//close server
		defer ts.Close()

		//construct cicil
		amm := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var order CheckoutRequest
		err := json.Unmarshal(orderData, &order)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errGetOrderURL := amm.GetCheckoutURL(order)
		if errGetOrderURL != nil {
			assert.Error(t, errGetOrderURL)
		}
	})
}

func TestCicil_SetCancelOrder(t *testing.T) {
	orderData := []byte(`{"po_number": "PO181112-123456","reason": "out of stock","cancelled_by": "tokocicil","total_amount":13119000,"transaction_id":"ORD10111808"}`)

	t.Run("Test Cicil Get Order URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"success","message":""}`))
		}))
		//close server
		defer ts.Close()

		//construct cicil
		cic := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var cancelOrder CancelOrderRequest
		err := json.Unmarshal(orderData, &cancelOrder)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errCancelOrder := cic.SetCancelOrder(cancelOrder)
		if errCancelOrder != nil {
			t.Errorf("SetCancelOrder() returned an error: %s", errCancelOrder.Error())
		}
	})
	t.Run("Test Negative Case Cicil Get Order URL", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"invalid parameter on order request"}`))

		}))
		//close server
		defer ts.Close()

		//construct cicil
		amm := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var cancelOrder CancelOrderRequest
		err := json.Unmarshal(orderData, &cancelOrder)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errCancelOrder := amm.SetCancelOrder(cancelOrder)
		if errCancelOrder != nil {
			assert.Error(t, errCancelOrder)
		}
	})
}

func TestCicil_UpdateStatus(t *testing.T) {
	orderData := []byte(`{"po_number": "PO181112-123456","po_status": "delivered","transaction_id":"ORD10111808"}`)

	t.Run("Test Cicil Update Status", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"status":"success","message":""}`))
		}))
		//close server
		defer ts.Close()

		//construct cicil
		amm := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var updateStatus UpdateStatusRequest
		err := json.Unmarshal(orderData, &updateStatus)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errUpdateStatus := amm.UpdateStatus(updateStatus)
		if errUpdateStatus != nil {
			t.Errorf("GetOrderURL() returned an error: %s", errUpdateStatus.Error())
		}
	})
	t.Run("Test Negative Case Cicil Update Status", func(t *testing.T) {
		ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			if r.Method != "POST" {
				t.Errorf("Expected ‘POST’ request, got ‘%s’", r.Method)
			}

			w.Header().Set("Content-Type", "application-json; charset=utf-8")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"message":"invalid parameter on order request"}`))

		}))
		//close server
		defer ts.Close()

		//construct cicil
		amm := New(ts.URL, APIKey, MerchantID, MerchantSecret,DefaultTimeout)
		var updateStatus UpdateStatusRequest
		err := json.Unmarshal(orderData, &updateStatus)
		if err != nil {
			t.Error("Cannot Unmarshal Order JSON data")
		}

		_, errUpdateStatus := amm.UpdateStatus(updateStatus)
		if errUpdateStatus != nil {
			assert.Error(t, errUpdateStatus)
		}
	})
}
