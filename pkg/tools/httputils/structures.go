package httputils

import (
	"context"

	"github.com/dgrijalva/jwt-go"
)

type tokenInfo struct {
	jwt.StandardClaims

	UserID int64
	Roles  []int
}

func NewUserContext(ctx context.Context) UserContext {
	if ctx == nil || ctx.Value("user_info") == nil {
		return UserContext{}
	}

	return ctx.Value("user_info").(UserContext)
}

type UserContext struct {
	token  string
	userID int64
}

func (u *UserContext) Token() string {
	return u.token
}

func (u *UserContext) SetToken(t string) {
	u.token = t
}

func (u *UserContext) UserID() int64 {
	return u.userID
}

func (u *UserContext) SetUserID(id int64) {
	u.userID = id
}
