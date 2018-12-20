package harbourpermissions

import (
	"log"
	"net/http"

	"github.com/FelixWieland/harbour-auth"

	"github.com/corneldamian/httpway"
)

func accessLogger(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming connection from %v", r.RemoteAddr)
	httpway.GetContext(r).Next(w, r)
}

func authCheck(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)
	w.Header().Set("Content-Type", "application/json")
	jwt := r.FormValue("jwt")
	if len(jwt) == 0 {
		//no jwt provided
	}
	claims, err := harbourauth.HarbourJWT(jwt).Decode(signKey, secret)
	if err != nil {
		//notloggedin
		w.Write([]byte("Error"))
	} else {
		ctx.Set("userid", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("issue", claims.Issuer)
		ctx.Next(w, r)
	}
}

func nautilusCheck(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)
	w.Header().Set("Content-Type", "application/json")

	userid := ctx.Get("userid")
	_ = userid
	ctx.Next(w, r)
}

/*
func julienHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
	}
}
*/
