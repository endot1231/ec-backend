CREATE TABLE shipments (
  id BIGINT PRIMARY KEY,
  order_id BIGINT,
  order_detail_id BIGINT,
  shipment_date TIMESTAMP,
  created_at TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (order_id) REFERENCES orders(id),
  FOREIGN KEY (order_detail_id) REFERENCES order_details(id)
);
