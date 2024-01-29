CREATE DATABASE clubone CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

USE clubone;

CREATE TABLE IF NOT EXISTS wikis
(
    id INT AUTO_INCREMENT PRIMARY KEY,
    creator INT(11),
    create_time DATETIME,
    update_time DATETIME,
    title VARCHAR(256),
    content TEXT,
    uuid CHAR(36)
) CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci;

CREATE INDEX idx_wikis_uuid ON wikis(uuid);
