CREATE TABLE users (
    id varchar(36) NOT NULL PRIMARY KEY,
    username varchar(36) NOT NULL,
    email varchar(255),
    password varchar(60)
);

CREATE TABLE roles (
    id INTEGER NOT NULL PRIMARY KEY,
    role varchar(36) NOT NULL
);

CREATE TABLE sessions (
    user_id varchar(36),
    token varchar(255),
    ip_address varchar(16),
    user_agent varchar(255),
    device_id varchar(255), 
    FOREIGN KEY (user_id) REFERENCES users(id)
);
