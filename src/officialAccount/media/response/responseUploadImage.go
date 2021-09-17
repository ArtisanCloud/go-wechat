package response

import (
	"github.com/ArtisanCloud/power-wechat/src/kernel/response"
)

type ResponseUploadImage struct {
	*response.ResponseOfficialAccount

	URL string `json:"url"`
}
