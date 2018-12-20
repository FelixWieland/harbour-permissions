package harbourpermissions

import (
	"net/http"

	"github.com/corneldamian/httpway"
)

//Route: /permissions
func permissions(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)
	_ = ctx
}

//Route: /permissions/ofUser/:userID
func permissionsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	_ = reqUserID
}

//Route: /permissions/ofUser/:userID/groups
func permissionGroupsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	_ = reqUserID
}

//Route: /permissions/getFrom/single/:permissionID/users
func usersWithPermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqPermissionID
}

//Route: /permissions/getFrom/single/:permissionID/groups
func groupsWithPermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqPermissionID
}

//Route: /permissions/getFrom/group/:groupID/users
func usersWithGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	_ = reqGroupID
}

//Route: /permissions/addTo/user/:userID/single/:permissionID
func addPermissionToUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqUserID
	_ = reqPermissionID
}

//Route: /permissions/addTo/user/:userID/group/:groupID
func addGroupToUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqGroupID := ctx.ParamByName("groupID")
	_ = reqUserID
	_ = reqGroupID
}

//Route: /permissions/addTo/group/:groupID/single/:permissionID
func addPermissionToGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqGroupID
	_ = reqPermissionID
}

//Route: /permissions/removeFrom/user/:userID/:groupORpermissionID
func removeGroupOrPermissionFromUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqPermissionOrGroupID := ctx.ParamByName("groupORpermissionID")
	_ = reqUserID
	_ = reqPermissionOrGroupID

}

//Route: /permissions/removeFrom/group/:groupID/:permissionID
func removePermissionFromGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupORpermissionID")
	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqGroupID
	_ = reqPermissionID

}

//Route: /permissions/create/group/:groupname
func createGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupname := ctx.ParamByName("groupname")
	_ = reqGroupname
}

//Route: /permissions/create/permission/:permissionCode
func createPermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionCode := ctx.ParamByName("permissionCode")
	_ = reqPermissionCode
}

//Route: /permissions/delete/group/:groupID
func deleteGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	_ = reqGroupID
}

//Route: /permissions/delete/permission/:permissionID
func deletePermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqPermissionID
}
