<script setup lang="ts">
import { ref } from "vue";
import { useColumns, form } from "./columns";

import "plus-pro-components/es/components/dialog-form/style/css";
import {
  type PlusColumn,
  type FieldValues,
  PlusDialogForm
} from "plus-pro-components";

const tableRef = ref();

const {
  loading,
  datacolumns,
  dataList,
  pagination,
  loadingConfig,
  adaptiveConfig,
  onSizeChange,
  onCurrentChange
} = useColumns();

const { columns, values, visible, handleAllConfirm, handleCancel } = form();
</script>

<template>
  <el-button> 新建产品 </el-button>
  <pure-table
    ref="tableRef"
    border
    adaptive
    :adaptiveConfig="adaptiveConfig"
    row-key="id"
    alignWhole="center"
    showOverflowTooltip
    stripe
    :loading="loading"
    :loading-config="loadingConfig"
    :data="
      dataList.slice(
        (pagination.currentPage - 1) * pagination.pageSize,
        pagination.currentPage * pagination.pageSize
      )
    "
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
    <template #form-footer="{ handleSubmit, handleReset }">
      <el-button type="primary" @click="handleAllConfirm(handleSubmit)"
        >提交</el-button
      >
      <el-button type="warning" @click="handleReset">重置</el-button>
      <el-button type="danger" @click="handleCancel()"> 返回</el-button>
    </template>
  </PlusDialogForm>
</template>
