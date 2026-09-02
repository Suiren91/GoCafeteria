-- id, name, description, price, stock, createdat, updatedat
CREATE TABLE menus (
    id bigint GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name text NOT NULL CONSTRAINT menus_name_len CHECK (
        char_length(name) BETWEEN 1 AND 255
    ),
    description text NOT NULL DEFAULT '',
    price_yen integer NOT NULL CONSTRAINT menus_price_non_negative CHECK (
        price_yen >= 0
    ),
    stock integer NOT NULL CONSTRAINT stock_price_non_negative CHECK (
        stock >= 0
    ),
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);
COMMENT ON COLUMN menus.price_yen IS '税抜き価格(円)';
COMMENT ON COLUMN menus.stock IS '当日提供可能な数'
