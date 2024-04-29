<template>
  <div class="pt-3 pl-5">
    <a-tabs v-model:activeKey="activeKey">
      <a-tab-pane key="labels" tab="标签">
        <EntityFieldLabelSetting />
      </a-tab-pane>
      <a-tab-pane key="form" tab="表单" force-render>Content of Tab Pane 2</a-tab-pane>
    </a-tabs>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import EntityFieldLabelSetting from './EntityFieldLabelSetting.vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const activeKey = ref()

onMounted(() => {
  if (!router.currentRoute.value.query.tab) {
    activeKey.value = 'labels'
  }
  activeKey.value = router.currentRoute.value.query.tab
})
</script>
