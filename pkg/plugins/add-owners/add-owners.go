package add_owners

import (
	"fmt"
	"github.com/jenkins-x/lighthouse/pkg/plugins"
	"github.com/jenkins-x/lighthouse/pkg/scmprovider"
	"github.com/sirupsen/logrus"
)

const (
	pluginName = "add-owners"
)

func createPlugin() plugins.Plugin {
	return plugins.Plugin{
		Description: "This plugin will help to add owners into repo",
		Commands: []plugins.Command{{
			Name: "add-owners",
			Arg: &plugins.CommandArg{
				Pattern: `[^\r\n]+`,
			},
			Action: plugins.
				Invoke(func(_ plugins.CommandMatch, pc plugins.Agent, event scmprovider.GenericCommentEvent) error {
					fmt.Println(event.Action.String())
					return handle(pc.SCMProviderClient, pc.Logger, &event)
				}),
		}},
	}
}

func handle(spc *scmprovider.Client, log *logrus.Entry, e *scmprovider.GenericCommentEvent) error {
	org := e.Repo.Namespace
	repo := e.Repo.Name
	num := e.Number
	fmt.Println("recive event", e.Repo)

	return spc.CreateComment(org, repo, num, e.IsPR, "hi " + org + "repo" + repo)
}

func init() {
	plugins.RegisterPlugin(pluginName, createPlugin())
}
