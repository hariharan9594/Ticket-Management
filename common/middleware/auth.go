package middleware


import (
	"net/http"
	
	"github.com/labstack/echo"
	"gopkg.in/dgrijalva/jwt-go.v2"
	"gitlab.com/vipindasvg/ticket-management/common"
)

const (
	ErrTokenExpired = "token expired"
	ErrParseToken   = "error parsing token"
	ErrInvalidToken = "invalid access token"
)

var (
	VerifyKey []byte
)

// Middleware for validating admin JWT tokens
func AdminRBAC(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		VerifyKey = []byte("TokenPassword")
		// validate the token
		token, err := jwt.ParseFromRequest(c.Request(), func(token *jwt.Token) (interface{}, error) {
			// Verify the token with public key, which is the counter part of private key
			return VerifyKey, nil
		})
		// token.Claims  map[UserInfo:map[Email:cstick@example.com Role:member] exp:1.569569521e+09]
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError: // JWT validation error
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					return echo.NewHTTPError(http.StatusUnauthorized, ErrTokenExpired)
				default:
					return echo.NewHTTPError(http.StatusInternalServerError, ErrParseToken)
				}
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, ErrParseToken)
			}
		}
		if token.Valid {
			userdata := token.Claims["UserInfo"].(map[string]interface{})
			if userdata["IsAdmin"] == true {
				return next(c)
			} else {
				return echo.NewHTTPError(http.StatusForbidden, "forbidden")
			}
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidToken)
		}
	}
}

// Middleware for validating user JWT tokens
func UserRBAC(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		VerifyKey = []byte("TokenPassword")
		// validate the token
		token, err := jwt.ParseFromRequest(c.Request(), func(token *jwt.Token) (interface{}, error) {
			// Verify the token with public key, which is the counter part of private key
			return VerifyKey, nil
		})
		// token.Claims  map[UserInfo:map[Email:cstick@example.com Role:member] exp:1.569569521e+09]
		if err != nil {
			switch err.(type) {
			case *jwt.ValidationError: // JWT validation error
				vErr := err.(*jwt.ValidationError)
				switch vErr.Errors {
				case jwt.ValidationErrorExpired:
					return echo.NewHTTPError(http.StatusUnauthorized, ErrTokenExpired)
				default:
					return echo.NewHTTPError(http.StatusInternalServerError, ErrParseToken)
				}
			default:
				return echo.NewHTTPError(http.StatusInternalServerError, ErrParseToken)
			}
		}
		if token.Valid {
			userdata := token.Claims["UserInfo"].(map[string]interface{})
			tid := userdata["Id"] // user id from JWT
			if userdata["IsAdmin"] == true {
				return next(c)
			} else {
				id, err := common.ParseUID(c.Request().URL.Path)
				if err != nil {
					return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
				}
				if float64(id) != tid {
					return echo.NewHTTPError(http.StatusForbidden, "forbidden")
				} else {
					return next(c)
				}
			}
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, ErrInvalidToken)
		}
	}
}
