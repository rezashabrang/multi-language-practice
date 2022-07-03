package main

import (
	"fmt"
	"strconv"
)

type Permissions struct {
	canSeeMessages      bool
	canDeleteMessages   bool
	canEditMessages     bool
	canKickMembers      bool
	canMakeMembersAdmin bool
	canAddMembers       bool
}

func SetUserPermissions(permissions *Permissions) int8 {
	// TODO: Implement
	return 0
}

func GetUserPermissions(permissionsMask int8) *Permissions {
	// TODO: Implement
	perm := Permissions {
		canSeeMessages: false,
		canDeleteMessages: false,
		canEditMessages: false,
		canKickMembers: false,
		canMakeMembersAdmin: false,
		canAddMembers: false,
	}
	stringPerm := strconv.Itoa(int(permissionsMask))
	for _, digit := range stringPerm {
		value, _ := strconv.Atoi(string(digit))
		fmt.Println(value)
	}
	
	return &perm
}

func main() {
	GetUserPermissions(01)
}