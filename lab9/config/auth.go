package config

const LoginUrl = "/user/login"

type AuthTestData struct {
	Login    string
	Password string
}

func GetValidLoginData() AuthTestData {
	return AuthTestData{
		Login:    "login.test",
		Password: "qwerty123",
	}
}

func GetInvalidLoginData() AuthTestData {
	return AuthTestData{
		Login:    "login.test",
		Password: "qwerty1234",
	}
}
