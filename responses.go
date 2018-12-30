package harbourpermissions

import "time"

//response for an unauthenticated user
type errorResponse struct {
	Type        string `json:"type"`
	Code        int    `json:"code"`
	Message     string `json:"message"`
	Description string `json:"description"`
}

//response header
type responseInformations struct {
	Type        string `json:"type"`
	GeneratedAt string `json:"generatedAt"`
	Code        int    `json:"code"`
}

type permission struct {
	Permissionid string `json:"permissionid"`
	Moduleid     string `json:"moduleid"`
	Description  string `json:"description"`
}

type permissionGroup struct {
	Groupid     string `json:"groupid"`
	Description string `json:"description"`
}

type permissionsRESP struct {
	responseInformations
	Permissions []permission `json:"permissions"`
}

type permissionsForUserRESP struct {
	responseInformations
	Userid          string       `json:"userid"`
	PermissionsUser []permission `json:"permissions"`
}

type permissionsGroupsOfUserRESP struct {
	responseInformations
	Userid     string            `json:"userid"`
	GroupsUser []permissionGroup `json:"groups"`
}

type usersWithPermissionRESP struct {
	responseInformations
	Permissionid string   `json:"permissionid"`
	Users        []string `json:"users"`
}

type groupsWithPermissionRESP struct {
	responseInformations
	Permissionid string            `json:"permissionid"`
	Groups       []permissionGroup `json:"groups"`
}

type usersWithGroupRESP struct {
	responseInformations
	Groupid string   `json:"groupid"`
	Users   []string `json:"users"`
}

type addedUserPermissionRESP struct {
	responseInformations
	State        bool   `json:"state"`
	Userid       string `json:"userid"`
	Permissionid string `json:"permissionid"`
}

type addedUserPermissionGroupRESP struct {
	responseInformations
	State   bool   `json:"state"`
	Userid  string `json:"userid"`
	Groupid string `json:"groupid"`
}

type addedPermissionToGroupRESP struct {
	responseInformations
	State        bool   `json:"state"`
	Groupid      string `json:"groupid"`
	Permissionid string `json:"permissionid"`
}

type deletedPermissionFromUserRESP struct {
	responseInformations
	State       bool   `json:"state"`
	DeletedID   string `json:"deletedid"`
	DeletedRows int    `json:"deletedRows"`
}

type deletedPermissionFromGroupRESP struct {
	responseInformations
	State       bool   `json:"state"`
	GroupID     string `json:"groupid"`
	DeletedID   string `json:"deletedid"`
	DeletedRows int    `json:"deletedRows"`
}

func newErrResponse(errorKey int, errorCode string, errorDescription string) errorResponse {
	return errorResponse{
		"error",
		errorKey,
		errorCode,
		errorDescription,
	}
}
func newResponseInformations(code int) responseInformations {
	return responseInformations{
		"Data",
		time.Now().String(),
		code,
	}
}
