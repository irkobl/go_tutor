package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

type authInfo struct {
	Usr string
	Pwd string
}

func (api *API) authJWT(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var auth authInfo

	err = json.Unmarshal(body, &auth)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if auth.Usr == "Usr" && auth.Pwd == "Pwd" {
		token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"usr": auth.Usr,
			"nbf": time.Now().Unix(),
		})

		tokenString, err := token.SignedString([]byte("secret-password"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte(tokenString))

	}

}
