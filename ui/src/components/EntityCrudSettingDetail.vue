<template>
  <div class="flex-col">
    <div class="pr-5 pt-5 flex justify-end">
      <a-switch
        v-model:checked="checked"
        checked-children="详情功能开启"
        un-checked-children="详情功能关闭"
        @change="onEnd"
      />
    </div>

    <div class="pt-5">
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
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from 'vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { get, save } from '@/api/systemsettings'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import draggable from 'vuedraggable'

const checked = ref(true)
const fields = ref([])
const config = ref({})

const store = useSystemSettingsStore()
const router = useRouter()

const getConfig = async () => {
  const key = `${router.currentRoute.value.params.name}_crud_detail`
  const resp = await get({
    key: key
  })
  if (resp.code != 0) {
    return
  }
  config.value = resp.data
}

const initData = () => {
  const schema = store.schemaMap[router.currentRoute.value.params.name]
  if (config.value.length != 0 && config.value.fields.length == schema.fields.length) {
    fields.value = config.value.fields
    checked.value = config.value.checked
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

  initData()
})

const drag = ref(false)

const onStart = () => {
  drag.value = true
}

const onEnd = async () => {
  drag.value = false
  const resp = await save({
    uri: `${router.currentRoute.value.params.name}_crud_detail`,
    data: {
      fields: fields.value,
      checked: checked.value
    },
    type: 'crud_detail'
  })

  if (resp.code != 0) {
    return message.error(resp.msg)
  }
  message.success('保存成功')
}
</script>
