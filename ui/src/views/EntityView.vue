<template>
  <a-breadcrumb class="">
    <a-breadcrumb-item>工作表</a-breadcrumb-item>
    <a-breadcrumb-item>{{ labelName }}</a-breadcrumb-item>
  </a-breadcrumb>

  <a-table :dataSource="dataSource" :columns="columns" class="my-5" />
</template>

<script setup>
import { computed, ref } from 'vue'
import { useRouter } from 'vue-router'
import { name2label } from '@/utils/helper'

const router = useRouter()

const labelName = computed(() => {
  return name2label(router.currentRoute.value.params.name)
})

const dataSource = ref([
  {
    key: '1',
    name: 'Mike',
    age: 32,
    address: '10 Downing Street'
  },
  {
    key: '2',
    name: 'John',
    age: 42,
    address: '10 Downing Street'
  }
])

const columns = ref([
  {
    title: 'Name',
    dataIndex: 'name',
    key: 'name'
  },
  {
    title: 'Age',
    dataIndex: 'age',
    key: 'age'
  },
  {
    title: 'Address',
    dataIndex: 'address',
    key: 'address'
  }
])
</script>
