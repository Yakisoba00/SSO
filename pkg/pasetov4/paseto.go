package pasetov4

import (
	"fmt"
	"time"

	"aidanwoods.dev/go-paseto"
)

type Authenticator interface {
	VerifyToken(token string) (data *paseto.Token, err error)
}

type Provider struct {
	publicKey     paseto.V4AsymmetricPublicKey
	privateKey    *paseto.V4AsymmetricSecretKey
	accessExpiry  time.Duration
	refreshExpiry time.Duration
	issuer        string
}

func NewAuthenticator(issuer, publicKeyHex string) (Authenticator, error) {
	pub, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		return nil, ErrInvalidKey
	}

	return &Provider{
		issuer:     issuer,
		publicKey:  pub,
		privateKey: nil,
	}, nil
}

func NewProvider(issuer, publicKeyHex, privateKeyHex string, accessTTL, refreshTTL time.Duration) (*Provider, error) {
	pub, err := paseto.NewV4AsymmetricPublicKeyFromHex(publicKeyHex)
	if err != nil {
		return nil, ErrInvalidKey
	}

	prt, err := paseto.NewV4AsymmetricSecretKeyFromHex(privateKeyHex)
	if err != nil {
		return nil, ErrInvalidKey
	}

	return &Provider{
		publicKey:     pub,
		privateKey:    &prt,
		accessExpiry:  accessTTL,
		refreshExpiry: refreshTTL,
		issuer:        issuer,
	}, nil
}

func (p *Provider) VerifyToken(token string) (data *paseto.Token, err error) {
	parser := paseto.NewParser()
	parser.AddRule(paseto.IssuedBy(p.issuer))

	data, err = parser.ParseV4Public(p.publicKey, token, nil)
	if err != nil {
		return nil, ErrInvalidToken
	}

	return data, nil
}

func (p *Provider) GeneratePair(accessClaims, refreshClaims map[string]interface{}) (access, refresh string, err error) {
	access, err = p.generateToken(p.accessExpiry, accessClaims)
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrTokenCreation, err)
	}
	refresh, err = p.generateToken(p.refreshExpiry, refreshClaims)
	if err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrTokenCreation, err)
	}

	return access, refresh, nil
}

func (p *Provider) Refresh(refreshToken string) (claims map[string]interface{}, err error) {
	token, err := p.VerifyToken(refreshToken)
	if err != nil {
		return nil, err
	}

	return token.Claims(), nil
}

func (p *Provider) GetTokenClaims(token string) (claims map[string]interface{}, err error) {
	verifiedToken, err := p.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	return verifiedToken.Claims(), nil
}

func (p *Provider) generateToken(duration time.Duration, claims map[string]interface{}) (token string, err error) {
	pasetoToken := paseto.NewToken()
	pasetoToken.SetIssuer(p.issuer)
	pasetoToken.SetExpiration(time.Now().Add(duration))

	for key, value := range claims {
		err = pasetoToken.Set(key, value)
		if err != nil {
			return "", fmt.Errorf("%w: %s", ErrClaimSetting, key)
		}
	}

	return pasetoToken.V4Sign(*p.privateKey, nil), nil
}
