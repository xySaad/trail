package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github-bot/models"
	"net/http"
)

const WEBHOOK_SECRET = "wxpx+#Jxxxxxsq>bLxxxxxx*vq"

func GithubSignature(c *models.Context) bool {
	// simulate valid signature
	// return true

	// Compute HMAC SHA256 of body
	mac := hmac.New(sha256.New, []byte(WEBHOOK_SECRET))
	mac.Write(c.BodyNoErr())
	expectedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))

	// Compare with GitHub signature
	if !hmac.Equal([]byte(expectedSignature), []byte(c.Header("X-Hub-Signature-256"))) {
		http.Error(c.Response, "invalid signature", http.StatusUnauthorized)
		return false
	}
	return true
}
