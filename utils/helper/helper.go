package helper

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/argon2"
)

const (
	ArgonTime    = 1
	ArgonMemory  = 64 * 1024
	ArgonThreads = 4
	ArgonKeyLen  = 32
)

func CreateJwtToken(id uint, email string, role string) (string, error) {
	claims := jwt.MapClaims{
		"email": email,
		"id":    id,
		"role":  role,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func VerifyJwtToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// generates a random salt of specified size
func GenerateRandomSalt(size int) (string, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(salt), nil
}

// hashes a password using Argon2id with a random salt
func HashPassword(password string) (string, string, error) {
	salt, err := GenerateRandomSalt(16)
	if err != nil {
		return "", "", err
	}

	decodedSalt, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return "", "", err
	}

	hash := argon2.IDKey([]byte(password), decodedSalt, ArgonTime, ArgonMemory, ArgonThreads, ArgonKeyLen)
	encodedHash := base64.RawStdEncoding.EncodeToString(hash)

	return encodedHash, salt, nil
}

// checks if the provided password matches the stored hash and salt
func VerifyPassword(password, salt, storedHash string) (bool, error) {
	decodedSalt, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false, fmt.Errorf("failed to decode salt: %w", err)
	}

	decodedStoredHash, err := base64.RawStdEncoding.DecodeString(storedHash)
	if err != nil {
		return false, fmt.Errorf("failed to decode stored hash: %w", err)
	}

	hash := argon2.IDKey([]byte(password), decodedSalt, ArgonTime, ArgonMemory, ArgonThreads, ArgonKeyLen)

	if len(hash) != len(decodedStoredHash) {
		return false, nil
	}

	for i := 0; i < len(hash); i++ {
		if hash[i] != decodedStoredHash[i] {
			return false, nil
		}
	}

	return true, nil
}
