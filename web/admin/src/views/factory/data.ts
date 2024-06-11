import dayjs from "dayjs";
import { clone } from "@pureadmin/utils";

const date = dayjs(new Date()).format("YYYY-MM-DD");

const tableData = [
  {
    date,
    name: "大列巴",
    brand: "麦小六",
    remake: "你好你你萨达xxxxxxx你好你你萨达xxxxxxx你好你你萨达xxxxxxx你好你你萨达xxxxxxx你好你你萨达xxxxxxx你好你你萨达xxxxxxx",
    barcode: "qqqqq",
    unit: 300,
    price: 4.5,
  },
  {
    date,
    name: "小列巴",
    brand: "麦小六",
    remake: "你好你你萨达xxxxxx",
    barcode: "333333333",
    unit: 90,
    price: 1.5,
  }
];


const tableDataImage = clone(tableData, true).map((item, index) =>
  Object.assign(item, {
    image: `https://pure-admin.github.io/pure-admin-table/imgs/${index + 1}.jpg`
  })
);

const tableDataSortable = clone(tableData, true).map((item, index) => {
  delete item.date;
  Object.assign(item, {
    date: `${dayjs(new Date()).format("YYYY-MM")}-${index + 1}`
  });
});

const tableDataDrag = clone(tableData, true).map((item, index) => {
  delete item.address;
  delete item.date;
  return Object.assign(item, {
    id: index + 1,
    date: `${dayjs(new Date()).format("YYYY-MM")}-${index + 1}`
  });
});

const tableDataEdit = clone(tableData, true).map((item, index) => {
  delete item.date;
  return Object.assign(item, {
    id: index + 1,
    date: `${dayjs(new Date()).format("YYYY-MM")}-${index + 1}`,
    address: "China",
    sex: index % 2 === 0 ? "男" : "女"
  });
});

export {
  tableData,
  tableDataDrag,
  tableDataEdit,
  tableDataImage,
  tableDataSortable
};
