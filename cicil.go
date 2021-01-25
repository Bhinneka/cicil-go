package cicil

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

// cicil entry struick for cicil
type cicil struct {
	MerchantID     string
	APIKey         string
	MerchantSecret string
	BaseURL        string
	client         *cicilHttpClient
	*logger
}

/*New Function, create cicil pointer
Required parameter :
1. Your MerchantID (this from Team cicil)
2. Your MerchantSecret (this from Team cicil)
3. Your APIKey (this from Team cicil)
3. BaseURL (hit to endpoint ex: https://sandbox-api.cicil.dev/v1 for sandbox.
this value based on https://docs.cicil.app/#introduction)
*/
func New(baseUrl string, apiKey string, merchantId string, merchantSecret string, timeout time.Duration) *cicil {
	httpRequest := newRequest(timeout)
	return &cicil{
		APIKey:         apiKey,
		MerchantID:     merchantId,
		MerchantSecret: merchantSecret,
		BaseURL:        baseUrl,
		client:         httpRequest,
		logger:         newLogger(),
	}
}

func (c *cicil) call(method string, path string, body io.Reader, v interface{}, headers map[string]string) error {
	c.info().Println("Starting http call..")
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	path = fmt.Sprintf("%s%s", c.BaseURL, path)
	return c.client.exec(method, path, body, v, headers)
}

func (c *cicil) GetToken() (digest string, dateNowReturn string) {
	dateNow := time.Now().Format(time.RFC1123)
	formattedMessage := fmt.Sprintf("%s%s", c.MerchantSecret, dateNow)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(c.APIKey))
	// Write Data to it
	h.Write([]byte(formattedMessage))

	// Get result and encode as hexadecimal string
	digestData := hex.EncodeToString(h.Sum(nil))
	basicData := fmt.Sprintf("%s:%s", c.MerchantID, digestData)
	output := b64.URLEncoding.EncodeToString([]byte(basicData))
	return output, dateNow
}

func (c *cicil) GetCheckoutURL(param CheckoutRequest) (resp CheckoutResponse, err error) {
	c.info().Println("Starting Get Order URL Cicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			c.error().Println(err.Error())
		}
	}()
	var getCheckoutURLResponse CheckoutResponse

	// get auth data
	getDiggest, getDate := c.GetToken()

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Date"] = getDate
	headers["Authorization"] = fmt.Sprintf("Basic %s", getDiggest)

	pathGetOrderURL := "po"
	//Marshal Order
	payload, errPayload := json.Marshal(param)
	if errPayload != nil {
		return getCheckoutURLResponse, err
	}

	err = c.call("POST", pathGetOrderURL, bytes.NewBuffer(payload), &getCheckoutURLResponse, headers)
	if err != nil {
		return getCheckoutURLResponse, err
	}
	if len(getCheckoutURLResponse.Message) > 0 {
		err = errors.New(getCheckoutURLResponse.Message)
		return getCheckoutURLResponse, err
	}

	return getCheckoutURLResponse, nil
}

func (c *cicil) SetCancelOrder (param CancelOrderRequest) (resp CancelOrderResponse, err error) {
	c.info().Println("Starting Set Cancel Order Cicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			c.error().Println(err.Error())
		}
	}()
	var setCancelOrderResponse CancelOrderResponse

	// get auth data
	getDiggest, getDate := c.GetToken()

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Date"] = getDate
	headers["Authorization"] = fmt.Sprintf("Basic %s", getDiggest)

	pathGetOrderURL := "po/cancel"
	//Marshal Order
	payload, errPayload := json.Marshal(param)
	if errPayload != nil {
		return setCancelOrderResponse, err
	}

	err = c.call("POST", pathGetOrderURL, bytes.NewBuffer(payload), &setCancelOrderResponse, headers)
	if err != nil {
		return setCancelOrderResponse, err
	}
	if len(setCancelOrderResponse.Message) > 0 {
		err = errors.New(setCancelOrderResponse.Message)
		return setCancelOrderResponse, err
	}

	return setCancelOrderResponse, nil
}

func (c *cicil) UpdateStatus (param UpdateStatusRequest) (resp UpdateStatusResponse, err error) {
	c.info().Println("Starting Set Cancel Order Cicil")
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic: %v", r)
		}
		if err != nil {
			c.error().Println(err.Error())
		}
	}()
	var updateStatusResponse UpdateStatusResponse

	// get auth data
	getDiggest, getDate := c.GetToken()

	// set header
	headers := make(map[string]string)
	headers["Content-Type"] = "application/json"
	headers["Date"] = getDate
	headers["Authorization"] = fmt.Sprintf("Basic %s", getDiggest)

	pathGetOrderURL := "po/update"
	//Marshal Order
	payload, errPayload := json.Marshal(param)
	if errPayload != nil {
		return updateStatusResponse, err
	}

	err = c.call("POST", pathGetOrderURL, bytes.NewBuffer(payload), &updateStatusResponse, headers)
	if err != nil {
		return updateStatusResponse, err
	}
	if len(updateStatusResponse.Message) > 0 {
		err = errors.New(updateStatusResponse.Message)
		return updateStatusResponse, err
	}

	return updateStatusResponse, nil
}