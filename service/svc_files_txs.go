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
type ISvcFilesTxs interface {
	CreateFile(file dto.FilesCreateDto) *dto.Problem
	GetFileById(id string) (interface{}, *dto.Problem)
	UpdateFile(id string, file dto.FilesUpdateDto) *dto.Problem
	DeleteFile(id string) *dto.Problem
	TransferFile(id string, userId string) *dto.Problem
	GetAllFiles() (interface{}, *dto.Problem)
	GetAllFilesByOwner(userId string) (interface{}, *dto.Problem)
	FilesHistory(id string) (interface{}, *dto.Problem)
}

type svcFilesTxs struct {
	repo *hlf.RepoFilesBlockchain
}

// endregion =============================================================================

// NewSvcBlockchainTxs instantiate the HLF blockchains transactions services
func NewSvcFilesTxs(pRepo *hlf.RepoFilesBlockchain) ISvcFilesTxs {
	return &svcFilesTxs{pRepo}
}

func (s *svcFilesTxs) CreateFile(file dto.FilesCreateDto) *dto.Problem {
	out, e := exec.Command("uuidgen").Output()
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}
	currentTime := time.Now()

	f := dto.Files{
		ID:        string(out),
		Name:      file.Name,
		Url:       file.Url,
		CreatedAt: currentTime.Format("2006.01.02 15:04:05"),
		Owner:     file.Owner,
		Size:      file.Size,
		Type:      file.Type,
	}
	e = (*s.repo).CreateFile(f)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	return nil
}

func (s *svcFilesTxs) GetFileById(id string) (interface{}, *dto.Problem) {
	item, err := (*s.repo).GetFileById(id)
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

func (s *svcFilesTxs) UpdateFile(id string, file dto.FilesUpdateDto) *dto.Problem {

	e := (*s.repo).UpdateFile(id, file)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	return nil
}

func (s *svcFilesTxs) DeleteFile(id string) *dto.Problem {

	e := (*s.repo).DeleteFile(id)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	return nil
}

func (s *svcFilesTxs) TransferFile(id string, userId string) *dto.Problem {

	e := (*s.repo).TransferFile(id, userId)
	if e != nil {
		return dto.NewProblem(iris.StatusBadGateway, schema.ErrBlockchainTxs, e.Error())
	}

	return nil
}

func (s *svcFilesTxs) GetAllFiles() (interface{}, *dto.Problem) {
	item, err := (*s.repo).GetAllFiles()
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

func (s *svcFilesTxs) GetAllFilesByOwner(userId string) (interface{}, *dto.Problem) {
	item, err := (*s.repo).GetAllFilesByOwner(userId)
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

func (s *svcFilesTxs) FilesHistory(id string) (interface{}, *dto.Problem) {
	item, err := (*s.repo).FilesHistory(id)
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
