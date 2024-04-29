<template>
  <a-form
    :model="formData"
    :label-col="{ span: 3 }"
    :wrapper-col="{ span: 18 }"
    ref="formRef"
    :rules="rules"
  >
    <a-form-item v-for="(item, index) in formData" :key="index" :label="index" :name="index">
      <a-input v-model:value="formData[index]" type="text" />
    </a-form-item>
    <a-form-item :wrapper-col="{ offset: 19 }">
      <a-button
        type="primary"
        html-type="submit"
        size="large"
        class="bg-blue-500"
        @click="onSubmit"
        :icon="h(SaveOutlined)"
        >保存</a-button
      >
    </a-form-item>
  </a-form>
</template>

<script setup>
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { onMounted, ref, h } from 'vue'
import { get } from '@/api/systemsettings'
import { useRouter } from 'vue-router'
import { SaveOutlined } from '@ant-design/icons-vue'
import { deepCopy } from '@/utils/helper'
import { message } from 'ant-design-vue'
import { save } from '@/api/systemsettings'

const store = useSystemSettingsStore()
const router = useRouter()

// 获取字段标签设置
// 获取字段配置
// 生成表单配置
const fileLabelConfig = ref({})
const formData = ref({})
const formRef = ref()
const rules = ref({})

const getFieldLabelConfig = async () => {
  const key = `${router.currentRoute.value.params.name}_field_label`
  const resp = await get({
    key: key
  })
  if (resp.code != 0) {
    return
  }
  fileLabelConfig.value = resp.data
}

const initFormData = () => {
  const schema = store.schemaMap[router.currentRoute.value.params.name]

  formData.value = schema.fields.reduce((pre, item) => {
    pre[item.name] = ''

    if (fileLabelConfig.value[item.name]) {
      pre[item.name] = fileLabelConfig.value[item.name]
    }
    return pre
  }, {})
}

const initRules = () => {
  const schema = store.schemaMap[router.currentRoute.value.params.name]

  rules.value = schema.fields.reduce((pre, item) => {
    pre[item.name] = [
      {
        required: false,
        message: '请输入标签',
        trigger: ['change', 'blur']
      }
    ]

    return pre
  }, {})
}

onMounted(async () => {
  await getFieldLabelConfig()
  initFormData()
  initRules()
})

const onSubmit = async () => {
  try {
    await formRef.value.validate()

    await save({
      uri: `${router.currentRoute.value.params.name}_field_label`,
      data: deepCopy(formData.value),
      type: 'field_label'
    })

    message.success('保存成功')
  } catch (err) {
    console.log(JSON.stringify(err))
  }
}
</script>
