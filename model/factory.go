package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

const TableNameBuild = "build"

type BuildData map[string]float64

// Build mapped from table <build>
type Build struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name         string    `gorm:"column:name;not null;comment:工艺名称" json:"name"`                     // 工艺名称
	Remake       string    `gorm:"column:remake;comment:备注" json:"remake"`                            // 备注
	Info         string    `gorm:"column:info;comment:制作流程" json:"info"`                              // 制作流程
	MaterialUnit float64   `gorm:"column:materialUnit;not null;comment:成型重量" json:"materialUnit"`     // 水重量g
	Water        float64   `gorm:"column:water;not null;comment:水重量g" json:"water"`                   // 水重量g
	Bakingtime   float64   `gorm:"column:bakingtime;not null;comment:烘焙时间" json:"bakingtime"`         // 水重量g
	Bakingtem    float64   `gorm:"column:bakingtem;not null;comment:烘焙温度" json:"bakingtem"`           // 水重量g
	BuildData    BuildData `gorm:"column:build_data;not null;comment:字典 材料id-使用克重" json:"build_data"` // 字典 材料id-使用克重
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (c BuildData) Value() (driver.Value, error) {
	b, err := json.Marshal(c)
	return string(b), err
}

func (c *BuildData) Scan(input interface{}) error {
	return json.Unmarshal(input.([]byte), c)
}

// TableName Build's table name
func (*Build) TableName() string {
	return TableNameBuild
}

const TableNameBrand = "brand"

// Brand mapped from table <brand>
type Brand struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string    `gorm:"column:name;not null;comment:品牌名称" json:"name"` // 品牌名称
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Brand's table name
func (*Brand) TableName() string {
	return TableNameBrand
}

const TableNameCustomer = "customer"

// Customer mapped from table <customer>
type Customer struct {
	ID           int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:客户ID" json:"id"` // 客户ID
	Name         string    `gorm:"column:name;not null;comment:客户名称" json:"name"`                  // 客户名称
	Phone        string    `gorm:"column:phone;comment:联系方式" json:"phone"`                         // 联系方式
	Address      string    `gorm:"column:address;comment:客户地址" json:"address"`                     // 客户地址
	Remake       string    `gorm:"column:remake;comment:简介" json:"remake"`                         // 简介
	CustomerType string    `gorm:"column:customer_type" json:"customer_type"`
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Customer's table name
func (*Customer) TableName() string {
	return TableNameCustomer
}

const TableNameDrpRecord = "drp_record"

// DrpRecord mapped from table <drp_record>
type DrpRecord struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:记录ID" json:"id"` // 记录ID
	GoodsID    int       `gorm:"column:Goods_id;not null;comment:产品名称" json:"Goods_id"`          // 产品名称
	Remake     string    `gorm:"column:remake;comment:备注" json:"remake"`                         // 备注
	Unit       string    `gorm:"column:unit;comment:产品单位(箱;个等)" json:"unit"`                     // 产品单位(箱;个等)
	Price      float64   `gorm:"column:price;comment:单价" json:"price"`                           // 单价
	DrpType    string    `gorm:"column:drp_type;not null" json:"drp_type"`
	Quantity   float64   `gorm:"column:quantity;not null;comment:数量" json:"quantity"`             // 数量
	CustomerID int       `gorm:"column:customer_id;not null;comment:关联到客户表id" json:"customer_id"` // 关联到客户表id
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName DrpRecord's table name
func (*DrpRecord) TableName() string {
	return TableNameDrpRecord
}

const TableNameDrp = "drp"

// Drp mapped from table <drp>
type Drp struct {
	ID         int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:记录ID" json:"id"` // 记录ID
	GoodsID    int       `gorm:"column:Goods_id;not null;comment:产品名称" json:"Goods_id"`          // 产品名称
	Remake     string    `gorm:"column:remake;comment:备注" json:"remake"`                         // 备注
	Unit       string    `gorm:"column:unit;comment:产品单位(箱;个等)" json:"unit"`                     // 产品单位(箱;个等)
	Price      float64   `gorm:"column:price;comment:单价" json:"price"`                           // 单价
	DrpType    string    `gorm:"column:drp_type;not null" json:"drp_type"`
	Quantity   float64   `gorm:"column:quantity;not null;comment:数量" json:"quantity"`             // 数量
	CustomerID int       `gorm:"column:customer_id;not null;comment:关联到客户表id" json:"customer_id"` // 关联到客户表id
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Drp's table name
func (*Drp) TableName() string {
	return TableNameDrp
}

const TableNameFinance = "finance"

// Finance mapped from table <finance>
type Finance struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true;comment:记录ID" json:"id"` // 记录ID
	Remake    string    `gorm:"column:remake;comment:备注" json:"remake"`                         // 备注
	Ftype     string    `gorm:"column:ftype;not null" json:"ftype"`
	Sftype    string    `gorm:"column:sftype;not null;comment:收支子类" json:"sftype"` // 收支子类
	Price     float64   `gorm:"column:price;not null;comment:价格" json:"price"`     // 价格
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Finance's table name
func (*Finance) TableName() string {
	return TableNameFinance
}

const TableNameInventory = "inventory"

// Inventory mapped from table <inventory>
type Inventory struct {
	Name      string    `gorm:"column:name;primaryKey;comment:产品名称" json:"name"` // 产品名称
	Brand     string    `gorm:"column:brand;not null;comment:品牌" json:"brand"`   // 品牌
	Ptype     string    `gorm:"column:ptype;not null" json:"ptype"`
	Unit      float64   `gorm:"column:unit;not null;comment:产品规格/g" json:"unit"`   // 产品规格/g
	Remake    string    `gorm:"column:remake;comment:备注" json:"remake"`            // 备注
	Stock     float64   `gorm:"column:stock;not null;comment:当前库存数量" json:"stock"` // 当前库存数量
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// TableName Inventory's table name
func (*Inventory) TableName() string {
	return TableNameInventory
}

const TableNameGoods = "goods"

type Goodstype string

const (
	Product  Goodstype = "成品"
	Material Goodstype = "原料"
)

// Goods mapped from table <goods>
type Goods struct {
	ID        int32     `gorm:"column:id;primaryKey;autoIncrement:true;comment:产品ID" json:"id"` // 产品ID
	Name      string    `gorm:"column:name;not null;comment:产品名称" json:"name"`                  // 产品名称
	Ptype     Goodstype `gorm:"column:ptype;not null" json:"ptype"`
	Brand     string    `gorm:"column:brand;not null;comment:品牌" json:"brand"`             // 品牌
	Remake    string    `gorm:"column:remake;comment:备注" json:"remake"`                    // 备注
	Barcode   string    `gorm:"column:barcode;not null;comment:产品条码" json:"barcode"`       // 产品条码
	Unit      float64   `gorm:"column:unit;comment:产品规格/g" json:"unit"`                    // 产品规格/g
	Price     float64   `gorm:"column:price;comment:产品价格" json:"price"`                    // 产品价格
	BuildID   int       `gorm:"column:build_id;default:-1;comment:关联制作工艺" json:"build_id"` // 关联制作工艺
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName Goods's table name
func (*Goods) TableName() string {
	return TableNameGoods
}

const TableNameSysLog = "sys_log"

// SysLog mapped from table <sys_log>
type SysLog struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Info      string    `gorm:"column:info;comment:变动信息" json:"info"` // 变动信息
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
}

// TableName SysLog's table name
func (*SysLog) TableName() string {
	return TableNameSysLog
}
