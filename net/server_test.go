package net

import (
    "github.com/golang/mock/gomock"
    "net"
    "testing"
)

func TestHandleConn(t *testing.T) {
    testCases := []struct {
       name    string
       mock    func(ctrl *gomock.Controller)net.Conn
       wantErr error
    }{
       {
          name: "read Error",
          mock:
       },
    }
    for _, tc := range testCases {
       t.Run(tc.name, func(t *testing.T) {
          err := handleConn()
       })
    }
}