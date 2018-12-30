package harbourpermissions

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/corneldamian/httpway"
)

//Route: /permissions
func permissions(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)
	_ = ctx

	rows, err := db.Query("SELECT * FROM harbour_permissions")
	if err != nil {
		err := "Error in selection of permission"
		log.Printf(err)
	}

	permissions := []permission{}

	for rows.Next() {
		var permissionid string
		var moduleid string
		var description string

		err = rows.Scan(&permissionid, &moduleid, &description)

		permissions = append(permissions, permission{
			permissionid,
			moduleid,
			description,
		})
	}
	data, _ := json.Marshal(permissionsRESP{
		newResponseInformations(1),
		permissions,
	})
	w.Write(data)
}

//Route: /permissions/ofUser/:userID
func permissionsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")

	rows, err := db.Query(sqlQuery(`
		SELECT h.permissionid, moduleid, description FROM harbour_permissions as h 
		JOIN harbour_userpermissions as u
		JOIN harbour_permissiongroups as g 
		WHERE
			u.userid = ? AND h.permissionid = u.permissionid
			OR
			g.groupid = u.groupid AND g.permissionid = h.permissionid
		`).prep(reqUserID))

	if err != nil {
		err := "Error in selection of permission for user: " + reqUserID
		log.Printf(err)
	}

	permissions := []permission{}

	for rows.Next() {
		var permissionid string
		var moduleid string
		var description string

		err = rows.Scan(&permissionid, &moduleid, &description)

		permissions = append(permissions, permission{
			permissionid,
			moduleid,
			description,
		})
	}

	data, _ := json.Marshal(permissionsForUserRESP{
		newResponseInformations(2),
		reqUserID,
		permissions,
	})
	w.Write(data)
}

//Route: /permissions/ofUser/:userID/groups
func permissionGroupsOfUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	_ = reqUserID

	rows, err := db.Query(sqlQuery(`
		SELECT h.groupid, description FROM harbour_permissiongroupdescr as h 
		JOIN harbour_userpermissions as u
		WHERE u.userid = ?
		AND u.groupid = h.groupid
		`).prep(reqUserID))

	if err != nil {
		err := "Error in selection of permission groups for user: " + reqUserID
		log.Printf(err)
	}

	permissionGroups := []permissionGroup{}

	for rows.Next() {
		var groupid string
		var description string

		err = rows.Scan(&groupid, &description)

		permissionGroups = append(permissionGroups, permissionGroup{
			groupid,
			description,
		})
	}

	data, _ := json.Marshal(permissionsGroupsOfUserRESP{
		newResponseInformations(3),
		reqUserID,
		permissionGroups,
	})
	w.Write(data)
}

//Route: /permissions/getFrom/single/:permissionID/users
func usersWithPermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqPermissionID

	rows, err := db.Query(sqlQuery(`
		SELECT userid FROM harbour_userpermissions
		WHERE permissionid = ?
		`).prep(reqPermissionID))

	if err != nil {
		err := "Error in selection of users with permission: " + reqPermissionID
		log.Printf(err)
	}

	users := []string{}

	for rows.Next() {
		var userid string

		err = rows.Scan(&userid)

		users = append(users, userid)
	}

	data, _ := json.Marshal(usersWithPermissionRESP{
		newResponseInformations(4),
		reqPermissionID,
		users,
	})
	w.Write(data)
}

//Route: /permissions/getFrom/single/:permissionID/groups
func groupsWithPermission(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqPermissionID

	rows, err := db.Query(sqlQuery(`
		SELECT h.groupid, g.description FROM harbour_permissiongroups as h
		JOIN harbour_permissiongroupdescr as g
		WHERE h.permissionid = ?
		`).prep(reqPermissionID))

	if err != nil {
		err := "Error in selection of groups with permission: " + reqPermissionID
		log.Printf(err)
	}

	groups := []permissionGroup{}

	for rows.Next() {
		var groupid string
		var description string

		err = rows.Scan(&groupid, &description)

		groups = append(groups, permissionGroup{
			groupid,
			description,
		})
	}

	data, _ := json.Marshal(groupsWithPermissionRESP{
		newResponseInformations(5),
		reqPermissionID,
		groups,
	})
	w.Write(data)
}

