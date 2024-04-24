<template>
  <div class="flex justify-between">
    <a-breadcrumb class="">
      <a-breadcrumb-item>工作表</a-breadcrumb-item>
      <a-breadcrumb-item>{{ labelName }}</a-breadcrumb-item>
    </a-breadcrumb>

    <div class="flex">
      <a-button type="primary" @click="handleAdd" :icon="h(PlusOutlined)">新增</a-button>
      <a-button type="default" @click="handleSave" class="ml-5" :icon="h(SettingOutlined)"
        >设置</a-button
      >
    </div>
  </div>

  <a-table
    :dataSource="dataSource"
    :columns="columns"
    class="my-5"
    size="middle"
    :scroll="{ x: 1500, y: 800 }"
    :pagination="pagination"
    @change="handleTableChange"
  />
</template>

<script setup>
import { computed, ref, reactive, onMounted, h } from 'vue'
import { useRouter, onBeforeRouteUpdate } from 'vue-router'
import { name2label } from '@/utils/helper'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { list } from '@/api/data'
import { PlusOutlined, SettingOutlined } from '@ant-design/icons-vue'

const store = useSystemSettingsStore()

const router = useRouter()

const labelName = computed(() => {
  return name2label(router.currentRoute.value.params.name)
})

const dataSource = ref([])

const getList = async (table) => {
  const req = {
    table: table,
    page: pagination.current,
    pageSize: pagination.pageSize
  }
  let sorter = sortedInfo.value
  if (sorter && sorter.order) {
    let order = sorter.order === 'ascend' ? 'asc' : 'desc'

    req.orderBy = `${sorter.field} ${order}`
  }

  const resp = await list(req)

  if (resp.code === 0) {
    dataSource.value = resp.data.data
    pagination.total = resp.data.totalRecords
  }
}

onMounted(async () => {
  console.log(router)
  await getList(router.currentRoute.value.params.name)
})

onBeforeRouteUpdate((to, from) => {
  // 当路由参数变化时重新获取数据
  if (to.params.name !== from.params.name) {
    sortedInfo.value = null
    getList(to.params.name)
  }
})

const sortedInfo = ref()

const columns = computed(() => {
  const sorted = sortedInfo.value || {}
  return store.getSchemaColumns(router.currentRoute.value.params.name).map((column) => {
    if (column.key === sorted.columnKey && sorted.order) {
      column.sortOrder = sorted.order
    } else {
      column.sortOrder = null
    }
    return column
  })
})

const pagination = reactive({
  total: 0,
  current: 1,
  pageSize: 20
})

const handleTableChange = (page, filters, sorter) => {
  pagination.current = page.current
  pagination.pageSize = page.pageSize
  sortedInfo.value = sorter
  getList(router.currentRoute.value.params.name)
}

const handleAdd = () => {}
const handleSave = () => {}
</script>
