package net

import (
    "errors"
    "net"
)

func Server(addr string) error {
    listener, err := net.Listen("tcp", addr)
    if err != nil {
       return err
    }

    for {
       conn, err := listener.Accept()
       if err != nil {
          return err
       }
       go func() {
          handleConn(conn)
       }()
    }
}

func handleConn(conn net.Conn) error {
    for {
       bs := make([]byte, 8)
       n, err := conn.Read(bs)
       if err != nil {
          return err
       }
       if n != 8 {
          return errors.New("micro: Not Read Enough")
       }
       res := handleMsg(bs)
       n, err = conn.Write(res)
       if err != nil {
          return err
       }
       if n != len(res) {
          return errors.New("micro: Not Write Enough")
       }
    }
}

func handleMsg(req []byte) []byte {
    res := make([]byte, 2*len(req))
    copy(res[:len(req)], req)
    copy(res[len(req):], req)
    return res
}