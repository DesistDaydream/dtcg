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
func (w *WishesClient) CreateList(name string) (*models.WishListCreateResp, error) {
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
