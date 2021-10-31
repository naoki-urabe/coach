package controllers

import (
	"coach/models"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	jwt "github.com/golang-jwt/jwt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	_ "os"
	_ "reflect"
	"strings"
	"time"
)

func getTokenHandler(privateKey *rsa.PrivateKey) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["sub"] = "525252"
	claims["name"] = "nao"
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func exportPEMStrToPublicKey(publicKeyPem string) *rsa.PublicKey {
	block, _ := pem.Decode([]byte(publicKeyPem))
	key, _ := x509.ParsePKCS1PublicKey(block.Bytes)
	return key
}

func exportPEMStrToPrivateKey(privateKeyPem string) *rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKeyPem))
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func exportPrivateKeyAsPEMStr(privateKey *rsa.PrivateKey) string {
	privateKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
		},
	))
	return privateKeyPem
}

func exportPublicKeyAsPEMStr(publicKey *rsa.PublicKey) string {
	publicKeyPem := string(pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(publicKey),
		},
	))
	return publicKeyPem
}

func validateJWTMiddleware(next http.Handler) http.Handler {
	publicKey := exportPEMStrToPublicKey(os.Getenv("PUBLICKEY"))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		if (*r).Method == "OPTIONS" {
			return
		}
		tokenString := strings.Split((*r).Header["Authorization"][0], " ")[1]
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})
		if err != nil {
			log.Fatal(err)
		}
		if token.Valid {
			next.ServeHTTP(w, r)
		} else {
			log.Fatal(err)
		}
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if (*r).Method == "OPTIONS" {
		return
	}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var user models.User
	if err := json.Unmarshal(reqBody, &user); err != nil {
		log.Fatal(err)
	}
	isLogin := models.FindUser(&user)
	privateKey := exportPEMStrToPrivateKey(user.PrivateKey)
	publicKey := user.PublicKey
	if user.PrivateKey == "" {
		log.Fatal(err)
	}
	os.Setenv("PUBLICKEY", publicKey)
	type Response struct {
		Token string `json:"token"`
		User  string `json:"user"`
	}
	response := Response{
		Token: getTokenHandler(privateKey),
		User:  user.Id,
	}
	json, _ := json.Marshal(response)
	if isLogin {
	    w.Write(json)
	} else {
	    w.WriteHeader(401)
	    w.Write(json)
	}
}
