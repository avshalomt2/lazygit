package main

import (
	"github.com/jesseduffield/lazygit/pkg/app"
)

// TODO When using F on commit to amend the staged changes into the commit, there's no work-in-progress indicator in the status bar (as exists when commiting/amending)
// TODO Do we really want to keep fetching periodically while committing/amending? The precommit hook output is being spammed in the middle by the fetch output
// TODO Why Amend error shows in a popup, but commit output shows in the command log?

// These values may be set by the build script via the LDFLAGS argument
var (
	commit      string
	date        string
	version     string
	buildSource = "unknown"
)

func main() {
	ldFlagsBuildInfo := &app.BuildInfo{
		Commit:      commit,
		Date:        date,
		Version:     version,
		BuildSource: buildSource,
	}

	app.Start(ldFlagsBuildInfo, nil)
}
