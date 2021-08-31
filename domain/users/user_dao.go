package users

import (
	"fmt"
	"github.com/aldo-zoop/bookstore-users-api/datasources/mysql/users_db"
	"github.com/aldo-zoop/bookstore-users-api/utils/date_utils"
	"github.com/aldo-zoop/bookstore-users-api/utils/errors"
	"log"
	"strings"
)

const(
	indexUniqueEmail = "email_UNIQUE"
	queryInsertUser = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
)

var usersDB = make(map[int64]*User)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}
	user.Email = result.Email
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.CreatedDate = result.CreatedDate
	return nil
}

func (user *User) Save() *errors.RestErr  {
	//if err := users_db.Client.Ping(); err != nil {
	//	panic(err)
	//}
	log.Print(fmt.Sprintf("%d", user.Id))
	stmt, err := users_db.Client.Prepare(queryInsertUser)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.CreatedDate = date_utils.GetNow(false)

	insertResult, err := stmt.Exec(user.FirstName, user.LastName, user.Email, user.CreatedDate)
	if err != nil {
		if strings.Contains(err.Error(), indexUniqueEmail) {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already exists", user.Email))
		}
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to save user: %s", err.Error()))
	}

	userId, err := insertResult.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("error when trying to get the last insert id: %s", err.Error()))
	}
	user.Id = userId

	//current := usersDB[user.Id]
	//if current != nil {
	//	if current.Email == user.Email {
	//		return errors.NewBadRequestError(fmt.Sprintf("Email %s already exists", user.Email))
	//	}
	//	return errors.NewBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	//}
	//user.CreatedDate = date_utils.GetNow(false)
	//usersDB[user.Id] = user
	return nil
}