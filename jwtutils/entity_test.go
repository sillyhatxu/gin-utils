package jwtutils

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNew(t *testing.T) {
	secretKey := "sillyhattest"
	client, err := New(secretKey)
	assert.Nil(t, err)
	assert.Equal(t, secretKey, client.SecretKey())
}

func TestCreateTokenAndParse(t *testing.T) {
	type MyCustomClaims struct {
		UserId   string `json:"userId"`
		UserName string `json:"userName"`
		RoleId   string `json:"roleId"`
		jwt.StandardClaims
	}
	secretKey := "sillyhattest"
	client, err := New(secretKey)
	assert.Nil(t, err)
	assert.Equal(t, secretKey, client.SecretKey())
	claim := &MyCustomClaims{
		UserId:         "test-user-id-0001",
		UserName:       "Silly Hat",
		RoleId:         "test-role-id-0152",
		StandardClaims: jwt.StandardClaims{},
	}
	token, err := client.CreateToken(claim)
	assert.Nil(t, err)
	assert.NotNil(t, token)
	var result MyCustomClaims
	err = client.ParseToken(token, &result)
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, claim.UserId, result.UserId)
	assert.Equal(t, claim.UserName, result.UserName)
	assert.Equal(t, claim.RoleId, result.RoleId)
}
