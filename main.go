package main

import (
	"bytes"
	"flag"
	"io"
	"net"
	"os"

	"github.com/pkg/errors"
	"github.com/xackery/golisten/helper"
)

var (
	//Version of golisten
	Version string
)

func main() {
	log := helper.Log()
	err := run()
	if err != nil {
		log.Error().Err(err).Msg("main")
		os.Exit(1)
	}
	os.Exit(0)
}

func run() error {
	log := helper.Log()
	var host string
	var isListen bool
	var protocol string
	flag.StringVar(&host, "host", "", "host:port to connect to or bind to")
	flag.StringVar(&protocol, "protocol", "tcp", "which protocol to use. options: tcp, udp")
	flag.BoolVar(&isListen, "listen", false, "bind to port and listen for connections?")
	flag.Parse()
	if host == "" {
		flag.Usage()
		os.Exit(1)
	}
	if !isListen {
		log.Info().Str("host", host).Str("protocol", protocol).Msg("connecting...")
		conn, err := net.Dial(protocol, host)
		if err != nil {
			return err
		}
		log.Info().Msg("connected! Trying to read introduction data...")
		reply := make([]byte, 128)
		_, err = conn.Read(reply)
		if err != nil {
			log.Warn().Err(err).Msg("failed to read messages from server. This can be ignored")
			return nil
		}
		log.Info().Bytes("reply", reply).Msg("success")
		return nil
	}
	log.Info().Str("host", host).Str("protocol", protocol).Msg("listening...")
	conn, err := net.Listen(protocol, host)
	if err != nil {
		return errors.Wrap(err, "listen")
	}
	for {
		conn, err := conn.Accept()
		if err != nil {
			return errors.Wrap(err, "accept connection")
		}
		log.Info().Str("remote host", conn.RemoteAddr().String()).Msg("got connection")
		io.Copy(conn, bytes.NewBufferString("connnection successful to golisten!"))
		conn.Close()
		log.Info().Str("remote host", conn.RemoteAddr().String()).Msg("replied and closing connection")
	}
}
