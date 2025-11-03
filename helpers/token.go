package helpers

import (
	"authentication/config"
	"authentication/models"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserID  uint `json:"user_id"`
	Email   string `json:"email"`
	IsAdmin bool `json:"is_admin"`
	jwt.StandardClaims
}

var jwtKey []byte

func SetJWTKey(key string) {
	jwtKey = []byte(key)
}

func GetJWTKey() []byte {
	return jwtKey
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

func GenerateToken(email string, userID uint, isAdmin bool) (string, string, error) {
    tokenExpiry := time.Now().Add(24 * time.Hour).Unix()
    refreshExpiry := time.Now().Add(7 * 24 * time.Hour).Unix()

    claims := &Claims{
        Email:   email,
        UserID:  userID,
        IsAdmin: isAdmin,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: tokenExpiry,
        },
    }

    access := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signedAccess, err := access.SignedString(jwtKey)
    if err != nil {
        return "", "", err
    }

    refreshClaims := &Claims{
        StandardClaims: jwt.StandardClaims{ExpiresAt: refreshExpiry},
    }
    refresh := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    signedRefresh, err := refresh.SignedString(jwtKey)
    if err != nil {
        return "", "", err
    }

    return signedAccess, signedRefresh, nil
}

func UpdateAllToken(signedToken, signedRefreshToken string, userID uint) error {
	var user models.User
	if err := config.DB.Model(&user).Where("id = ?", userID).Updates(map[string]interface{}{
		"token":         signedToken,
		"refresh_token": signedRefreshToken,
		"updated_at":    time.Now(),
	}).Error; err != nil {
		return err
	}
	return nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func VerifyPassword(hashedPwd, plainPwd string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(plainPwd))
	return err == nil, err
}
