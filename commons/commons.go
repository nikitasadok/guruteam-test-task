package commons

import (
	"testTaskGuru/globals"
)

func IsValidToken(userToken string) bool{
	return userToken == globals.ServerToken
}
