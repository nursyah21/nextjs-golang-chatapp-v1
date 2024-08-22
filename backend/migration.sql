CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);
-- Create the messages table
CREATE TABLE IF NOT EXISTS messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    send_id INTEGER NOT NULL REFERENCES users (id),
    receive_id INTEGER NOT NULL REFERENCES users (id),
    message VARCHAR(512) NOT NULL,
    file_link VARCHAR(255),
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Create the refresh_tokens table
CREATE TABLE IF NOT EXISTS refresh_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL REFERENCES users (id),
    token VARCHAR(255) NOT NULL,
    createdAt DATETIME DEFAULT CURRENT_TIMESTAMP,
    updatedAt DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER IF NOT EXISTS update_users_updated_at
AFTER UPDATE ON users
WHEN old.updatedAt <> current_timestamp
BEGIN
    UPDATE users
    SET updatedAt = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS update_messages_updated_at
AFTER UPDATE ON messages
WHEN old.updatedAt <> current_timestamp
BEGIN
    UPDATE messages
    SET updatedAt = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;

CREATE TRIGGER IF NOT EXISTS update_refresh_tokens_updated_at
AFTER UPDATE ON refresh_tokens
WHEN old.updatedAt <> current_timestamp
BEGIN
    UPDATE refresh_tokens
    SET updatedAt = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;