-- @block
CREATE TABLE Users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    uuid VARCHAR(255) NOT NULL UNIQUE,
    refresh_token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- @block
INSERT INTO users (email, password, uuid)
VALUES(
    'test@gmail.com',
    'testpassword123',
    'uidasfiausdf'
)

-- @block
SELECT * FROM users;

-- @block
DROP TABLE users;

-- @block
CREATE TABLE Quotes(
    id INT AUTO_INCREMENT,
    quote VARCHAR(255),
    author VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    owner_uuid VARCHAR(255) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (owner_uuid) REFERENCES Users(uuid)
)

-- @block
DROP TABLE Quotes;

-- @block
INSERT INTO Quotes (owner_uuid, quote, author)
VALUES
    ('uidasfiausdf', 'some random quote', 'random'),
    ('uidasfiausdf', 'some fasd quote', 'random'),
    ('uidasfiausdf', 'some dfsaf quote', 'random');

-- @block
SELECT * FROM Users
INNER JOIN Quotes
ON Quotes.owner_uuid = Users.uuid;

-- @block
SELECT quote, author, created_at FROM quotes WHERE quotes.owner_uuid = "uidasfiausdf";

-- @block
UPDATE users SET refresh_token = "test" WHERE email = "ewfwfsdf@gmail.com";