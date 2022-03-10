package user_service_test

import (
	"api-with-mongodb/models"
	"api-with-mongodb/services/user_service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
	"time"
)

var userId string

func TestInsertOne(t *testing.T) {

	oid := primitive.NewObjectID()
	userId = oid.Hex()

	user := models.User{
		ID:    oid,
		Name:  "TestInsertOne",
		Email: "TestInsertOne@test.com",
		Roles: []string{
			"TestInsertOne",
		},
		CreatedAt: time.Now(),
	}

	err := user_service.InsertOne(user)
	if err != nil {
		t.Error("Erro de persitência", err)
		t.Fail()
	} else {
		t.Log("Usuário criado com sucesso")
		t.Log(user.ID)
		userId = user.ID.Hex()
	}

}

func TestFindMany(t *testing.T) {

	users, err := user_service.FindMany()

	if err != nil {
		t.Error("Erro de leitura", err)
		t.Fail()
		return
	}
	if len(users) == 0 {
		t.Log("Nenhum usuário encontrado")
		return
	}

	t.Log("Usuário(s) encontrado(s): ", len(users))
}

func TestFindOneByID(t *testing.T) {

	users, err := user_service.FindMany()
	if err != nil {
		t.Error("Erro de leitura", err)
		t.Fail()
		return
	}
	if len(users) == 0 {
		t.Log("Nenhum usuário encontrado")
		return
	}

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

func TestUpdateOne_NameField(t *testing.T) {

	user := models.User{
		Name: "TestUpdateOne_NameField",
	}

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
		t.Log(userId)
	}

}

func TestUpdateOne_EmailField(t *testing.T) {

	user := models.User{
		Email: "TestUpdateOne_NameField@test.com",
	}

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
		t.Log(userId)
	}

}

func TestUpdateOne_RolesField(t *testing.T) {

	user := models.User{
		Roles: []string{
			"TestUpdateOne_RolesField",
		},
	}

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
		t.Log(userId)
	}

}

func TestUpdateOne_ActiveField(t *testing.T) {

	user := models.User{
		Active: new(bool),
	}
	*user.Active = true

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
		t.Log(userId)
	}

}

func TestUpdateOne_AllFieldsEditableByClient(t *testing.T) {

	user := models.User{
		Name:  "TestUpdateOne_AllFieldsEditableByClient",
		Email: "TestUpdateOne_AllFieldsEditableByClient@test.com",
		Roles: []string{
			"TestUpdateOne_AllFieldsEditableByClient",
		},
		Active:    new(bool),
		UpdatedAt: new(time.Time),
	}
	*user.Active = true
	*user.UpdatedAt = time.Now()

	err := user_service.UpdateOne(user, userId)
	if err != nil {
		t.Error("Erro de atualização", err)
		t.Fail()
	} else {
		t.Log("Usuário atualizado com sucesso")
		t.Log(userId)
	}

}

func TestDeleteOne(t *testing.T) {

	err := user_service.DeleteOne(userId)
	if err != nil {
		t.Error("Erro de exclusão", err)
		t.Fail()
	} else {
		t.Log("Usuário excluído com sucesso")
		t.Log(userId)
	}

}
