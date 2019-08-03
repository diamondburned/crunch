/*

crunch - utilities for taking bytes out of things
copyright (c) 2019 superwhiskers <whiskerdev@protonmail.com>

this program is free software: you can redistribute it and/or modify
it under the terms of the gnu lesser general public license as published by
the free software foundation, either version 3 of the license, or
(at your option) any later version.

this program is distributed in the hope that it will be useful,
but without any warranty; without even the implied warranty of
merchantability or fitness for a particular purpose.  see the
gnu lesser general public license for more details.

you should have received a copy of the gnu lesser general public license
along with this program.  if not, see <https://www.gnu.org/licenses/>.

*/

package crunch

import (
	"unsafe"
)

// ReadBits stores the next n bits from the specified offset without modifying the internal offset value in out
func (b *MiniBuffer) ReadBitsC(out *uint64, off, n int64) {

	var (
		bout byte
	)

	for i := int64(0); i < n; i++ {
		b.ReadBit(&bout, off+i)
		*out = (*out << uint64(1)) | uint64(bout)
	}
}

// SetBits sets the next n bits from the specified offset without modifying the internal offset value
func (b *MiniBuffer) SetBitsC(off int64, data uint64, n int64) {

	for i := int64(0); i < n; i++ {
		if byte((data>>uint64(n-i-1))&1) == 0 {

			b.ClearBit(off + i)

		} else {

			b.SetBit(off + i)

		}
	}

}

// ClearAllBits sets all of the buffer's bits to 0
func (b *MiniBuffer) ClearAllBitsC() {

	var (
		n = int64(len(b.buf))
	)

	for i := int64(0); i < n; i++ {
		b.buf[i] = 0
	}

}

// SetAllBits sets all of the buffer's bits to 1
func (b *MiniBuffer) SetAllBitsC() {

	var (
		n = int64(len(b.buf))
	)
	for i := int64(0); i < n; i++ {
		b.buf[i] = 0xFF
	}

}

// FlipAllBits flips all of the buffer's bits
func (b *MiniBuffer) FlipAllBitsC() {

	var (
		n = int64(len(b.buf))
	)
	for i := int64(0); i < n; i++ {
		b.buf[i] = ^b.buf[i]
	}

}

/* byte buffer methods */

// WriteBytes writes bytes to the buffer at the specified offset without modifying the internal offset value
func (b *MiniBuffer) WriteBytesC(off int64, data []byte) {

	/* i'm just leaving this here incase this new method proves to be slower in some edge cases */
	/*var (
		i = int64(0)
		n = int64(len(data))
	)
	{
	write_loop:
		b.buf[off+i] = data[i]
		i++
		if i < n {

			goto write_loop

		}
	}*/

	var (
		p = uintptr(off) + b.obuf
		n = int64(len(data))
	)
	for i := int64(0); i < n; i++ {
		*(*byte)(unsafe.Pointer(p)) = data[i]
		p++
	}

}

// WriteU16LE writes a slice of uint16s to the buffer at the specified offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU16LEC(off int64, data []uint16) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*2)] = byte(data[i])
		b.buf[off+int64(1+(i*2))] = byte(data[i] >> 8)
	}

}

// WriteU16BE writes a slice of uint16s to the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU16BEC(off int64, data []uint16) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*2)] = byte(data[i] >> 8)
		b.buf[off+int64(1+(i*2))] = byte(data[i])
	}

}

// WriteU32LE writes a slice of uint32s to the buffer at the specified offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU32LEC(off int64, data []uint32) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*4)] = byte(data[i])
		b.buf[off+int64(1+(i*4))] = byte(data[i] >> 8)
		b.buf[off+int64(2+(i*4))] = byte(data[i] >> 16)
		b.buf[off+int64(3+(i*4))] = byte(data[i] >> 24)
	}

}

// WriteU32BE writes a slice of uint32s to the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU32BEC(off int64, data []uint32) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*4)] = byte(data[i] >> 24)
		b.buf[off+int64(1+(i*4))] = byte(data[i] >> 16)
		b.buf[off+int64(2+(i*4))] = byte(data[i] >> 8)
		b.buf[off+int64(3+(i*4))] = byte(data[i])
	}

}

