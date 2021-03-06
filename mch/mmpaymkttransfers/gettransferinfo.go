package mmpaymkttransfers

import (
	"github.com/chanxuehong/wechat.v2/mch/core"
)

// 查询企业付款.
//  NOTE: 请求需要双向证书
func GetTransferInfo(clt *core.Client, req map[string]string) (resp map[string]string, err error) {
	return clt.PostXML("https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo", req)
}
