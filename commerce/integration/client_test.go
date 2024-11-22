package integration

import (
	"testing"
)

func TestSubscriptionListByInstanceID(t *testing.T) {
	t.Parallel()
	// userID := "user_123"
	// config := &clerk.ClientConfig{}
	// config.HTTPClient = &http.Client{
	// 	Transport: &clerktest.RoundTripper{
	// 		T:      t,
	// 		Method: http.MethodDelete,
	// 		Out:    json.RawMessage(fmt.Sprintf(`{"id":"%s","object":"external_account"}`, externalAccountID)),
	// 		Path:   fmt.Sprintf("/v1/users/%s/external_accounts/%s", userID, externalAccountID),
	// 	},
	// }
	// client := NewClient(config)
	// externalAccount, err := client.DeleteExternalAccount(context.Background(), &DeleteExternalAccountParams{
	// 	UserID: userID,
	// 	ID:     externalAccountID,
	// })
	// require.NoError(t, err)
	// require.Equal(t, externalAccountID, externalAccount.ID)
	// require.Equal(t, "external_account", externalAccount.Object)
}
