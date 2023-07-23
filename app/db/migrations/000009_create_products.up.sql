CREATE TABLE products (
  id BIGINT PRIMARY KEY,
  shop_id BIGINT,
  product_category_id BIGINT,
  product_brand_id BIGINT,
  name VARCHAR(255),
  description TEXT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (shop_id) REFERENCES shops(id),
  FOREIGN KEY (product_category_id) REFERENCES product_categories(id),
  FOREIGN KEY (product_brand_id) REFERENCES product_brands(id)
);