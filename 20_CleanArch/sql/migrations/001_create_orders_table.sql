-- Migration: Create orders table
-- Description: Tabela para armazenar pedidos do sistema

CREATE TABLE IF NOT EXISTS orders (
    id VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL,
    tax FLOAT NOT NULL,
    final_price FLOAT NOT NULL,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

