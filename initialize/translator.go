package initialize

import (
	"github.com/childelins/go-gin-api/global"
	"github.com/childelins/go-gin-api/pkg/translator"
)

func InitTrans(locale string) error {
	trans, err := translator.NewTrans(locale)
	if err != nil {
		return err
	}

	global.Trans = trans
	return nil
}
