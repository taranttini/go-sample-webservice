-- show databases
CREATE DATABASE `inventorydb`;

CREATE TABLE `inventorydb`.`products` (
    `productId` INT NOT NULL AUTO_INCREMENT,
    `manufacturer` VARCHAR(255) NOT NULL,
    `sku` VARCHAR(60) NOT NULL,
    `upc` VARCHAR(60) NOT NULL,
    `pricePerUnit` DECIMAL(13,2) NOT NULL,
    `quantityOnHand` INT NOT NULL,
    `productName` VARCHAR(255) NOT NULL,
    PRIMARY KEY (`productId`)
);

SELECT * FROM inventorydb.products