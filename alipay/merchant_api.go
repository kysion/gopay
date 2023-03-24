package alipay

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/kysion/gopay"
)

// alipay.trade.royalty.relation.bind(分账关系绑定)
// 文档地址：https://opendocs.alipay.com/open/02c7hq
func (a *Client) TradeRelationBind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationBindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.bind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationBindResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.royalty.relation.unbind(分账关系解绑)
// 文档地址：https://opendocs.alipay.com/open/02c7hr
func (a *Client) TradeRelationUnbind(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRelationUnbindResponse, err error) {
	err = bm.CheckEmptyError("receiver_list", "out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.unbind"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeRelationUnbindResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.royalty.relation.batchquery(分账关系查询)
// 文档地址：https://opendocs.alipay.com/open/02c7hs
func (a *Client) TradeRelationBatchQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeRoyaltyRelationRes, err error) {
	err = bm.CheckEmptyError("out_request_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.royalty.relation.batchquery"); err != nil {
		return nil, err
	}
	fmt.Println("-----------")
	fmt.Println(string(bs))

	aliRsp = new(TradeRoyaltyRelationRes)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}

	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.order.settle(统一收单交易结算接口)
// 文档地址：https://opendocs.alipay.com/open/02j2bt
func (a *Client) TradeOrderSettle(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleResponse, err error) {
	err = bm.CheckEmptyError("out_request_no", "trade_no", "royalty_parameters")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}

// alipay.trade.order.settle.query(交易分账查询接口)
// 文档地址：https://opendocs.alipay.com/open/02pj6l
func (a *Client) TradeOrderSettleQuery(ctx context.Context, bm gopay.BodyMap) (aliRsp *TradeOrderSettleQueryResponse, err error) {
	err = bm.CheckEmptyError("settle_no")
	if err != nil {
		return nil, err
	}
	var bs []byte
	if bs, err = a.doAliPay(ctx, bm, "alipay.trade.order.settle.query"); err != nil {
		return nil, err
	}
	aliRsp = new(TradeOrderSettleQueryResponse)
	if err = json.Unmarshal(bs, aliRsp); err != nil || aliRsp.Response == nil {
		return nil, fmt.Errorf("[%w], bytes: %s", gopay.UnmarshalErr, string(bs))
	}
	if err = bizErrCheck(aliRsp.Response.ErrorResponse); err != nil {
		return aliRsp, err
	}
	signData, signDataErr := a.getSignData(bs, aliRsp.AlipayCertSn)
	aliRsp.SignData = signData
	return aliRsp, a.autoVerifySignByCert(aliRsp.Sign, signData, signDataErr)
}
