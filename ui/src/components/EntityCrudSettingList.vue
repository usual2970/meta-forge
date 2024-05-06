<template>
  <draggable
    v-model="fields"
    group="people"
    @start="onStart"
    @end="onEnd"
    item-key="name"
    class="border-t border-r border-l mb-3 mr-3"
  >
    <template #item="{ element }">
      <div class="border-b p-2 flex px-3 cursor-move">
        <HolderOutlined class="font-bold text-lg" />
        <div class="text-base text-slate-700 ml-2">{{ element.name }}</div>
        <div class="grow"></div>
        <a-switch v-model:checked="element.visible" class="w-12" @change="onEnd" />
      </div>
    </template>
  </draggable>
</template>

<script setup>
import draggable from 'vuedraggable'
import { HolderOutlined } from '@ant-design/icons-vue'

import { onMounted, ref } from 'vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { get, save } from '@/api/systemsettings'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'

const store = useSystemSettingsStore()

const router = useRouter()

const fields = ref([])

const config = ref([])

const getConfig = async () => {
  const key = `${router.currentRoute.value.params.name}_crud_list`
  const resp = await get({
    key: key
  })
  if (resp.code != 0) {
    return
  }
  config.value = resp.data
}

const initFields = () => {
  const schema = store.schemaMap[router.currentRoute.value.params.name]
  if (config.value.length != 0 && config.value.length == schema.fields.length) {
    fields.value = config.value
    return
  }

  fields.value = schema.fields.map((item) => {
    return {
      name: item.name,
      visible: true
    }
  })
}

onMounted(async () => {
  await getConfig()

  initFields()

  // 如果没有配置，则使用默认配置
})

const drag = ref(false)

const onStart = () => {
  drag.value = true
}

const onEnd = async () => {
  drag.value = false
  const resp = await save({
    uri: `${router.currentRoute.value.params.name}_crud_list`,
    data: fields.value,
    type: 'crud_list'
  })

  if (resp.code != 0) {
    return message.error(resp.msg)
  }
  message.success('保存成功')
}
</script>
