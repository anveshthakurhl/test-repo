package version

import (
	"fmt"
	"runtime"

	"github.com/coreos/go-semver/semver"
)

var Name = "monopoly"

var Description = "monopoly - waypoint service"

// Version is the main version number that is being run at the moment.
var Version = "0.15.62"

// Prerelease is a prerelease marker for the version. If this is "" (empty string)
// then it means that it is a final release. Otherwise, this is a prerelease
// such as "dev" (in development), "beta", "rc1", etc.
var Prerelease = "stage"

// SemVer is an instance of version.Version. This has the secondary
// benefit of verifying during tests and init time that our version is a
// proper semantic version, which should always be the case.
var SemVer *semver.Version

func init() {
	SemVer = semver.New(Version)
}

// String returns the complete version string, including prerelease
func String() string {
	s := fmt.Sprintf("%s %s %s", runtime.GOOS, runtime.GOARCH, runtime.Version())
	if Prerelease != "" {
		return fmt.Sprintf("%s-%s %s", Version, Prerelease, s)
	}

	return fmt.Sprintf("%s %s", Version, s)
}
