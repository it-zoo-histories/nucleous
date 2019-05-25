package routes

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"nucleous/dao"
	"nucleous/enhancer"
	"nucleous/middlewares"
	"nucleous/models"
	"nucleous/payloads"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

/*UserRoute - маршрут для изменения данных пользователя и его удаления*/
type UserRoute struct {
	TokenDao         *dao.TokenDAO
	UserDao          *dao.UserDAO
	JWTMiddle        *middlewares.JWTChecker
	EResponser       *enhancer.Responser
	UserVerification map[string]map[string]interface{} // верификационные коды пользователей
}

const (
	createUser         = "/user/create"
	confirmAccount     = "/user/confirm"
	findAllUsers       = "/user/all"
	updateUserSettings = "/user"
	removeUserByID     = "/user/remove"
	nameServer         = "[NUCLEOUS: ROUTES]: "
)

/*confirmAccount - подтверждение аккаунта*/
func (route *UserRoute) confirmAccount(w http.ResponseWriter, r *http.Request) {
	userid := r.URL.Query().Get("userid")

	// fmt.Println(nameServer+"Handled userid: ", userid)

	if strings.TrimSpace(userid) == "" {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.ConfirmUser",
			"code":    "userid is empty",
		}, "application/json")
		return
	}

	user, err := route.UserDao.FindUserByUserID(userid)
	if err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.ConfirmUser",
			"code":    "user bd",
			"error":   err.Error(),
		}, "application/json")
		return
	}
	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]interface{}{
		"status":  "success",
		"context": "UserRoute.ConfirmUser",
		"data":    "confirmation mail was sent",
		"timer":   time.Minute * 5,
	}, "application/json")

	go route.sendVerificationMail(user, userid)

	return
}

func (route *UserRoute) sendVerificationMail(user *models.User, userid string) {
	sender := enhancer.NewSender("s6galaxyru@gmail.com", "AppLudKo")
	reciever := []string{user.Email}
	subject := "Verification your account"
	testCode := 0

	if err := route.checkExistCode(userid); err != nil {
		testCode = route.UserVerification[userid]["code"].(int)
	} else {
		testCode = route.generatorVerificationPassPhrase()
		route.UserVerification[userid] = map[string]interface{}{
			"timeCreated":  time.Now(),
			"timeDuration": time.Minute * 5,
			"code":         testCode,
		}
	}

	fmt.Println("Verificated need user: ", route.UserVerification)

	message := `
	<!DOCTYPE HTML PULBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
	<html>
	<head>
	<meta http-equiv="content-type" content="text/html"; charset=ISO-8859-1">
	</head>
	<body>hi, please write this number query in special field of client side: <br>
	<span style="background: black; color: white; padding: 15px; display: block; margin: auto; text-align: center; font-size: 1.9em;">` + strconv.Itoa(testCode) + `</span>
	<div class="moz-signature"><i><br>
	<br>
	ChatBotAI<br>
	team: IT-Zoo<br>
	<i></div>
	</body>
	</html>
	`
	bodyMessage := sender.WriteHTMLEmail(reciever, subject, message)
	sender.SendMail(reciever, subject, bodyMessage)
}

func (route *UserRoute) generatorVerificationPassPhrase() int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(20000)%20000 + 10000
	// fmt.Println("Random value for verification: ", value)
	return value
}

func (route *UserRoute) checkExistCode(userid string) error {
	for key, value := range route.UserVerification {
		if key == userid {
			creation := value["timeCreated"]
			dureation := value["timeDuration"]
			timeEspared := creation.(time.Time).Add(dureation.(time.Duration))

			if timeEspared.Before(time.Now()) {
				delete(route.UserVerification, userid)
				return errors.New("code expired")
			}
		}
	}
	return nil
}

/*confirmCodeUser - подтверждение кода пользователя*/
func (route *UserRoute) confirmCodeUser(w http.ResponseWriter, r *http.Request) {
	var confirmPayload payloads.ConfirmCodeUserPayload

	if err := json.NewDecoder(r.Body).Decode(&confirmPayload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.ConfirmCode",
			"error":   err.Error(),
		}, "application/json")
		return
	}

	if err2 := route.checkCodeUser(confirmPayload); err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.ConfirmCode",
			"error":   err2.Error(),
		}, "application/json")
		return
	}

	if err3 := route.UserDao.UpdateUserVerification(confirmPayload.Userid, true); err3 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.ConfirmCode",
			"error":   err3.Error(),
		}, "application/json")
		return
	}

	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status":  "success verificated",
		"context": "UserRoute.ConfirmCode",
	}, "application/json")
	return
}

func (route *UserRoute) checkCodeUser(codePayload payloads.ConfirmCodeUserPayload) error {
	for key, value := range route.UserVerification {
		if key == codePayload.Userid {
			creation := value["timeCreated"]
			dureation := value["timeDuration"]
			timeEspared := creation.(time.Time).Add(dureation.(time.Duration))

			fmt.Println("Time Created: ", creation.(time.Time).Format("Mon Jan _2 15:04:05 2006"))
			fmt.Println("Time Espired: ", creation.(time.Time).Add(dureation.(time.Duration)).Format("Mon Jan _2 15:04:05 2006"))
			fmt.Println("Time Now: ", time.Now().Format("Mon Jan _2 15:04:05 2006"))

			if timeEspared.Before(time.Now()) {
				return errors.New("code expired")
			}

			if value["code"].(int) == codePayload.Code {
				delete(route.UserVerification, codePayload.Userid)
				return nil
			} else {
				return errors.New("error code value! Can you get new code")
			}
		}
	}
	return errors.New("code not exist in memory")
}

