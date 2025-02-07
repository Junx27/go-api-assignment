package main

import (
	"bytes"
	"encoding/json"
	"go-api-assignment/handler"
	"go-api-assignment/mock"
	"go-api-assignment/model"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddUserHandler(t *testing.T) {
	mockRepo := mock.NewMockUserRepository()
	handler := handler.NewUserHandler(mockRepo)

	user := model.User{ID: 1, Name: "John", Age: 25}
	jsonData, _ := json.Marshal(user)

	req, err := http.NewRequest("POST", "/add-user", bytes.NewReader(jsonData))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.AddUserHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var responseUser model.User
	json.NewDecoder(rr.Body).Decode(&responseUser)
	assert.Equal(t, user, responseUser)
}

func TestGetUserByIDHandler(t *testing.T) {
	mockRepo := mock.NewMockUserRepository()
	handler := handler.NewUserHandler(mockRepo)

	user := model.User{ID: 1, Name: "John", Age: 25}
	mockRepo.AddUser(user)

	req, err := http.NewRequest("GET", "/get-user?id=1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler.GetUserByIDHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var responseUser model.User
	json.NewDecoder(rr.Body).Decode(&responseUser)
	assert.Equal(t, user, responseUser)
}
