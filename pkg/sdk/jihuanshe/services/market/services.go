package market

import (
	"strconv"

	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	"github.com/DesistDaydream/dtcg/pkg/sdk/utils"
)

type MarketClient struct {
	client *core.Client
}

func NewMarketClient(client *core.Client) *MarketClient {
	return &MarketClient{
		client: client,
	}
}

// 更新 Token
func (m *MarketClient) AuthUpdateTokenPost() (*models.CommonSuccessResp, error) {
	var commonSuccessResp models.CommonSuccessResp
	uri := "/api/market/auth/update-push-token"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.UpdateTokenPostReqBody{
			PushDevice: "ios",
			Token:      m.client.Token,
		},
	}

	err := m.client.Request(uri, &commonSuccessResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &commonSuccessResp, nil
}

// 添加我在卖的商品
func (m *MarketClient) SellersProductsAdd(productsAddRequestBody *models.ProductsAddReqBody) (*models.CommonSuccessResp, error) {
	var productsAddResp models.CommonSuccessResp
	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method:  "POST",
		ReqBody: productsAddRequestBody,
	}

	err := m.client.Request(uri, &productsAddResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsAddResp, nil
}

// 列出我在卖的商品
func (m *MarketClient) SellersProductsList(page int, keyword, onSale, sorting string) (*models.ProductsListResp, error) {
	var productsResp models.ProductsListResp

	uri := "/api/market/sellers/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.ProductsListReqQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Keyword:    keyword,
			OnSale:     onSale,
			Page:       numToString(page),
			Sorting:    "published_at_desc",
			Rarity:     "",
		}),
	}

	err := m.client.Request(uri, &productsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsResp, nil
}

// 更新我在卖的商品
func (m *MarketClient) SellersProductsUpdate(productsUpdateRequestBody *models.ProductsUpdateReqBody, productID string) (*models.ProductsUpdateResp, error) {
	var productsUpdateResp models.ProductsUpdateResp
	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "PUT",
		ReqBody: productsUpdateRequestBody,
	}

	err := m.client.Request(uri, &productsUpdateResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsUpdateResp, nil
}

// 删除我在卖的商品
func (m *MarketClient) SellersProductsDel(productID string) (*models.CommonSuccessResp, error) {
	var productsDelResp models.CommonSuccessResp

	uri := "/api/market/sellers/products/" + productID

	reqOpts := &core.RequestOption{
		Method:  "DELETE",
		ReqBody: &models.ProductsDelReqBody{},
	}

	err := m.client.Request(uri, &productsDelResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productsDelResp, nil
}

// 获取提现日志
func (m *MarketClient) SellersWithdrawLogGet(page int) (*models.WithdrawResp, error) {
	var withdrawResp models.WithdrawResp
	uri := "/api/market/sellers/withdraw/withdrawLogs"

	reqOpts := &core.RequestOption{
		Method:   "GET",
		ReqQuery: map[string]string{"page": strconv.Itoa(page)},
	}

	err := m.client.Request(uri, &withdrawResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &withdrawResp, nil
}

// 获取用户订单列表(买入)
func (m *MarketClient) OrderList(page int) (*models.BuyerOrdersListResp, error) {
	var buyerOrders models.BuyerOrdersListResp

	uri := "/api/market/orders"
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderListReqQuery{
			Page:   strconv.Itoa(page),
			Status: "complete",
			Token:  m.client.Token,
		}),
	}

	err := m.client.Request(uri, &buyerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &buyerOrders, nil
}

// 获取用户订单详情(买入)
func (m *MarketClient) OrderGet(orderID int) (*models.OrderByBuyerGetResp, error) {
	var orderProducts models.OrderByBuyerGetResp

	uri := "/api/market/orders/" + strconv.Itoa(orderID)
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderGetReqQuery{
			Token: m.client.Token,
		}),
	}

	err := m.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取用户订单列表（卖出）
func (m *MarketClient) SellerOrderList(page int) (*models.SellerOrderListResp, error) {
	var sellerOrders models.SellerOrderListResp

	uri := "/api/market/sellers/orders"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderListReqQuery{
			Page:   strconv.Itoa(page),
			Status: "complete",
			Token:  m.client.Token,
		}),
	}

	err := m.client.Request(uri, &sellerOrders, reqOpts)
	if err != nil {
		return nil, err
	}

	return &sellerOrders, nil
}

