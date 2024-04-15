package redisx

import (
	"context"
)

type Set interface {
	Redis
	Add(vs ...any) error
	Rem(vs ...any) error
	IsMember(v any) bool
	IsMembers(vs ...any) []bool
	Len() int64
}

func NewSet(rc Client, key string) Set {
	return &setImpl{rc: rc, key: key}
}

type setImpl struct {
	rc  Client
	key string
}

func (s *setImpl) Keys(p string) ([]string, error) {
	cmd := s.rc.Keys(context.Background(), p)
	return cmd.Val(), cmd.Err()
}

func (s *setImpl) Add(vs ...any) error {
	cmd := s.rc.SAdd(context.Background(), s.key, vs...)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (s *setImpl) Rem(vs ...any) error {
	cmd := s.rc.SRem(context.Background(), s.key, vs...)
	if cmd.Err() != nil {
		return cmd.Err()
	}
	return nil
}

func (s *setImpl) IsMember(v any) bool {
	cmd := s.rc.SIsMember(context.Background(), s.key, v)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	return cmd.Val()
}

func (s *setImpl) IsMembers(vs ...any) []bool {
	cmd := s.rc.SMIsMember(context.Background(), s.key, vs...)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	return cmd.Val()
}

func (s *setImpl) Len() int64 {
	cmd := s.rc.SCard(context.Background(), s.key)
	if cmd.Err() != nil {
		panic(cmd.Err())
	}
	return cmd.Val()
}
