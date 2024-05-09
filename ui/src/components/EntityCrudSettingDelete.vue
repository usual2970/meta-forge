<template>
  <div class="flex-col">
    <div class="pr-5 pt-5 flex justify-end">
      <a-switch
        v-model:checked="checked"
        checked-children="详情功能开启"
        un-checked-children="详情功能关闭"
        @change="onChange"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { save, get } from '@/api/systemsettings'

const router = useRouter()
const checked = ref(true)

onMounted(async () => {
  const resp = await get({
    key: `${router.currentRoute.value.params.name}_crud_delete`
  })
  if (resp.code == 0) {
    checked.value = resp.data.checked
  }
})

const onChange = async () => {
  const resp = await save({
    uri: `${router.currentRoute.value.params.name}_crud_delete`,
    data: {
      checked: checked.value
    },
    type: 'crud_delete'
  })

  if (resp.code != 0) {
    return message.error(resp.msg)
  }
  message.success('保存成功')
}
</script>
