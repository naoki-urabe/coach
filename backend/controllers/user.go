package controllers

import (
	"coach/models"
	"crypto/rand"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func registerUser(w http.ResponseWriter, r *http.Request) {
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
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	privateKeyPemStr := exportPrivateKeyAsPEMStr(privateKey)
	publicKeyPemStr := exportPublicKeyAsPEMStr(publicKey)
	user.PrivateKey = privateKeyPemStr
	user.PublicKey = publicKeyPemStr
	models.InsertUser(&user)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	w.Write(responseBody)
}