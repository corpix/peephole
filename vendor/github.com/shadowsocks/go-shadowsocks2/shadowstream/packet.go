package shadowstream

import (
	"crypto/rand"
	"errors"
	"io"
	"net"
	"sync"
)

// ErrShortPacket means the packet is too short to be a valid encrypted packet.
var ErrShortPacket = errors.New("short packet")

// Pack encrypts plaintext using stream cipher s and a random IV.
// Returns a slice of dst containing random IV and ciphertext.
// Ensure len(dst) >= s.IVSize() + len(plaintext).
func Pack(dst, plaintext []byte, s Cipher) ([]byte, error) {
	if len(dst) < s.IVSize()+len(plaintext) {
		return nil, io.ErrShortBuffer
	}
	iv := dst[:s.IVSize()]
	_, err := io.ReadFull(rand.Reader, iv)
	if err != nil {
		return nil, err
	}

	s.Encrypter(iv).XORKeyStream(dst[len(iv):], plaintext)
	return dst[:len(iv)+len(plaintext)], nil
}

// Unpack decrypts pkt using stream cipher s.
// Returns a slice of dst containing decrypted plaintext.
func Unpack(dst, pkt []byte, s Cipher) ([]byte, error) {
	if len(pkt) < s.IVSize() {
		return nil, ErrShortPacket
	}

	if len(dst) < len(pkt)-s.IVSize() {
		return nil, io.ErrShortBuffer
	}
	iv := pkt[:s.IVSize()]
	s.Decrypter(iv).XORKeyStream(dst, pkt[len(iv):])
	return dst[:len(pkt)-len(iv)], nil
}

type packetConn struct {
	net.PacketConn
	Cipher
	buf        []byte
	sync.Mutex // write lock
}

// NewPacketConn wraps a net.PacketConn with stream cipher encryption/decryption.
func NewPacketConn(c net.PacketConn, ciph Cipher) net.PacketConn {
	return &packetConn{PacketConn: c, Cipher: ciph, buf: make([]byte, 64*1024)}
}

func (c *packetConn) WriteTo(b []byte, addr net.Addr) (int, error) {
	c.Lock()
	defer c.Unlock()
	buf, err := Pack(c.buf, b, c.Cipher)
	if err != nil {
		return 0, err
	}
	_, err = c.PacketConn.WriteTo(buf, addr)
	return len(b), err
}

func (c *packetConn) ReadFrom(b []byte) (int, net.Addr, error) {
	n, addr, err := c.PacketConn.ReadFrom(b)
	if err != nil {
		return n, addr, err
	}
	b, err = Unpack(b, b[:n], c.Cipher)
	return len(b), addr, err
}
