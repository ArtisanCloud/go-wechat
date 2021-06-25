package featureUnit

import (
	"github.com/ArtisanCloud/go-libs/fmt"
	"testing"
)

func Test_ExternalContact_Get_Contact_Way(t *testing.T) {


	response := Work.ContactWay.Get("008dc067bf677e5f03df89ce49bea25a")

	if response == nil || response.ResponseWX == nil {
		t.Error("response nil")
	} else if response.ErrCode != 0 {
		t.Error("response error message as :", response.ErrMSG)
	}

	fmt.Dump(response)

}