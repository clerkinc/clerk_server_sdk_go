package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/clerktest"
	"github.com/stretchr/testify/require"
)

func TestUserClientCreate(t *testing.T) {
	t.Parallel()
	id := "user_123"
	username := "username"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			In:     json.RawMessage(fmt.Sprintf(`{"username":"%s"}`, username)),
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","username":"%s"}`, id, username)),
			Method: http.MethodPost,
			Path:   "/v1/users",
		},
	}
	client := NewClient(config)
	user, err := client.Create(context.Background(), &CreateParams{
		Username: clerk.String(username),
	})
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, username, *user.Username)
}

func TestUserClientList_Request(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Method: http.MethodGet,
			Query: &url.Values{
				"limit":                 []string{"1"},
				"offset":                []string{"2"},
				"order_by":              []string{"-created_at"},
				"email_address":         []string{"foo@bar.com", "baz@bar.com"},
				"organization_id":       []string{"org_123", "org_456"},
				"email_address_query":   []string{"@bar.com"},
				"name_query":            []string{"foobar"},
				"created_at_before":     []string{"1730333164378"},
				"created_at_after":      []string{"1730333164378"},
				"last_active_at_before": []string{"1730333164378"},
				"last_active_at_after":  []string{"1730333164378"},
				"banned":                []string{"false"},
			},
		},
	}
	client := NewClient(config)
	params := &ListParams{
		EmailAddresses:    []string{"foo@bar.com", "baz@bar.com"},
		OrderBy:           clerk.String("-created_at"),
		OrganizationIDs:   []string{"org_123", "org_456"},
		EmailAddressQuery: clerk.String("@bar.com"),
		NameQuery:         clerk.String("foobar"),
	}
	params.Limit = clerk.Int64(1)
	params.Offset = clerk.Int64(2)
	params.CreatedAtBefore = clerk.Int64(1730333164378)
	params.CreatedAtAfter = clerk.Int64(1730333164378)
	params.LastActiveAtBefore = clerk.Int64(1730333164378)
	params.LastActiveAtAfter = clerk.Int64(1730333164378)
	params.Banned = clerk.Bool(false)
	_, err := client.List(context.Background(), params)
	require.NoError(t, err)
}

func TestUserClientList_Response(t *testing.T) {
	t.Parallel()
	usersJSON := `[{"object":"user","id":"user_123"}]`
	countJSON := `{"object":"total_count","total_count":5}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "count") {
			_, err := w.Write([]byte(countJSON))
			require.NoError(t, err)
			return
		}
		_, err := w.Write([]byte(usersJSON))
		require.NoError(t, err)
	}))
	defer ts.Close()

	config := &clerk.ClientConfig{}
	config.URL = clerk.String(ts.URL)
	config.HTTPClient = ts.Client()
	client := NewClient(config)
	list, err := client.List(context.Background(), &ListParams{})
	require.NoError(t, err)
	require.Equal(t, int64(5), list.TotalCount)
	require.Equal(t, 1, len(list.Users))
	require.Equal(t, "user_123", list.Users[0].ID)
}

func TestUserClientCount(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(`{"object":"total_count","total_count":10}`),
			Method: http.MethodGet,
			Path:   "/v1/users/count",
			Query: &url.Values{
				"limit":         []string{"1"},
				"offset":        []string{"2"},
				"order_by":      []string{"-created_at"},
				"email_address": []string{"foo@bar.com", "baz@bar.com"},
			},
		},
	}
	client := NewClient(config)
	params := &ListParams{
		EmailAddresses: []string{"foo@bar.com", "baz@bar.com"},
		OrderBy:        clerk.String("-created_at"),
	}
	params.Limit = clerk.Int64(1)
	params.Offset = clerk.Int64(2)
	totalCount, err := client.Count(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, "total_count", totalCount.Object)
	require.Equal(t, int64(10), totalCount.TotalCount)
}

func TestUserClientGet(t *testing.T) {
	t.Parallel()
	id := "user_123"
	username := "username"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","username":"%s"}`, id, username)),
			Method: http.MethodGet,
			Path:   "/v1/users/" + id,
		},
	}
	client := NewClient(config)
	user, err := client.Get(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, username, *user.Username)
}

func TestUserClientDelete(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","deleted":true}`, id)),
			Method: http.MethodDelete,
			Path:   "/v1/users/" + id,
		},
	}
	client := NewClient(config)
	user, err := client.Delete(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.True(t, user.Deleted)
}

func TestUserClientUpdate(t *testing.T) {
	t.Parallel()
	id := "user_123"
	username := "username"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			In:     json.RawMessage(fmt.Sprintf(`{"username":"%s"}`, username)),
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","username":"%s"}`, id, username)),
			Method: http.MethodPatch,
			Path:   "/v1/users/" + id,
		},
	}
	client := NewClient(config)
	user, err := client.Update(context.Background(), id, &UpdateParams{
		Username: clerk.String(username),
	})
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, username, *user.Username)
}

type testFile struct {
	bytes.Reader
}

