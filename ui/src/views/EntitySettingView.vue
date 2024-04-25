<template>
  <div class="flex justify-between">
    <div class="text-2xl text-slate-700">{{ labelName }}</div>
  </div>
  <div class="flex justify-between mt-2">
    <a-breadcrumb class="">
      <a-breadcrumb-item>
        <router-link :to="backLink">{{ parentName }}</router-link></a-breadcrumb-item
      >
      <a-breadcrumb-item>{{ labelName }}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>

  <div class="mt-5 flex">
    <div>
      <a-menu :items="menuItems" class="min-w-56 border rounded" />
    </div>

    <a-card :title="labelName" :bordered="false" class="ml-5 grow">
      <template #extra>
        <a href="#">More</a>
      </template>
      <p>card content</p>
    </a-card>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { computed, h } from 'vue'
import { name2label } from '@/utils/helper'
import { SettingOutlined } from '@ant-design/icons-vue'

const router = useRouter()

const labels = {
  crud: '增删改查',
  dict: '字典',
  field: '字段'
}

const labelName = computed(() => {
  return labels[router.currentRoute.value.params.type] + '设置'
})

const parentName = computed(() => {
  return name2label(router.currentRoute.value.params.name)
})

const backLink = computed(() => {
  return `/entity/${router.currentRoute.value.params.name}`
})

const menuItems = [
  {
    label: '增删改查',
    key: 'basic',
    icon: () => h(SettingOutlined),
    path: `/entity/${router.currentRoute.value.params.name}/setting/basic`
  },
  { type: 'divider' },
  {
    label: '字典',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '字段',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'field',
    icon: 'SettingOutlined',
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  }
]
</script>
