package organizationinvitation

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/clerk/clerk-sdk-go/v2/clerktest"
	"github.com/stretchr/testify/require"
)

func TestOrganizationInvitationClientCreate(t *testing.T) {
	t.Parallel()
	id := "orginv_123"
	organizationID := "org_123"
	emailAddress := "foo@bar.com"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			In:     json.RawMessage(fmt.Sprintf(`{"email_address":"%s", "expires_in_days": 1}`, emailAddress)),
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","email_address":"%s","organization_id":"%s", "expires_at": 1}`, id, emailAddress, organizationID)),
			Method: http.MethodPost,
			Path:   "/v1/organizations/" + organizationID + "/invitations",
		},
	}
	client := NewClient(config)
	invitation, err := client.Create(context.Background(), &CreateParams{
		OrganizationID: organizationID,
		EmailAddress:   clerk.String(emailAddress),
		ExpiresInDays:  clerk.Int64(1),
	})
	require.NoError(t, err)
	require.Equal(t, id, invitation.ID)
	require.Equal(t, organizationID, invitation.OrganizationID)
	require.Equal(t, emailAddress, invitation.EmailAddress)
	require.Equal(t, int64(1), *invitation.ExpiresAt)
}

func TestOrganizationInvitationClientCreate_Error(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Status: http.StatusBadRequest,
			Out: json.RawMessage(`{
  "errors":[{
		"code":"create-error-code"
	}],
	"clerk_trace_id":"create-trace-id"
}`),
		},
	}
	client := NewClient(config)
	_, err := client.Create(context.Background(), &CreateParams{})
	require.Error(t, err)
	apiErr, ok := err.(*clerk.APIErrorResponse)
	require.True(t, ok)
	require.Equal(t, "create-trace-id", apiErr.TraceID)
	require.Equal(t, 1, len(apiErr.Errors))
	require.Equal(t, "create-error-code", apiErr.Errors[0].Code)
}

func TestOrganizationInvitationClientList(t *testing.T) {
	t.Parallel()
	organizationID := "org_123"
	id := "orginv_123"
	statuses := []string{"pending", "accepted"}
	limit := int64(10)
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"data":[{"id":"%s","object":"organization_invitation","email_address":"string","role":"string","organization_id":"%s","status":"string","public_metadata":{},"private_metadata":{},"expires_at":1,"created_at":0,"updated_at":0}],"total_count":1}`, id, organizationID)),
			Method: http.MethodGet,
			Path:   "/v1/organizations/" + organizationID + "/invitations",
			Query: &url.Values{
				"limit":  []string{fmt.Sprintf("%d", limit)},
				"status": statuses,
			},
		},
	}
	client := NewClient(config)
	response, err := client.List(context.Background(), &ListParams{
		OrganizationID: organizationID,
		ListParams: clerk.ListParams{
			Limit: clerk.Int64(limit),
		},
		Statuses: &statuses,
	})
	require.NoError(t, err)
	require.Len(t, response.OrganizationInvitations, 1)
	require.Equal(t, id, response.OrganizationInvitations[0].ID)
	require.Equal(t, organizationID, response.OrganizationInvitations[0].OrganizationID)
	require.Equal(t, int64(1), *response.OrganizationInvitations[0].ExpiresAt)
	require.Equal(t, int64(1), response.TotalCount)
}

func TestOrganizationInvitationClientList_Error(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Status: http.StatusBadRequest,
			Out: json.RawMessage(`{
				"errors":[{
					"code":"list-error-code"
				}],
				"clerk_trace_id":"list-trace-id"
			}`),
		},
	}
	client := NewClient(config)
	_, err := client.List(context.Background(), &ListParams{OrganizationID: "org_123"})
	require.Error(t, err)
	apiErr, ok := err.(*clerk.APIErrorResponse)
	require.True(t, ok)
	require.Equal(t, "list-trace-id", apiErr.TraceID)
	require.Equal(t, 1, len(apiErr.Errors))
	require.Equal(t, "list-error-code", apiErr.Errors[0].Code)
}

func TestOrganizationInvitationClientGet(t *testing.T) {
	t.Parallel()
	organizationID := "org_123"
	id := "orginv_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","object":"organization_invitation","email_address":"string","role":"string","role_name":"string","organization_id":"%s","status":"string","public_metadata":{},"private_metadata":{},"expires_at": 1,"created_at":0,"updated_at":0}`, id, organizationID)),
			Method: http.MethodGet,
			Path:   "/v1/organizations/" + organizationID + "/invitations/" + id,
		},
	}
	client := NewClient(config)
	response, err := client.Get(context.Background(), &GetParams{
		OrganizationID: organizationID,
		ID:             id,
	})
	require.NoError(t, err)
	require.Equal(t, id, response.ID)
	require.Equal(t, organizationID, response.OrganizationID)
	require.Equal(t, "string", response.RoleName)
	require.Equal(t, "string", response.Role)
	require.Equal(t, int64(1), *response.ExpiresAt)
}

