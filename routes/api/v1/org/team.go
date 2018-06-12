// Copyright 2016 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package org

import (
	api "github.com/gogs/go-gogs-client"

	"github.com/gogs/gogs/pkg/context"
	"github.com/gogs/gogs/routes/api/v1/convert"
	"github.com/gogs/gogs/models"
	"github.com/gogs/gogs/models/errors"
)

func ListTeams(c *context.APIContext) {
	org := c.Org.Organization
	if err := org.GetTeams(); err != nil {
		c.Error(500, "GetTeams", err)
		return
	}

	apiTeams := make([]*api.Team, len(org.Teams))
	for i := range org.Teams {
		apiTeams[i] = convert.ToTeam(org.Teams[i])
	}
	c.JSON(200, apiTeams)
}

func AddMemberToOrganizationTeam(c *context.APIContext) {
	org := c.Org.Organization

	if !org.IsOwnedBy(c.User.ID) {
		c.Status(403)

		return
	}

	t, err := org.GetTeam(c.Params(":teamname"))

	if err != nil {
		c.Error(500, "AddMemberToOrganization", err)
	}

	u, err := models.GetUserByName(c.Params(":username"))

	if err != nil {
		if errors.IsUserNotExist(err) {
			c.Status(404)
		} else {
			c.Error(500, "AddMemberToOrganization", err)
		}

		return
	}

	err = t.AddMember(u.ID)

	if err != nil {
		c.Error(500, "AddMemberToOrganization", err)

		return
	}

	c.JSON(200, true)
}

func RemoveMemberFromOrganizationTeam(c *context.APIContext) {
	org := c.Org.Organization

	if !org.IsOwnedBy(c.User.ID) {
		c.Status(403)

		return
	}

	t, err := org.GetTeam(c.Params(":teamname"))

	if err != nil {
		c.Error(500, "AddMemberToOrganization", err)
	}

	u, err := models.GetUserByName(c.Params(":username"))

	if err != nil {
		if errors.IsUserNotExist(err) {
			c.Status(404)
		} else {
			c.Error(500, "AddMemberToOrganization", err)
		}

		return
	}

	err = t.RemoveMember(u.ID)

	if err != nil {
		c.Error(500, "AddMemberToOrganization", err)

		return
	}

	c.JSON(200, true)
}
