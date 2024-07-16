<script setup lang="ts">
import { ref } from "vue";
defineOptions({
  name: "brand"
});
import { useColumns } from "./brand";

import "plus-pro-components/es/components/dialog-form/style/css";

const tableRef = ref();

const {
  editMap,
  columns,
  dataList,
  onEdit,
  onSave,
  onCancel,
  onCloseCallBackClick,
  loading,
  pagination,
  loadingConfig,
  adaptiveConfig,
  onCurrentChange
} = useColumns();
</script>

<template>
  <div>
    <pure-table
      ref="tableRef"
      :adaptiveConfig="adaptiveConfig"
      row-key="id"
      showOverflowTooltip
      :loading="loading"
      :loading-config="loadingConfig"
      :data="dataList"
      :columns="columns"
      :pagination="pagination"
      @page-current-change="onCurrentChange"
    >
      <template #operation="{ row, index }">
        <div class="button-container">
          <el-button
            v-if="!editMap[index]?.editable"
            class="reset-margin"
            link
            type="primary"
            @click="onEdit(row, index)"
          >
            修改
          </el-button>
          <el-button
            v-if="!editMap[index]?.editable"
            class="reset-margin"
            link
            type="primary"
            @click="onCloseCallBackClick(index, row)"
          >
            删除
          </el-button>
          <div v-else>
            <el-button
              class="reset-margin"
              link
              type="primary"
              @click="onSave(row, index)"
            >
              保存
            </el-button>
            <el-button class="reset-margin" link @click="onCancel(index)">
              取消
            </el-button>
          </div>
        </div>
      </template>
    </pure-table>
  </div>
</template>

<style scoped>
.button-container {
  display: flex;
  gap: 10px; /* 控制按钮之间的间距 */
}

.reset-margin {
  margin: 0;
}
</style>
