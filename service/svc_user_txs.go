package service

import (
	"os/exec"
	"time"

	"github.com/ic-matcom/api.dapp/lib"
	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/kataras/iris/v12"
)

// region ======== SETUP =================================================================

// ISvcBlockchainTxs Blockchain transactions service interface
type ISvcUserTxs interface {
	CreateUser(user dto.UserBlockchainCreate) *dto.Problem
	GetUserById(id string) (interface{}, *dto.Problem)
	DeleteUser(id string) *dto.Problem
	GetAllUsers() (interface{}, *dto.Problem)
}

type svcUsersTxs struct {
	repo *hlf.RepoUserBlockchain
}

// endregion =============================================================================

// NewSvcBlockchainTxs instantiate the HLF blockchains transactions services
func NewSvcUserTxs(pRepo *hlf.RepoUserBlockchain) ISvcUserTxs {
	return &svcUsersTxs{pRepo}
}

func (s *svcUsersTxs) CreateUser(user dto.UserBlockchainCreate) *dto.Problem {
	out, e := exec.Command("uuidgen").Output()
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}
	currentTime := time.Now()

	id := string(out)
	u := dto.UserBlockchain{
		ID:        id[:len(id)-1],
		Name:      user.Name,
		CreatedAt: currentTime.Format("2006.01.02 15:04:05"),
	}
	ccErr, e := (*s.repo).CreateUser(u)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	result := lib.DecodePayload(ccErr)

	m, ok := result.(error)
	if ok {
		return dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, m.Error())
	}

	return nil
}

func (s *svcUsersTxs) GetUserById(id string) (interface{}, *dto.Problem) {
	item, err := (*s.repo).GetUserById(id)
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}

	result := lib.DecodePayload(item)

	m, ok := result.(interface{})
	if !ok {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())
	}

	return m, nil
}

func (s *svcUsersTxs) DeleteUser(id string) *dto.Problem {

	ccErr, e := (*s.repo).DeleteUser(id)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	result := lib.DecodePayload(ccErr)

	m, ok := result.(error)
	if ok {
		return dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, m.Error())
	}

	return nil
}

func (s *svcUsersTxs) GetAllUsers() (interface{}, *dto.Problem) {
	item, err := (*s.repo).GetAllUsers()
	if err != nil {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrBuntdb, err.Error())
	}

	result := lib.DecodePayload(item)

	m, ok := result.(interface{})
	if !ok {
		return nil, dto.NewProblem(iris.StatusExpectationFailed, schema.ErrDecodePayloadTx, err.Error())
	}

	return m, nil
}