func TestOrganizationInvitationClientGet_Error(t *testing.T) {
	t.Parallel()
	organizationID := "org_123"
	id := "orginv_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Status: http.StatusBadRequest,
			Out: json.RawMessage(`{
				"errors":[{
					"code":"get-error-code"
				}],
				"clerk_trace_id":"get-trace-id"
			}`),
		},
	}
	client := NewClient(config)
	_, err := client.Get(context.Background(), &GetParams{
		OrganizationID: organizationID,
		ID:             id,
	})
	require.Error(t, err)
	apiErr, ok := err.(*clerk.APIErrorResponse)
	require.True(t, ok)
	require.Equal(t, "get-trace-id", apiErr.TraceID)
	require.Equal(t, 1, len(apiErr.Errors))
	require.Equal(t, "get-error-code", apiErr.Errors[0].Code)
}

func TestOrganizationInvitationClientRevoke(t *testing.T) {
	t.Parallel()
	organizationID := "org_123"
	id := "orginv_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			In:     json.RawMessage(`{"requesting_user_id": "user_123"}`),
			Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","object":"organization_invitation","email_address":"string","role":"string","organization_id":"%s","status":"string","public_metadata":{},"private_metadata":{},"created_at":0,"updated_at":0}`, id, organizationID)),
			Method: http.MethodPost,
			Path:   "/v1/organizations/" + organizationID + "/invitations/" + id + "/revoke",
		},
	}
	client := NewClient(config)
	response, err := client.Revoke(context.Background(), &RevokeParams{
		OrganizationID:   organizationID,
		RequestingUserID: clerk.String("user_123"),
		ID:               id,
	})
	require.NoError(t, err)
	require.Equal(t, id, response.ID)
	require.Equal(t, organizationID, response.OrganizationID)
}

func TestOrganizationInvitationClientRevoke_Error(t *testing.T) {
	t.Parallel()
	organizationID := "org_123"
	id := "orginv_123"
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Status: http.StatusBadRequest,
			Out: json.RawMessage(`{
				"errors":[{
					"code":"revoke-error-code"
				}],
				"clerk_trace_id":"revoke-trace-id"
			}`),
		},
	}
	client := NewClient(config)
	_, err := client.Revoke(context.Background(), &RevokeParams{
		OrganizationID: organizationID,
		ID:             id,
	})
	require.Error(t, err)
	apiErr, ok := err.(*clerk.APIErrorResponse)
	require.True(t, ok)
	require.Equal(t, "revoke-trace-id", apiErr.TraceID)
	require.Equal(t, 1, len(apiErr.Errors))
	require.Equal(t, "revoke-error-code", apiErr.Errors[0].Code)
}

func TestOrganizationInvitationClientListFromInstance(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Out:    json.RawMessage(`{"data":[{"id":"orginv_123","object":"organization_invitation","email_address":"string","role":"string","organization_id":"org_123","status":"string","public_metadata":{},"private_metadata":{},"created_at":0,"updated_at":0}],"total_count":1}`),
			Method: http.MethodGet,
			Path:   "/v1/organization_invitations",
			Query: &url.Values{
				"limit":    []string{"10"},
				"order_by": []string{"-created_at"},
				"query":    []string{"query"},
				"status":   []string{"pending", "accepted"},
			},
		},
	}
	client := NewClient(config)
	response, err := client.ListFromInstance(context.Background(), &ListFromInstanceParams{
		Statuses: &[]string{"pending", "accepted"},
		Query:    clerk.String("query"),
		OrderBy:  clerk.String("-created_at"),
		ListParams: clerk.ListParams{
			Limit: clerk.Int64(10),
		},
	})
	require.NoError(t, err)
	require.Len(t, response.OrganizationInvitations, 1)
	require.Equal(t, int64(1), response.TotalCount)
	require.Equal(t, "orginv_123", response.OrganizationInvitations[0].ID)
}

func TestOrganizationInvitationClientListFromInstance_Error(t *testing.T) {
	t.Parallel()
	config := &clerk.ClientConfig{}
	config.HTTPClient = &http.Client{
		Transport: &clerktest.RoundTripper{
			T:      t,
			Status: http.StatusBadRequest,
			Out: json.RawMessage(`{
				"errors":[{
					"code":"list-from-instance-error-code"
				}],
				"clerk_trace_id":"list-from-instance-trace-id"
			}`),
		},
	}
	client := NewClient(config)
	_, err := client.ListFromInstance(context.Background(), &ListFromInstanceParams{})
	require.Error(t, err)
	apiErr, ok := err.(*clerk.APIErrorResponse)
	require.True(t, ok)
	require.Equal(t, "list-from-instance-trace-id", apiErr.TraceID)
	require.Equal(t, 1, len(apiErr.Errors))
	require.Equal(t, "list-from-instance-error-code", apiErr.Errors[0].Code)
}
