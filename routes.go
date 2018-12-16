package harbourpermissions

import (
	"log"
	"net/http"
)

func permissionsOfUser(w http.ResponseWriter, r *http.Request) {
	log.Printf("Incoming connection from %v", r.RemoteAddr)
}