func (_ *testFile) Close() error {
	return nil
}

func TestUserClientUpdateProfileImage(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s"}`, userID)),
			Method: http.MethodPost,
			Path:   "/v1/users/" + userID + "/profile_image",
		},
	}
	client := NewClient(config)
	user, err := client.UpdateProfileImage(context.Background(), userID, &UpdateProfileImageParams{
		File: &testFile{Reader: *bytes.NewReader([]byte{})},
	})
	require.NoError(t, err)
	require.Equal(t, userID, user.ID)
}

func TestUserClientDeleteProfileImage(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s"}`, userID)),
			Method: http.MethodDelete,
			Path:   "/v1/users/" + userID + "/profile_image",
		},
	}
	client := NewClient(config)
	user, err := client.DeleteProfileImage(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, userID, user.ID)
}

func TestUserClientUpdateMetadata(t *testing.T) {
	t.Parallel()
	id := "user_123"
	metadata := json.RawMessage(`{"foo":"bar"}`)
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			In:     json.RawMessage(fmt.Sprintf(`{"private_metadata":%s}`, string(metadata))),
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","private_metadata":%s}`, id, string(metadata))),
			Method: http.MethodPatch,
			Path:   "/v1/users/" + id + "/metadata",
		},
	}
	client := NewClient(config)
	user, err := client.UpdateMetadata(context.Background(), id, &UpdateMetadataParams{
		PrivateMetadata: &metadata,
	})
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.JSONEq(t, string(metadata), string(user.PrivateMetadata))
}

func TestUserClientListOAuthAccessTokens(t *testing.T) {
	t.Parallel()
	id := "user_123"
	provider := "oauth_custom"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T: t,
			Out: json.RawMessage(fmt.Sprintf(`{
"data":[{
	"external_account_id":"eac_2dYS7stz9bgxQsSRvNqEAHhuxvW",
	"provider":"%s",
	"token":"the-token"
}],
"total_count":1
}`,
				provider)),
			Method: http.MethodGet,
			Path:   "/v1/users/" + id + "/oauth_access_tokens/" + provider,
			Query: &url.Values{
				"paginated": []string{"true"},
			},
		},
	}
	client := NewClient(config)
	list, err := client.ListOAuthAccessTokens(context.Background(), &ListOAuthAccessTokensParams{
		ID:       id,
		Provider: provider,
	})
	require.NoError(t, err)
	require.Equal(t, int64(1), list.TotalCount)
	require.Equal(t, 1, len(list.OAuthAccessTokens))
	require.Equal(t, "eac_2dYS7stz9bgxQsSRvNqEAHhuxvW", list.OAuthAccessTokens[0].ExternalAccountID)
	require.Equal(t, provider, list.OAuthAccessTokens[0].Provider)
}

func TestUserClientDeleteMFA(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"user_id":"%s"}`, id)),
			Method: http.MethodDelete,
			Path:   "/v1/users/" + id + "/mfa",
		},
	}
	client := NewClient(config)
	mfa, err := client.DeleteMFA(context.Background(), &DeleteMFAParams{
		ID: id,
	})
	require.NoError(t, err)
	require.Equal(t, id, mfa.UserID)
}

func TestUserClientBan(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"object":"user","id":"%s"}`, id)),
			Method: http.MethodPost,
			Path:   "/v1/users/" + id + "/ban",
		},
	}
	client := NewClient(config)
	user, err := client.Ban(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, "user", user.Object)
}

func TestUserClientUnban(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"object":"user","id":"%s"}`, id)),
			Method: http.MethodPost,
			Path:   "/v1/users/" + id + "/unban",
		},
	}
	client := NewClient(config)
	user, err := client.Unban(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, "user", user.Object)
}

func TestUserClientLock(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"object":"user","id":"%s"}`, id)),
			Method: http.MethodPost,
			Path:   "/v1/users/" + id + "/lock",
		},
	}
	client := NewClient(config)
	user, err := client.Lock(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, "user", user.Object)
}

