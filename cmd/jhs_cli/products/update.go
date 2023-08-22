package products

import (
	"fmt"
	"sync"

	dbmodels "github.com/DesistDaydream/dtcg/internal/database/models"
	"github.com/DesistDaydream/dtcg/pkg/handler"
	"github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/market/models"
	pmodels "github.com/DesistDaydream/dtcg/pkg/sdk/jihuanshe/services/products/models"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UpdateFlags struct {
	SellerUserID string // 集换社卖家 ID
	CurSaleState string // 当前商品的售卖状态
	ExpSaleState int    // 期望商品变成哪种售卖状态
	Remark       string // 商品备注
}

var updateFlags UpdateFlags

func UpdateCommand() *cobra.Command {

	updateProductsCmd := &cobra.Command{
		Use:   "update",
		Short: "更新商品",
		// Run:   updatePrice,
		PersistentPreRun: updatePersistentPreRun,
	}

	updateProductsCmd.PersistentFlags().StringVarP(&updateFlags.SellerUserID, "seller-user-id", "u", "934972", "卖家用户ID。")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.CurSaleState, "cur-sale-state", "1", "当前售卖状态。即获取什么状态的商品。1: 售卖。0: 下架")
	updateProductsCmd.PersistentFlags().IntVar(&updateFlags.ExpSaleState, "exp-sale-state", 1, "期望的售卖状态。")
	updateProductsCmd.PersistentFlags().StringVar(&updateFlags.Remark, "remark", "", "商品备注信息")

	updateProductsCmd.AddCommand(
		UpdatePriceCommand(),
		UpdateImageCommand(),
		UpdateSaleStateCommand(),
		UpdateQuantityCommand(),
	)

	return updateProductsCmd
}

func updatePersistentPreRun(cmd *cobra.Command, args []string) {
	if productsFlags.SetPrefix == nil && productsFlags.CardVersionID == 0 {
		logrus.Fatalln("请指定要更新的卡牌，可以使用 dtcg_cli card-set list 子命令获取卡包名称；或者直接指定卡牌的 card_version_id。")
	}
}

type NeedHandleProducts struct {
	count    int
	products []Product
}

type Product struct {
	card           *dbmodels.CardPrice
	product        *models.ProductData
	defaultProduct *pmodels.DefaultProduct
	productID      int
	condition      int
	// 该商品所属卡牌的部分信息
	cardVersionID int
	cardNameCN    string
	// 更新商品用的新数据，并不是一定会用上。主要用于不同更新场景时使用
	img       string
	onSale    int
	price     string
	quantity  int
	isDefautl int // TODO: 添加使用是否是默认商品的逻辑
}

// 生成需要更新的商品信息
func genNeedUpdateProducts(cards *dbmodels.CardsPrice) *NeedHandleProducts {
	var wg sync.WaitGroup
	defer wg.Wait()
	var lock sync.Mutex // 并发中有对数组的 append 操作，加锁保证并发安全

	var needHandleProducts NeedHandleProducts

	// 逐一生成待处理卡牌的商品信息
	for _, card := range cards.Data {
		wg.Add(1)

		go func(card dbmodels.CardPrice) {
			defer wg.Done()

			// 用于记录待处理的卡牌的商品是否已更新
			isUpdate := 0

			// 使用 /api/market/sellers/products 接口通过卡牌关键字(即卡牌编号)获取到该卡牌的商品列表
			products, err := handler.H.JhsServices.Market.SellersProductsList(1, card.Serial, updateFlags.CurSaleState, "published_at_desc")
			if err != nil {
				logrus.Errorf("获取 %v 卡牌的商品失败: %v", card.ScName, err)
				updateFailCount++
				return
			}
			// 判断一下这个这个卡牌有几个商品，若商品为0，则继续处理下一个
			if len(products.Data) == 0 {
				logrus.Errorf("%v %v 卡牌没有任何版本可供处理的商品，跳过", card.CardVersionID, card.ScName)
				updateSkip++
				return
			}

			// 通过卡牌编号获取到的商品信息不是唯一的，这个编号的卡有多个版本（包含异画），所以需要通过 card_version_id 对比以确定唯一的 product_id(商品ID)
			for _, p := range products.Data {
				// 只有当商品的 card_version_id 与当前想要处理的卡牌的 card_version_id 相同时，则可以确定这个商品就是想要更新的卡牌的版本的商品
				if p.CardVersionID == card.CardVersionID {
					// 由于 go 语言中，for 每次迭代的值都是存储到同一个内存空间中的，所以想要引用 for 中 value 的内存空间，需要再手动创建一个
					cardCopy := card

					lock.Lock() // append 切片在并发中不安全，加个锁
					needHandleProducts.products = append(needHandleProducts.products, Product{
						card:           &cardCopy,
						product:        &p,
						defaultProduct: nil,
						productID:      p.ProductID,
						condition:      p.Condition,
						cardVersionID:  p.CardVersionID,
						cardNameCN:     p.CardNameCN,
						img:            p.CardVersionImage,
						onSale:         p.OnSale,
						price:          p.Price,
						quantity:       p.Quantity,
					})
					lock.Unlock()

					logrus.WithFields(logrus.Fields{
						"原始价格": card.AvgPrice,
						"商品价格": p.Price,
						"商品数量": p.Quantity,
					}).Debugf("检查匹配到的商品: 【%v】【%v】【 %v %v 】【 %v 】", p.ProductID, p.CardVersionID, card.Serial, p.CardNameCN, p.CardVersionRarity)

					// 这是一个类似进度条的方法。
					// TODO: 需要整理一下 \033[2K 这是什么意思
					fmt.Printf("\r\033[2K 已生成 %v 个待处理商品 \033[0m", len(needHandleProducts.products))

					needHandleProducts.count++

					isUpdate = 1
					break
				} else {
					logrus.Debugf("当前商品 [%v %v %v %v] 与期望处理的商品 [%v %v %v %v] 不匹配，跳过",
						p.CardVersionID, p.CardVersionNumber, p.CardNameCN, p.CardVersionRarity,
						card.CardVersionID, card.Serial, card.ScName, card.Rarity)
					continue
				}
			}
			// 挺尴尬的做法，通过卡牌名称获取到的商品可能是该卡牌的其它版本的商品(各种异画)。。。o(╯□╰)o
			// 所以需要一个有状态的数据来记录待更新卡牌是否获取到商品
			if isUpdate == 0 {
				logrus.Errorf("%v %v 卡牌存在其它版本的商品，但没有当前版本商品可以更新", card.CardVersionID, card.ScName)
				updateSkip++
			}
		}(card)
	}

	return &needHandleProducts
}

