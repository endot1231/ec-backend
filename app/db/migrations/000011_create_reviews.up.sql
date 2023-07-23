CREATE TABLE reviews (
  id BIGINT PRIMARY KEY,
  product_id BIGINT,
  user_id BIGINT,
  contents VARCHAR(255),
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (user_id) REFERENCES users(id)
);