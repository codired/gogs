// Copyright 2015 The Gogs Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package admin

import (
	api "github.com/gogs/go-gogs-client"

	"github.com/gogs/gogs/pkg/context"
	"github.com/gogs/gogs/routes/api/v1/org"
	"github.com/gogs/gogs/routes/api/v1/user"
	"github.com/gogs/gogs/models"
)

// https://github.com/gogs/go-gogs-client/wiki/Administration-Organizations#create-a-new-organization
func CreateOrg(c *context.APIContext, form api.CreateOrgOption) {
	org.CreateOrgForUser(c, form, user.GetUserByParams(c))
}

func DeleteOrg(c *context.APIContext) {
	u, err := models.GetOrgByName(c.Params(":name"))
	if err != nil {
		c.Error(500, "GetOrgByName", err)
		return
	}

	err = models.DeleteOrganization(u)
	if err != nil {
		c.Error(500, "DeleteOrganization", err)
		return
	}

	c.Status(204)
}
