-- Insert test users
INSERT INTO users (username, email, password)
VALUES ('john_doe', 'john@example.com', '$2a$10$1234567890123456789012'),
       ('jane_doe', 'jane@example.com', '$2a$10$1234567890123456789012'),
       ('test_user', 'test@example.com', '$2a$10$1234567890123456789012');

-- Insert test reviews (используем category_id=1 для 'Products')
INSERT INTO reviews (title, content, rating, category_id, user_id)
VALUES ('Great Product', 'This product exceeded my expectations!', 5, 1, 1),
       ('Excellent Service', 'The customer service was outstanding.', 8, 2, 1),
       ('Could be better', 'There is room for improvement.', 6, 1, 2),
       ('Amazing Experience', 'I would definitely recommend this!', 10, 3, 3);