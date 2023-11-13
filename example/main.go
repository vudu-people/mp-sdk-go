package main

import (
	"fmt"

	mpsdkgo "github.com/vudu-people/mp-sdk-go"
)

func main() {

	cfg := mpsdkgo.NewConfiguration(mpsdkgo.WithDebug(false), mpsdkgo.WithToken("TEST-11233"))

	client := mpsdkgo.NewApiClient(cfg)

	stores, err := client.Store.GetStores("12345", 10)
	if err != nil {
		panic(err)
	}

	fmt.Println(stores.Paging.Total)

	for _, store := range stores.Results {
		fmt.Println(store.Name)
	}

}
