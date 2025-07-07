CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY NOT NULL,
    username VARCHAR(50) NOT NULL,
    email VARCHAR(100) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
INSERT IGNORE INTO users (username, email) VALUES
('john_doe', 'john_doe@test.com'),
('jane_doe', 'jane_doe@test.com');
CREATE TABLE IF NOT EXISTS tasks (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status VARCHAR(50),
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
INSERT INTO tasks (title, description, status, user_id) VALUES
('Task 1', 'Description for Task 1', 'pending', 1),
('Task 2', 'Description for Task 2', 'in-progress', 2),
('Task 3', 'Description for Task 3', 'completed', 1);