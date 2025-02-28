package fundApp

import (
	"context"
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/fundApp/request"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/payment/fundApp/response"
	payment "github.com/ArtisanCloud/PowerWeChat/v3/src/payment/kernel"

	"net/http"
)

type Client struct {
	*payment.BaseClient
}

func NewClient(app *payment.ApplicationPaymentInterface) (*Client, error) {
	baseClient, err := payment.NewBaseClient(app)
	if err != nil {
		return nil, err
	}
	return &Client{
		baseClient,
	}, nil
}

// 发起转账
// https://pay.weixin.qq.com/doc/v3/merchant/4012716434
func (comp *Client) TransferBills(ctx context.Context, data *request.RequestTransferBills) (*response.ResponseTransferBills, error) {

	result := &response.ResponseTransferBills{}

	params, err := object.StructToHashMap(data)
	if err != nil {
		return nil, err
	}

	endpoint := comp.Wrap("/v3/fund-app/mch-transfer/transfer-bills")
	_, err = comp.SafeRequestV3(ctx, endpoint, nil, http.MethodPost, params, nil, result)

	return result, err
}
