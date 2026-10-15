BEGIN;

---items---

CREATE TABLE items (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,
    stock BIGINT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_items_price ON items (price);
CREATE INDEX idx_items_name  ON items (name);

---orders---

CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    total BIGINT NOT NULL,
    status VARCHAR(20) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_orders_user_id    ON orders (user_id);
CREATE INDEX idx_orders_status     ON orders (status);
CREATE INDEX idx_orders_created_at ON orders (created_at DESC);

---order_items---

CREATE TABLE order_items (
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL REFERENCES orders(id),
    item_id BIGINT NOT NULL REFERENCES items(id),
    price BIGINT NOT NULL CHECK (price >= 0),
    quantity BIGINT NOT NULL CHECK (quantity >= 0)
)

CREATE INDEX idx_line_items_order_id ON line_items (order_id);
CREATE INDEX idx_line_items_item_id  ON line_items (item_id);

CREATE TABLE carts (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    item_id BIGINT NOT NULL REFERENCES items(id),
    quantity BIGINT NOT NULL DEFAULT 1 CHECK (quantity > 0),
    UNIQUE (user_id, item_id)
);

CREATE INDEX idx_user_id_item_id ON carts (user_id, item_id);

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL UNIQUE,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
);