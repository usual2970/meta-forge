<template>
  <div class="pt-3 pl-5">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="table" tab="列表">
        <EntityCrudSettingList />
      </a-tab-pane>
      <a-tab-pane key="detail" tab="详情" force-render>
        <EntityCrudSettingDetail />
      </a-tab-pane>
      <a-tab-pane key="delete" tab="删除" force-render>Content of Tab Pane 3</a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import EntityCrudSettingList from './EntityCrudSettingList.vue'
import EntityCrudSettingDetail from './EntityCrudSettingDetail.vue'

const router = useRouter()

const activeKey = ref()

onMounted(() => {
  if (!router.currentRoute.value.query.tab) {
    activeKey.value = 'table'
  }
  activeKey.value = router.currentRoute.value.query.tab
})
</script>
