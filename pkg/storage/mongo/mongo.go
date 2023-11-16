package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	databaseName   = "data"
	collectionName = "posts"
)

// Хранилище данных.
type Storage struct {
	db *mongo.Client
}

// Конструктор, принимает строку подключения к БД.
func New(constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	s := Storage{
		db: client,
	}
	return &s, nil
}

// AddPost добовляет новую статью в базу.
func (s *Storage) AddPost(p storage.Post) error {
	data := []interface{}{p} // Конвертируем структуру в массив
	collection := s.db.Database(databaseName).Collection(collectionName)
	for _, post := range data {
		_, err := collection.InsertOne(context.Background(), post)
		if err != nil {
			return err
		}
	}
	return nil
}

// Posts возвращает список статей из БД.
func (s *Storage) Posts() ([]storage.Post, error) {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())
	var data []storage.Post
	for cur.Next(context.Background()) {
		var l storage.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

// Обновление статьи по id (Title, Content).
func (s *Storage) UpdatePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.M{"id": bson.M{"$eq": p.ID}}
	update := bson.M{"$set": bson.M{"title": p.Title, "content": p.Content}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

// Удаление статьи по id.
func (s *Storage) DeletePost(p storage.Post) error {
	collection := s.db.Database(databaseName).Collection(collectionName)
	filter := bson.M{"id": bson.M{"$eq": p.ID}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
