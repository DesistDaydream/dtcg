package wishes

import (
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/core"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/wishes/models"
)

type WishesClient struct {
	client *core.Client
}

func NewWishesClient(client *core.Client) *WishesClient {
	return &WishesClient{
		client: client,
	}
}

// 创建清单
func (w *WishesClient) CreateWashList(name string) (*models.WishListCreateResp, error) {
	var wishListCreateResp models.WishListCreateResp
	uri := "/api/market/wish-list"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.WishListCreateReqBody{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Name:       name,
		},
	}

	err := w.client.Request(uri, &wishListCreateResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &wishListCreateResp, nil
}

// 列出清单
func (w *WishesClient) List() {}

// 更新清单
func (w *WishesClient) Update() {
	// TODO: 好像没有更新清单的逻辑，只有向清单中添加卡牌和删除清单中的卡牌这两种逻辑。
}

// 删除清单
func (w *WishesClient) Del() {}

// 获取清单详情
func (w *WishesClient) Get(wishListID string) (*models.WishListGetResp, error) {
	var wishListGetResp models.WishListGetResp
	uri := "/api/market/wishes"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.WishListGetReqQuery{
			GameKey:    "dgm",
			GameSubKey: "sc",
			Page:       "1",
			WishListID: wishListID,
		}),
	}

	err := w.client.Request(uri, &wishListGetResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &wishListGetResp, nil
}

// 向清单中添加卡牌
func (w *WishesClient) Add(cardVersionID, ignoreCardVersion, quantity, remark, wishListID string) (*models.WishesAddResp, error) {
	var wishesAddResp models.WishesAddResp
	uri := "/api/market/wishes/add"

	reqOpts := &core.RequestOption{
		Method: "POST",
		ReqBody: &models.WishesAddReqBody{
			CardVersionID:     cardVersionID,
			GameKey:           "dgm",
			IgnoreCardVersion: ignoreCardVersion,
			Quantity:          quantity,
			Remark:            remark,
			WishListID:        wishListID,
		},
	}

	err := w.client.Request(uri, &wishesAddResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &wishesAddResp, nil
}

// 列出官方推荐的清单
func (w *WishesClient) ListWishListRecommend() (*models.WishListRecommendResp, error) {
	var wishListRecommendResp models.WishListRecommendResp
	uri := "/api/market/wish-list/recommend"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.WishListRecommendReqQuery{
			GameKey:     "dgm",
			GameSubKey:  "sc",
			IsRecommend: "0",
			Page:        "1",
		}),
	}

	err := w.client.Request(uri, &wishListRecommendResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return &wishListRecommendResp, nil
}

// 一键匹配清单
func (w *WishesClient) WishListMatch(wishListID string) (models.WishListMatchResultsResp, error) {
	var wishListMatchResultsResp models.WishListMatchResultsResp
	uri := "/api/market/wishes/match"

	reqOpts := &core.RequestOption{
		Method: "GET",
		ReqQuery: core.StructToMapStr(&models.WishListMatchResultsReqQuery{
			GameKey:           "dgm",
			GameSubKey:        "sc",
			IgnoreCardVersion: "0",
			ShowMatchDetails:  "1",
			WishListID:        wishListID,
		}),
	}

	err := w.client.Request(uri, &wishListMatchResultsResp, reqOpts)
	if err != nil {
		return nil, err
	}

	return wishListMatchResultsResp, nil
}
