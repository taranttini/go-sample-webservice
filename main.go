package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

var productList []Product

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

func init() {
	jsonProductList := `[
		{
			"productId": 1,
			"manufacturer": "Big Box Company",
			"sku": "4561qHJK",
			"upc": "414654444566",
			"pricePerUnit": "12.99",
			"quantityOnHand": 28,
			"productName": "Gizmo"
		},
		{
			"productId": 2,
			"manufacturer": "Small Company",
			"sku": "4561qABC",
			"upc": "414654444533",
			"pricePerUnit": "2.99",
			"quantityOnHand": 5,
			"productName": "Pencil"
		},
		{
			"productId": 3,
			"manufacturer": "Medium Company",
			"sku": "4561qZYZ",
			"upc": "414654444511",
			"pricePerUnit": "5.98",
			"quantityOnHand": 13,
			"productName": "Photo"
		}
	]`
	err := json.Unmarshal([]byte(jsonProductList), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}
	return highestID + 1
}

func findProductByID(productID int) (*Product, int) {
	for i, product := range productList {
		if product.ProductID == productID {
			return &product, i
		}
	}
	return nil, 0
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJson, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJson)
	case http.MethodPost:
		// add a new product to the list
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	default:
		w.Write([]byte("method not defined!!!"))
	}
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, "products/")
	productID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	product, listItemIndex := findProductByID(productID)
	if product == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		// return single product
		productJSON, err := json.Marshal(product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "applicaction/json")
		w.Write(productJSON)

	case http.MethodPut:
		// update product in the list
		var updatedProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &updatedProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if updatedProduct.ProductID != productID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		product = &updatedProduct
		productList[listItemIndex] = *product
		w.WriteHeader(http.StatusOK)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}

}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called!!!"})
	http.HandleFunc("/bar", barHandler)
	http.HandleFunc("/product-to-json", productToJsonHandler)
	http.HandleFunc("/json-to-product", jsonToProductHandler)
	http.HandleFunc("/products", productsHandler)
	http.HandleFunc("/products/", productHandler)
	http.ListenAndServe(":5000", nil)
}
