package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	//sdk "github.com/cosmos/cosmos-sdk/types"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	//"github.com/Wanxiang-Blockchain-Hackathon-2020/N06-admin/x/admin/internal/types"
	"net/http"
)

// PostProposalReq defines the properties of a proposal request's body.
type PostMintReq struct {
	BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`
	Minter  string       `json:"minter" yaml:"minter"` // Address of the minter
}

func mintHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req PostMintReq
		//if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
		//	return
		//}
		//
		//baseReq := req.BaseReq.Sanitize()
		//if !baseReq.ValidateBasic(w) {
		//	return
		//}
		//
		//sender, err := sdk.AccAddressFromBech32(baseReq.From)
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}
		//
		//minter, err := sdk.AccAddressFromBech32(req.Minter)
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}
		//
		//// create the message
		//msg := types.NewMsgRegisterDocter(sender, minter)
		//err = msg.ValidateBasic()
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}

		//utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
