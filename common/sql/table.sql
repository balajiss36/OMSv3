CREATE TABLE orders (
    ID TEXT PRIMARY KEY,
    CustomerID TEXT,
    Status TEXT,
    PaymentLink TEXT
);

CREATE TABLE items (
    ID TEXT SERIAL PRIMARY KEY,
    PriceID TEXT REFERENCES orders(ID),
    Name TEXT,
    Quantity INT
);