// WriteU64LE writes a slice of uint64s to the buffer at the specfied offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU64LEC(off int64, data []uint64) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*8)] = byte(data[i])
		b.buf[off+int64(1+(i*8))] = byte(data[i] >> 8)
		b.buf[off+int64(2+(i*8))] = byte(data[i] >> 16)
		b.buf[off+int64(3+(i*8))] = byte(data[i] >> 24)
		b.buf[off+int64(4+(i*8))] = byte(data[i] >> 32)
		b.buf[off+int64(5+(i*8))] = byte(data[i] >> 40)
		b.buf[off+int64(6+(i*8))] = byte(data[i] >> 48)
		b.buf[off+int64(7+(i*8))] = byte(data[i] >> 56)
	}

}

// WriteU64BE writes a slice of uint64s to the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) WriteU64BEC(off int64, data []uint64) {

	var (
		n = len(data)
	)
	for i := 0; i < n; i++ {
		b.buf[off+int64(i*8)] = byte(data[i] >> 56)
		b.buf[off+int64(1+(i*8))] = byte(data[i] >> 48)
		b.buf[off+int64(2+(i*8))] = byte(data[i] >> 40)
		b.buf[off+int64(3+(i*8))] = byte(data[i] >> 32)
		b.buf[off+int64(4+(i*8))] = byte(data[i] >> 24)
		b.buf[off+int64(5+(i*8))] = byte(data[i] >> 16)
		b.buf[off+int64(6+(i*8))] = byte(data[i] >> 8)
		b.buf[off+int64(7+(i*8))] = byte(data[i])
	}

}

// ReadU16LE reads a slice of uint16s from the buffer at the specified offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU16LEC(out *[]uint16, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint16(b.buf[off+(i*2)]) |
			uint16(b.buf[off+(1+(i*2))])<<8
	}

}

// ReadU16BE reads a slice of uint16s from the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU16BEC(out *[]uint16, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint16(b.buf[off+(1+(i*2))]) |
			uint16(b.buf[off+(i*2)])<<8
	}

}

// ReadU32LE reads a slice of uint32s from the buffer at the specified offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU32LEC(out *[]uint32, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint32(b.buf[off+(i*4)]) |
			uint32(b.buf[off+(1+(i*4))])<<8 |
			uint32(b.buf[off+(2+(i*4))])<<16 |
			uint32(b.buf[off+(3+(i*4))])<<24
	}

}

// ReadU32BE reads a slice of uint32s from the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU32BEC(out *[]uint32, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint32(b.buf[off+(3+(i*4))]) |
			uint32(b.buf[off+(2+(i*4))])<<8 |
			uint32(b.buf[off+(1+(i*4))])<<16 |
			uint32(b.buf[off+(i*4)])<<24
	}

}

// ReadU64LE reads a slice of uint64s from the buffer at the specified offset in little-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU64LEC(out *[]uint64, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint64(b.buf[off+(i*8)]) |
			uint64(b.buf[off+(1+(i*8))])<<8 |
			uint64(b.buf[off+(2+(i*8))])<<16 |
			uint64(b.buf[off+(3+(i*8))])<<24 |
			uint64(b.buf[off+(4+(i*8))])<<32 |
			uint64(b.buf[off+(5+(i*8))])<<40 |
			uint64(b.buf[off+(6+(i*8))])<<48 |
			uint64(b.buf[off+(7+(i*8))])<<56
	}

}

// ReadU64BE reads a slice of uint64s from the buffer at the specified offset in big-endian without modifying the internal offset value
func (b *MiniBuffer) ReadU64BEC(out *[]uint64, off, n int64) {

	for i := int64(0); i < n; i++ {
		(*out)[i] = uint64(b.buf[off+(7+(i*8))]) |
			uint64(b.buf[off+(6+(i*8))])<<8 |
			uint64(b.buf[off+(5+(i*8))])<<16 |
			uint64(b.buf[off+(4+(i*8))])<<24 |
			uint64(b.buf[off+(3+(i*8))])<<32 |
			uint64(b.buf[off+(2+(i*8))])<<40 |
			uint64(b.buf[off+(1+(i*8))])<<48 |
			uint64(b.buf[off+(i*8)])<<56
	}

}
