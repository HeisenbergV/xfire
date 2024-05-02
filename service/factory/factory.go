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

func (f *Factory) GetProduct(id int) (*model.Product, error) {
	var b model.Product
	err := global.DB.First(&b, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &b, nil
}
func (f *Factory) GetProductList(info request.PageInfo, order string, desc bool) (list interface{}, total int64, err error) {
	db := global.DB.Model(&model.Product{})
	var plist []model.Product
	err = db.Count(&total).Error

	if err != nil {
		return plist, total, err
	}
	if order == "" {
		err = db.Order("name").Find(&plist).Error
		return plist, total, err
	}

	offset := info.PageSize * (info.Page - 1)
	db = db.Limit(info.PageSize).Offset(offset)
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

func (f *Factory) CreateProduct(product model.Product) error {
	//无此品牌
	if err := global.DB.Model(&model.Brand{}).Where("id=?", product.BrandID).First(&model.Brand{}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}
	//产品名称+品牌 重复
	var duplicateProduct model.Product
	if err := global.DB.Where("name = ? and brand_id=?", product.Name, product.BrandID).First(&duplicateProduct).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	}

	return global.DB.Create(product).Error
}

func (f *Factory) DelProduct(id []int) error {
	err := global.DB.Where("id in(?)", id).Delete(&model.Product{}).Error
	return err
}

func (f *Factory) UpdateProduct(product model.Product) error {
	return global.DB.Model(&product).Where("id=?", product.ID).Updates(&model.Product{
		Name: product.Name, BrandID: product.BrandID, Remake: product.Remake,
		Barcode: product.Barcode, Specification: product.Specification, MaterialUnit: product.MaterialUnit,
		Price: product.Price, BuildID: product.BuildID,
	}).Error
}

// material
// 原材料增减品类,修改品类信息
func (f *Factory) AddMaterial(material model.Material) error {
	var m model.Material
	if err := global.DB.First(&m, "name = ? and brand = ?", material.Name, material.Brand).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return global.DB.Create(&material).Error
		} else {
			return errors.New("add material fail")
		}
	} else {
		return errors.New("add  material fail")
	}
}

func (f *Factory) DelMaterial(id int) error {
	return global.DB.Where("id=?", id).Delete(&model.Material{}).Error
}

// 只有材料进货与消耗才会修改quantity
func (f *Factory) UpdateMaterial(material model.Material) error {
	return global.DB.Where("id=?", material.ID).Updates(&model.Material{
		Name:  material.Name,
		Brand: material.Brand,
		Info:  material.Info,
		Unit:  material.Unit,
		Price: material.Price,
	}).Error
}

// 原材料消耗与 使用
func (f *Factory) ReplenishConsumeMaterial(id int, quantity float64) error {
	var material model.Material
	if err := global.DB.First(&material, "id = ? ", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("material find error")
		}
	} else {
		// global.DB.Where("id=?", material.ID).Update("quantity", material.Quantity+quantity)
	}
	return errors.New("not this material")
}

// 展示材料价格、剩余
func (f *Factory) GetMaterial(id int) (m *model.Material, err error) {
	err = global.DB.Where("id=?", id).First(&m).Error
	return
}

// 展示材料价格、剩余
func (f *Factory) GetMaterials(id []int) (m []model.Material, err error) {
	err = global.DB.Find(&m, id).Error
	return
}

func (f *Factory) GetMaterialList() (mlist []model.Material, err error) {
	err = global.DB.Find(&mlist).Error
	return
}
