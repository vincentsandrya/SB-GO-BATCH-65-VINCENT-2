-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS Bioskop (
    ID     SERIAL PRIMARY KEY,
    Nama   VARCHAR(200) NOT NULL,
    Lokasi VARCHAR(500),
    Rating DECIMAL(3,2)
);

-- +migrate StatementEnd