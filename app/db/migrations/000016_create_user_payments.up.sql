CREATE TABLE payments (
  id BIGINT PRIMARY KEY,
  user_id BIGINT,
  order_id BIGINT,
  amount INT,
  payment_date TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id),
  FOREIGN KEY (order_id) REFERENCES orders(id)
);