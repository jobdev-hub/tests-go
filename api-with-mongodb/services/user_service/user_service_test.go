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

	err := user_service.Create(user)
	if err != nil {
		t.Error("Erro de persitência", err)
		t.Fail()
	} else {
		t.Log("Usuário criado com sucesso")
	}

}

func TestRead(t *testing.T) {

	users, err := user_service.Read()
	if err != nil {
		t.Error("Erro de leitura", err)
		t.Fail()
	}
	if len(users) == 0 {
		t.Log("Nenhum usuário encontrado")
	} else {
		t.Log("Usuário(s) encontrado(s): ", len(users))
	}

}

func TestUpdate(t *testing.T) {

	user := models.User{
		Name:  "Nome Teste Atualizado",
		Email: "email@atualizado.com",
	}

	err := user_service.Update(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
	}

}

func TestDelete(t *testing.T) {

	err := user_service.Delete(userId)
	if err != nil {
		t.Error("Erro de exclusão", err)
		t.Fail()
	} else {
		t.Log("Usuário excluído com sucesso")
	}

}
