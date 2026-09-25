package share

import (
	"errors"
	"hospital-a/constants"
	"hospital-a/entities"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(constants.EnvKeys.JwtSecret)

type Claims struct {
	StaffId      int    `json:"staffId"`
	HospitalName string `json:"hospitalName"`
	jwt.RegisteredClaims
}

func GenerateToken(staff *entities.Staff) (string, error) {
	claims := Claims{
		StaffId:      staff.ID,
		HospitalName: staff.HospitalName,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtSecret)
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(t *jwt.Token) (interface{}, error) {
			if t.Method != jwt.SigningMethodES256 {
				return nil, errors.New("Invalid signing method")
			}

			return jwtSecret, nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
