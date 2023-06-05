package modeltests

import (
	"log"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/SyarifuanEfendi/go-test/api/models"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T)  {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
	}
	assert.Equal(t, len(*users), 2)
}

func TestSaveUser(t *testing.T)  {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	newUser := models.User{
		ID: 1,
		Email: "efendi@gmail.com",
		Nickname: "efendi",
		Password: "password",
	}
	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)
}

func TestGetUserByID(t *testing.T)  {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table: %v", err)
	}
	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting one user: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)
}

func TestUpdateAUser(t *testing.T)  {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	userUpdate := models.User{
		ID: 1,
		Nickname: "Syarifuan",
		Email: "Syarifuan@gmail.com",
		Password: "password",
	}
	updateUser, err := userUpdate.UpdateAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
	}
	assert.Equal(t, updateUser.ID, userUpdate.ID)
	assert.Equal(t, updateUser.Email, userUpdate.Email)
	assert.Equal(t, updateUser.Nickname, userUpdate.Nickname)
}

func TestDeleteAUser(t *testing.T) {
	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user: %v\n", err)
	}

	isDeleted, err := userInstance.DeleteAUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
	}

	assert.Equal(t, isDeleted, int16(1))
}