package handler

import (
	"errors"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/rodzinkaPLN/GrindingOptimalization/api/db"
)

var (
	errUserTokenMissing = errors.New("failed to get user token")
	errUserClaimMissing = errors.New("failed to get user claims")
	errUserUIDMissing   = errors.New("failed to get user uid")
)

type CrudHandler struct {
	db *db.DB
}

func NewCrudHandler(dbc *db.DB) *CrudHandler {
	return &CrudHandler{
		db: dbc,
	}
}

func getUID(c echo.Context) (uuid.UUID, error) {
	token, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return uuid.Nil, errUserTokenMissing
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.Nil, errUserClaimMissing
	}
	uid, ok := claims["uid"].(string)
	if !ok {
		return uuid.Nil, errUserUIDMissing
	}
	return uuid.Parse(uid)
}

// BASE FOR EP
// type requestCreateUserFile struct {
// }

// type responseCreateUserFile struct {
// }

// func (h *CrudHandler) CreateUserFile(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	uid, err := getUID(c)
// 	if err != nil {
// 		return err
// 	}
// 	var req requestCreateUserFile
// 	if err := c.Bind(&req); err != nil {
// 		return err
// 	}

// 	if err := c.Validate(req); err != nil {
// 		return err
// 	}

// 	// TODO: LOGIC
// 	return c.JSON(http.StatusCreated, responseCreateUserFile{})
// }
