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
    <a-menu
      :items="menuItems"
      class="min-w-56 border rounded shadow-lg"
      v-model:selectedKeys="selectedKeys"
      @click="onMenuClick"
    />

    <div class="ml-5 grow border rounded flex flex-col">
      <div class="flex border-b justify-between items-center">
        <div class="text-xl text-slate-700 p-3">{{ labelName }}</div>
        <a-button :icon="h(QuestionCircleOutlined)" type="link" v-if="selectedKeys[0] == 'dict'"
          >解释字典设置</a-button
        >
      </div>
      <div class="flex-grow">
        <!-- 字典设置 -->
        <div class="pt-7" v-if="selectedKeys[0] === 'dict'">
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
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { computed, h, onMounted, ref } from 'vue'
import { name2label, deepCopy } from '@/utils/helper'
import { SettingOutlined, QuestionCircleOutlined, SaveOutlined } from '@ant-design/icons-vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { message } from 'ant-design-vue'

const store = useSystemSettingsStore()
const router = useRouter()

const labels = {
  crud: '增删改查',
  dict: '字典',
  field: '字段',
  relation: '关联'
}

const labelName = computed(() => {
  return labels[router.currentRoute.value.params.type] + '设置'
})

const parentName = computed(() => {
  return store.getLabel(router.currentRoute.value.params.name)
})

const backLink = computed(() => {
  return `/entity/${router.currentRoute.value.params.name}`
})

const selectedKeys = ref(['crud'])

onMounted(async () => {
  selectedKeys.value = [router.currentRoute.value.params.type]
  let dictKey = `${router.currentRoute.value.params.name}_dict`

  dictFormData.value = deepCopy(store.dict[dictKey])
})

const menuItems = [
  {
    label: '增删改查',
    key: 'crud',
    icon: () => h(SettingOutlined),
    path: `/entity/${router.currentRoute.value.params.name}/setting/basic`
  },
  { type: 'divider' },
  {
    label: '字典',
    key: 'dict',
    icon: () => h(SettingOutlined),
    path: `/entity/${router.currentRoute.value.params.name}/setting/dict`
  },
  { type: 'divider' },
  {
    label: '字段',
    key: 'field',
    icon: () => h(SettingOutlined),
    path: `/entity/${router.currentRoute.value.params.name}/setting/field`
  },
  { type: 'divider' },
  {
    label: '关联',
    key: 'relation',
    icon: () => h(SettingOutlined),
    path: `/entity/${router.currentRoute.value.params.name}/setting/relation`
  }
]

const dictFormRef = ref()

const dictFormData = ref({
  plural: '',
  singular: ''
})

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

const onMenuClick = (e) => {
  router.push({
    name: 'entity-setting',
    params: {
      name: router.currentRoute.value.params.name,
      type: e.key
    }
  })
}
</script>
