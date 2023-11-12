CREATE TABLE
    todos (
        id BIGINT AUTO_INCREMENT NOT NULL,
        title VARCHAR(255) NOT NULL,
        is_completed TINYINT NOT NULL DEFAULT 0,
        created_at TIMESTAMP NOT NULL,
        PRIMARY KEY (id)
    );