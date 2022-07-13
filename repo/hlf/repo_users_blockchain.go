package hlf

import (
	ccfuncnames "github.com/ic-matcom/api.dapp/schema/ccFuncNames"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service/utils"
)

// region ======== SETUP =================================================================

type RepoUserBlockchain interface {
	CreateUser(user dto.UserBlockchain) ([]byte, error)
	GetUserById(id string) ([]byte, error)
	DeleteUser(id string) ([]byte, error)
	GetAllUsers() ([]byte, error)
}

// endregion =============================================================================

func NewRepoUserBlockchain(SvcConf *utils.SvcConfig) RepoUserBlockchain {
	return newRepoBlockchain(SvcConf)
}

func (r *repoBlockchain) CreateUser(user dto.UserBlockchain) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactUserCreate,
		string(user.ID), string(user.Name), string(user.CreatedAt))
	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) GetUserById(id string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	user, e := contract.SubmitTransaction(ccfuncnames.ContactUserRead, id)
	if e != nil {
		return nil, e
	}

	return user, nil
}

func (r *repoBlockchain) DeleteUser(id string) ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	ccErr, e := contract.SubmitTransaction(ccfuncnames.ContactUserDelete, string(id))
	if e != nil {
		return nil, e
	}

	return ccErr, nil
}

func (r *repoBlockchain) GetAllUsers() ([]byte, error) {
	gw, _, contract, e := r.getSDKComponents(r.ChannelName, ccfuncnames.ContractNameCC1, false)
	if e != nil {
		return nil, e
	}
	defer gw.Close()

	//strArgs, _ := jsoniter.Marshal(ID)

	res, e := contract.SubmitTransaction(ccfuncnames.ContactUserGetAll)
	if e != nil {
		return nil, e
	}

	return res, nil
}
