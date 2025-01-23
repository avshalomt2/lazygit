package filter_by_path

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

// FIXME Complete this

var FilterByStatusTest = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Filter files by status (unstaged / staged / untracked / conflicted)",
	ExtraCmdArgs: []string{},
	Skip:         false,
	SetupConfig: func(config *config.AppConfig) {
	},
	SetupRepo: func(shell *Shell) {
		commonSetup(shell)
	},
	Run: func(t *TestDriver, keys config.KeybindingConfig) {
		t.Views().Files().
			IsFocused().
			Press(keys.Files.OpenStatusFilter)

		t.ExpectPopup().Menu().
			Title(Equals("Filtering")).
			Select(Contains("Staged files")).
			Confirm()
        t.Views().Files().

		// See postFilterTest function
	},
})
