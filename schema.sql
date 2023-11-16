--Схема БД для сайта новостей.

DROP TABLE IF EXISTS posts, authors;

-- авторы статей
CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- статьи
CREATE TABLE posts (
    id SERIAL PRIMARY KEY,
    author_id INTEGER REFERENCES authors(id) NOT NULL,
    title TEXT  NOT NULL,
    content TEXT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);

-- наполнение БД начальными данными
INSERT INTO authors (id, name) VALUES (0, 'Автор не указан');