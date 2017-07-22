package gost

import (
	"net"
)

// Client is a proxy client.
type Client struct {
	Connector   Connector
	Transporter Transporter
}

// NewClient creates a proxy client.
// A client is divided into two layers: connector and transporter.
// Connector is responsible for connecting to the destination address through this proxy.
// Transporter performs a handshake with this proxy.
func NewClient(c Connector, tr Transporter) *Client {
	return &Client{
		Connector:   c,
		Transporter: tr,
	}
}

// Dial connects to the target address
func (c *Client) Dial(addr string) (net.Conn, error) {
	return net.Dial(c.Transporter.Network(), addr)
}

// Handshake performs a handshake with the proxy.
// The conn should be an connection to this proxy.
func (c *Client) Handshake(conn net.Conn) (net.Conn, error) {
	return c.Transporter.Handshake(conn)
}

// Connect connects to the address addr via the proxy.
// The conn should be an connection to this proxy.
func (c *Client) Connect(conn net.Conn, addr string) (net.Conn, error) {
	return c.Connector.Connect(conn, addr)
}

// DefaultClient is a standard HTTP proxy client
var DefaultClient = NewClient(HTTPConnector(nil), TCPTransporter())

// Dial connects to the address addr via the DefaultClient.
func Dial(addr string) (net.Conn, error) {
	return DefaultClient.Dial(addr)
}

// Handshake performs a handshake via the DefaultClient
func Handshake(conn net.Conn) (net.Conn, error) {
	return DefaultClient.Handshake(conn)
}

// Connect connects to the address addr via the DefaultClient.
func Connect(conn net.Conn, addr string) (net.Conn, error) {
	return DefaultClient.Connect(conn, addr)
}

// Connector is responsible for connecting to the destination address
type Connector interface {
	Connect(conn net.Conn, addr string) (net.Conn, error)
}

// Transporter is responsible for handshaking with the proxy server.
type Transporter interface {
	Network() string
	Handshake(conn net.Conn) (net.Conn, error)
}

type tcpTransporter struct {
}

func TCPTransporter() Transporter {
	return &tcpTransporter{}
}

func (tr *tcpTransporter) Network() string {
	return "tcp"
}

func (tr *tcpTransporter) Handshake(conn net.Conn) (net.Conn, error) {
	return conn, nil
}