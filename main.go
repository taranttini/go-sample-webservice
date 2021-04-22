package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called!!!"))
}

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func productToJsonHandler(w http.ResponseWriter, r *http.Request) {
	product := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		PricePerUnit:   "12.99",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	productJSON, err := json.Marshal(product)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte(string(productJSON)))
}

func jsonToProductHandler(w http.ResponseWriter, r *http.Request) {
	productJSON := `{
		"productId": 123,
		"manufacturer": "Big Box Company",
		"sku": "4561qHJK",
		"upc": "414654444566",
		"pricePerUnit": "12.99",
		"quantityOnHand": 28,
		"productName": "Gizmo"
	  }`

	product := Product{}
	err := json.Unmarshal([]byte(productJSON), &product)
	if err != nil {
		log.Fatal(err)
	}
	w.Write([]byte("json unmarshal product: " + product.ProductName))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called!!!"})
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/product-to-json", productToJsonHandler)
	http.HandleFunc("/json-to-product", jsonToProductHandler)
	http.ListenAndServe(":5000", nil)
}
