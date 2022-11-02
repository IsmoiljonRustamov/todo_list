CREATE TABLE to_do_list (
    id SERIAL PRIMARY KEY,
    title Varchar(50),
    descriptions text,
    assignee VARCHAR(30),
    status BOOLEAN,
    deadline DATE,
    created_at timestamp DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp,
    deleted_at timestamp 
);