<template>
  <a-breadcrumb class="">
    <a-breadcrumb-item>工作表</a-breadcrumb-item>
    <a-breadcrumb-item>{{ labelName }}</a-breadcrumb-item>
  </a-breadcrumb>

  <a-table :dataSource="dataSource" :columns="columns" class="my-5" />
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter, onBeforeRouteUpdate } from 'vue-router'
import { name2label } from '@/utils/helper'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { list } from '@/api/data'

const store = useSystemSettingsStore()

const router = useRouter()

const labelName = computed(() => {
  return name2label(router.currentRoute.value.params.name)
})

const dataSource = ref([])

const getList = async (table) => {
  const resp = await list({
    table: table,
    page: 1,
    pageSize: 1000
  })

  if (resp.code === 0) {
    dataSource.value = resp.data.data
  }
}

onBeforeRouteUpdate((to, from) => {
  // 当路由参数变化时重新获取数据
  if (to.params.name !== from.params.name) {
    getList(to.params.name)
  }
})

getList(router.currentRoute.value.params.name)

const columns = computed(() => {
  return store.getSchemaColumns(router.currentRoute.value.params.name)
})
</script>
