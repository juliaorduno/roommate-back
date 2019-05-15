package models

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"github.com/joho/godotenv"
	"../db"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type User struct {
	gorm.Model
	Email string `json:"email"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

type UserForm struct {
	Email string `sql:"unique" json:"email"`
	Password string `json:"password"`
	Token string `json:"token"`
	FullName  string     `json:"full_name"`
	Username  string     `json:"username"`
}

type UserModel struct{}

//Validate incoming user details...
/*func (account *Account) Validate() (map[string] interface{}, bool) {

	if !strings.Contains(account.Email, "@") {
		return u.Message(false, "Email address is required"), false
	}

	if len(account.Password) < 6 {
		return u.Message(false, "Password is required"), false
	}

	//Email must be unique
	temp := &Account{}

	//check for errors and duplicate emails
	err := GetDB().Table("accounts").Where("email = ?", account.Email).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection error. Please retry"), false
	}
	if temp.Email != "" {
		return u.Message(false, "Email address already in use by another user."), false
	}

	return u.Message(false, "Requirement passed"), true
}*/

func (m *UserModel) Create(user *User) (err error) {

	/*if resp, ok := account.Validate(); !ok {
		return resp
	}*/

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	
	if err := db.DB.Create(user).Error; err != nil {
		return err
	}

	var myEnv map[string]string
	myEnv, err = godotenv.Read()

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(myEnv["TOKEN_PASSWORD"]))
	user.Token = tokenString

	user.Password = "" //delete password

	return nil
}

func Login(email, password string) bool {

	user := &User{}
	
	if err := db.DB.Where("email = ?", email).First(user).Error; err != nil {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return false
	}
	//Worked! Logged In
	user.Password = ""

	return true
}