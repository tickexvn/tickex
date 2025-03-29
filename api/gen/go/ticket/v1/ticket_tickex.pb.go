// Code generated by protoc-gen-go-tickex. DO NOT EDIT.
// Copyright 2025 Duc-Hung Ho

package ticket

import (
	"fmt"

	"github.com/tickexvn/tickex/api/gen/go/stdx/v1"
)

var (
	_ stdx.Empty
)

const ascii = `
 _______     __          
/_  __(_)___/ /_______ __	TICKEX // TICKET
 / / / / __/  '_/ -_) \ /	--------------
/_/ /_/\__/_/\_\\__/_\_\	tickex.ticket.v1

`

// PrintASCII the ASCII art to the console.
func PrintASCII() {
	fmt.Print(ascii)
}

// HasRoleAtTicketService_CreateTicket checks if the role has access to the method
func HasRoleAtTicketService_CreateTicket(role stdx.Role) bool {
	roleMap := make(map[stdx.Role]bool, 1)
	roleMap[stdx.Role_ROLE_ADMIN] = true

	hasRole, ok := roleMap[role]
	if !ok {
		return false
	}

	return hasRole
}

// HasPermissionAtTicketService_CreateTicket checks if the permission has access to the method
func HasPermissionAtTicketService_CreateTicket(permission stdx.Permission) bool {
	permissionMap := make(map[stdx.Permission]bool, 1)
	permissionMap[stdx.Permission_PERMISSION_CREATE] = true

	hasPermission, ok := permissionMap[permission]
	if !ok {
		return false
	}

	return hasPermission
}

// HasRoleAtTicketService_EditTicket checks if the role has access to the method
func HasRoleAtTicketService_EditTicket(role stdx.Role) bool {
	roleMap := make(map[stdx.Role]bool, 2)
	roleMap[stdx.Role_ROLE_ADMIN] = true
	roleMap[stdx.Role_ROLE_SELLER] = true

	hasRole, ok := roleMap[role]
	if !ok {
		return false
	}

	return hasRole
}

// HasPermissionAtTicketService_EditTicket checks if the permission has access to the method
func HasPermissionAtTicketService_EditTicket(permission stdx.Permission) bool {
	permissionMap := make(map[stdx.Permission]bool, 2)
	permissionMap[stdx.Permission_PERMISSION_UPDATE_ANY] = true
	permissionMap[stdx.Permission_PERMISSION_UPDATE_OWN] = true

	hasPermission, ok := permissionMap[permission]
	if !ok {
		return false
	}

	return hasPermission
}
