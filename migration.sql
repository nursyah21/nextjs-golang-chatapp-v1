.open database.db

-- Create the users table
DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username VARCHAR(60) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    avatar VARCHAR(255),
    createdAt DATE NOT NULL,
    updatedAt DATE NOT NULL
);

-- Create the messages table
DROP TABLE IF EXISTS messages;
CREATE TABLE messages (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    send_id INTEGER NOT NULL,
    receive_id INTEGER NOT NULL,
    message VARCHAR(512) NOT NULL,
    hide BOOLEAN DEFAULT FALSE,
    file_link VARCHAR(255),
    room BOOLEAN DEFAULT FALSE,
    createdAt DATE NOT NULL,
    updatedAt DATE NOT NULL
);

-- Create the refresh_tokens table
DROP TABLE IF EXISTS refresh_tokens;
CREATE TABLE refresh_tokens (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    token VARCHAR(255) NOT NULL
);