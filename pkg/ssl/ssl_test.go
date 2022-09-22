package ssl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseSSLCertificate(t *testing.T) {
	_, err := ParseSSLCertificate([]byte("asdsd"))
	
	assert.Nil(t, err)
	
}

func TestVerifyHostName(t *testing.T) {

}
