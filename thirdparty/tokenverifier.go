package thirdparty

import (
	"context"

	oidc "github.com/coreos/go-oidc/v3/oidc"

	"github.com/pkg/errors"
)

// Verifier a interface hold the third party verifier
type Verifier interface {
	VerifyToken(ctx context.Context, token string) (*oidc.IDToken, error)
}

// OidcVerifier oidc verifier
type OidcVerifier struct {
	Provider *oidc.Provider
}

// VerifyToken verify idToken of oauth proto with oidc
func (o *OidcVerifier) VerifyToken(ctx context.Context, token string) (*oidc.IDToken, error) {
	if o.Provider == nil {
		return nil, errors.New("invalid provider")
	}
	if token == "" {
		return nil, errors.New("invalid token")
	}
	oidcConfig := &oidc.Config{SkipClientIDCheck: true}
	verifier := o.Provider.Verifier(oidcConfig)
	t, err := verifier.Verify(ctx, token)
	if err != nil {
		return nil, errors.Wrap(err, "verify token failed")
	}
	return t, nil
}
