package cmd

import (
	"context"
	"redis-like/storage"
	"strconv"
	"strings"
)

type SetCmd struct {
	key   []byte
	value []byte
	ex    bool
	exTtl int64
	px    bool
	pxTtl int64
	nx    bool
	xx    bool
}

func (s *SetCmd) Init(bs [][]byte) error {
	s.key = bs[0]
	s.value = bs[1]
	for i := 2; i < len(bs); i = i + 2 {
		temp := string(bs[i])
		if strings.Compare(temp, "ex") == 0 {
			ttl, err := strconv.ParseInt(string(bs[i+1]), 10, 64)
			if err != nil {
				return err
			}
			s.ex = true
			s.exTtl = ttl
		} else if strings.Compare(temp, "px") == 0 {
			ttl, err := strconv.ParseInt(string(bs[i+1]), 10, 64)
			if err != nil {
				return err
			}
			s.px = true
			s.pxTtl = ttl
		} else if strings.Compare(temp, "nx") == 0 {
			s.nx = true
		} else if strings.Compare(temp, "xx") == 0 {
			s.xx = true
		}
	}
	return nil
}

func (s *SetCmd) Deal(ctx context.Context) []byte {
	storage := storage.StorageInstance()
	err := storage.Set(context.Background(), s.key, s.value)
	if err == nil {
		return OK
	} else {
		return CommonErr
	}
}
