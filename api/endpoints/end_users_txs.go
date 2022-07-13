package endpoints

import (
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
type HUserTxs struct {
	response *utils.SvcResponse
	service  *service.ISvcUserTxs
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
func NewUserTxsHandler(app *iris.Application, mdwAuthChecker *context.Handler, svcR *utils.SvcResponse, svcC *utils.SvcConfig) HUserTxs {

	// --- VARS SETUP ---
	repo := hlf.NewRepoUserBlockchain(svcC)
	svc := service.NewSvcUserTxs(&repo)
	// registering protected / guarded router
	h := HUserTxs{svcR, &svc}
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

		guardTxsRouter.Post("/user", h.CreateUser)
		guardTxsRouter.Delete("/user/{id:string}", h.DeleteUser)
		guardTxsRouter.Get("/user/{id:string}", h.GetUserById)
		guardTxsRouter.Get("/user", h.GetAllUsers)
	}

	return h
}

// region ======== ENDPOINT HANDLERS DEV =================================================

// CreateUser
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	tx				body	dto.UserBlockchainCreate		true	"Test data"
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/user [post]
func (h HUserTxs) CreateUser(ctx iris.Context) {
	var request dto.UserBlockchainCreate

	// unmarshalling the json and check
	if err := ctx.ReadJSON(&request); err != nil {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: err.Error()}, &ctx)
		return
	}

	problem := (*h.service).CreateUser(request)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// DeleteUser
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(tomoko@gmail.com)
// @Success 200 {object} interface{} "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/user/{id} [delete]
func (h HUserTxs) DeleteUser(ctx iris.Context) {

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	problem := (*h.service).DeleteUser(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData("res", &ctx)
}

// GetUserById
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Param	id				path	string					true	"ID"	Format(string) default(tomoko@gmail.com)
// @Success 200 {object} dto.UserBlockchain 	 "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/user/{id} [get]
func (h HUserTxs) GetUserById(ctx iris.Context) {

	id := ctx.Params().GetString("id")
	if id == "" {
		(*h.response).ResErr(&dto.Problem{Status: iris.StatusBadRequest, Title: schema.ErrProcParam, Detail: schema.ErrDetInvalidField}, &ctx)
		return
	}

	res, problem := (*h.service).GetUserById(id)
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(res, &ctx)
}

// GetAllUsers
// @Tags Txs.eVote
// @Security ApiKeyAuth
// @Accept  json
// @Produce json
// @Param	Authorization	header	string 			        true 	"Insert access token" default(Bearer <Add access token here>)
// @Success 200 {object} []dto.UserBlockchain "OK"
// @Failure 401 {object} dto.Problem "err.unauthorized"
// @Failure 400 {object} dto.Problem "err.processing_param"
// @Failure 502 {object} dto.Problem "err.bad_gateway"
// @Failure 504 {object} dto.Problem "err.network"
// @Router /txs/user [get]
func (h HUserTxs) GetAllUsers(ctx iris.Context) {
	res, problem := (*h.service).GetAllUsers()
	if problem != nil {
		(*h.response).ResErr(problem, &ctx)
		return
	}

	(*h.response).ResOKWithData(res, &ctx)
}
