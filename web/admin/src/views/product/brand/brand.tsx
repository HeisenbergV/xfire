import type {
  LoadingConfig,
  AdaptiveConfig,
  PaginationProps
} from "@pureadmin/table";
import { ref, onMounted, reactive } from "vue";
import { delay } from "@pureadmin/utils";
import { message } from "@/utils/message";
import dayjs from "dayjs";
import { clone, delObjectProperty } from "@pureadmin/utils";

import { addDialog } from "@/components/ReDialog";

import { getBrandList, updateBrand, delBrand, Brand } from "@/api/product";

const tableData = ref<Brand[]>([]);
const dataList = ref(clone(tableData, true));

function handleDelete(index: number, row: Brand) {
  delBrand(row.id).then(response => {
    if (response.code == 0) {
      message(`删除成功`);
      dataList.value.splice(index, 1);
    } else {
      message(`删除失败`);
    }
  });
}

function onCloseCallBackClick(s: number, row: Brand) {
  addDialog({
    title: "警告",
    closeCallBack: ({ options, index, args }) => {
      console.log(options, index, args);
      if (args?.command === "sure") {
        handleDelete(s, row);
      }
    },
    contentRenderer: () => <p>确认要删除？</p>
  });
}

export function useColumns() {
  const editMap = ref({});

  const loading = ref(true);
  const columns: TableColumnList = [
    {
      label: "品牌",
      prop: "name",
      cellRenderer: ({ row, index }) => (
        <>
          {editMap.value[index]?.editable ? (
            <el-input v-model={row.name} />
          ) : (
            <p>{row.name}</p>
          )}
        </>
      )
    },
    {
      label: "创建时间",
      prop: "created_at",
      cellRenderer: ({ row }) => (
        <div style="display: flex; align-items: center">
          <iconify-icon-online icon="ep:timer" />
          <span style="margin-left: 10px">
            {dayjs(row.created_at).format("YYYY-MM-DD")}
          </span>
        </div>
      )
    },
    {
      label: "操作",
      fixed: "right",
      slot: "operation"
    }
  ];

  function onEdit(row, index) {
    editMap.value[index] = Object.assign({ ...row, editable: true });
  }

  function onSave(row, index) {
    editMap.value[index].editable = false;
    updateBrand(row.id, row.name).then(response => {
      if (response.code == 0) {
        message(`修改成功`);
      } else {
        message(`修改失败`);
      }
    });
  }

  function onCancel(index) {
    editMap.value[index].editable = false;
    dataList.value[index] = delObjectProperty(editMap.value[index], "editable");
  }

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
        " style="stroke-width: 4px; fill: rgba(0, 0, 0, 0)" />
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
      let data = await getBrandList({
        page: pagination.currentPage,
        pageSize: pagination.pageSize,
        orderkey: "id"
      });

      pagination.total = data.data.total;

      dataList.value = data.data.list;
      console.log(tableData);
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
    editMap,
    columns,
    dataList,
    onEdit,
    onSave,
    onCancel,
    onCloseCallBackClick,
    pagination,
    loadingConfig,
    adaptiveConfig,
    onCurrentChange
  };
}
