# Table of Contents

- [API Routes](#API-Routes)

**API Routes**

1. Every User:
/permissions
/permissions/ofUser/:userID
/permissions/ofUser/:userID/groups
 
/permissions/getFrom/single/:permissionID/users
/permissions/getFrom/single/:permissionID/groups
/permissions/getFrom/group/:groupID/users

 2. Admins:
/permissions/addTo/user/:userID/single/:permissionID
/permissions/addTo/user/:userID/group/:groupID
/permissions/addTo/group/:groupID/single/:permissionID
 
/permissions/removeFrom/user/:userID/:groupORpermissionID
/permissions/removeFrom/group/:groupID/:permissionID
 
/permissions/create/group/:groupname?groupparameters
/permissions/create/permission/:permissionCode?permissionparameters

/permissions/delete/group/:groupID
/permissions/delete/permission/:permissionID