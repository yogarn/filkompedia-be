CREATE TABLE users (
    id varchar(36) NOT NULL PRIMARY KEY,
    username varchar(36) NOT NULL,
    email varchar(255),
    password varchar(60),
    role_id integer NOT NULL DEFAULT 0
);

CREATE TABLE roles (
    id INTEGER NOT NULL PRIMARY KEY,
    role varchar(36) NOT NULL
);

INSERT INTO roles (id, role) VALUES 
(0, 'user'),
(1, 'admin');

CREATE TABLE sessions (
    user_id varchar(36),
    token varchar(255),
    ip_address varchar(16),
    expires_at timestamp,
    user_agent varchar(255),
    device_id varchar(255), 
    FOREIGN KEY (user_id) REFERENCES users(id)
);



