package glog

import (
	"testing"
	"github.com/golang/glog"
	"errors"
)

func TestGlog0(t *testing.T)  {
	err := errors.New("hi")
	glog.Error(err)
}