// TODO: 下面这个接口与 genNeedUpdateProducts 接口各有优缺点，还有什么其他的好用的接口么，可以通过卡牌的唯一ID获取到商品信息~~o(╯□╰)o
// 生成需要更新的商品信息
func genNeedUpdateProductsWithBySellerCardVersionId(cards *dbmodels.CardsPrice) *NeedHandleProducts {
	var wg sync.WaitGroup
	defer wg.Wait()
	var lock sync.Mutex // 并发中有对数组的 append 操作，加锁保证并发安全

	var needHandleProducts NeedHandleProducts

	for _, card := range cards.Data {
		wg.Add(1)

		go func(card dbmodels.CardPrice) {
			defer wg.Done()
			// 使用 /api/market/products/bySellerCardVersionId 接口时提交卖家 ID 和 card_version_id，即可获得唯一指定卡牌的商品信息，而不用其他逻辑来判断该卡牌是原画还是异画。
			// 然后，只需要遍历修改这些商品即可。
			// 但是，该接口只能获取到在售的商品，已经下架的商品无法获取到。所以想要修改下架后的商品价格或者让商品的状态变为在售或下架，是不可能的。
			// 后来，官方添加了一个 default_product 的字段，这里也可以获得 product_id、价格、等 信息。并且就算是**下架的商品**也可以看到，只不过如果一张卡牌有多个商品，则 defautl_product 只能看到 1 件商品。
			// 但是感觉用起来还是有点奇怪，待补充...
			p, err := handler.H.JhsServices.Products.Get(fmt.Sprint(card.CardVersionID), updateFlags.SellerUserID)
			if err != nil {
				logrus.Errorf("获取 %v 卡牌的商品失败: %v", card.ScName, err)
			}

			// 由于 go 语言中，for 每次迭代的值都是存储到同一个内存空间中的，所以想要引用 for 中 value 的内存空间，需要再手动创建一个
			cardCopy := card

			lock.Lock() // append 切片在并发中不安全，加个锁
			needHandleProducts.products = append(needHandleProducts.products, Product{
				card:           &cardCopy,
				product:        nil,
				defaultProduct: &p.DefaultProduct,
				productID:      p.DefaultProduct.ProductID,
				condition:      p.DefaultProduct.Condition,
				cardVersionID:  card.CardVersionID,
				cardNameCN:     p.CardNameCN,
				img:            p.CardVersionImage,
				onSale:         1,
				price:          fmt.Sprintf("%.2f", p.DefaultProduct.Price),
				quantity:       p.DefaultProduct.Quantity,
			})
			lock.Unlock()

			logrus.WithFields(logrus.Fields{
				"原始价格": card.AvgPrice,
				"商品价格": p.DefaultProduct.Price,
				"商品数量": p.DefaultProduct.Quantity,
			}).Debugf("检查匹配到的商品: 【%v】【%v】【 %v %v 】【 %v 】", p.DefaultProduct.ProductID, card.CardVersionID, card.Serial, p.CardNameCN, p.CardVersionRarity)

			fmt.Printf("\r\033[2K 已生成 %v 个待处理商品 \033[0m", len(needHandleProducts.products))

			needHandleProducts.count++
		}(card)
	}

	return &needHandleProducts
}

// 更新商品
func updateRun(p *Product) {
	// 生成备注信息
	var remark string
	if updateFlags.Remark != "" {
		remark = updateFlags.Remark
	} else {
		if p.defaultProduct != nil {
			remark = p.defaultProduct.Remark
		} else if p.product != nil {
			remark = p.product.Remark
		}
	}

	// 使用 /api/market/sellers/products/{product_id} 接口更新商品信息
	resp, err := handler.H.JhsServices.Market.SellersProductsUpdate(&models.ProductsUpdateReqBody{
		AuthenticatorID:         "",
		Grading:                 "",
		Condition:               fmt.Sprint(p.condition),
		Default:                 "1",
		OnSale:                  p.onSale,
		Price:                   p.price,
		ProductCardVersionImage: p.img,
		Quantity:                fmt.Sprint(p.quantity),
		Remark:                  remark,
	}, fmt.Sprint(p.productID))
	if err != nil {
		logrus.Errorf("商品 %v %v 修改失败：%v", p.productID, p.cardNameCN, err)
		updateFailCount++
	} else {
		logrus.Infof("商品 %v %v 修改成功：%v", p.productID, p.cardNameCN, resp)
		updateSuccessCount++
	}
}
