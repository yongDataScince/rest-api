CREATE TABLE users 
(
    id SERIAL NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);

CREATE TABLE todo_list
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_list
(
    id SERIAL NOT NULL UNIQUE,
    user_id int references users(id) on delete cascade not null,
    todo_id int references todo_list(id) on delete cascade not null
);

CREATE TABLE todo_items
(
    id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done boolean not null default false
);

CREATE TABLE lists_items
(
    id SERIAL NOT NULL UNIQUE,
    user_id int references todo_items(id) on delete cascade not null,
    todo_id int references todo_list(id) on delete cascade not null
);


