-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE IF NOT EXISTS tbl_company (
    company_id SERIAL,
    company_name VARCHAR(50) NOT NULL,
    code VARCHAR(50) NOT NULL,
    country VARCHAR(50),
    website VARCHAR(50),
    phone VARCHAR(50),
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    PRIMARY KEY (company_id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS tbl_company;
