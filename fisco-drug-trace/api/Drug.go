package api

import (
	"fisco-drug-trace/model"
	"fisco-drug-trace/sdk"
	"fisco-drug-trace/serializer"
	"github.com/gin-gonic/gin"
	"math/big"
)

type DrugVo struct {
	Id                    *big.Int `json:"id" `
	Producer              string   `json:"producer" `
	ProducerId            *big.Int `json:"producerId" `
	ProductionDate        *big.Int `json:"productionDate" `
	ProductionDateStr     string   `json:"productionDateStr" `
	Flow                  string   `json:"flow" `
	TransportationStr     string   `json:"transportationStr" `
	Dealer                string   `json:"dealer" `
	DealerId              *big.Int `json:"dealerId" `
	DrugAcceptanceTime    *big.Int `json:"drugAcceptanceTime" `
	DrugStorageConditions string   `json:"drugStorageConditions" `
	DrugStorageLocation   string   `json:"drugStorageLocation" `
	Buyer                 string   `json:"buyer" `
	BuyerId               *big.Int `json:"buyerId" `
	BuyTime               *big.Int `json:"buyTime" `
}

type DrugListAo struct {
	serializer.PageAo
	Id *big.Int `json:"id" `
}

// 药品列表
func DrugList(c *gin.Context) {
	ao := new(DrugListAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		count, err := sdk.Contract.GetDrugCount(big.NewInt(0).SetUint64(uint64(user.ID)), ao.Id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		var drugVos []DrugVo
		if count.Cmp(big.NewInt(0)) == 0 {
			c.JSON(200, serializer.BuildPage(drugVos, 0))
			return
		}

		drugs, err := sdk.Contract.GetDrugs(big.NewInt(0).SetInt64(int64(ao.Page)), big.NewInt(0).SetInt64(int64(ao.PageSize)), big.NewInt(0).SetUint64(uint64(user.ID)), ao.Id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		for _, drug := range drugs {
			if drug.Id.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			var producer *model.User
			producer, err = getUser(drug.Producer)
			if err != nil {
				c.JSON(200, ErrorResponse(err))
				return
			}
			var dealer string
			if drug.Dealer != nil && drug.Dealer.Cmp(big.NewInt(0)) != 0 {
				dealerUser, err := getUser(drug.Dealer)
				if err != nil {
					c.JSON(200, ErrorResponse(err))
					return
				}
				dealer = dealerUser.Nickname
			} else {
				dealer = ""
			}
			var buyer string
			if drug.Buyer != nil && drug.Buyer.Cmp(big.NewInt(0)) != 0 {
				buyerUser, err := getUser(drug.Buyer)
				if err != nil {
					c.JSON(200, ErrorResponse(err))
					return
				}
				buyer = buyerUser.Nickname
			} else {
				buyer = ""
			}

			drugVo := DrugVo{
				Id:                    drug.Id,
				Producer:              producer.Nickname,
				ProductionDate:        drug.ProductionDate,
				ProductionDateStr:     drug.ProductionDateStr,
				Flow:                  drug.Flow,
				TransportationStr:     drug.TransportationStr,
				Dealer:                dealer,
				DrugAcceptanceTime:    drug.DrugAcceptanceTime,
				DrugStorageConditions: drug.DrugStorageConditions,
				DrugStorageLocation:   drug.DrugStorageLocation,
				Buyer:                 buyer,
				BuyTime:               drug.BuyTime,
			}
			drugVos = append(drugVos, drugVo)
		}

		c.JSON(200, serializer.BuildPage(drugVos, count.Int64()))
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// 生产商生成药品
type ProduceDrugAo struct {
	ProductionDateStr string `json:"productionDateStr" `
}

func ProduceDrug(c *gin.Context) {
	ao := new(ProduceDrugAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		_, _, err := sdk.Contract.AddDrug(big.NewInt(0).SetUint64(uint64(user.ID)), ao.ProductionDateStr)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 销售给经销商
type SellDrugAo struct {
	DrugId            *big.Int `json:"drugId" `
	Dealer            *big.Int `json:"dealer" `
	TransportationStr string   `json:"transportationStr" `
}

func SellDrug(c *gin.Context) {
	ao := new(SellDrugAo)
	if err := c.ShouldBind(ao); err == nil {
		_, _, err := sdk.Contract.Sale(ao.DrugId, ao.Dealer, ao.TransportationStr)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 经销商签收
type DealerAcceptDrugAo struct {
	DrugId                *big.Int `json:"drugId" `
	DrugStorageConditions string   `json:"drugStorageConditions" `
	DrugStorageLocation   string   `json:"drugStorageLocation" `
}

func DealerAcceptDrug(c *gin.Context) {
	ao := new(DealerAcceptDrugAo)
	if err := c.ShouldBind(ao); err == nil {
		_, _, err := sdk.Contract.Accept(ao.DrugId, ao.DrugStorageConditions, ao.DrugStorageLocation)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

// 分页查询可购买的药品
func DrugPayableList(c *gin.Context) {
	ao := new(DrugListAo)
	if err := c.ShouldBind(ao); err == nil {
		count, err := sdk.Contract.GetPayableDrugCount(ao.Id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		var drugVos []DrugVo
		if count.Cmp(big.NewInt(0)) == 0 {
			c.JSON(200, serializer.BuildPage(drugVos, 0))
			return
		}

		drugs, err := sdk.Contract.GetPayableDrugs(big.NewInt(0).SetInt64(int64(ao.Page)), big.NewInt(0).SetInt64(int64(ao.PageSize)), ao.Id)
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		for _, drug := range drugs {
			if drug.Id.Cmp(big.NewInt(0)) == 0 {
				continue
			}
			var producer *model.User
			producer, err = getUser(drug.Producer)
			if err != nil {
				c.JSON(200, ErrorResponse(err))
				return
			}
			var dealer string
			if drug.Dealer != nil && drug.Dealer.Cmp(big.NewInt(0)) != 0 {
				dealerUser, err := getUser(drug.Dealer)
				if err != nil {
					c.JSON(200, ErrorResponse(err))
					return
				}
				dealer = dealerUser.Nickname
			} else {
				dealer = ""
			}
			var buyer string
			if drug.Buyer != nil && drug.Buyer.Cmp(big.NewInt(0)) != 0 {
				buyerUser, err := getUser(drug.Buyer)
				if err != nil {
					c.JSON(200, ErrorResponse(err))
					return
				}
				buyer = buyerUser.Nickname
			} else {
				buyer = ""
			}

			drugVo := DrugVo{
				Id:                    drug.Id,
				Producer:              producer.Nickname,
				ProductionDate:        drug.ProductionDate,
				ProductionDateStr:     drug.ProductionDateStr,
				Flow:                  drug.Flow,
				TransportationStr:     drug.TransportationStr,
				Dealer:                dealer,
				DrugAcceptanceTime:    drug.DrugAcceptanceTime,
				DrugStorageConditions: drug.DrugStorageConditions,
				DrugStorageLocation:   drug.DrugStorageLocation,
				Buyer:                 buyer,
				BuyTime:               drug.BuyTime,
			}
			drugVos = append(drugVos, drugVo)
		}

		c.JSON(200, serializer.BuildPage(drugVos, count.Int64()))
	} else {
		c.JSON(200, ErrorResponse(err))
	}

}

// 购买药品
type BuyDrugAo struct {
	DrugId *big.Int `json:"drugId" `
}

func BuyDrug(c *gin.Context) {
	ao := new(BuyDrugAo)
	if err := c.ShouldBind(ao); err == nil {
		user := CurrentUser(c)
		_, _, err := sdk.Contract.Buy(ao.DrugId, big.NewInt(0).SetUint64(uint64(user.ID)))
		if err != nil {
			c.JSON(200, ErrorResponse(err))
			return
		}
		c.JSON(200, serializer.Success(nil))
	} else {
		c.JSON(200, ErrorResponse(err))
	}
}

func getUser(id *big.Int) (*model.User, error) {
	user, err := model.GetUser(id.Uint64())
	if err != nil {
		return nil, err
	}
	return &user, nil
}
