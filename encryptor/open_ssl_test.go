package encryptor

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestOpenSSL_options(t *testing.T) {
	ctx := &OpenSSL{
		password: "foo(872",
		salt:     false,
		base64:   true,
		pbkdf2:	  true,
	}

	opts := strings.Join(ctx.options(), " ")
	assert.Equal(t, opts, "aes-256-cbc -base64 -pbkdf2 -pass pass:foo(872")

	ctx.salt = true
	opts = strings.Join(ctx.options(), " ")
	assert.Equal(t, opts, "aes-256-cbc -base64 -salt -pbkdf2 -pass pass:foo(872")

	ctx.base64 = false
	opts = strings.Join(ctx.options(), " ")
	assert.Equal(t, opts, "aes-256-cbc -salt -pbkdf2 -pass pass:foo(872")

	ctx.pbkdf2 = false
	opts = strings.Join(ctx.options(), " ")
	assert.Equal(t, opts, "aes-256-cbc -salt -pass pass:foo(872")
}
