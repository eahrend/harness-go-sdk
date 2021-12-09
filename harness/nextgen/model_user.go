/*
 * CD NextGen API Reference
 *
 * This is the Open Api Spec 3 for the NextGen Manager. This is under active development. Beware of the breaking change with respect to the generated code stub
 *
 * API version: 3.0
 * Contact: contact@harness.io
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package nextgen

type User struct {
	Uuid                             string               `json:"uuid"`
	AppId                            string               `json:"appId"`
	CreatedBy                        *EmbeddedUser        `json:"createdBy,omitempty"`
	CreatedAt                        int64                `json:"createdAt,omitempty"`
	LastUpdatedBy                    *EmbeddedUser        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedAt                    int64                `json:"lastUpdatedAt"`
	Name                             string               `json:"name,omitempty"`
	ExternalUserId                   string               `json:"externalUserId,omitempty"`
	GivenName                        string               `json:"givenName,omitempty"`
	FamilyName                       string               `json:"familyName,omitempty"`
	Email                            string               `json:"email,omitempty"`
	CompanyName                      string               `json:"companyName,omitempty"`
	AccountName                      string               `json:"accountName,omitempty"`
	UserGroups                       []UserGroup          `json:"userGroups,omitempty"`
	Accounts                         []Account            `json:"accounts,omitempty"`
	PendingAccounts                  []Account            `json:"pendingAccounts,omitempty"`
	SupportAccounts                  []Account            `json:"supportAccounts,omitempty"`
	LastLogin                        int64                `json:"lastLogin,omitempty"`
	FirstLogin                       bool                 `json:"firstLogin,omitempty"`
	Password                         []string             `json:"password,omitempty"`
	Token                            string               `json:"token,omitempty"`
	TwoFactorJwtToken                string               `json:"twoFactorJwtToken,omitempty"`
	EmailVerified                    bool                 `json:"emailVerified,omitempty"`
	PasswordExpired                  bool                 `json:"passwordExpired,omitempty"`
	UserLocked                       bool                 `json:"userLocked,omitempty"`
	StatsFetchedOn                   int64                `json:"statsFetchedOn,omitempty"`
	LastAccountId                    string               `json:"lastAccountId,omitempty"`
	DefaultAccountId                 string               `json:"defaultAccountId,omitempty"`
	LastAppId                        string               `json:"lastAppId,omitempty"`
	Disabled                         bool                 `json:"disabled,omitempty"`
	Imported                         bool                 `json:"imported,omitempty"`
	UserLockoutInfo                  *UserLockoutInfo     `json:"userLockoutInfo,omitempty"`
	RateLimitProtection              *RateLimitProtection `json:"rateLimitProtection,omitempty"`
	TwoFactorAuthenticationEnabled   bool                 `json:"twoFactorAuthenticationEnabled,omitempty"`
	TwoFactorAuthenticationMechanism string               `json:"twoFactorAuthenticationMechanism,omitempty"`
	OauthProvider                    string               `json:"oauthProvider,omitempty"`
	ReportedSegmentTracks            []string             `json:"reportedSegmentTracks,omitempty"`
	UtmInfo                          *UtmInfo             `json:"utmInfo,omitempty"`
	AccountIds                       []string             `json:"accountIds,omitempty"`
}