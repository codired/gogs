package repo

import (
	"github.com/gogits/gogs/models"
	"github.com/gogits/gogs/pkg/context"
	"github.com/gogits/gogs/pkg/setting"
)

func GetDiff(c *context.APIContext) {
	var userName = c.Params(":username")
	var repoName = c.Params(":reponame")
	var commitId = c.Params(":sha")

	diff, err := models.GetDiffCommit(models.RepoPath(userName, repoName),
		commitId,
		setting.Git.MaxGitDiffLines,
		setting.Git.MaxGitDiffLineCharacters,
		setting.Git.MaxGitDiffFiles)

	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	c.JSON(200, diff)
}
