package harbourpermissions

import (
	"log"
	"net/http"

	"github.com/FelixWieland/harbour-auth"

	"github.com/corneldamian/httpway"
)

/*
func checkAuth(h http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, rt httprouter.Params) {
		log.Printf("Incoming connection from %v", r.RemoteAddr)

		s, err1 := r.Cookie("session")
		if err1 == nil {
			log.Printf("%v", s.Value)
		}

		w.Header().Set("Content-Type", "application/json")
		_, err := isLoggedin(w, r, rt)
		if err != nil {
			//notloggedin
			apiErrorHandler(w, r, rt, err)
		} else {
			isLoggedin(w, r, rt)
		}
	}
}
*/

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
	claims, err := harbourauth.HarbourJWT(jwt).Decode(signKey)
	if err != nil {
		//notloggedin
	} else {
		ctx.Set("userid", claims.UserID)
		ctx.Set("username", claims.Username)
		ctx.Set("issue", claims.Issuer)
		ctx.Next(w, r)
	}
}

/*
func julienHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		// do stuff
	}
}
*/
