package cowww

import (
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type ServerConfig struct {
	Host string
	Port string
}

type Server struct {
	Config ServerConfig
}

func NewServer(c ServerConfig) *Server {
	return &Server{
		Config: c,
	}
}

func (s *Server) Start() {
	l, err := net.Listen("tcp", s.Config.Host+":"+s.Config.Port)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Printf("Listening on %s:%s", s.Config.Host, s.Config.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT)
	go func() {
		<-c

		log.Println("Received SIGINT signal. Stopping server...")
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
		}

		go handleRequest(conn)
	}
}
