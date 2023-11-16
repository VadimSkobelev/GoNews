package mongo

import (
	"GoNews/pkg/storage"
	"log"
	"os"
	"testing"
)

var s *Storage

func TestMain(m *testing.M) {
	constr := "mongodb://localhost:27017/"
	var err error
	s, err = New(constr)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(m.Run())
}

func TestStorage_AddPost(t *testing.T) {
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{"Create post", s, args{storage.Post{ID: 2, Title: "Unit Test-01", Content: "Testing"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Posts(t *testing.T) {
	tests := []struct {
		name    string
		s       *Storage
		want    []storage.Post
		wantErr bool
	}{
		{"All posts", s, []storage.Post{}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Storage.Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("Storage.Posts() = %v, want %v", got, tt.want)
			// }
			t.Log(got)
		})
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{"Update post with id=2", s, args{storage.Post{ID: 2, Title: "Update test-05 modified", Content: "tests!!!-05 modified"}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		s       *Storage
		args    args
		wantErr bool
	}{
		{"Delete post with id=2", s, args{storage.Post{ID: 2}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("Storage.DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
