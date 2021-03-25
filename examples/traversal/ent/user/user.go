// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeManage holds the string denoting the manage edge name in mutations.
	EdgeManage = "manage"
	// Table holds the table name of the user in the database.
	Table = "users"
	// PetsTable is the table the holds the pets relation/edge.
	PetsTable = "pets"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pets"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
	// FriendsTable is the table the holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// GroupsTable is the table the holds the groups relation/edge. The primary key declared below.
	GroupsTable = "group_users"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// ManageTable is the table the holds the manage relation/edge.
	ManageTable = "groups"
	// ManageInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	ManageInverseTable = "groups"
	// ManageColumn is the table column denoting the manage relation/edge.
	ManageColumn = "group_admin"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}

var (
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"group_id", "user_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
