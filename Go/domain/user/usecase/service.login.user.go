package usecase

import (
	"context"
	"errors"
	"time"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt"
)

type LoginUserRequest struct {
	Phone string			`json:"phone"`
	Password	string			`json:"password"`
}

type MyClaims struct {
	jwt.StandardClaims
	Name string `json:"name"`
	Phone string `json:"phone"`
	Role	string `json:"role"`
}

var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256

func (s *Service) Login(ctx context.Context, loginData LoginUserRequest) (*string, error) {

	user, err := s.users.FindUserByPhone(ctx, loginData.Phone)
	if err != nil {
		return nil, err
	}

	// TODO: encode password

	if user.Password != loginData.Password {
		return nil, errors.New("password incorrect")
	}

	claims := MyClaims{
    StandardClaims: jwt.StandardClaims{
        ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
    },
    Name: user.Name,
    Phone: user.Phone,
    Role: user.Role,
	}

	token := jwt.NewWithClaims(
    JWT_SIGNING_METHOD,
    claims,
	)

	jwtSecret := ctx.Value("JWT_SECRET").(string)
	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

// func (s *Service) AuthenticateToken(ctx context.Context, token string) (*string, error) {

// 	jwtSecret := ctx.Value("JWT_SECRET").(string)
// 	newToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
// 		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 				return nil, fmt.Errorf("Signing method invalid")
// 		} else if method != JWT_SIGNING_METHOD {
// 				return nil, fmt.Errorf("Signing method invalid")
// 		}
	
// 		return []byte(jwtSecret), nil
// 	})
	
// 	if err != nil {
// 		return nil, err
// 	}
// 	fmt.Println("token:", newToken)
// }
