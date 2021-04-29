package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/MinhWalker/store_oauth-api/src/domain/users"
	"github.com/MinhWalker/store_utils-go/rest_errors"
	"github.com/go-resty/resty/v2"
	"time"
)

var(
	client = resty.New().SetTimeout(100*time.Millisecond)
)

type RestUsersRepository interface {
	LoginUser(string, string) (*users.User, rest_errors.RestErr)
}

type usersRepository struct{
}

func NewRestUsersRepository() RestUsersRepository {
	return &usersRepository{}
}

func (u *usersRepository) LoginUser(email string, password string) (*users.User, rest_errors.RestErr) {
	request := users.UserLoginRequest{
		Email: email,
		Password: password,
	}

	response, _ := client.R().SetBody(request).Post("http://localhost:8081/users/login")

	if response == nil {
		fmt.Println(response)
		return nil, rest_errors.NewInternalServerError("invalid rest client response when trying to login user!", errors.New("restclient error"))
	}
	if response.StatusCode() >= 300 {
		fmt.Println(response.String())
		var restErr rest_errors.RestErr
		err := json.Unmarshal(response.Body(), &restErr)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("invalid error interface when trying to login user!", err)
		}
		return nil, restErr
	}

	var user users.User
	if err := json.Unmarshal(response.Body(), &user); err != nil {
		return nil, rest_errors.NewInternalServerError("error when trying to unmarshal users response", errors.New("json parsing error"))
	}
	return &user, nil
}

//TODO: testing
//func (u *usersRepository) LoginUserTest(res *httptest.ResponseRecorder) (*users.User, *errors.RestErr) {
//	response := res.Result()
//	body, _ := ioutil.ReadAll(response.Body)
//
//	fmt.Println(string(body))
//
//	if response == nil || body == nil {
//		return nil, errors.NewInternalServerError("invalid rest client response when trying to login user!")
//	}
//	if response.StatusCode >= 300 {
//		fmt.Println(response.Body)
//		var restErr errors.RestErr
//		err := json.Unmarshal(body, &restErr)
//		if err != nil {
//			return nil, errors.NewInternalServerError("invalid error interface when trying to login user!")
//		}
//		return nil, &restErr
//	}
//
//	var user users.User
//	if err := json.Unmarshal(body, &user); err != nil {
//		return nil, errors.NewInternalServerError("error when trying to unmarshal users response")
//	}
//	return &user, nil
//}

