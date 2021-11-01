package controllers

import (
	"coach/models"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/json"
	"fmt"
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
	if user.Id == "" || user.Pw == "" {
		w.WriteHeader(406)
		return
	}
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	privateKeyPemStr := exportPrivateKeyAsPEMStr(privateKey)
	publicKeyPemStr := exportPublicKeyAsPEMStr(publicKey)
	p := []byte(user.Pw)
	sha256 := sha256.Sum256(p)
	user.Pw = fmt.Sprintf("%x", sha256)
	user.PrivateKey = privateKeyPemStr
	user.PublicKey = publicKeyPemStr
	isSuccess := models.InsertUser(&user)
	responseBody, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	if isSuccess {
		w.Write(responseBody)
	} else {
		w.WriteHeader(409)
	}
}
