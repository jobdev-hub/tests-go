package user_service_test

import (
	"api-with-mongodb/models"
	"api-with-mongodb/services/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestCreate(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := models.User{
		ID:        oid,
		Name:      "Nome Teste",
		Email:     "email@teste.com",
		CreatedAt: time.Now(),
	}

	err := user_service.InsertOne(user)
	if err != nil {
		t.Error("Erro de persitência", err)
		t.Fail()
	} else {
		t.Log("Usuário criado com sucesso")
	}

}

func TestRead(t *testing.T) {

	users, err := user_service.FindMany()
	if HandleRead(t, err, users) {
		return
	}

	t.Log("Usuário(s) encontrado(s): ", len(users))
}

func TestReadByID(t *testing.T) {

	users, err := user_service.FindMany()
	if HandleRead(t, err, users) {
		return
	}

	userId = users[0].ID.Hex()

	user, err := user_service.FindOneByID(userId)
	if err != nil {
		t.Error("Erro de leitura", err)
		t.Fail()
		return
	}

	if user.ID.Hex() != userId {
		t.Error("Usuário não encontrado")
		t.Fail()
		return
	}

	t.Log("Usuário encontrado: ", user)
}

func TestUpdate(t *testing.T) {

	user := models.User{
		Name:  "Nome Teste Atualizado",
		Email: "email@atualizado.com",
	}

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
	}

}

func TestDelete(t *testing.T) {

	err := user_service.DeleteOne(userId)
	if err != nil {
		t.Error("Erro de exclusão", err)
		t.Fail()
	} else {
		t.Log("Usuário excluído com sucesso")
	}

}

func HandleRead(t *testing.T, err error, users models.Users) bool {
	if err != nil {
		t.Error("Erro de leitura", err)
		t.Fail()
		return true
	}
	if len(users) == 0 {
		t.Log("Nenhum usuário encontrado")
		return true
	}
	return false
}
