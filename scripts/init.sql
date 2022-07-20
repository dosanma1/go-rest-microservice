CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
	product_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
	name VARCHAR NOT NULL,
	quantity INT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL
);


INSERT INTO products
("name", quantity, created_at, updated_at)
VALUES('sun_cream_30', 10, NOW(), NOW());

INSERT INTO products
("name", quantity, created_at, updated_at)
VALUES('sun_cream_50', 20, NOW(), NOW());

INSERT INTO products
("name", quantity, created_at, updated_at)
VALUES('shampooo', 25, NOW(), NOW());
