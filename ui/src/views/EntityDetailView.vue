<template>
  <div class="flex justify-between">
    <div class="text-2xl text-slate-700">{{ router.currentRoute.value.params.id }}</div>
    <div class="flex"></div>
  </div>
  <div class="flex justify-between mt-2">
    <a-breadcrumb class="">
      <a-breadcrumb-item>
        <router-link :to="backLink">{{ parentName }}</router-link>
      </a-breadcrumb-item>
      <a-breadcrumb-item>{{ router.currentRoute.value.params.id }}</a-breadcrumb-item>
    </a-breadcrumb>
  </div>
  <a-divider></a-divider>
  <div class="w-full">
    <a-form>
      <a-row>
        <a-col :span="12" v-for="(item, index) in formData" :key="index">
          <a-form-item
            :label="fieldLabels[index] ? fieldLabels[index] : index"
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 19 }"
          >
            <a-input v-model:value="formData[index]" />
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
  </div>
</template>

<script setup>
import { computed, onMounted, ref } from 'vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'
import { useRouter } from 'vue-router'
import { detail } from '@/api/data'
import { get } from '@/api/systemsettings'

const router = useRouter()

const store = useSystemSettingsStore()

const parentName = computed(() => {
  return store.getLabel(router.currentRoute.value.params.name)
})

const backLink = computed(() => {
  return `/entity/${router.currentRoute.value.params.name}`
})
const formData = ref({})

const initFormData = async () => {
  const idField = store.getSchemaIdField(router.currentRoute.value.params.name)
  const res = await detail({
    table: router.currentRoute.value.params.name,
    filter: `${idField.name}=${router.currentRoute.value.params.id}`
  })
  formData.value = res.data
}

const fieldLabels = ref({})

const initFieldLabels = async (name) => {
  const res = await get({
    key: `${name}_field_label`
  })
  if (res.code == 0) {
    fieldLabels.value = res.data
  }
}

onMounted(async () => {
  await initFormData()

  await initFieldLabels(router.currentRoute.value.params.name)
})
</script>
