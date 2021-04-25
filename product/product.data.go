package product

import (
	"context"
	"database/sql"
	"inventoryservice/database"
	"log"
	"time"
)

func getProduct(productID int) (*Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	row := database.DbConn.QueryRowContext(ctx, `
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
		log.Panicln(err)
		return nil, err
	}
	return product, nil
}

func GetTopTenProducts() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `
		SELECT productId,
			manufacturer,
			sku,
			upc,
			pricePerUnit,
			quantityOnHand,
			productName
		FROM products
		ORDER BY quantityOnHand DESC LIMIT 10
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

func removeProduct(productID int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `DELETE FROM products WHERE productId = ?`, productID)
	if err != nil {
		return err
	}
	return nil
}

func getProductList() ([]Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	results, err := database.DbConn.QueryContext(ctx, `
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, err := database.DbConn.ExecContext(ctx, `
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
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	result, err := database.DbConn.ExecContext(ctx, `
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
