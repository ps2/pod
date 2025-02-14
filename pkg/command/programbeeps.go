package command

import (
	"github.com/avereha/pod/pkg/response"
	log "github.com/sirupsen/logrus"
)

	type ProgramBeeps struct {
	Seq uint8
	ID  []byte
}

func UnmarshalProgramBeeps(data []byte) (*ProgramBeeps, error) {
	ret := &ProgramBeeps{}
	// TODO deserialize this command
	log.Debugf("ProgramBeeps, 0x1e, received, data %x", data)
	return ret, nil
}

func (g *ProgramBeeps) GetResponse() (response.Response, error) {
	// TODO improve responses
	return &response.GeneralStatusResponse{}, nil
}

func (g *ProgramBeeps) SetHeaderData(seq uint8, id []byte) error {
	g.ID = id
	g.Seq = seq
	return nil
}

func (g *ProgramBeeps) GetHeaderData() (uint8, []byte, error) {
	return g.Seq, g.ID, nil
}
