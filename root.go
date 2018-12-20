package harbourpermissions

import (
	"crypto/rsa"
	"database/sql"
	"net/http"

	harbourauth "github.com/FelixWieland/harbour-auth"
	"github.com/corneldamian/httpway"
)

var signKey *rsa.PrivateKey
var server *httpway.Server
var db *sql.DB
var secret string

const (
	privKeyPath = "keys/app.rsa" //openssl genrsa -out app.rsa 1024
)

//Start starts the API Server
func Start() {

	signKey, _ = harbourauth.LoadAsPrivateRSAKey("")
	credentials := loadCredentials("../auth.json")
	secret = "demoSecret"

	if ldb, err := connectToDB(credentials.toString()); err == nil {
		db = ldb
		defer db.Close()
	} else {
		println("Cant connect to Database")
	}

	router := httpway.New()

	public := router.Middleware(accessLogger)
	private := public.Middleware(authCheck)
	nautilus := private.Middleware(nautilusCheck)

	//TEST PRIVATE ROUTES*/
	private.POST("/pvt", testJWTLogin)

	/*FOR EVERY USER*/
	private.POST("/permissions", permissions) //Get a list of all permissions
	private.POST("/permissions/ofUser/:userID", permissionsOfUser)
	private.POST("/permissions/ofUser/:userID/groups", permissionGroupsOfUser)

	private.POST("/permissions/getFrom/single/:permissionID/users", usersWithPermission)
	private.POST("/permissions/getFrom/single/:permissionID/groups", groupsWithPermission)
	private.POST("/permissions/getFrom/group/:groupID/users", usersWithGroup)

	/*FOR NAUTLIUS USER*/
	nautilus.POST("/permissions/addTo/user/:userID/single/:permissionID", addPermissionToUser)
	nautilus.POST("/permissions/addTo/user/:userID/group/:groupID", addGroupToUser)
	nautilus.POST("/permissions/addTo/group/:groupID/single/:permissionID", addPermissionToGroup)

	nautilus.POST("/permissions/removeFrom/user/:userID/:groupORpermissionID", removeGroupOrPermissionFromUser)
	nautilus.POST("/permissions/removeFrom/group/:groupID/:permissionID", removePermissionFromGroup)

	nautilus.POST("/permissions/create/group/:groupname", createGroup)
	nautilus.POST("/permissions/create/permission/:permissionCode", createPermission)

	nautilus.POST("/permissions/delete/group/:groupID", deleteGroup)
	nautilus.POST("/permissions/delete/permission/:permissionID", deletePermission)

	http.ListenAndServe(":5001", router)

	server = httpway.NewServer(nil)
	server.Addr = ":5001"
	server.Handler = router

	server.Start()
}

func showPublicStats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Nothing to see here"))
}
