-- @block
CREATE TABLE Users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    uuid VARCHAR(255) NOT NULL UNIQUE,
    token VARCHAR(255) UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- @block
INSERT INTO users (email, password)
VALUES(
    'test@gmail.com',
    'testpassword123'
)

-- @block
SELECT * FROM users;

-- @block
DROP TABLE users;