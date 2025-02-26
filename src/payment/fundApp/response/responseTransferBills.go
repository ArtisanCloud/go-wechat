package response

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel/response"
	"time"
)

// https://work.weixin.qq.com/api/doc/90000/90135/90278

type ResponseTransferBills struct {
	response.ResponsePayment
	OutBillNo      string    `json:"out_bill_no"`
	TransferBillNo string    `json:"transfer_bill_no"`
	CreateTime     time.Time `json:"create_time"`
	State          string    `json:"state"`
	FailReason     string    `json:"fail_reason"`
	PackageInfo    string    `json:"package_info"`
}
