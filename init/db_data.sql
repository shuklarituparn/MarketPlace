CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    CONSTRAINT users_password_check CHECK (LENGTH(password) >= 8)
);

CREATE TABLE ads (
   id SERIAL PRIMARY KEY,
   user_id INTEGER REFERENCES users(id),
   title VARCHAR(100) UNIQUE NOT NULL,
   ad_text VARCHAR(500) NOT NULL,
   image_address VARCHAR(255) NOT NULL,
   price NUMERIC NOT NULL
);

INSERT INTO users (email, password)
VALUES
    ('user1@example.com', '$2a$14$51yspoPT/KCJtkl002RYNOmlpqctFdoGXpxgXlVulaPuo/1PN..2.'),
    ('user2@example.com', '$2a$14$Kxtyd8B3uS4WcC0nAOntM.MyN461l6itLjb3Rkco7xc47RgQ7bJE2'),
    ('admin@example.com', '$2a$14$KM/hkTUGDfIXzfaI3IbAXOA0YhAWfgTkQbDJZrd9tJEcUzzD2fUZG');

INSERT INTO ads (user_id, title, ad_text, image_address, price)
VALUES
    (1, 'Car for sale', 'Selling my car, good condition', 'https://example.com/car.jpg', 10000.00),
    (1, 'Furniture for sale', 'Selling furniture, great condition', 'https://example.com/furniture.jpg', 500.00),
    (2, 'Electronics for sale', 'Various electronics for sale', 'https://example.com/electronics.jpg', 2000.00),
    (3, 'Luxury watch for sale', 'Rare luxury watch for sale', 'https://example.com/watch.jpg', 5000.00);