/*createUser - создание пользователя*/
func (route *UserRoute) createUser(w http.ResponseWriter, r *http.Request) {
	var newUser payloads.CreateUserPayload
	var bufferForAvatar bytes.Buffer

	r.ParseForm()

	src, _, err := r.FormFile("avatar")
	if err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.CreateUser",
			"code":    "can not recognize avatar",
			"error":   err.Error(),
		}, " multipart/form-data;")
		return
	}
	_, err2 := bufferForAvatar.ReadFrom(src)

	if err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.CreateUser",
			"code":    "can not read buffer from reseived bytes image",
			"error":   err2.Error(),
		},
			"multipart/form-data;",
		)
		return
	}
	encodedAvatar := base64.StdEncoding.EncodeToString(bufferForAvatar.Bytes())

	newUser.Avatar = "data:image/png;base64, " + encodedAvatar
	fmt.Println(r.Form)

	for key, value := range r.Form {
		switch key {
		case "username":
			if err := route.validateFields(w, r, value[0]); err != nil {
				return
			}
			newUser.Username = value[0]
			break
		case "password":
			if err := route.validateFields(w, r, value[0]); err != nil {
				return
			}
			newUser.Password = value[0]
			break
		case "email":
			if err := route.validateFields(w, r, value[0]); err != nil {
				return
			}
			newUser.Email = value[0]
			break
		default:
			break
		}
	}

	// insert new data to database
	err3 := route.UserDao.CreateNewUser(newUser)
	if err3 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"context": "UserDao.CreateNewUser",
			"error":   err3.Error(),
		},
			"multipart/form-data;")
		return
	}

	route.EResponser.ResponseWithError(w, r, http.StatusOK, map[string]string{
		"status":  "success",
		"context": "UserRoute.CreateUser",
	},
		"multipart/form-data;")
	return
}

func (route *UserRoute) validateFields(w http.ResponseWriter, r *http.Request, field string) error {
	if strings.TrimSpace(field) == "" {
		route.EResponser.ResponseWithError(w, r, http.StatusOK, map[string]string{
			"status":  "error",
			"context": "UserRoute.CreateUser",
			"error":   "you email field does not be empty",
		},
			"multipart/form-data;")
		return errors.New("empty field")
	}
	return nil
}

func (route *UserRoute) updateUserSettings(w http.ResponseWriter, r *http.Request) {
	var payload payloads.UserUpdatePayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.UpdateUserSettings",
		},
			"application/json",
		)
		return
	}

	newUser, err2 := route.UserDao.UpdateUser(&payload)
	if err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.UpdateUserSettings",
			"code":    err2.Error(),
		},
			"application/json",
		)
		return
	}

	jsonUser, err3 := json.Marshal(newUser)
	if err3 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusInternalServerError, map[string]string{
			"status":  "error",
			"context": "UserRoute.UpdateUserSettings",
			"code":    err3.Error(),
		},
			"application/json",
		)
		return
	}

	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status":  "success updated",
		"data":    string(jsonUser),
		"context": "UserRoute.RemoveUserByID",
	},
		"application/json",
	)
	return

}

func (route *UserRoute) removeUserByID(w http.ResponseWriter, r *http.Request) {
	var payload payloads.RemoveUserPayload

	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.RemoveUserByID",
			"code":    err.Error(),
		},
			"application/json",
		)
		return
	}

	if err2 := route.UserDao.RemoveUserByID(&payload); err2 != nil {
		route.EResponser.ResponseWithError(w, r, http.StatusNotAcceptable, map[string]string{
			"status":  "error",
			"context": "UserRoute.RemoveUserByID",
			"code":    err2.Error(),
		},
			"application/json",
		)
		return
	}

	route.EResponser.ResponseWithJSON(w, r, http.StatusOK, map[string]string{
		"status":  "success removing user",
		"context": "UserRoute.RemoveUserByID",
	},
		"application/json",
	)
	return
}

/*InitRoute - инициализация роутера*/
func (route *UserRoute) InitRoute(tokens *dao.TokenDAO, users *dao.UserDAO, jwtMiddle *middlewares.JWTChecker) *UserRoute {
	route.TokenDao = tokens
	route.UserDao = users
	route.JWTMiddle = jwtMiddle
	route.UserVerification = map[string]map[string]interface{}{}
	return route
}

/*RoutesSetting - конфигурация роутера для маршрутов авторизации\регистрации*/
func (route *UserRoute) RoutesSetting(router *mux.Router) *mux.Router {
	router.HandleFunc(createUser, route.createUser).Methods("POST")
	router.HandleFunc(confirmAccount, route.confirmAccount).Methods("GET")
	router.HandleFunc(confirmAccount, route.confirmCodeUser).Methods("POST")
	router.HandleFunc(updateUserSettings, route.JWTMiddle.JWTMiddleware(route.removeUserByID)).Methods("UPDATE")
	router.HandleFunc(removeUserByID, route.JWTMiddle.JWTMiddleware(route.removeUserByID)).Methods("DELETE")
	return router
}
