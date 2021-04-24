package product

import (
	"database/sql"
	"inventoryservice/database"
)

func getProduct(productID int) (*Product, error) {
	row := database.DbConn.QueryRow(`
		SELECT productId,
			manufacturer,
			sku,
			upc,
			pricePerUnit,
			quantityOnHand,
			productName
		FROM products
		WHERE productId = ?
	`, productID)

	product := &Product{}
	err := row.Scan(
		&product.ProductID,
		&product.Manufacturer,
		&product.Sku,
		&product.Upc,
		&product.PricePerUnit,
		&product.QuantityOnHand,
		&product.ProductName)
	if err == sql.ErrNoRows {

		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return product, nil
}

func removeProduct(productID int) error {
	_, err := database.DbConn.Query(`DELETE FROM products WHERE productId = ?`, productID)
	if err != nil {
		return err
	}
	return nil
}

func getProductList() ([]Product, error) {
	results, err := database.DbConn.Query(`
		SELECT productId,
			manufacturer,
			sku,
			upc,
			pricePerUnit,
			quantityOnHand,
			productName
		FROM products
	`)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	products := make([]Product, 0)

	for results.Next() {
		var product Product
		results.Scan(
			&product.ProductID,
			&product.Manufacturer,
			&product.Sku,
			&product.Upc,
			&product.PricePerUnit,
			&product.QuantityOnHand,
			&product.ProductName)
		products = append(products, product)
	}
	return products, nil
}

func updateProduct(product Product) error {
	_, err := database.DbConn.Exec(`
		UPDATE products SET
			manufacturer=?,
			sku=?,
			upc=?,
			pricePerUnit=CAST(? AS DECIMAL(13,2)),
			quantityOnHand=?,
			productName=?
		WHERE productId=?`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
		product.ProductID,
	)
	if err != nil {
		return err
	}
	return nil
}

func insertProduct(product Product) (int, error) {
	result, err := database.DbConn.Exec(`
		INSERT INTO products (
			manufacturer,
			sku,
			upc,
			pricePerUnit,
			quantityOnHand,
			productName
		) VALUES (?,?,?,?,?,?)`,
		product.Manufacturer,
		product.Sku,
		product.Upc,
		product.PricePerUnit,
		product.QuantityOnHand,
		product.ProductName,
	)
	if err != nil {
		return 0, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return 0, nil
	}
	return int(insertID), nil
}
