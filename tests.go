package harbourpermissions

import "net/http"

func testJWTLogin(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}
