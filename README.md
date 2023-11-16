# Go API client for mpsdk

Mercadopago api SDK for golang 


## Example 

```go
func main() {

    // create config
	cfg := mpsdkgo.NewConfiguration(mpsdkgo.WithDebug(true), mpsdkgo.WithToken(os.Getenv("MP_TOKEN")))

    // init clint
	client := mpsdkgo.NewApiClient(cfg)

	// get payment methods
	paymentMethods, err := client.PaymentMethods.GetPaymentMethods()
	if err != nil {
		panic(err)
	}

	fmt.Println(paymentMethods)

	return
```