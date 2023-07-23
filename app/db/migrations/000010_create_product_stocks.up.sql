CREATE TABLE product_stocks (
  id BIGINT PRIMARY KEY,
  product_id BIGINT,
  product_color_id BIGINT,
  product_size_id BIGINT,
  description VARCHAR(255),
  price INT,
  stock_quantity INT,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (product_id) REFERENCES products(id),
  FOREIGN KEY (product_color_id) REFERENCES product_colors(id),
  FOREIGN KEY (product_size_id) REFERENCES product_sizes(id)
);
