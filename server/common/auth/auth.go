package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwk"
	"github.com/lestrrat-go/jwx/jwt"
)

type JWT struct {
	key *rsa.PrivateKey
}

var jwtInst *JWT

func GetJwtInstance() *JWT {
	if jwtInst == nil {
		initJwtInstance()
	}

	return jwtInst
}

func initJwtInstance() error {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Panicln("Failed to generate private key: " + err.Error())
		return err
	}

	jwtInst = &JWT{key}

	return nil
}

func GenJwtToken(id uint) string {
	token, _ := GetJwtInstance().NewToken(id)

	return token
}

func (jwtInst *JWT) NewToken(id uint) (string, error) {
	key := jwtInst.key

	token := jwt.New()
	token.Set("id", id)
	token.Set("exp", time.Now().Add(time.Hour*24).Unix())

	// signing a token (using jwk)
	jwkKey, err := jwk.New(key)
	if err != nil {
		log.Panicln("failed to create a JWK key: " + err.Error())
		return "", err
	}

	signed, err := jwt.Sign(token, jwa.RS256, jwkKey)
	if err != nil {
		log.Panicln("failed to sign the token: " + err.Error())
		return "", err
	}

	return string(signed), nil
}

type TokenClaims struct {
	ID uint `json:"id"`
}

func (jwtInst *JWT) VerifyTokenPayload(payload string) (TokenClaims, bool) {
	token, err := jwt.Parse(
		[]byte(payload),
		jwt.WithValidate(true),
		jwt.WithVerify(jwa.RS256, &jwtInst.key.PublicKey),
	)

	claims := TokenClaims{}

	if err != nil {
		log.Panicln("failed to parse JWT token: " + err.Error())
		return claims, false
	}

	buf, err := json.Marshal(token)
	if err != nil {
		log.Panicln("failed to generate JSON: " + err.Error())
		return claims, false
	}

	if err := json.Unmarshal(buf, &claims); err != nil {
		log.Panicln("failed to unmarshal JSON: " + err.Error())
		return claims, false
	}

	return claims, true
}

type TokenPayloadParam struct {
	Token string `json:"access_token" binding:"required"`
}

func (jwtInst *JWT) ParseToken(c *gin.Context) (string, error) {
	token := c.Request.Header["Authorization"]
	if len(token) > 0 {
		return strings.TrimPrefix(token[0], "Bearer"), nil
	}

	tokenPayloadParam := TokenPayloadParam{}
	b := binding.Default(c.Request.Method, c.ContentType())
	err := c.ShouldBindWith(&tokenPayloadParam, b)

	return tokenPayloadParam.Token, err
}
