CREATE TABLE orders (
  id BIGINT PRIMARY KEY,
  shop_id BIGINT,
  user_id BIGINT,
  order_date TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (shop_id) REFERENCES shops(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);