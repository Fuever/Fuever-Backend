package secret

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAuthentication(t *testing.T) {
	userID := 114514
	InitTokenCache()
	token := GenerateTokenAndCache(userID)
	assert.True(t, Authentication(userID, token))
	assert.False(t, Authentication(1919810, token))
	token1 := GenerateTokenAndCache(1145141919810)
	assert.True(t, Authentication(1145141919810, token1))
	assert.False(t, Authentication(1145141919810, token))
}
