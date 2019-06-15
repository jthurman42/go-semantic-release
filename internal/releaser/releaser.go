package releaser

import (
	"fmt"
	"os"

	"github.com/Nightapes/go-semantic-release/internal/releaser/github"
	"github.com/Nightapes/go-semantic-release/internal/shared"

	"github.com/Nightapes/go-semantic-release/pkg/config"
)

// Releasers struct type
type Releasers struct {
	config *config.ReleaseConfig
}

// Releaser interface for providers
type Releaser interface {
	ValidateConfig() error
	CreateRelease(*shared.ReleaseVersion, *shared.GeneratedChangelog) error
	UploadAssets(assets []config.Asset) error
	GetCommitURL() string
	GetCompareURL(oldVersion, newVersion string) string
}

// New initialize a Relerser
func New(c *config.ReleaseConfig) *Releasers {
	return &Releasers{
		config: c,
	}
}

//GetReleaser returns an initialized releaser
func (r *Releasers) GetReleaser() (Releaser, error) {
	switch r.config.Release {
	case github.GITHUB:
		return github.New(&r.config.GitHubProvider)
	}
	return nil, fmt.Errorf("Could not initialize a releaser from this type: %s", r.config.Release)
}

func checkIfAssetsExists(assets []config.Asset) error {
	var missingAssets []string
	for _, asset := range assets {

		if _, err := os.Stat(asset.Name); err != nil {
			missingAssets = append(missingAssets, asset.Name)
		}
	}

	if len(missingAssets) != 0 {
		return fmt.Errorf("Could not find specified Asset: %+v ", assets)
	}

	return nil

}
