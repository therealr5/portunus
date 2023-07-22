/*******************************************************************************
* Copyright 2019 Stefan Majewsky <majewsky@gmx.net>
* SPDX-License-Identifier: GPL-3.0-only
* Refer to the file "LICENSE" for details.
*******************************************************************************/

package core

import (
	"encoding/json"
	"reflect"
	"sort"
	"strconv"
)

// Group represents a single group of users. Membership in a group implicitly
// grants its Permissions to all users in that group.
type Group struct {
	Name             string           `json:"name"`
	LongName         string           `json:"long_name"`
	MemberLoginNames GroupMemberNames `json:"members"`
	Permissions      Permissions      `json:"permissions"`
	PosixGID         *PosixID         `json:"posix_gid,omitempty"`
}

// Cloned returns a deep copy of this user.
func (g Group) Cloned() Group {
	logins := g.MemberLoginNames
	g.MemberLoginNames = make(GroupMemberNames)
	for name, isMember := range logins {
		if isMember {
			g.MemberLoginNames[name] = true
		}
	}
	if g.PosixGID != nil {
		val := *g.PosixGID
		g.PosixGID = &val
	}
	return g
}

// ContainsUser checks whether this group contains the given user.
func (g Group) ContainsUser(u User) bool {
	return g.MemberLoginNames[u.LoginName]
}

// IsEqualTo is a type-safe wrapper around reflect.DeepEqual().
func (g Group) IsEqualTo(other Group) bool {
	return reflect.DeepEqual(g, other)
}

// GroupMemberNames is the type of Group.MemberLoginNames.
type GroupMemberNames map[string]bool

// MarshalJSON implements the json.Marshaler interface.
func (g GroupMemberNames) MarshalJSON() ([]byte, error) {
	names := make([]string, 0, len(g))
	for name, isMember := range g {
		if isMember {
			names = append(names, name)
		}
	}
	sort.Strings(names)
	return json.Marshal(names)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (g *GroupMemberNames) UnmarshalJSON(data []byte) error {
	var names []string
	err := json.Unmarshal(data, &names)
	if err != nil {
		return err
	}
	*g = make(map[string]bool)
	for _, name := range names {
		(*g)[name] = true
	}
	return nil
}

////////////////////////////////////////////////////////////////////////////////

// PosixID represents a POSIX user or group ID.
type PosixID uint16

func (id PosixID) String() string {
	return strconv.FormatUint(uint64(id), 10)
}
