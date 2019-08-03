package crunch

import (
	"testing"
)

func BenchmarkMiniBufferCstyleWriteBytes(b *testing.B) {

	b.ReportAllocs()

	buf := &MiniBuffer{}
	NewMiniBuffer(&buf, []byte{0x00, 0x00, 0x00, 0x00})

	for n := 0; n < b.N; n++ {

		buf.WriteBytesC(0x00, []byte{0x01, 0x02})

	}

}

func BenchmarkMiniBufferCstyleWriteU32LE(b *testing.B) {

	b.ReportAllocs()

	buf := &MiniBuffer{}
	NewMiniBuffer(&buf, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	for n := 0; n < b.N; n++ {

		buf.WriteU32LEC(0x00, []uint32{0x01, 0x02})

	}

}

func BenchmarkMiniBufferCstyleReadU32LE(b *testing.B) {

	b.ReportAllocs()

	buf := &MiniBuffer{}
	NewMiniBuffer(&buf, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00})

	out := []uint32{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}
	for n := 0; n < b.N; n++ {

		buf.ReadU32LEC(&out, 0x00, 2)

	}

	_ = out

}

func BenchmarkMiniBufferCstyleReadBits(b *testing.B) {

	b.ReportAllocs()

	buf := &MiniBuffer{}
	NewMiniBuffer(&buf, []byte{0x00, 0x00, 0x00, 0x00})

	var out uint64
	for n := 0; n < b.N; n++ {

		buf.ReadBitsC(&out, 0x00, 2)

	}

	_ = out

}
