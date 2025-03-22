DROP TABLE IF EXISTS checkouts;

ALTER TABLE carts ADD COLUMN checkout_id UUID;
ALTER TABLE carts ADD FOREIGN KEY (checkout_id) REFERENCES checkouts(id);
