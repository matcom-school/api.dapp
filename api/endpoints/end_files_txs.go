package endpoints

import (
	"fmt"

	"github.com/ic-matcom/api.dapp/repo/hlf"
	"github.com/ic-matcom/api.dapp/schema"
	"github.com/ic-matcom/api.dapp/schema/dto"
	"github.com/ic-matcom/api.dapp/service"
	"github.com/ic-matcom/api.dapp/service/utils"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	"github.com/kataras/iris/v12/hero"
)

// HBlockchainTxs  endpoint handler struct for HLF blockchain transactions
type HFilesTxs struct {
	response *utils.SvcResponse
	service  *service.ISvcFilesTxs
}

// NewBlockchainTxsHandler create and register the handler for HLF blockchain transactions (txs)
//
// - app [*iris.Application] ~ Iris App instance
//
// - MdwAuthChecker [*context.Handler] ~ Authentication checker middleware
//
// - svcR [*utils.SvcResponse] ~ GrantIntentResponse service instance
//
// - svcC [utils.SvcConfig] ~ Configuration service instance
func NewFilesTxsHandler(app *iris.Application, mdwAuthChecker *context.Handler, svcR *utils.SvcResponse, svcC *utils.SvcConfig) HFilesTxs {

	// --- VARS SETUP ---
	repo := hlf.NewRepoFileBlockchain(svcC)
	svc := service.NewSvcFilesTxs(&repo)
	// registering protected / guarded router
	h := HFilesTxs{svcR, &svc}
	//repoUsers := db.NewRepoUsers(svcC)

	// registering unprotected router
	//authRouter := app.Party("/txs") // unauthorized
	//{
	//
	//}

	// registering protected / guarded router
	guardTxsRouter := app.Party("/txs")
	{
		// --- GROUP / PARTY MIDDLEWARES ---
		guardTxsRouter.Use(*mdwAuthChecker)

		// --- DEPENDENCIES ---
		hero.Register(DepObtainUserDid)
		//hero.Register(repoUsers)

		// --- REGISTERING ENDPOINTS ---

		// identity contract
		// we use the hero handler to inject the depObtainUserDid dependency. If we don't need to inject any dependencies we jus call guardTxsRouter.Get("/identity/identity/{id:string}", h.Identity_DevPopulate)

		guardTxsRouter.Post("/file", h.CreateFile)
		guardTxsRouter.Delete("/file/{id:string}", h.DeleteFile)
		guardTxsRouter.Patch("/file/{id:string}", h.UpdateFile)
		guardTxsRouter.Patch("/transfer/file/{id:string}", h.TransferFile)
		guardTxsRouter.Get("/file", h.GetAllFiles)
		guardTxsRouter.Get("/file/{id:string}", h.GetFileById)
		guardTxsRouter.Get("/history/file/{id:string}", h.FilesHistory)
	}

	return h
}

// region ======== ENDPOINT HANDLERS DEV =================================================

// CreateFile
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	tx				body	dto.FilesCreateDto		true	"Test data"
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/file [post]
func (h HFilesTxs) CreateFile(ctx iris.Context) {
	var request dto.FilesCreateDto

	// unmarshalling the json and check
	if err := ctx.ReadJSON(&request); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	problem := (*h.service).CreateFile(request)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// UpdateFile
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(mockAssect1)
// @Param	tx				body	dto.FilesUpdateDto		true	"Test data"
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/file/{id} [patch]
func (h HFilesTxs) UpdateFile(ctx iris.Context) {
	fmt.Printf("Aqui estoy")

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	var request dto.FilesUpdateDto
	// unmarshalling the json and check
	if err := ctx.ReadJSON(&request); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	problem := (*h.service).UpdateFile(id, request)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// DeleteFile
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(mockAssect1)
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/file/{id} [delete]
func (h HFilesTxs) DeleteFile(ctx iris.Context) {

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	problem := (*h.service).DeleteFile(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// TransferFile
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(mockAssect1)
// @Param	tx				body	dto.FileTransferDto		true	"Test data"
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/transfer/file/{id} [patch]
func (h HFilesTxs) TransferFile(ctx iris.Context) {

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	var request dto.FileTransferDto
	// unmarshalling the json and check
	if err := ctx.ReadJSON(&request); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	problem := (*h.service).TransferFile(id, request.UserId)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// GetFileById
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(mockAssect1)
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/file/{id} [get]
func (h HFilesTxs) GetFileById(ctx iris.Context) {

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	res, problem := (*h.service).GetFileById(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(res, &ctx)
}

// GetAllFiles
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	owner			query	string					false	"ownerID"	Format(string) default(tomoko@gmail.com)
// @Success 200 {object} []dto.Files "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/file [get]
func (h HFilesTxs) GetAllFiles(ctx iris.Context) {
	owner := ctx.URLParam("owner")

	if owner != "" {
		res, problem := (*h.service).GetAllFilesByOwner(owner)
		if problem != nil {
			(*h.response).ResErr(problem, &ctx)
			return
		}

		(*h.response).ResOKWithData(res, &ctx)
	} else {
		res, problem := (*h.service).GetAllFiles()
		if problem != nil {
			(*h.response).ResErr(problem, &ctx)
			return
		}

		(*h.response).ResOKWithData(res, &ctx)
	}
}

// FilesHistory
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(mockAssect1)
// @Success 200 {object} []dto.FileHistoryDto "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/history/file/{id} [get]
func (h HFilesTxs) FilesHistory(ctx iris.Context) {
	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	res, problem := (*h.service).FilesHistory(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(res, &ctx)
}
