import type {
  LoadingConfig,
  AdaptiveConfig,
  PaginationProps
} from "@pureadmin/table";
import { ref, onMounted, reactive } from "vue";
import { clone, delay } from "@pureadmin/utils";
import { message } from "@/utils/message";
import forms from "./showProduct.vue";

import {
  addDialog,
  closeDialog,
  updateDialog,
  closeAllDialog
} from "@/components/ReDialog";

import "plus-pro-components/es/components/dialog-form/style/css";
import {
  type PlusColumn,
  type FieldValues,
} from "plus-pro-components";
import { getProductList,delProduct, Mianbao ,getProductBuildInfo} from "@/api/product";

const visible = ref(false);
const values = ref<FieldValues>({});
function handleEdit (index: number, row )  {
    visible.value = true;
    values.value = row
};
const dataList = ref<Mianbao[]>([]);

function handleDelete(index: number, row: Mianbao) {
  delProduct(row.id).then(response => {
    if (response.code == 0) {
      message(`删除成功`);
      dataList.value.splice(index, 1);
    } else {
     message(`删除失败`);
     }
   });
};


function onCloseCallBackClick(s: number, row: Mianbao) {
  addDialog({
    title: "警告",
    closeCallBack: ({ options, index, args }) => {
      console.log(options, index, args);
      let text = "";
      if (args?.command === "sure") {
        handleDelete(s,row)
      } 
    },
    contentRenderer: () => <p>确认要删除？</p>
  });
}

function onShowProduct(row: Mianbao) {
   getProductBuildInfo(row.id).then(v => {
      addDialog({
    width: "50%",
    title: row.name + "制作工艺",
    contentRenderer: () => forms,
    props: {
      // 赋默认值
      formInline: {
        M: row,
        BuildInfo: v,
        user: "菜虚鲲",
        region: "浙江"
      }
    }
  });
  });
}


export function form() {
  let columns: PlusColumn[] = [
  {
      label: "条码",
      width: 80,
      prop: "barcode",
      valueType: "copy",
    },
    {
      label: "名称",
      width: 80,
      prop: "name",
      valueType: "copy",
    },
    {
      label: "品牌",
      width: 80,
      prop: "brand",
      valueType: "copy",
    },
   {
      label: "制作工艺",
      width: 80,
      prop: "build_id",
     valueType: "cascader",
      options: [
      {
        value: "0",
        label: "陕西",
      },
      {
        value: "1",
        label: "山西",
      }
    ]
    },
    {
      label: "重量",
      prop: "unit",
      tooltip: "单位g",
      valueType: "input-number",
    },
    {
      label: "单价",
    prop: "price",
    valueType: "input-number",
    },
     {
    label: "简介",
    prop: "remake",
    valueType: "textarea",
    fieldProps: {
      maxlength: 10,
      showWordLimit: true,
      autosize: { minRows: 2, maxRows: 4 }
    }
  }
  ];

  const handleAllConfirm = async (handleSubmit: () => Promise<boolean>) => {
  const isPass = await handleSubmit()
    isPass && (visible.value = false)
  }

 function handleCancel() {
    visible.value = false;
  }

  return {
    columns,
    visible,
    handleAllConfirm,
    handleCancel,
    values
  }
}

export function useColumns() {
  const loading = ref(true);
  const datacolumns: TableColumnList = [
     {
      label: "条码",
      prop: "barcode"
    },{
      label: "产品",
      prop: "name",
      cellRenderer: ({ row, index }) => (
     <el-tag onClick={() => onShowProduct(row)}>
          <p style="text-decoration: underline;">{ row.name}</p>
            </el-tag>
      )
    },{
      label: "品牌",
      prop: "brand"
    },
    {
      label: "简介",
      prop: "remake"
    },
    {
      label: "重量/g",
      prop: "unit"
    },  {
      label: "单价",
      prop: "price"
    },  {
       cellRenderer: ({ index, row }) => (
        <>
          <el-button size="small" onClick={() => handleEdit(index + 1, row)}>
            修改
          </el-button>
          <el-button
            size="small"
            type="danger"
            onClick={() => onCloseCallBackClick(index,row)}
          >
            删除
          </el-button>
        </>
      )
    }
  ];

  /** 分页配置 */
  const pagination = reactive<PaginationProps>({
    pageSize: 20,
    currentPage: 1,
    pageSizes: [20, 40, 60],
    total: 0,
    align: "right",
    background: true,
    small: false
  });

    
  /** 加载动画配置 */
  const loadingConfig = reactive<LoadingConfig>({
    text: "正在加载第一页...",
    viewBox: "-10, -10, 50, 50",
    spinner: `
        <path class="path" d="
          M 30 15
          L 28 17
          M 25.61 25.61
          A 15 15, 0, 0, 1, 15 30
          A 15 15, 0, 1, 1, 27.99 7.5
          L 15 15
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)"/>
      `
    // svg: "",
    // background: rgba()
  });

  /** 撑满内容区自适应高度相关配置 */
  const adaptiveConfig: AdaptiveConfig = {
    /** 表格距离页面底部的偏移量，默认值为 `96` */
    offsetBottom: 110
    /** 是否固定表头，默认值为 `true`（如果不想固定表头，fixHeader设置为false并且表格要设置table-layout="auto"） */
    // fixHeader: true
    /** 页面 `resize` 时的防抖时间，默认值为 `60` ms */
    // timeout: 60
    /** 表头的 `z-index`，默认值为 `100` */
    // zIndex: 100
  };

  function onCurrentChange(val) {
    loadingConfig.text = `正在加载第${val}页...`;
    loading.value = true;
   
    delay(600).then(() => {
      getCardListData();
    });
  }
  
  const getCardListData = async () => {
    try {
        let data = await getProductList({
          page: pagination.currentPage,
          pageSize: pagination.pageSize,
          ptype: "成品",
          orderKey: "barcode"
        });
      
        pagination.total = data.data.total;
        dataList.value = data.data.list;
        loading.value = false;
    } finally {
      setTimeout(() => {}, 500);
    }
  };

  onMounted(() => {
    delay(600).then(() => {
      getCardListData();
    });
  });

  return {
    loading,
    datacolumns,
    dataList,
    pagination,
    loadingConfig,
    adaptiveConfig,
    onCurrentChange
  };
}
