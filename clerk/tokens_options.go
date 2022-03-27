package clerk

import "time"

// VerifyTokenOption describes a functional parameter for the VerifyToken method
type VerifyTokenOption func(*verifyTokenOptions) error

// WithAuthorizedParty allows to set the authorized parties to check against the azp claim of the session token
func WithAuthorizedParty(parties ...string) VerifyTokenOption {
	return func(o *verifyTokenOptions) error {
		authorizedParties := make(map[string]struct{})
		for _, party := range parties {
			authorizedParties[party] = struct{}{}
		}

		o.authorizedParties = authorizedParties
		return nil
	}
}

// WithLeeway allows to set a custom leeway that gives some extra time to the token to accomodate for clock skew, etc.
func WithLeeway(leeway time.Duration) VerifyTokenOption {
	return func(o *verifyTokenOptions) error {
		o.leeway = leeway
		return nil
	}
}
