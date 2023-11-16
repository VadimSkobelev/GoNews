package postgres

import (
	"GoNews/pkg/storage"
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

// Хранилище данных.
type Storage struct {
	db *pgxpool.Pool
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	db, err := pgxpool.Connect(context.Background(), constr)
	if err != nil {
		return nil, err
	}
	s := Storage{
		db: db,
	}
	return &s, nil
}

// Posts возвращает список статей из БД.
func (s *Storage) Posts() ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT 
			id,
			title,
			content,
			author_id,
			(SELECT name FROM authors WHERE id=author_id),		
			created_at
		FROM posts
		ORDER BY id;
	`,
	)
	if err != nil {
		return nil, err
	}
	var posts []storage.Post
	// итерирование по результату выполнения запроса
	// и сканирование каждой строки в переменную
	for rows.Next() {
		var p storage.Post
		err = rows.Scan(
			&p.ID,
			&p.Title,
			&p.Content,
			&p.AuthorID,
			&p.AuthorName,
			&p.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		// добавление переменной в массив результатов
		posts = append(posts, p)

	}
	// ВАЖНО не забыть проверить rows.Err()
	return posts, rows.Err()
}

// AddPost добовляет новую статью в базу.
func (s *Storage) AddPost(p storage.Post) error {
	err := s.db.QueryRow(context.Background(), `
		INSERT INTO posts (title, content, author_id)
		VALUES ($1, $2, $3) RETURNING id;
		`,
		p.Title,
		p.Content,
		p.AuthorID,
	).Scan(&p.ID)
	return err
}

// Обновление статьи по id (Title, Content).
func (s *Storage) UpdatePost(p storage.Post) error {
	row := s.db.QueryRow(context.Background(), `
		UPDATE 
			posts
		SET 
			title=$1,
			content=$2
		WHERE 
			id=$3
			RETURNING id;
		`,
		p.Title,
		p.Content,
		p.ID,
	)
	err := row.Scan(
		&p.ID,
	)
	if err != nil {
		return err
	}
	return nil
}

// Удаление статьи по id.
func (s *Storage) DeletePost(p storage.Post) error {
	_, err := s.db.Exec(context.Background(), `DELETE FROM posts WHERE id=$1;`, p.ID)
	if err != nil {
		return err
	}
	return nil
}
