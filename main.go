package main

import (
	"os"

	"code.cloudfoundry.org/cli/cf/terminal"
	"code.cloudfoundry.org/cli/cf/trace"

	"code.cloudfoundry.org/cli/plugin"
	"code.cloudfoundry.org/cli/plugin/models"
)

// WhoamiCmd struct for the plugin
type WhoamiCmd struct {
	ui terminal.UI
}

// GetMetadata shows metadata for the plugin
func (c *WhoamiCmd) GetMetadata() plugin.PluginMetadata {
	return plugin.PluginMetadata{
		Name: "Whoami Plugin",
		Version: plugin.VersionType{
			Major: 0,
			Minor: 1,
			Build: 0,
		},
		MinCliVersion: plugin.VersionType{
			Major: 6,
			Minor: 12,
			Build: 0,
		},
		Commands: []plugin.Command{
			{
				Name:     "whoami",
				HelpText: "Displays informarion about the currently logged in user",
				UsageDetails: plugin.Usage{
					Usage: "cf whoami",
				},
			},
		},
	}
}

func main() {
	plugin.Start(new(WhoamiCmd))
}

// Run will be executed when cf whoami gets invoked
func (c *WhoamiCmd) Run(cliConnection plugin.CliConnection, args []string) {
	if args[0] != "whoami" {
		return
	}

	var err error
	var hasAPIEndpoint bool
	var apiEndpoint string
	var isLoggedIn bool
	var username string
	var hasOrg bool
	var hasSpace bool
	var org plugin_models.Organization
	var space plugin_models.Space

	traceLogger := trace.NewLogger(os.Stdout, true, os.Getenv("CF_TRACE"), "")
	c.ui = terminal.NewUI(os.Stdin, os.Stdout, terminal.NewTeePrinter(os.Stdout), traceLogger)

	if hasAPIEndpoint, err = cliConnection.HasAPIEndpoint(); err != nil {
		c.ui.Failed(err.Error())
	}

	if !hasAPIEndpoint {
		c.ui.Failed("You have to set your api endpoint first with `cf api YOUR_API_URL`")
	}

	if isLoggedIn, err = cliConnection.IsLoggedIn(); err != nil {
		c.ui.Failed(err.Error())
	}

	if !isLoggedIn {
		c.ui.Failed("Nobody is logged in")
	}

	if apiEndpoint, err = cliConnection.ApiEndpoint(); err != nil {
		c.ui.Failed(err.Error())
	}

	if username, err = cliConnection.Username(); err != nil {
		c.ui.Failed(err.Error())
	}
	if len(username) == 0 {
		c.ui.Failed("Interesting - you are logged in, but your username is empty. Mysterious!")
	}

	c.ui.Say("You are logged in as %s at %s", terminal.EntityNameColor(username), terminal.EntityNameColor(apiEndpoint))

	if hasOrg, err = cliConnection.HasOrganization(); err != nil {
		c.ui.Failed(err.Error())
	}
	if hasSpace, err = cliConnection.HasSpace(); err != nil {
		c.ui.Failed(err.Error())
	}

	if hasOrg && hasSpace {
		if org, err = cliConnection.GetCurrentOrg(); err != nil {
			c.ui.Failed(err.Error())
		}
		if space, err = cliConnection.GetCurrentSpace(); err != nil {
			c.ui.Failed(err.Error())
		}

		c.ui.Say("You are targeting the %s space in the %s org", terminal.EntityNameColor(space.Name), terminal.EntityNameColor(org.Name))
	}
}
