package details

import (
	"log"
	"net"
	"os"
)

func GetHostname() (string, error) {
	hostname, error := os.Hostname()
	return hostname, error
}

func GetIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}
