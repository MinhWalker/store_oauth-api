package rest

import (
	"encoding/json"
	"fmt"
	"github.com/MinhWalker/store_oauth-api/src/domain/users"
	"github.com/MinhWalker/store_oauth-api/src/utils/errors"
	"gopkg.in/go-resty/resty.v2"
	"io/ioutil"
	"net/http/httptest"
	"time"
)

var(
	client = resty.New().SetTimeout(100*time.Millisecond).SetHostURL("http://localhost:8080")
)

type RestUserRepository interface {
	LoginUser(string, string) (*users.User, *errors.RestErr)
	LoginUserTest(*httptest.ResponseRecorder) (*users.User, *errors.RestErr)
}

type usersRepository struct {
}

func NewRepository() RestUserRepository {
	return &usersRepository{}
}

func (u *usersRepository) LoginUser(email, password string) (*users.User, *errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}

	response, _ := client.R().SetBody(request).Post("/users/login")

	if response == nil || response.Result() == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user!")
	}
	if response.StatusCode() >= 300 {
		fmt.Println(response.String())
		var restErr errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user!")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}

//TODO: testing
func (u *usersRepository) LoginUserTest(res *httptest.ResponseRecorder) (*users.User, *errors.RestErr) {
	response := res.Result()
	body, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(body))

	if response == nil || body == nil {
		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user!")
	}
	if response.StatusCode >= 300 {
		fmt.Println(response.Body)
		var restErr errors.RestErr
		err := json.Unmarshal(body, &restErr)
		if err != nil {
			return nil, errors.NewInternalServerError("invalid error interface when trying to login user!")
		}
		return nil, &restErr
	}

	var user users.User
	if err := json.Unmarshal(body, &user); err != nil {
		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
	}
	return &user, nil
}

