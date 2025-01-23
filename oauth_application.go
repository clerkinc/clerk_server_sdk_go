package clerk

type OAuthApplication struct {
	APIResource
	Object                string  `json:"object"`
	ID                    string  `json:"id"`
	InstanceID            string  `json:"instance_id"`
	Name                  string  `json:"name"`
	ClientID              string  `json:"client_id"`
	ClientSecret          *string `json:"client_secret,omitempty"`
	Public                bool    `json:"public"`
	Scopes                string  `json:"scopes"`
	CallbackURL           string  `json:"callback_url"`
	DiscoveryURL          string  `json:"discovery_url"`
	AuthorizeURL          string  `json:"authorize_url"`
	TokenFetchURL         string  `json:"token_fetch_url"`
	UserInfoURL           string  `json:"user_info_url"`
	TokenIntrospectionURL string  `json:"token_introspection_url"`
	CreatedAt             int64   `json:"created_at"`
	UpdatedAt             int64   `json:"updated_at"`
}

type OAuthApplicationList struct {
	APIResource
	OAuthApplications []*OAuthApplication `json:"data"`
	TotalCount        int64               `json:"total_count"`
}
