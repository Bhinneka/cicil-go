package main

import (
	"encoding/json"
	"fmt"
	"github.com/abdulahwahdi/cicil-go"
	"time"
)

const (
	BaseURL        = "https://sandbox-api.cicil.dev/v1"
    MerchantID     = "YOUR_MERCHANT_ID"
    MerchantSecret = "YOUR_MERCHANT_SERCRET"
    APIKey         = "YOUR_API_KEY"
	DefaultTimeout = 20 * time.Second
)

func main() {
	newCicil := cicil.New(BaseURL,APIKey, MerchantID, MerchantSecret,DefaultTimeout)
	SampleCheckout(newCicil)
}

func SampleCheckout(cic cicil.CicilService) {
	dataSample := `{
					"transaction": {
						"total_amount": 13119000,
						"transaction_id": "ORD10111808",
						"item_list": [
							{
								"item_id":"SKU101112",
								"type": "product",
								"name":"Notebook CICIL C12",
								"price":12999000,
								"category":"laptop",
								"url":"https://www.tokocicil.com/product/sku101112",
								"quantity":1,
								"seller_id":"tokocicil-official"
							},
							{
								"item_id":"SKU131415",
								"type": "product",
								"name":"Sticker Aja",
								"price":60000,
								"category":"accessories",
								"url":"https://www.tokocicil.com/product/sku131415",
								"quantity":2
							},
							{
								"item_id":"insurance",
								"type": "fee",
								"name":"Insurance Fee",
								"price":5000,
								"quantity":1
							},
							{
								"item_id":"shipment_cost",
								"type": "shipment_cost",
								"name":"Shipping Fee",
								"price":45000,
								"quantity":1
							},
							{
								"item_id":"CICIL1212",
								"type": "discount",
								"name":"Promo Harbolnas CICIL",
								"price":-50000,
								"quantity":1
							}
						]
					},
					"buyer": {
						"fullname": "John Doe",
						"email": "john.doe@mail.com",
						"phone": "085322984060",
						"address": "Jl. Sd Inpres RT.003/RW.006 No.174A 13950 Cakung, Pulogebang",
						"city": "Jakarta Timur",
						"district": "JK",
						"postal_code": "11630",
						"company": "Cicil",
						"country": "ID"
					},
					"shipment": {
						"shipment_provider": "Flat rate",
						"shipping_price": 40000,
						"shipping_tax": 0,
						"name": "John Doe",
						"address": "Jl. Sd Inpres RT.003/RW.006 No.174A 13950 Cakung, Pulogebang",
						"city": "Jakarta Timur",
						"district": "Jakarta Timur",
						"postal_code": "11630",
						"phone": "085322984060",
						"company": "Cicil",
						"country": "ID"
					},
					"push_url": "https://api.tokocicil.com/update",
					"redirect_url": "https://toko.cicil.dev",
					"back_url": "https://toko.cicil.dev"
				}`
	dataSampleStruct := cicil.CheckoutRequest{}
	json.Unmarshal([]byte(dataSample), &dataSampleStruct)
	respGetCheckoutUrl, errGetCheckoutUrl := cic.GetCheckoutURL(dataSampleStruct)
	if errGetCheckoutUrl != nil {
		fmt.Println(errGetCheckoutUrl.Error())
	}
	fmt.Println(respGetCheckoutUrl)
}
