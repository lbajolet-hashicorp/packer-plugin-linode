package linode

import (
	"fmt"
	"net/http"

	"github.com/linode/linodego"
	"github.com/linode/packer-plugin-linode/version"
	"golang.org/x/oauth2"
)

func newLinodeClient(pat string) linodego.Client {
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: pat})

	oauthTransport := &oauth2.Transport{
		Source: tokenSource,
	}
	oauth2Client := &http.Client{
		Transport: oauthTransport,
	}

	client := linodego.NewClient(oauth2Client)

	projectURL := "https://www.packer.io"
	userAgent := fmt.Sprintf("Packer/%s (+%s) linodego/%s",
		version.PluginVersion.FormattedVersion(), projectURL, linodego.Version)

	client.SetUserAgent(userAgent)
	return client
}
