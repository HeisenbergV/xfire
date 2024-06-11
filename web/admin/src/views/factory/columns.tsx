import type {
  LoadingConfig,
  AdaptiveConfig,
  PaginationProps
} from "@pureadmin/table";
import { tableData } from "./data";
import { ref, onMounted, reactive } from "vue";
import { clone, delay } from "@pureadmin/utils";
import { message } from "@/utils/message";

import "plus-pro-components/es/components/dialog-form/style/css";
import {
  type PlusColumn,
  type FieldValues,
} from "plus-pro-components";


const visible = ref(false);
const values = ref<FieldValues>({});
function handleEdit (index: number, row )  {
    visible.value = true;
    values.value = row
};

function handleDelete (index: number, row )  {
    message(`您删除了第 ${index} 行，数据为：${JSON.stringify(row)}`);
};


export function form() {
  let columns: PlusColumn[] = [
    {
      label: "名称",
      width: 120,
      prop: "name",
      valueType: "copy",
      tooltip: "名称最多显示6个字符"
    },

    {
      label: "时间",
      prop: "time",
      valueType: "date-picker"
    },
    {
      label: "数量",
      prop: "number",
      valueType: "input-number",
      fieldProps: { precision: 2, step: 2 }
    },
    {
      label: "地区",
      prop: "place",
      tooltip: "请精确到门牌号",
      fieldProps: {
        placeholder: "请精确到门牌号"
      }
    },
    {
      label: "要求",
      prop: "demand",
      valueType: "checkbox",
      options: [
        {
          label: "四六级",
          value: "0"
        },
        {
          label: "计算机二级证书",
          value: "1"
        },
        {
          label: "普通话证书",
          value: "2"
        }
      ]
    },
    {
      label: "梦想",
      prop: "gift",
      valueType: "radio",
      options: [
        {
          label: "诗",
          value: "0"
        },
        {
          label: "远方",
          value: "1"
        },
        {
          label: "美食",
          value: "2"
        }
      ]
    },
    {
      label: "到期时间",
      prop: "endTime",
      valueType: "date-picker",
      fieldProps: {
        type: "datetimerange",
        startPlaceholder: "请选择开始时间",
        endPlaceholder: "请选择结束时间"
      }
    },
    {
      label: "说明",
      prop: "desc",
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
    message("fff");
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
  const dataList = ref([]);
  const loading = ref(true);
  const datacolumns: TableColumnList = [
     {
      label: "条码",
      prop: "barcode"
    }, {
      label: "产品",
      prop: "name"
    }, {
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
            Edit
          </el-button>
          <el-button
            size="small"
            type="danger"
            onClick={() => handleDelete(index + 1, row)}
          >
            Delete
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

  function onSizeChange(val) {
    console.log("onSizeChange", val);
  }

  function onCurrentChange(val) {
    loadingConfig.text = `正在加载第${val}页...`;
    loading.value = true;
    delay(600).then(() => {
      loading.value = false;
    });
  }

  onMounted(() => {
    delay(600).then(() => {
      const newList = [];
      newList.push(clone(tableData, true));
      newList.flat(Infinity).forEach((item, index) => {
        dataList.value.push({ id: index, ...item });
      });
      pagination.total = dataList.value.length;
      loading.value = false;
    });
  });

  return {
    loading,
    datacolumns,
    dataList,
    pagination,
    loadingConfig,
    adaptiveConfig,
    onSizeChange,
    onCurrentChange
  };
}
