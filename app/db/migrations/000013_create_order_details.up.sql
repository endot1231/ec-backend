CREATE TABLE order_details (
  id BIGINT PRIMARY KEY,
  order_id BIGINT,
  product_stock_id BIGINT,
  price BIGINT,
  quantity BIGINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (product_stock_id) REFERENCES product_stocks(id)
);