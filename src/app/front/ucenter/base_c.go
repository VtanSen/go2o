/**
 * Copyright 2015 @ S1N1 Team.
 * name : default_c.go
 * author : jarryliu
 * date : -- :
 * description :
 * history :
 */
package ucenter

import (
	"github.com/atnet/gof/web"
	"github.com/atnet/gof/web/mvc"
	"go2o/src/core/domain/interface/member"
	"go2o/src/core/domain/interface/partner"
	"go2o/src/core/service/goclient"
	"net/url"
	"go2o/src/core/service/dps"
	"go2o/src/cache"
)

var _ mvc.Filter = new(baseC)

type baseC struct {
}

func (this *baseC) Requesting(ctx *web.Context) bool {
	//验证是否登陆
	s := ctx.Session().Get("member")
	if s != nil {
		if m := s.(*member.ValueMember); m != nil {
			ctx.Items["member"] = m
			ctx.Items["member_ptId"] = dps.MemberService.GetRelation(m.Id).Reg_PtId
			return true
		}
	}
	ctx.ResponseWriter.Write([]byte("<script>window.parent.location.href='/login?return_url=" +
	url.QueryEscape(ctx.Request.URL.String()) + "'</script>"))
	return false
}

func (this *baseC) RequestEnd(ctx *web.Context) {
}

// 获取商户
func (this *baseC) GetPartner(ctx *web.Context) *partner.ValuePartner {
	var partnerId int = ctx.Items["member_ptId"].(int)
	pt := cache.GetValuePartnerCache(partnerId)
	if pt == nil{
		var err error
		if pt,err = dps.PartnerService.GetPartner(partnerId);err == nil{
			cache.SetValuePartnerCache(partnerId,pt)
		}
	}
	return pt
}

// 获取会员
func (this *baseC) GetMember(ctx *web.Context) *member.ValueMember {
	return ctx.Items["member"].(*member.ValueMember)
}

// 获取商户的站点设置
func (this *baseC) GetSiteConf(partnerId int, secret string) (*partner.SiteConf, error) {
	return goclient.Partner.GetSiteConf(partnerId, secret)
}
