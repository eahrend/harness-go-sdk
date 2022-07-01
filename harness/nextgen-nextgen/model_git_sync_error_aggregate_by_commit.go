/*
 * Harness NextGen Software Delivery Platform API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// This contains a list of Git Sync Error details for a given Commit Id
type GitSyncErrorAggregateByCommit struct {
	// Commit Id
	GitCommitId string `json:"gitCommitId,omitempty"`
	// The number of active errors in a commit
	FailedCount int32 `json:"failedCount,omitempty"`
	// Git Sync Config Id.
	RepoId string `json:"repoId,omitempty"`
	// Name of the branch.
	BranchName string `json:"branchName,omitempty"`
	// Commit Message to use for the merge commit.
	CommitMessage string `json:"commitMessage,omitempty"`
	// This is the time at which the Git Sync error was logged
	CreatedAt int64 `json:"createdAt,omitempty"`
	// This has the list of Git Sync errors corresponding to a specific Commit Id
	ErrorsForSummaryView []GitSyncError `json:"errorsForSummaryView,omitempty"`
}
