CREATE TABLE categories
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    created_at  TIMESTAMP DEFAULT NOW(),
    updated_at  TIMESTAMP DEFAULT NOW()
);

-- Добавляем базовые категории
INSERT INTO categories (name) VALUES 
    ('Products'),
    ('Services'),
    ('Restaurants'),
    ('Hotels'),
    ('Other');
