package harbourpermissions

import "database/sql"

//UserID is a own type for the userid
type UserID string

// HasPermissionDB checks if a UserID has a specific permission with a database connection
func (UserID UserID) HasPermissionDB(db *sql.DB, permissionid string) bool {
	return false
}

// HasPermissionWEB checks if a UserID has a specific permission with a http service
func (UserID UserID) HasPermissionWEB(service string, permissionid string) bool {
	return false
}

// IsNautilusDB checks if a User has Nautilus permissions with a database connection
func (UserID UserID) IsNautilusDB(db *sql.DB) bool {
	return true
}

// IsNautilusWEB checks if a User has Nautilus permissions with a http service
func (UserID UserID) IsNautilusWEB(service string) bool {
	return false
}