//Route: /permissions/getFrom/group/:groupID/users
func usersWithGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	_ = reqGroupID

	rows, err := db.Query(sqlQuery(`
		SELECT userid FROM harbour_userpermissions
		WHERE groupid = ?
		`).prep(reqGroupID))

	if err != nil {
		err := "Error in selection of users with group: " + reqGroupID
		log.Printf(err)
	}

	users := []string{}

	for rows.Next() {
		var userid string

		err = rows.Scan(&userid)

		users = append(users, userid)
	}

	data, _ := json.Marshal(usersWithGroupRESP{
		newResponseInformations(6),
		reqGroupID,
		users,
	})
	w.Write(data)
}

//Route: /permissions/addTo/user/:userID/single/:permissionID
func addPermissionToUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqPermissionID := ctx.ParamByName("permissionID")

	//FIRST CHECK IF USERID AND PERMISSIONID ARE VIABLE
	userIDExists := rowExists(sqlQuery(`
					SELECT userid FROM harbour_userauth
					WHERE userid = ?
					`).prep(reqUserID))
	permissionIDExits := rowExists(sqlQuery(`
					SELECT permissionid FROM harbour_permissions
					WHERE permissionid = ?`).prep(reqPermissionID))

	if !userIDExists || !permissionIDExits {
		data, _ := json.Marshal(newErrResponse(1, "IdDontExist", "UserID or GroupID dont exist"))
		w.Write(data)
		return
	}

	//NOW CHECK IF USER ALREADY HAVE THIS PERMISSION
	userpermissionExits := rowExists(sqlQuery(`
					SELECT permissionid FROM harbour_userpermissions
					WHERE permissionid = ?
					AND userid = ?`).prep(reqPermissionID, reqUserID))

	if userpermissionExits {
		data, _ := json.Marshal(newErrResponse(2, "UserAlreadyHavePermission", "The User already have this permission"))
		w.Write(data)
		return
	}

	stmt, err := db.Prepare("INSERT INTO harbour_userpermissions SET permissionid=?, groupid=?, userid=?")
	if err != nil {
		log.Printf("%v", err.Error())
	}
	_, err = stmt.Exec(reqPermissionID, "", reqUserID)

	// if there is an error inserting, "handle" it
	if err != nil {
		err := "A Error occured while adding a userpermission"
		log.Printf("%v", err)
		data, _ := json.Marshal(addedUserPermissionGroupRESP{
			newResponseInformations(7),
			false,
			"",
			"",
		})
		w.Write(data)
		return
	}

	data, _ := json.Marshal(addedUserPermissionRESP{
		newResponseInformations(7),
		true,
		reqUserID,
		reqPermissionID,
	})
	w.Write(data)
}

//Route: /permissions/addTo/user/:userID/group/:groupID
func addGroupToUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqGroupID := ctx.ParamByName("groupID")

	//FIRST CHECK IF USERID AND GROUPID ARE VIABLE

	userIDExists := rowExists(sqlQuery(`
					SELECT userid FROM harbour_userauth
					WHERE userid = ?
					`).prep(reqUserID))
	groupIDExits := rowExists(sqlQuery(`
					SELECT groupid FROM harbour_permissiongroupsdescr
					WHERE groupid = ?`).prep(reqGroupID))

	if !userIDExists || !groupIDExits {
		data, _ := json.Marshal(newErrResponse(1, "IdDontExist", "UserID or GroupID dont exist"))
		w.Write(data)
		return
	}

	//NOW CHECK IF USER ALREADY HAVE THIS GROUP
	usergroupExits := rowExists(sqlQuery(`
					SELECT groupid FROM harbour_userpermissions
					WHERE groupid = ?
					AND userid = ?`).prep(reqGroupID, reqUserID))

	if usergroupExits {
		data, _ := json.Marshal(newErrResponse(2, "UserAlreadyHavePermission", "The User already have this group"))
		w.Write(data)
		return
	}

	stmt, err := db.Prepare("INSERT INTO harbour_userpermissions SET permissionid=?, groupid=?, userid=?")
	_, err = stmt.Exec("", reqGroupID, reqUserID)

	// if there is an error inserting, "handle" it
	if err != nil {
		err := "A Error occured while adding a userpermissiongroup"
		log.Printf("%v", err)
		data, _ := json.Marshal(addedUserPermissionGroupRESP{
			newResponseInformations(8),
			false,
			"",
			"",
		})
		w.Write(data)
		return
	}

	data, _ := json.Marshal(addedUserPermissionGroupRESP{
		newResponseInformations(8),
		true,
		reqUserID,
		reqGroupID,
	})
	w.Write(data)
}

