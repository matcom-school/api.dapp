package hlf

import (
	"strconv"

	ccfuncnames "github.com/ic-matcom/api.dapp/schema/ccFuncNames"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service/utils"
)

// region ======== SETUP =================================================================

type RepoFilesBlockchain interface {
	CreateFile(file dto.Files) ([]byte, error)
	GetFileById(id string) ([]byte, error)
	UpdateFile(id string, file dto.FilesUpdateDto) ([]byte, error)
	DeleteFile(id string) ([]byte, error)
	TransferFile(id string, userId string) ([]byte, error)
	GetAllFiles() ([]byte, error)
	GetAllFilesByOwner(userId string) ([]byte, error)
	FilesHistory(id string) ([]byte, error)
}

// endregion =============================================================================

func NewRepoFileBlockchain(SvcConf *utils.SvcConfig) RepoFilesBlockchain {
	return newRepoBlockchain(SvcConf)
}

func (r *repoBlockchain) CreateFile(file dto.Files) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)
	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactFileCreate,
		string(file.ID),
		string(file.Url),
		string(file.Name),
		string(file.CreatedAt),
		string(strconv.Itoa(file.Size)),
		string(file.Owner),
		string(file.Type))
	//res, e := contract.SubmitTransaction(ccfuncnames.CC1ReadAsset, string(strArgs))
	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) GetFileById(id string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	file, e := contract.SubmitTransaction(ccfuncnames.ContactFileGetById, string(id))
	if e != nil {
		return nil, e
	}

	return file, nil
}

func (r *repoBlockchain) UpdateFile(id string, file dto.FilesUpdateDto) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactFileUpdate,
		string(id), string(file.Name), string(file.Url),
		string(strconv.Itoa(file.Size)),
		string(file.Type))

	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) DeleteFile(id string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactFileDelete, string(id))
	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) TransferFile(id string, userId string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactFileTransfer, string(id), string(userId))
	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) GetAllFiles() ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	res, e := contract.SubmitTransaction(ccfuncnames.ContactFileGetAll)
	if e != nil {
		return nil, e
	}

	return res, nil
}

func (r *repoBlockchain) GetAllFilesByOwner(userId string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	res, e := contract.SubmitTransaction(ccfuncnames.ContactFileGetAllByOwner, string(userId))
	if e != nil {
		return nil, e
	}

	return res, nil
}

func (r *repoBlockchain) FilesHistory(id string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	res, e := contract.SubmitTransaction(ccfuncnames.ContactFileHistory, string(id))
	if e != nil {
		return nil, e
	}

	return res, nil
}
