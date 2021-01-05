package connector

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"gopirate.com/instances"
	"io"
	"io/ioutil"
	"net"
	"os"
)

//Endpoint ...
type Endpoint struct {
	Host string
	Port int
}

func (endpoint *Endpoint) String() string {
	return fmt.Sprintf("%s:%d", endpoint.Host, endpoint.Port)
}

//SSHtunnel ...
type SSHtunnel struct {
	Local  *Endpoint
	Server *Endpoint
	Remote *Endpoint
	Config *ssh.ClientConfig
}

//Start ...
func (tunnel *SSHtunnel) Start() error {
	listener, err := net.Listen("tcp", tunnel.Local.String())
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go tunnel.forward(conn)
	}
}

func (tunnel *SSHtunnel) forward(localConn net.Conn) {
	serverConn, err := ssh.Dial("tcp", tunnel.Server.String(), tunnel.Config)
	if err != nil {
		fmt.Printf("Server dial error: %s\n", err)
		return
	}

	remoteConn, err := serverConn.Dial("tcp", tunnel.Remote.String())
	if err != nil {
		fmt.Printf("Remote dial error: %s\n", err)
		return
	}

	copyConn := func(writer, reader net.Conn) {
		defer writer.Close()
		defer reader.Close()

		_, err := io.Copy(writer, reader)
		if err != nil {
			fmt.Printf("io.Copy error: %s", err)
		}
	}

	go copyConn(localConn, remoteConn)
	go copyConn(remoteConn, localConn)
}

//SSHAgent ...
func SSHAgent() ssh.AuthMethod {

	if sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers)
	}
	return nil
}

//PublicKeyFile ...
func PublicKeyFile(file string) ssh.AuthMethod {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil
	}
	return ssh.PublicKeys(key)
}

//CreateTunnel ...
func CreateTunnel(remote Instance) {
	//we need to manage multi tunnels concurunt connections
	for _, v := range remote.IP {
		fmt.Println("connecting to", string(v))
		localEndpoint := &Endpoint{
			Host: "localhost",
			Port: 9966,
		}

		serverEndpoint := &Endpoint{
			Host: string(v),
			Port: 22,
		}

		remoteEndpoint := &Endpoint{
			Host: "localhost",
			Port: 6699,
		}

		sshConfig := &ssh.ClientConfig{
			User: "ec2-user",
			Auth: []ssh.AuthMethod{
				PublicKeyFile("/home/ahmed/.goPirate/gonhuntKey.pem")},
			// Auth: []ssh.AuthMethod{
			// 	SSHAgent(),
			// },
		}

		tunnel := &SSHtunnel{
			Config: sshConfig,
			Local:  localEndpoint,
			Server: serverEndpoint,
			Remote: remoteEndpoint,
		}

		tunnel.Start()
	}

}