//Route: /permissions/addTo/group/:groupID/single/:permissionID
func addPermissionToGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	reqPermissionID := ctx.ParamByName("permissionID")
	_ = reqGroupID
	_ = reqPermissionID

	//FIRST CHECK IF PERMISSIONID AND GROUPID ARE VIABLE

	permissionIDExits := rowExists(sqlQuery(`
					SELECT permissionid FROM harbour_permissions
					WHERE permissionid = ?`).prep(reqPermissionID))
	groupIDExits := rowExists(sqlQuery(`
					SELECT groupid FROM harbour_permissiongroupsdescr
					WHERE groupid = ?`).prep(reqGroupID))

	if !permissionIDExits || !groupIDExits {
		data, _ := json.Marshal(newErrResponse(1, "IdDontExist", "UserID or GroupID dont exist"))
		w.Write(data)
		return
	}

	//NOW CHECK IF GROUP ALREADY HAVE THIS PERMSSION
	grouppermissionExists := rowExists(sqlQuery(`
					SELECT permissionid FROM harbour_permissiongroups
					WHERE groupid = ?
					AND permissionid = ?`).prep(reqGroupID, reqPermissionID))

	if grouppermissionExists {
		data, _ := json.Marshal(newErrResponse(2, "GroupAlreadyHavePermission", "The Group already have this permission"))
		w.Write(data)
		return
	}

	stmt, err := db.Prepare("INSERT INTO harbour_permissiongroups SET groupid=?, permissionid=?")
	_, err = stmt.Exec(reqGroupID, reqPermissionID)

	// if there is an error inserting, "handle" it
	if err != nil {
		err := "A Error occured while adding a permssion to a group"
		log.Printf("%v", err)
		data, _ := json.Marshal(addedUserPermissionGroupRESP{
			newResponseInformations(9),
			false,
			"",
			"",
		})
		w.Write(data)
	}

	data, _ := json.Marshal(addedPermissionToGroupRESP{
		newResponseInformations(9),
		true,
		reqGroupID,
		reqPermissionID,
	})
	w.Write(data)

}

//Route: /permissions/removeFrom/user/:userID/:groupORpermissionID
func removeGroupOrPermissionFromUser(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqUserID := ctx.ParamByName("userID")
	reqPermissionOrGroupID := ctx.ParamByName("groupORpermissionID")

	stmt, _ := db.Prepare(`
					DELETE FROM harbour_userpermissions
					WHERE userid=?
					AND (
						groupid=?
						OR
						permissionid=?
					)`)
	res, _ := stmt.Exec(reqUserID, reqPermissionOrGroupID, reqPermissionOrGroupID)

	affRows, _ := res.RowsAffected()
	if affRows == 0 {
		//deleted nothing
		data, _ := json.Marshal(deletedPermissionFromUserRESP{
			newResponseInformations(10),
			false,
			"",
			int(affRows),
		})
		w.Write(data)
		return
	}

	data, _ := json.Marshal(deletedPermissionFromUserRESP{
		newResponseInformations(10),
		true,
		reqPermissionOrGroupID,
		int(affRows),
	})
	w.Write(data)
}

//Route: /permissions/removeFrom/group/:groupID/:permissionID
func removePermissionFromGroup(w http.ResponseWriter, r *http.Request) {
	ctx := httpway.GetContext(r)

	reqGroupID := ctx.ParamByName("groupID")
	reqPermissionID := ctx.ParamByName("permissionID")

	stmt, _ := db.Prepare(`
					DELETE FROM harbour_permissiongroups
					WHERE groupid=?
					AND permissionid=?
				`)
	res, _ := stmt.Exec(reqGroupID, reqPermissionID)

	affRows, _ := res.RowsAffected()
	if affRows == 0 {
		//deleted nothing
		data, _ := json.Marshal(deletedPermissionFromGroupRESP{
			newResponseInformations(10),
			false,
			"",
			"",
			int(affRows),
		})
		w.Write(data)
		return
	}

	data, _ := json.Marshal(deletedPermissionFromGroupRESP{
		newResponseInformations(10),
		true,
		reqGroupID,
		reqPermissionID,
		int(affRows),
	})
	w.Write(data)

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
