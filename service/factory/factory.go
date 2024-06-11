package factory

import (
	"errors"
	"fmt"
	"xfire/global"
	"xfire/model"
	"xfire/model/request"

	"gorm.io/gorm"
)

type Factory struct{}

// brand
func (f *Factory) GetBrand(id int) (*model.Brand, error) {
	var b model.Brand
	err := global.DB.First(&b, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (f *Factory) GetBrandList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	db := global.DB.Model(&model.Brand{})
	var blist []model.Brand
	err = db.Count(&total).Error

	if err != nil {
		return blist, total, err
	}
	if order == "" {
		err = db.Order("name").Find(&blist).Error
		return blist, total, err
	}

	offset := info.PageSize * (info.Page - 1)
	db = db.Limit(info.PageSize).Offset(offset)
	var OrderStr string
	// 设置有效排序key 防止sql注入
	orderMap := make(map[string]bool, 5)
	orderMap["id"] = true
	orderMap["name"] = true
	if orderMap[order] {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	} else { // didn't match any order key in `orderMap`
		err = fmt.Errorf("非法的排序字段: %v", order)
		return blist, total, err
	}

	err = db.Order(OrderStr).Find(&blist).Error
	return blist, total, err
}

func (f *Factory) CreateBrand(brand model.Brand) error {
	return global.DB.Create(&brand).Error
}

func (f *Factory) DelBrand(id []int) error {
	err := global.DB.Where("id IN(?)", id).Delete(&model.Brand{}).Error
	return err
}

func (f *Factory) UpdateBrand(brand model.Brand) error {
	var duplicateBrand model.Brand
	if err := global.DB.Where("name = ?", brand.Name).First(&duplicateBrand).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return global.DB.Where("id=?", brand.ID).Updates(&model.Brand{Name: brand.Name}).Error
}

// build
// bdata map[int]float64 其中int表示原材料id， float64表示使用数量
func (f *Factory) GetBuildList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	db := global.DB.Model(&model.Build{})
	var blist []model.Build
	err = db.Count(&total).Error

	if err != nil {
		return blist, total, err
	}

	if order == "" {
		err = db.Order("name").Find(&blist).Error
	}

	if err != nil {
		return blist, total, err
	}

	offset := info.PageSize * (info.Page - 1)
	db = db.Limit(info.PageSize).Offset(offset)
	var OrderStr string

	// 设置有效排序key 防止sql注入
	orderMap := make(map[string]bool, 5)
	orderMap["id"] = true
	orderMap["name"] = true
	if !orderMap[order] {
		err = fmt.Errorf("非法的排序字段: %v", order)
		return blist, total, err

	}

	if desc {
		OrderStr = order + " desc"
	} else {
		OrderStr = order
	}

	err = db.Order(OrderStr).Find(&blist).Error
	if err != nil {
		return blist, total, err
	}

	return blist, total, err
}

func (f *Factory) CreateBuild(build model.Build) error {
	return global.DB.Create(&build).Error
}

func (f *Factory) UpdateBuild(build model.Build) error {
	var duplicateBuild model.Build
	if err := global.DB.Where("name = ?", build.Name).First(&duplicateBuild).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return global.DB.Where("id=?", build.ID).Updates(&model.Build{Name: build.Name,
		Remake: build.Remake, Info: build.Info, Water: build.Water, Bakingtime: build.Bakingtime,
		Bakingtem: build.Bakingtem, BuildData: build.BuildData,
	}).Error
}

func (f *Factory) GetBuildInfo(id int) (*model.Build, error) {
	var b model.Build
	err := global.DB.Where("id=?", id).First(&b).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (f *Factory) DelBuild(id []int) error {
	err := global.DB.Where("id in(?)", id).Delete(&model.Build{}).Error
	return err
}

// customer
func (f *Factory) GetCustomer(id int) (*model.Customer, error) {
	var b model.Customer
	err := global.DB.First(&b, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func (f *Factory) CreateCustomer(customer model.Customer) error {
	return global.DB.Create(customer).Error
}

func (f *Factory) DelCustomer(id int) error {
	err := global.DB.Where("id=?", id).Delete(&model.Customer{}).Error
	return err
}

func (f *Factory) UpdateCustomer(customer model.Customer) error {
	return global.DB.Model(&customer).Where("id=?", customer.ID).Updates(model.Customer{
		Name:         customer.Name,
		Phone:        customer.Phone,
		Address:      customer.Address,
		Remake:       customer.Remake,
		CustomerType: customer.CustomerType,
	}).Error
}

func (f *Factory) GetGoodsInfo(id int) (model.Goods, error) {
	var b model.Goods
	err := global.DB.First(&b, "id =? ", id).Error
	return b, err
}

func (f *Factory) GetGoods(id []int) ([]model.Goods, error) {
	var b []model.Goods
	err := global.DB.Find(&b, "id in(?) ", id).Error
	return b, err
}

func (f *Factory) GetGoodsList(ptype model.Goodstype, info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	db := global.DB.Model(&model.Goods{}).Where("ptype=?", ptype)
	var plist []model.Goods
	err = db.Count(&total).Error

	if err != nil {
		return plist, total, err
	}

	offset := info.PageSize * (info.Page - 1)
	db = db.Limit(info.PageSize).Offset(offset)

	if order == "" {
		err = db.Order("name").Find(&plist).Error
		return plist, total, err
	}

	var OrderStr string
	// 设置有效排序key 防止sql注入
	orderMap := make(map[string]bool, 5)
	orderMap["id"] = true
	orderMap["name"] = true
	orderMap["brand_id"] = true
	orderMap["barcode"] = true
	if orderMap[order] {
		if desc {
			OrderStr = order + " desc"
		} else {
			OrderStr = order
		}
	} else { // didn't match any order key in `orderMap`
		err = fmt.Errorf("非法的排序字段: %v", order)
		return plist, total, err
	}

	err = db.Order(OrderStr).Find(&plist).Error

	return plist, total, err
}

func (f *Factory) CreateGoods(Goods model.Goods) error {
	//无此品牌
	if err := global.DB.Model(&model.Brand{}).Where("name=?", Goods.Brand).First(&model.Brand{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	//产品名称+品牌 重复
	var duplicateGoods model.Goods
	if err := global.DB.Where("name = ? and brand=?", Goods.Name, Goods.Brand).First(&duplicateGoods).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return global.DB.Create(Goods).Error
}

func (f *Factory) DelGoods(id []int) error {
	err := global.DB.Where("id in(?)", id).Delete(&model.Goods{}).Error
	return err
}

func (f *Factory) UpdateGoods(Goods model.Goods) error {
	return global.DB.Model(&Goods).Where("id=?", Goods.ID).Updates(&model.Goods{
		Name: Goods.Name, Brand: Goods.Brand, Remake: Goods.Remake,
		Barcode: Goods.Barcode, Unit: Goods.Unit,
		Price: Goods.Price, BuildID: Goods.BuildID,
	}).Error
}

func (f *Factory) ReduceGoods(id int, quantity float64) error {
	return global.DB.Model(&model.Goods{}).Where("id = ? ", id).Update("stock", gorm.Expr("stock- ?", quantity)).Error
}

func (f *Factory) AddGoods(id int, quantity float64) error {
	return global.DB.Model(&model.Goods{}).Where("id = ? ", id).Update("stock", gorm.Expr("stock+ ?", quantity)).Error
}
