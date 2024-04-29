<template>
  <div>Entity Field Label Setting</div>
</template>

<script setup>
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { onMounted, ref } from 'vue'
import { get } from '@/api/systemsettings'
import { useRouter } from 'vue-router';

const store = useSystemSettingsStore()
const router=useRouter()

// 获取字段标签设置
// 获取字段配置
// 生成表单配置
const fileLabelConfig = ref({})

onMounted(async () => {
    const key=`${router}`
  const resp = await get({
    key: 'entity_field_label'
  })
})
</script>
