package dada

import "ixingkong.com/qiyi-tech/dada/schema"

// GetCityCodeList 获取达达配送开通城市列表信息
func (c *Client) GetCityCodeList() (result []*schema.CityCode, err error) {
	err = c.Request("/api/cityCode/list", nil, &result)
	return result, err
}

// BusinessMap 达达快送品类
var BusinessMap = map[int64]string{
	1:  "食品小吃",
	2:  "饮料",
	3:  "鲜花绿植",
	5:  "其他",
	8:  "文印票务",
	9:  "便利店",
	13: "水果生鲜",
	19: "同城电商",
	20: "医药",
	21: "蛋糕",
	24: "酒品",
	25: "小商品市场",
	26: "服装",
	27: "汽修零配",
	28: "数码家电",
	29: "小龙虾/烧烤",
	31: "超市",
	51: "火锅",
	53: "个护美妆",
	55: "母婴",
	57: "家居家纺",
	59: "手机",
	61: "家装",
	63: "成人用品",
	65: "校园",
	66: "高端市场",
}

// GetBusinessMap 获取达达快送品类映射表
func (c *Client) GetBusinessMap() map[int64]string {
	return BusinessMap
}

// AddMerchant 注册商户帐号
func (c *Client) AddMerchant(req *schema.AddMerchantRequest) (result int64, err error) {
	err = c.Request("/merchantApi/merchant/add", req, &result)
	return result, err
}

// AddShop 创建门店
func (c *Client) AddShop(req []*schema.AddShopItem) (result *schema.AddShopResp, err error) {
	err = c.Request("/api/shop/add", req, &result)
	return result, err
}

// UpdateShop 更新门店
func (c *Client) UpdateShop(req *schema.UpdateShopRequest) (result bool, err error) {
	err = c.Request("/api/shop/update", req, &result)
	if err != nil {
		return false, err
	}
	return true, nil
}

// GetShop 查询门店详情
func (c *Client) GetShop(originShopID string) (result *schema.Shop, err error) {
	err = c.Request("/api/shop/detail", map[string]string{"origin_shop_id": originShopID}, &result)
	return result, err
}
