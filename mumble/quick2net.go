package mumble

import (
	"context"
	"net"
	"time"

	"github.com/quic-go/quic-go"
)

type QuickToNetListenerAdapter struct {
	lis quic.Listener
}

func (q QuickToNetListenerAdapter) Accept() (net.Conn, error) {
	qcon, err := q.lis.Accept(context.Background())

	return QuickToNetConnAdapter{qcon}, err
}

func (q QuickToNetListenerAdapter) Close() error {
	return q.lis.Close()
}

func (q QuickToNetListenerAdapter) Addr() net.Addr {
	return q.lis.Addr()
}

type QuickToNetConnAdapter struct {
	conn quic.Connection
}

// Read reads data from the connection.
// Read can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetReadDeadline.
func (q QuickToNetConnAdapter) Read(b []byte) (int, error) {
	b2, err := q.conn.ReceiveMessage()
	copy(b, b2)

	return len(b2), err
}

// Write writes data to the connection.
// Write can be made to time out and return an error after a fixed
// time limit; see SetDeadline and SetWriteDeadline.
func (q QuickToNetConnAdapter) Write(b []byte) (int, error) {
	err := q.conn.SendMessage(b)

	return len(b), err
}

// Close closes the connection.
// Any blocked Read or Write operations will be unblocked and return errors.
func (q QuickToNetConnAdapter) Close() error {
	return q.conn.CloseWithError(0, "")
}

// LocalAddr returns the local network address, if known.
func (q QuickToNetConnAdapter) LocalAddr() net.Addr {
	return q.conn.LocalAddr()
}

// RemoteAddr returns the remote network address, if known.
func (q QuickToNetConnAdapter) RemoteAddr() net.Addr {
	return q.conn.RemoteAddr()
}

// SetDeadline sets the read and write deadlines associated
// with the connection. It is equivalent to calling both
// SetReadDeadline and SetWriteDeadline.
//
// A deadline is an absolute time after which I/O operations
// fail instead of blocking. The deadline applies to all future
// and pending I/O, not just the immediately following call to
// Read or Write. After a deadline has been exceeded, the
// connection can be refreshed by setting a deadline in the future.
//
// If the deadline is exceeded a call to Read or Write or to other
// I/O methods will return an error that wraps os.ErrDeadlineExceeded.
// This can be tested using errors.Is(err, os.ErrDeadlineExceeded).
// The error's Timeout method will return true, but note that there
// are other possible errors for which the Timeout method will
// return true even if the deadline has not been exceeded.
//
// An idle timeout can be implemented by repeatedly extending
// the deadline after successful Read or Write calls.
//
// A zero value for t means I/O operations will not time out.
func (q QuickToNetConnAdapter) SetDeadline(t time.Time) error {
	return nil
}

// SetReadDeadline sets the deadline for future Read calls
// and any currently-blocked Read call.
// A zero value for t means Read will not time out.
func (q QuickToNetConnAdapter) SetReadDeadline(t time.Time) error {
	return nil
}

// SetWriteDeadline sets the deadline for future Write calls
// and any currently-blocked Write call.
// Even if write times out, it may return n > 0, indicating that
// some of the data was successfully written.
// A zero value for t means Write will not time out.
func (q QuickToNetConnAdapter) SetWriteDeadline(t time.Time) error {
	return nil
}
