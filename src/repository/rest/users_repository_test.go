package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MinhWalker/store_oauth-api/src/domain/users"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	os.Exit(m.Run())
}

func TestLoginUserTimeoutFromApi(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("POST", "http://localhost:8080/users/login",
		httpmock.NewStringResponder(200, `{"error": false, "errorMessage": "", "result": []}`))

	repository := usersRepository{}

	user, err := repository.LoginUser("wewwew1", "12345")

	fmt.Println(err)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user!", err.Message)
}

func TestLoginUserInvalidErrorInterface(t *testing.T) {

}

func TestLoginUserInvalidLoginCredentials(t *testing.T) {

}

func TestLoginUserInvalidUserJsonResponse(t *testing.T) {

}

func TestLoginUserInvalidNoError(t *testing.T) {

}

func TestList_Handle(t *testing.T) {
	request := users.UserLoginRequest{
		Email: "wewwew1",
		Password: "12345",
	}

	response := users.User{
		Id:        1,
		FirstName: "test",
		LastName:  "test",
		Email:     "test",
	}

	reqUser, _ := json.Marshal(request)
	resUser, _ := json.Marshal(response)

	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqUser))
	res := httptest.NewRecorder()
	res.Body.Write(resUser)
	res.Header().Set("X-Request-Id", req.Header.Get("X-Request-Id"))

	repository := usersRepository{}

	user, err := repository.LoginUserTest(res)

	fmt.Println(err)
	assert.Nil(t, user)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "invalid rest client response when trying to login user!", err.Message)
}


