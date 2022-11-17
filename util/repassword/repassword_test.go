package repassword

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func TestRePassword(t *testing.T) {
	rand.Seed(time.Now().Unix())
	password := "关注嘉然 顿顿解馋"
	passwordHash := GeneratePasswordHash(password)
	assert.True(t, strings.Count(strings.SplitN(password, "$", 2)[0], "$") == 0)
	assert.True(t, CheckPasswordHash(password, passwordHash))
	assert.False(t, CheckPasswordHash(password, passwordHash+","))
}
