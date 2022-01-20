package auth

import (
	"database/sql"
	"net/http"
	_middlewares "sirclo/restapi/layered/delivery/middlewares"
	_entity "sirclo/restapi/layered/entity"
)

type AuthRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (ar *AuthRepository) Login(username string, password string) (string, int) {
	// query := fmt.Sprintf("SELECT id, password FROM users WHERE name='%v'", username)

	result, err := ar.db.Query("SELECT id, password FROM users WHERE name= ?", username)

	if err != nil {
		return "get user failed", http.StatusInternalServerError
	}

	defer result.Close()

	eligibles := []_entity.User{}
	user := _entity.User{}

	for result.Next() {
		if err := result.Scan(&user.Id, &user.Password); err != nil {
			return "geat user failed", http.StatusInternalServerError
		}
		eligibles = append(eligibles, user)
	}

	if len(eligibles) == 0 {
		return "user does not exist", http.StatusUnauthorized
	}

	notMatched := true

	for i := 0; i < len(eligibles) && notMatched; i++ {
		if eligibles[i].Password == password {
			notMatched = false
			user.Id = eligibles[i].Id
			user.Name = eligibles[i].Name
		}
	}

	if notMatched {
		return "password incorrect", http.StatusUnauthorized
	}

	token, err := _middlewares.CreateToken(user.Id, user.Name)

	if err != nil {
		return "token creation failed", http.StatusInternalServerError
	}

	return token, http.StatusOK
}
