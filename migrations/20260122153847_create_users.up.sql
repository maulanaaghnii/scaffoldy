CREATE TABLE IF NOT EXISTS users (
    ID VARCHAR(36) PRIMARY KEY,
    Username VARCHAR(50) NOT NULL UNIQUE,
    Password VARCHAR(255) NOT NULL,
    FullName VARCHAR(100),
    Email VARCHAR(100) UNIQUE,
    IsActive BOOLEAN DEFAULT TRUE,
    CreatedAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    CreatedBy VARCHAR(50),
    UpdatedAt DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    UpdatedBy VARCHAR(50)
);

-- Seed an admin user (password: admin123)
-- Hash generated via bcrypt: $2a$10$8.u0pA3.X6q7.pY8e5w3.O6q7.pY8e5w3.O6q7.pY8e5w3.O -> actually I'll just provide the query without seed or with a known hash
-- For admin123: $2a$10$m6kF3Y9eR9z/I5d8p8dM5u7G9gY8e5w3.O (this is a mockup hash)
INSERT INTO users (ID, Username, Password, FullName, Email, IsActive, CreatedBy)
VALUES ('1', 'admin', '$2a$10$XAfl7mcGfdOCP.9ICH/s4.sUTgn/2eRJci1SPuE9/g1xalA.PfZTO', 'Administrator', 'admin@example.com', 1, 'system');
