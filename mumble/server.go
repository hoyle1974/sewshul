package mumble

import (
	"crypto/rsa"
	"fmt"
	"io"
	"net/http"

	"github.com/quic-go/quic-go"
	"github.com/quic-go/quic-go/http3"
	"google.golang.org/protobuf/proto"

	"github.com/google/uuid"
	pb "github.com/hoyle1974/sewshul/proto"
	"github.com/hoyle1974/sewshul/services"
)

/*
Mumble is a gossip protocol build on top fo SPDY designed to propagate chunks of data to groups of users

The data looks like this

	id - unique id
	owner - unique id of the owner
	timestamp - the time the data was created
	distro_list - set of unique ids of who the data should be distributed to
	data - binary blob of data, this could be a post, a like, a comment on a post
	signature - the signature of all the data, signed with the owners public key
*/

type MumbleServer interface {
	Start() error
	Distribute([]byte, []services.AccountId) (uuid.UUID, error)
}

type mumbleServer struct {
	ownerId       services.AccountId
	pubKey        *rsa.PublicKey
	privKey       *rsa.PrivateKey
	distroStorage DistroStorage
}

func NewMumbleServer(ownerId services.AccountId, pubKey *rsa.PublicKey, privKey *rsa.PrivateKey) MumbleServer {
	return &mumbleServer{
		ownerId:       ownerId,
		pubKey:        pubKey,
		privKey:       privKey,
		distroStorage: NewDistroStorage(ownerId),
	}
}

func (m *mumbleServer) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/distribute", func(writer http.ResponseWriter, req *http.Request) {
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Printf("error reading body while handling /distribute: %s\n", err.Error())
		}
		mm := &pb.MumbleMurmer{}
		proto.Unmarshal(body, mm)
		m.ReceiveMurmer(mm)
	})

	quicConf := &quic.Config{}

	server := http3.Server{
		Handler:    mux,
		Addr:       "localhost:8911",
		QuicConfig: quicConf,
	}

	return server.ListenAndServe()

}

func (m *mumbleServer) ReceiveMurmer(mm *pb.MumbleMurmer) {
	// We have recieved a murmer, figure out what to do next
	data := MumbleDataFromPB(mm)
	m.distroStorage.Store(data)
}

func (m *mumbleServer) Distribute(data []byte, distroList []services.AccountId) (uuid.UUID, error) {
	mdata, err := NewMumbleData(m.ownerId, distroList, data, m.privKey)

	// Store the message and a background process will work on distributing it
	m.distroStorage.Store(mdata)

	return mdata.id, err
}
