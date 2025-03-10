CREATE TABLE books (
    book_id VARCHAR(36) PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    author VARCHAR(255) NOT NULL,
    release_date DATE NOT NULL,
    price DECIMAL NOT NULL CHECK(price > 1000)
);