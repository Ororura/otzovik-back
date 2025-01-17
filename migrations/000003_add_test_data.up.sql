-- Insert test users
INSERT INTO users (username, email, password)
VALUES 
    ('john_doe', 'john@example.com', '$2a$10$1234567890123456789012'),
    ('jane_doe', 'jane@example.com', '$2a$10$1234567890123456789012'),
    ('test_user', 'test@example.com', '$2a$10$1234567890123456789012');

-- Insert test reviews
INSERT INTO reviews (title, content, rating, user_id)
VALUES
    ('Great Product', 'This product exceeded my expectations!', 5, 1),
    ('Excellent Service', 'The customer service was outstanding.', 1, 1),
    ('Could be better', 'There is room for improvement.', 2, 1),
    ('Amazing Experience', 'I would definitely recommend this!', 10, 1);