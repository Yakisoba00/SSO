package pasetov4

import (
	"testing"
	"time"
)

const (
	testIssuer = "rewind-sso"
	// Ключи в hex формате (валидные для V4)
	pubHex = "fe515220ef3fcaef3ebc9f46f888af00027876d7a4b109b52356bd3ecd0d8eb3"
	prtHex = "6358cfc5929b4863a5952d208597622e1b2f2a80b75a51fba30c4c36d514c746fe515220ef3fcaef3ebc9f46f888af00027876d7a4b109b52356bd3ecd0d8eb3"
)

var (
	accessClaims  = map[string]interface{}{"user_id": "123", "role": "admin"}
	refreshClaims = map[string]interface{}{"user_id": "123", "session_id": "abc"}
)

func setupAuthenticator(t *testing.T) Authenticator {
	authenticator, err := NewAuthenticator(testIssuer, pubHex)
	if err != nil {
		t.Fatalf("Failed to create authenticator: %v", err)
	}

	return authenticator
}

func setupProvider(t *testing.T) *Provider {
	aTTL := 15 * time.Minute
	rTTL := 24 * time.Hour
	provider, err := NewProvider(testIssuer, pubHex, prtHex, aTTL, rTTL)
	if err != nil {
		t.Fatalf("Failed to create Provider: %v", err)
	}

	return provider
}

func TestProvider_GeneratePair(t *testing.T) {
	p := setupProvider(t)

	_, _, err := p.GeneratePair(accessClaims, refreshClaims)
	if err != nil {
		t.Fatalf("Error GeneratePair: %v", err)
	}
}

func TestProvider_VerifyAndClaims(t *testing.T) {
	p := setupProvider(t)
	expectedID := "user-99"

	claims := map[string]interface{}{"user_id": expectedID}
	tokenStr, _ := p.generateToken(time.Minute, claims)

	token, err := p.VerifyToken(tokenStr)
	if err != nil {
		t.Fatalf("VerifyToken провалился: %v", err)
	}

	extractedClaims := token.Claims()
	if extractedClaims["user_id"] != expectedID {
		t.Errorf("Ожидался user_id %s, получен %v", expectedID, extractedClaims["user_id"])
	}

	directClaims, err := p.GetTokenClaims(tokenStr)
	if err != nil {
		t.Fatalf("GetTokenClaims провалился: %v", err)
	}
	if directClaims["user_id"] != expectedID {
		t.Errorf("GetTokenClaims вернул неверные данные")
	}
}

func TestProvider_Refresh(t *testing.T) {
	p := setupProvider(t)
	expectedSessionID := "unique-session-id"

	claims := map[string]interface{}{"session_id": expectedSessionID}
	oldRefresh, _ := p.generateToken(time.Hour, claims)

	extractedClaims, err := p.Refresh(oldRefresh)
	if err != nil {
		t.Fatalf("Refresh провалился: %v", err)
	}

	if extractedClaims["session_id"] != expectedSessionID {
		t.Errorf("Ожидался session_id %s, получен %v", expectedSessionID, extractedClaims["session_id"])
	}
}

func TestProvider_InvalidIssuer(t *testing.T) {
	aTTL, rTTL := time.Minute, time.Hour

	pWrong, _ := NewProvider("fake-issuer", pubHex, prtHex, aTTL, rTTL)
	tokenFromWrong, _ := pWrong.generateToken(time.Minute, nil)

	pMain := setupProvider(t)
	_, err := pMain.VerifyToken(tokenFromWrong)

	if err == nil {
		t.Error("VerifyToken должен был выдать ошибку из-за неверного Issuer")
	}
}

func TestAuthenticator(t *testing.T) {
	p := setupProvider(t)

	access, refresh, err := p.GeneratePair(accessClaims, refreshClaims)
	if err != nil {
		t.Fatalf("Error GeneratePair: %v", err)
	}

	a := setupAuthenticator(t)

	acc, err := a.VerifyToken(access)
	if err != nil {
		t.Errorf("Error VerifyToken: %v", err)
	}
	t.Log(acc)

	ref, err := a.VerifyToken(refresh)
	if err != nil {
		t.Errorf("Error VerifyToken: %v", err)
	}
	t.Log(ref)
}
