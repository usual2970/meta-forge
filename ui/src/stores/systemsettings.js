import { defineStore } from 'pinia'
import { computed, ref, h } from 'vue'
import { batchGet, getByType, save } from '@/api/systemsettings'
import { HomeOutlined, FileOutlined } from '@ant-design/icons-vue'

import { name2label } from '@/utils/helper'

export const useSystemSettingsStore = defineStore('systemsettings', () => {
  const settings = ref({
    '@hasInitialized': 0,
    schemas: {}
  })

  const dict = ref({})
  const fieldLabel = ref({})

  const hasInitialized = computed(() => {
    return settings.value['@hasInitialized']
  })
  const schemas = computed(() => {
    return settings.value.schemas
  })

  const schemaMap = computed(() => {
    return schemas.value.reduce((acc, cur) => {
      acc[cur.name] = cur
      return acc
    }, {})
  })

  const getSchemaColumns = (name) => {
    return schemaMap.value[name].fields.map((item, index) => {
      let rs = {
        title: name2label(item.name),
        key: item.name,
        dataIndex: item.name
      }
      if (index < 3) {
        rs.fixed = 'left'
      }

      if (item.type == 'number') {
        rs.width = 80
        rs.sorter = (a, b) => {
          return a[item.name] - b[item.name]
        }
      } else {
        rs.width = 200
      }

      return rs
    })
  }

  const getSchemaIdField = (name) => {
    return schemaMap.value[name].fields.find((item) => item.isId)
  }

  const menuItems = computed(() => {
    const entities = []

    for (const entity of settings.value.schemas) {
      let label = name2label(entity.name)
      let dictKey = `${entity.name}_dict`
      if (dict.value[dictKey]) {
        label = dict.value[dictKey]['plural']
      }

      entities.push({
        key: entity.name,
        label: label,
        title: entity.name
      })
    }

    return [
      {
        key: 'home',
        icon: () => h(HomeOutlined),
        label: '首页',
        title: '首页'
      },
      {
        key: 'entity',
        icon: () => h(FileOutlined),
        label: '工作表',
        title: '工作表',
        children: entities
      }
    ]
  })

  const getSettings = async () => {
    const resp = await batchGet({ keys: '@hasInitialized,schemas' })
    if (resp.code != 0) {
      settings.value['@hasInitialized'] = -1
    } else {
      if (!resp['data']['@hasInitialized']) {
        settings.value['@hasInitialized'] = -1
      } else {
        settings.value['@hasInitialized'] = 1
      }
      settings.value.schemas = !resp.data.schemas ? {} : resp.data.schemas
    }

    if (settings.value['@hasInitialized'] == 1) {
      const dictData = await getByType({ type: 'dict' })
      if (resp.code == 0) {
        dict.value = dictData.data
      }
    }

    return settings.value
  }

  const saveDict = async (data) => {
    const resp = await save(data)
    if (resp.code == 0) {
      dict.value[data.uri] = data.data
    }
  }

  const saveFieldLabel = async (data) => {
    const resp = await save(data)
    if (resp.code == 0) {
      fieldLabel.value[data.uri] = data.data
    }
  }

  const getLabel = (name) => {
    const labelKey = `${name}_dict`
    const data = dict.value[labelKey]
    if (data) {
      return data.plural
    }
    return name2label(name)
  }

  return {
    settings,
    getSettings,
    hasInitialized,
    schemas,
    menuItems,
    schemaMap,
    getSchemaColumns,
    saveDict,
    dict,
    getLabel,
    saveFieldLabel,
    getSchemaIdField
  }
})
