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
          <EntityDictSetting />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router'
import { computed, h, onMounted, ref } from 'vue'
import { SettingOutlined, QuestionCircleOutlined } from '@ant-design/icons-vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import EntityDictSetting from '@/components/EntityDictSetting.vue'

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
