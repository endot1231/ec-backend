CREATE TABLE user_favorites (
  id BIGINT PRIMARY KEY,
  user_id BIGINT,
  product_stock_id BIGINT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (product_stock_id) REFERENCES product_stocks(id)
);
