package datastore

import (
	"auth/model"
	"fmt"
)

type Datastore interface {
	AddUser(user *model.User) error
	FindUser(email, password string) (map[string]interface{}, error)
	Close()
	FindUsers() []model.User
	FindUserById(id string) *model.User
	AddPost(id uint, todo *model.Todo) error
	GetPosts(id uint) []*model.Todo
	GetPostByPriority(id uint, priority string) []*model.Todo
}

func NewDatastore(dataStoreType string) Datastore {
	if dataStoreType == "postgres" {
		fmt.Println("Connecting to a postgres database...")
		return createPostgresDatastore()
	}
	return nil
}
