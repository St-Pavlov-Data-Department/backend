package buildinfo

var (
	GitCommitHash string
	GoVersion     string
)

type BuildInfo struct {
	GitCommitHash string `json:"git_commit_hash"`
	GoVersion     string `json:"go_version"`
}

func New() *BuildInfo {
	return &BuildInfo{
		GitCommitHash: GitCommitHash,
		GoVersion:     GoVersion,
	}
}
