package util_test

import (
	"github.com/axetroy/go-server/src/util"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	secret string
)

func TestGenerate2FASecret(t *testing.T) {
	secret, err := util.Generate2FASecret("101645075095748608")
	assert.Nil(t, err)
	assert.Len(t, secret, 32)
}

func TestVerify2FA(t *testing.T) {
	_, err := util.Generate2FASecret("101645075095748608")
	assert.Nil(t, err)
	assert.False(t, util.Verify2FA("101645075095748608", "12345678"))
}
