<script setup lang="ts">
import { ref } from "vue";
defineOptions({
  name: "product"
});
import { useColumns, form } from "./product";

import "plus-pro-components/es/components/dialog-form/style/css";
import { PlusDialogForm } from "plus-pro-components";

const tableRef = ref();

const {
  loading,
  dataList,
  datacolumns,
  pagination,
  loadingConfig,
  adaptiveConfig,
  onCurrentChange
} = useColumns();

const { columns, values, visible, handleAllConfirm, handleCancel } = form();
</script>

<template>
  <div>
    <pure-table
      :adaptiveConfig="adaptiveConfig"
      :loading="loading"
      :loading-config="loadingConfig"
      :data="dataList"
      :columns="datacolumns"
      :pagination="pagination"
      @page-current-change="onCurrentChange"
    >
    </pure-table>
    <PlusDialogForm
      v-model:visible.sync="visible"
      v-model="values"
      :form="{ columns, hasFooter: true }"
      :dialog="{ hasFooter: false }"
    >
      <template #form-footer="{ handleSubmit }">
        <el-button type="primary" @click="handleAllConfirm(handleSubmit)"
          >提交</el-button
        >
        <el-button type="danger" @click="handleCancel()"> 返回</el-button>
      </template>
    </PlusDialogForm>
  </div>
</template>
