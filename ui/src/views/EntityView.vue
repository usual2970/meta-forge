<template>
  <div class="flex justify-between">
    <div class="text-2xl text-slate-700">{{ dict.plural }}</div>
    <div class="flex">
      <a-button type="primary" @click="handleAdd" :icon="h(PlusOutlined)">新增</a-button>
      <a-button type="default" @click="handleSetting" class="ml-5" :icon="h(SettingOutlined)"
        >设置</a-button
      >
    </div>
  </div>
  <div class="flex justify-between mt-2">
    <a-breadcrumb class="">
      <a-breadcrumb-item>{{ dict.plural }}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <a-affix :offset-top="0">
    <a-table
      :dataSource="dataSource"
      :columns="columns"
      class="my-5"
      size="middle"
      :scroll="{ x: 1500, y: 800 }"
      :pagination="pagination"
      @change="handleTableChange"
    >
      <template #bodyCell="{ column, record }">
        <template v-if="column.key === 'action'">
          <span>
            <a @click="handleView(record)"><EyeOutlined /></a>
            <a-divider type="vertical" />
            <a><DeleteOutlined /></a>
          </span>
        </template>
      </template>
    </a-table>
  </a-affix>
</template>

<script setup>
import { computed, ref, reactive, onMounted, h } from 'vue'
import { useRouter, onBeforeRouteUpdate } from 'vue-router'

import { useSystemSettingsStore } from '@/stores/systemsettings'
import { list } from '@/api/data'
import {
  PlusOutlined,
  SettingOutlined,
  EyeOutlined,
  DeleteOutlined,
  FormOutlined
} from '@ant-design/icons-vue'
import { batchGet, get } from '@/api/systemsettings'
import { name2label } from '@/utils/helper'

const dict = ref({})

const initDict = async (name) => {
  const dictKey = `${name}_dict`
  const res = await batchGet({
    keys: dictKey
  })
  const dictData = res.data[dictKey]
  if (!dictData) {
    dict.value = {
      plural: name2label(name),
      singular: name2label(name)
    }
  } else {
    dict.value = dictData
  }
}

const store = useSystemSettingsStore()

const router = useRouter()

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
  await getList(router.currentRoute.value.params.name)
  await initDict(router.currentRoute.value.params.name)
  await initFieldLabels(router.currentRoute.value.params.name)
})

onBeforeRouteUpdate(async (to, from) => {
  // 当路由参数变化时重新获取数据
  if (to.params.name !== from.params.name) {
    sortedInfo.value = null
    await getList(to.params.name)
    await initDict(to.params.name)
    await initFieldLabels(to.params.name)
  }
})

const sortedInfo = ref()

const fieldLabels = ref({})

const initFieldLabels = async (name) => {
  const res = await get({
    key: `${name}_field_label`
  })
  if (res.code == 0) {
    fieldLabels.value = res.data
  }
}

const columns = computed(() => {
  const sorted = sortedInfo.value || {}
  const rs = store.getSchemaColumns(router.currentRoute.value.params.name).map((column) => {
    if (column.key === sorted.columnKey && sorted.order) {
      column.sortOrder = sorted.order
    } else {
      column.sortOrder = null
    }
    if (fieldLabels.value[column.key]) {
      column.title = fieldLabels.value[column.key]
    }
    return column
  })

  rs.push({
    title: '操作',
    dataIndex: 'action',
    key: 'action',
    fixed: 'right',
    width: 70,
    scopedSlots: { customRender: 'operation' }
  })

  return rs
})

const idField = computed(() => {
  return store.getSchemaIdField(router.currentRoute.value.params.name)
})

const handleView = (record) => {
  router.push({
    name: 'entity-detail',
    params: {
      name: router.currentRoute.value.params.name,
      id: record[idField.value.name]
    }
  })
}

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
const handleSetting = () => {
  router.push({
    name: 'entity-setting',
    params: {
      name: router.currentRoute.value.params.name,
      type: 'crud'
    }
  })
}
</script>
