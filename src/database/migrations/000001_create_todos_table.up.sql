CREATE TABLE todos (
    id varchar(36) PRIMARY KEY,
    title TEXT NOT NULL,
    is_completed TINYINT NOT NULL, 
    created_at TIMESTAMP NOT NULL
);
