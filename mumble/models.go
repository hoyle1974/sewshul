package mumble

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/gob"
	"net"
	"time"

	"github.com/google/uuid"
	"github.com/hoyle1974/sewshul/services"
	"google.golang.org/protobuf/types/known/timestamppb"

	pb "github.com/hoyle1974/sewshul/proto"
)

type MumbleAddress struct {
	addr net.Addr
}

type MumbleData struct {
	id         uuid.UUID
	owner_id   services.AccountId
	timestamp  time.Time
	distroList []services.AccountId
	data       []byte
	signature  []byte
}

func NewMumbleData(ownerId services.AccountId, distroList []services.AccountId, data []byte, privKey *rsa.PrivateKey) (MumbleData, error) {
	m := MumbleData{
		id:         uuid.New(),
		owner_id:   ownerId,
		timestamp:  time.Now(),
		distroList: distroList,
		data:       data,
		signature:  nil,
	}

	buffer := bytes.Buffer{}
	enc := gob.NewEncoder(&buffer)
	enc.Encode(m.id)
	enc.Encode(m.owner_id.String())
	enc.Encode(m.timestamp.UnixNano())
	for _, d := range distroList {
		enc.Encode(d.String())
	}
	enc.Encode(data)

	msgHash := sha256.New()
	_, err := msgHash.Write(buffer.Bytes())
	if err != nil {
		return MumbleData{}, err
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, privKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return MumbleData{}, err
	}

	m.signature = signature

	return m, nil
}

func (m MumbleData) ToPB() pb.MumbleMurmer {
	return pb.MumbleMurmer{
		Id:         m.id.String(),
		OwnerId:    m.owner_id.String(),
		Timestamp:  timestamppb.New(m.timestamp),
		DistroList: services.AccountIdsToStrings(m.distroList),
		Data:       m.data,
		Signature:  m.signature,
	}
}

func MumbleDataFromPB(mm *pb.MumbleMurmer) MumbleData {
	id, err := uuid.Parse(mm.GetId())
	if err != nil {
		panic(err)
	}
	return MumbleData{
		id:         id,
		owner_id:   services.NewAccountId(mm.GetOwnerId()),
		timestamp:  mm.Timestamp.AsTime(),
		distroList: services.StringsToAccountIds(mm.GetDistroList()),
		data:       mm.GetData(),
		signature:  mm.GetSignature(),
	}
}
