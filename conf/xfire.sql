use xfire;

DROP TABLE IF EXISTS brand;
CREATE TABLE brand (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL UNIQUE COMMENT '品牌名称',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS customer;
CREATE TABLE customer (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '客户ID',
    name VARCHAR(255) NOT NULL COMMENT '客户名称',
    phone VARCHAR(20) COMMENT '联系方式',
    address VARCHAR(255) COMMENT '客户地址',
    remake VARCHAR(255) COMMENT '简介',
    customer_type ENUM('厂家直销', '经销商','原料厂家'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS goods;
CREATE TABLE goods (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '产品ID',
    name VARCHAR(255) NOT NULL COMMENT '产品名称',
    ptype ENUM('原料', '成品')  NOT NULL,
    brand VARCHAR(255) NOT NULL COMMENT '品牌',
    remake VARCHAR(255) COMMENT '备注',
    barcode VARCHAR(255) NOT NULL UNIQUE COMMENT '产品条码',
    unit DECIMAL(10, 2) COMMENT '产品规格/g',
    price DECIMAL(10, 2) COMMENT '产品价格',
    build_id INT COMMENT '关联制作工艺' DEFAULT -1,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS inventory;
CREATE TABLE inventory (
    name VARCHAR(255) PRIMARY KEY NOT NULL COMMENT '产品名称',
    brand VARCHAR(255) NOT NULL COMMENT '品牌',
    ptype ENUM('原料', '成品')  NOT NULL,
    unit DECIMAL(10, 2) NOT NULL COMMENT '产品规格/g',
    remake VARCHAR(255) COMMENT '备注',
    stock DECIMAL(10, 2) COMMENT '当前库存数量'  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS build;
CREATE TABLE build (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL COMMENT '工艺名称',
    remake VARCHAR(255) COMMENT '备注',
    info VARCHAR(1000) COMMENT '制作流程',
    materialUnit DECIMAL(10, 2) COMMENT '成型重量/g',
    bakingtime  DECIMAL(10, 2)  NOT NULL COMMENT'烘焙时间',
    bakingtem  DECIMAL(10, 2)  NOT NULL COMMENT'烘焙温度',
    water  DECIMAL(10, 2)  NOT NULL COMMENT'水重量g',
    build_data JSON NOT NULL COMMENT '字典 材料name-使用克重',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS drp;
CREATE TABLE drp (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '记录ID',
    Goods_id INT NOT NULL COMMENT '产品名称',
    remake VARCHAR(255) COMMENT '备注',
    unit VARCHAR(50) COMMENT '产品单位(箱;个等)',
    price DECIMAL(10, 2) COMMENT '单价',
    drp_type  ENUM('生产', '进货', '出货','退货','残次','样品')  NOT NULL,
    quantity DECIMAL(10, 2) COMMENT '数量'   NOT NULL,
    customer_id INT COMMENT '关联到客户表id'  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


DROP TABLE IF EXISTS finance;
CREATE TABLE finance (
    id INT PRIMARY KEY AUTO_INCREMENT COMMENT '记录ID',
    remake VARCHAR(255) COMMENT '备注',
    ftype ENUM('支出', '收入')  NOT NULL ,
    sftype ENUM('损失', '盈利', '进货', '日常','投资')  NOT NULL COMMENT '收支子类',
    price DECIMAL(10, 2) COMMENT '价格'  NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DROP TABLE IF EXISTS sys_log;
CREATE TABLE sys_log (
    id INT PRIMARY KEY AUTO_INCREMENT,
    info VARCHAR(255) COMMENT '变动信息',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);