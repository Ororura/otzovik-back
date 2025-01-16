CREATE TABLE reviews
(
    id          SERIAL PRIMARY KEY,
    text        TEXT NOT NULL,
    date_create TIMESTAMP DEFAULT NOW(),
    date_update TIMESTAMP DEFAULT NOW(),
    author_id   INT,
    CONSTRAINT fk_author_id FOREIGN KEY (author_id) REFERENCES users (id)

);

INSERT INTO reviews (id, text, author_id)
VALUES (1, 'IAUShdiuahsgd', 1)