package request

type RequestPayTransferToPocket struct {
}

type TransferSceneReportInfo struct {
	InfoType    string `json:"info_type,omitempty"`
	InfoContent string `json:"info_content,omitempty"`
}

type RequestTransferBills struct {
	Appid                    string                    `json:"appid,omitempty"`
	OutBillNo                string                    `json:"out_bill_no,omitempty"`
	TransferSceneId          string                    `json:"transfer_scene_id,omitempty"`
	Openid                   string                    `json:"openid,omitempty"`
	UserName                 string                    `json:"user_name,omitempty"`
	TransferAmount           int                       `json:"transfer_amount,omitempty"`
	TransferRemark           string                    `json:"transfer_remark,omitempty"`
	NotifyUrl                string                    `json:"notify_url,omitempty"`
	UserRecvPerception       string                    `json:"user_recv_perception,omitempty"`
	TransferSceneReportInfos []TransferSceneReportInfo `json:"transfer_scene_report_infos,omitempty"`
}
