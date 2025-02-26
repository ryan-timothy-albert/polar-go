// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package components

import (
	"encoding/json"
	"fmt"
)

// Permission - The permission level to grant. Read more about roles and their permissions on [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
type Permission string

const (
	PermissionPull     Permission = "pull"
	PermissionTriage   Permission = "triage"
	PermissionPush     Permission = "push"
	PermissionMaintain Permission = "maintain"
	PermissionAdmin    Permission = "admin"
)

func (e Permission) ToPointer() *Permission {
	return &e
}
func (e *Permission) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pull":
		fallthrough
	case "triage":
		fallthrough
	case "push":
		fallthrough
	case "maintain":
		fallthrough
	case "admin":
		*e = Permission(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Permission: %v", v)
	}
}

// BenefitGitHubRepositoryProperties - Properties for a benefit of type `github_repository`.
type BenefitGitHubRepositoryProperties struct {
	// The owner of the repository.
	RepositoryOwner string `json:"repository_owner"`
	// The name of the repository.
	RepositoryName string `json:"repository_name"`
	// The permission level to grant. Read more about roles and their permissions on [GitHub documentation](https://docs.github.com/en/organizations/managing-user-access-to-your-organizations-repositories/managing-repository-roles/repository-roles-for-an-organization#permissions-for-each-role).
	Permission Permission `json:"permission"`
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RepositoryID *string `json:"repository_id,omitempty"`
}

func (o *BenefitGitHubRepositoryProperties) GetRepositoryOwner() string {
	if o == nil {
		return ""
	}
	return o.RepositoryOwner
}

func (o *BenefitGitHubRepositoryProperties) GetRepositoryName() string {
	if o == nil {
		return ""
	}
	return o.RepositoryName
}

func (o *BenefitGitHubRepositoryProperties) GetPermission() Permission {
	if o == nil {
		return Permission("")
	}
	return o.Permission
}

func (o *BenefitGitHubRepositoryProperties) GetRepositoryID() *string {
	if o == nil {
		return nil
	}
	return o.RepositoryID
}
