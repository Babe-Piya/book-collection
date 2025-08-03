CREATE TABLE IF NOT EXISTS book_collection (
    id              SERIAL  PRIMARY KEY,
    book_name       text,
    type            VARCHAR(100),
    volume          INT,
    price           DECIMAL(10, 2),
    created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
