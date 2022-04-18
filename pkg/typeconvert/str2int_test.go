package typeconvert

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStr2int(t *testing.T) {
	r := Str2int("91", 0)
	assert.Equal(t, 91, r, "string 转换 int 失败.")
}

func BenchmarkStr2int1(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Str2int("91", 0)
	}
}
