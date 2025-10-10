package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github-bot/types"
	"net/http"

	"github.com/xySaad/trail"
)

const WEBHOOK_SECRET = "wxpx+#Jxxxxxsq>bLxxxxxx*vq"

func GithubSignature(c *trail.Context[types.Dependecies]) bool {
	// simulate valid signature
	if c.Dep.SkipSignature {
		return true
	}
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