// 获取用户订单详情（卖出）
func (m *MarketClient) SellerOrderGet(orderID int) (*models.OrderBySellerGetResp, error) {
	var orderProducts models.OrderBySellerGetResp

	uri := "/api/market/sellers/orders/" + strconv.Itoa(orderID)
	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.OrderGetReqQuery{
			Token: m.client.Token,
		}),
	}

	err := m.client.Request(uri, &orderProducts, reqOpts)
	if err != nil {
		return nil, err
	}

	return &orderProducts, nil
}

// 获取商品的“在售”列表
func (m *MarketClient) CardVersionsProductsGet(cardVersionID string, page int) (*models.ProductSellersGetResp, error) {
	var productSellers models.ProductSellersGetResp
	uri := "/api/market/card-versions/products"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: utils.StructToMapStr(&models.ProductSellersGetReqQuery{
			CardVersionID: cardVersionID,
			Condition:     "1",
			GameKey:       "dgm",
			Page:          strconv.Itoa(page),
		}),
	}

	err := m.client.Request(uri, &productSellers, reqOpts)
	if err != nil {
		return nil, err
	}

	return &productSellers, nil
}

// 获取卡包信息
func (m *MarketClient) GetPacks(packID int, page int) (*models.PacksGetResp, error) {
	var packsGetResp models.PacksGetResp
	uri := "/api/market/packs/" + strconv.Itoa(packID)

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqBody: &models.PacksGetReq{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Page:       numToString(page),
		},
	}

	err := m.client.RequestWithEncrypt(uri, &packsGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &packsGetResp, nil
}

// 列出卡牌
func (m *MarketClient) ListCardVersions(packID, categoryID int, page int) (*models.CardVersionsListResp, error) {
	var cardVersionListResp models.CardVersionsListResp
	uri := "/api/market/card-versions"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqBody: &models.CardVersionsListReqBody{
			PackID:           numToString(packID),
			CategoryID:       numToString(categoryID),
			GameKey:          "dgm",
			GameSubKey:       "sc",
			Page:             numToString(page),
			Rarity:           "",
			Sorting:          "number",
			SortingPriceType: "product",
		},
	}

	err := m.client.RequestWithEncrypt(uri, &cardVersionListResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardVersionListResp, nil
}

// 获取卡牌信息
func (m *MarketClient) GetCardVersions(cardVersion int) (*models.CardVersionGetResp, error) {
	var cardVersionGetResp models.CardVersionGetResp
	uri := "/api/market/card-versions/" + strconv.Itoa(cardVersion)

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqBody: &models.CardVersionGetReqBody{
			GameKey:    "dgm",
			GameSubKey: "sc",
		},
	}

	err := m.client.RequestWithEncrypt(uri, &cardVersionGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardVersionGetResp, nil
}

// 获取卡牌基本信息
func (m *MarketClient) GetCardVersionsBaseInfo(cardVersionID int) (*models.CardVersionsBaseInfoResp, error) {
	var cardVersionsBaseInfoResp models.CardVersionsBaseInfoResp
	uri := "/api/market/card-versions/get-base-info"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqBody: &models.CardVersionsBaseInfoGetReqBody{
			CardVersionID: numToString(cardVersionID),
			GameKey:       "dgm",
			GameSubKey:    "sc",
		},
	}

	err := m.client.RequestWithEncrypt(uri, &cardVersionsBaseInfoResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardVersionsBaseInfoResp, nil
}

// 获取卡牌价格历史
func (m *MarketClient) GetCardVersionsPriceHistory(cardVersionID int) (*models.CardVersionsPriceHistoryResp, error) {
	var cardVersionsPriceHistoryResp models.CardVersionsPriceHistoryResp
	uri := "/api/market/card-versions/price-history"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqBody: &models.CardVersionsBaseInfoGetReqBody{
			CardVersionID: numToString(cardVersionID),
			GameKey:       "dgm",
			GameSubKey:    "sc",
		},
	}

	err := m.client.RequestWithEncrypt(uri, &cardVersionsPriceHistoryResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &cardVersionsPriceHistoryResp, nil
}

func numToString(num int) string {
	var str string

	if num == 0 {
		str = ""
	} else {
		str = strconv.Itoa(num)
	}

	return str
}
