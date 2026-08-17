-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    category_id INTEGER NOT NULL,
    description VARCHAR(255),
    image_url VARCHAR(255),
    release_year INTEGER,
    price BIGINT,
    total_page INTEGER,
    thickness VARCHAR(50),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by VARCHAR(50),
    modified_at TIMESTAMP,
    modified_by VARCHAR,

    CONSTRAINT fk_books_category
        FOREIGN KEY (category_id)
        REFERENCES categories(id)
);

-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin

DROP TABLE IF EXISTS books;

-- +migrate StatementEnd