func TestUserClientUnlock(t *testing.T) {
	t.Parallel()
	id := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"object":"user","id":"%s"}`, id)),
			Method: http.MethodPost,
			Path:   "/v1/users/" + id + "/unlock",
		},
	}
	client := NewClient(config)
	user, err := client.Unlock(context.Background(), id)
	require.NoError(t, err)
	require.Equal(t, id, user.ID)
	require.Equal(t, "user", user.Object)
}

func TestUserClientListOrganizationMemberships(t *testing.T) {
	t.Parallel()
	membershipID := "orgmem_123"
	organizationID := "org_123"
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T: t,
			Out: json.RawMessage(fmt.Sprintf(`{
"data": [{
	"id":"%s",
	"organization":{"id":"%s"},
	"public_user_data":{"user_id":"%s"}
}],
"total_count": 1
}`,
				membershipID, organizationID, userID)),
			Method: http.MethodGet,
			Path:   "/v1/users/" + userID + "/organization_memberships",
			Query: &url.Values{
				"limit":  []string{"1"},
				"offset": []string{"2"},
			},
		},
	}
	client := NewClient(config)
	params := &ListOrganizationMembershipsParams{}
	params.Limit = clerk.Int64(1)
	params.Offset = clerk.Int64(2)
	list, err := client.ListOrganizationMemberships(context.Background(), userID, params)
	require.NoError(t, err)
	require.Equal(t, membershipID, list.OrganizationMemberships[0].ID)
	require.Equal(t, organizationID, list.OrganizationMemberships[0].Organization.ID)
	require.Equal(t, userID, list.OrganizationMemberships[0].PublicUserData.UserID)
}

func TestUserClientListOrganizationInvitations(t *testing.T) {
	t.Parallel()
	invitationID := "orginv_123"
	organizationID := "org_123"
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T: t,
			Out: json.RawMessage(fmt.Sprintf(`{
				"data": [{
					"id":"%s",
					"organization_id":"%s"
				}],
				"total_count": 1
				}`,
				invitationID, organizationID)),
			Method: http.MethodGet,
			Path:   "/v1/users/" + userID + "/organization_invitations",
			Query: &url.Values{
				"limit":  []string{"1"},
				"offset": []string{"2"},
				"status": []string{"accepted"},
			},
		},
	}
	client := NewClient(config)
	params := &ListOrganizationInvitationsParams{
		Statuses: &[]string{"accepted"},
		UserID:   userID,
	}
	params.Limit = clerk.Int64(1)
	params.Offset = clerk.Int64(2)
	list, err := client.ListOrganizationInvitations(context.Background(), params)
	require.NoError(t, err)
	require.Equal(t, invitationID, list.OrganizationInvitations[0].ID)
	require.Equal(t, organizationID, list.OrganizationInvitations[0].OrganizationID)
}

func TestUserClientDeletePasskey(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	passkeyIdentificationID := "idn_345"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s"}`, passkeyIdentificationID)),
			Method: http.MethodDelete,
			Path:   "/v1/users/" + userID + "/passkeys/" + passkeyIdentificationID,
		},
	}
	client := NewClient(config)
	passkey, err := client.DeletePasskey(context.Background(), userID, passkeyIdentificationID)
	require.NoError(t, err)
	require.Equal(t, passkeyIdentificationID, passkey.ID)
}
func TestUserClientDeleteWeb3Wallet(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	web3WalletIdentificationID := "idn_345"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s"}`, web3WalletIdentificationID)),
			Method: http.MethodDelete,
			Path:   "/v1/users/" + userID + "/web3_wallets/" + web3WalletIdentificationID,
		},
	}
	client := NewClient(config)
	web3Wallet, err := client.DeleteWeb3Wallet(context.Background(), userID, web3WalletIdentificationID)
	require.NoError(t, err)
	require.Equal(t, web3WalletIdentificationID, web3Wallet.ID)
}

func TestUserClientCreateTOTP(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Method: http.MethodPost,
			Out:    json.RawMessage(`{"backup_codes":[],"created_at":1725548779338,"id":"totp_id","object":"totp","secret":"secret","updated_at":1725548779338,"uri":"otpauth://totp/","verified":false}`),
			Path:   fmt.Sprintf("/v1/users/%s/totp", userID),
		},
	}
	client := NewClient(config)
	totp, err := client.CreateTOTP(context.Background(), userID)
	require.NoError(t, err)
	require.Empty(t, totp.ID)
	require.Empty(t, totp.Secret)
	require.Empty(t, totp.URI)
	require.Equal(t, totp.Object, "totp")
}

func TestUserClientDeleteTOTP(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Method: http.MethodDelete,
			Out:    json.RawMessage(fmt.Sprintf(`{"user_id": "%s"}`, userID)),
			Path:   fmt.Sprintf("/v1/users/%s/totp", userID),
		},
	}
	client := NewClient(config)
	totp, err := client.DeleteTOTP(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, totp.UserID, userID)
}

func TestUserClientDeleteBackupCode(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Method: http.MethodDelete,
			Out:    json.RawMessage(fmt.Sprintf(`{"user_id": "%s"}`, userID)),
			Path:   fmt.Sprintf("/v1/users/%s/backup_code", userID),
		},
	}
	client := NewClient(config)
	totp, err := client.DeleteBackupCode(context.Background(), userID)
	require.NoError(t, err)
	require.Equal(t, totp.UserID, userID)
}

func TestUserClientDeleteExternalAccount(t *testing.T) {
	t.Parallel()
	userID := "user_123"
	externalAccountID := "eac_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Method: http.MethodDelete,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","object":"external_account"}`, externalAccountID)),
			Path:   fmt.Sprintf("/v1/users/%s/external_accounts/%s", userID, externalAccountID),
		},
	}
	client := NewClient(config)
	externalAccount, err := client.DeleteExternalAccount(context.Background(), &DeleteExternalAccountParams{
		UserID: userID,
		ID:     externalAccountID,
	})
	require.NoError(t, err)
	require.Equal(t, externalAccountID, externalAccount.ID)
	require.Equal(t, "external_account", externalAccount.Object)
}
