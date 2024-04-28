<template>
  <a-form
    :model="dictFormData"
    :label-col="{ span: 3 }"
    :wrapper-col="{ span: 18 }"
    :rules="dictRules"
    ref="dictFormRef"
  >
    <a-form-item label="复数" name="plural">
      <a-input v-model:value="dictFormData.plural" type="text" />
    </a-form-item>
    <a-form-item label="单数" name="singular">
      <a-input v-model:value="dictFormData.singular" type="text" />
    </a-form-item>
    <a-form-item :wrapper-col="{ offset: 19 }">
      <a-button
        type="primary"
        html-type="submit"
        size="large"
        class="bg-blue-500"
        @click="onDictSubmit"
        :icon="h(SaveOutlined)"
        >保存</a-button
      >
    </a-form-item>
  </a-form>
</template>
<script setup>
import { reactive, h, ref } from 'vue'
import { SaveOutlined } from '@ant-design/icons-vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { useRouter } from 'vue-router'
import { message } from 'ant-design-vue'
import { deepCopy } from '@/utils/helper'
const props = defineProps({
  formData: {
    type: Object,
    default: () => ({})
  }
})

const store = useSystemSettingsStore()
const router = useRouter()
const dictFormData = ref({ ...props.formData })

const dictFormRef = ref()

const dictRules = {
  plural: [{ required: true, message: '请输入复数名称', trigger: ['change', 'blur'] }],
  singular: [{ required: true, message: '请输入单数名称', trigger: ['change', 'blur'] }]
}

const onDictSubmit = async () => {
  try {
    await dictFormRef.value.validate()

    await store.saveDict({
      uri: `${router.currentRoute.value.params.name}_dict`,
      data: deepCopy(dictFormData.value),
      type: 'dict'
    })

    message.success('保存成功')
  } catch (err) {
    console.log(JSON.stringify(err))
  }
}
</script>
