package secureString

import (
	"encoding/json"
	"errors"
	"math/rand"
	"sync"
	"time"
)

type SecureString struct {
	val []rune
	key int
	sec bool
	mu  sync.Mutex
}

// NewSecureString secures the sensitive string with a random key and returns it as a SecureString.
//
// Note: It is not cryptographically safe.
func NewSecureString(str string) SecureString {
	s := SecureString{}
	s.key = s.newKey()
	s.val = s.xorEn([]rune(str))
	s.sec = true
	return s
}

func (s *SecureString) UnmarshalJSON(data []byte) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var v interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.key = s.newKey()
	s.val = s.xorEn([]rune(v.(string)))
	s.sec = true
	return nil
}

// Set secures the given sensitive string with a new random key and returns it as a SecureString.
func (s *SecureString) Set(str string) *SecureString {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.sec = false
	s.key = s.newKey()
	s.val = s.xorEn([]rune(str))
	s.sec = true
	return s
}

// Get returns the secured sensitive string.
func (s *SecureString) Get() (string, error) {
	if !s.sec {
		return "", errors.New("no secure value")
	}
	return string(s.xorDe()), nil
}

// KeyChange secures the sensitive string with a new random key.
//
// Note: It is not cryptographically safe.
func (s *SecureString) KeyChange() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	str, err := s.Get()
	if err != nil {
		return err
	}

	s.key = s.newKey()
	s.val = s.xorEn([]rune(str))
	return nil
}

func (s *SecureString) newKey() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(int(^uint(0) >> 1))
}

func (s *SecureString) xorEn(val []rune) []rune {
	res := make([]rune, len(val))
	for i, v := range val {
		m := int32((i % 20) + 1)
		res[i] = (v * m) ^ int32(s.key)
	}
	return res
}

func (s *SecureString) xorDe() []rune {
	res := make([]rune, len(s.val))
	for i, v := range s.val {
		m := int32((i % 20) + 1)
		res[i] = (v ^ int32(s.key)) / m
	}
	return res
}
