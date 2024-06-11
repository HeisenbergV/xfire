<script setup lang="ts">
import { ref, computed } from "vue";
import { useColumns, form } from "./columns";

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
  onSizeChange,
  onCurrentChange
} = useColumns();

const { columns, values, visible, handleAllConfirm, handleCancel } = form();
</script>

<template>
  <pure-table
    ref="tableRef"
    border
    adaptive
    :adaptiveConfig="adaptiveConfig"
    row-key="id"
    showOverflowTooltip
    stripe
    :loading="loading"
    :loading-config="loadingConfig"
    :data="dataList"
    :columns="datacolumns"
    :pagination="pagination"
    @page-size-change="onSizeChange"
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
</template>
