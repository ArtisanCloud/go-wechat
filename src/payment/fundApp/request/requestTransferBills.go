package request

type RequestPayTransferToPocket struct {
}

type TransferSceneReportInfo struct {
	InfoType    string `json:"info_type"`
	InfoContent string `json:"info_content"`
}

type RequestTransferBills struct {
	Appid                    string                    `json:"appid"`
	OutBillNo                string                    `json:"out_bill_no"`
	TransferSceneId          string                    `json:"transfer_scene_id"`
	Openid                   string                    `json:"openid"`
	UserName                 string                    `json:"user_name"`
	TransferAmount           int                       `json:"transfer_amount"`
	TransferRemark           string                    `json:"transfer_remark"`
	NotifyUrl                string                    `json:"notify_url"`
	UserRecvPerception       string                    `json:"user_recv_perception"`
	TransferSceneReportInfos []TransferSceneReportInfo `json:"transfer_scene_report_infos"`
}
