CREATE TABLE carts (
    id VARCHAR(36) PRIMARY KEY,
    user_id varchar(36) NOT NULL,
    book_id VARCHAR(36) NOT NULL,
    amount INTEGER NOT NULL CHECK(amount > 0),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (book_id) REFERENCES books(book_id)
);