CREATE TABLE shop_users (
  id BIGINT PRIMARY KEY,
  shop_id BIGINT,
  name VARCHAR(255),
  email VARCHAR(255),
  email_verified_at TIMESTAMP,
  password VARCHAR(255),
  remember_token VARCHAR(100),
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  deleted_at TIMESTAMP,
  FOREIGN KEY (shop_id) REFERENCES shops(id)
);