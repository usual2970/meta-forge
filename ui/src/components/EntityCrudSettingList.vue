<template>
  <draggable
    v-model="myArray"
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
        <a-switch v-model:checked="checked" class="w-12" />
      </div>
    </template>
  </draggable>
</template>

<script setup>
import draggable from 'vuedraggable'
import { HolderOutlined } from '@ant-design/icons-vue'

import { onMounted, ref } from 'vue'
import { useSystemSettingsStore } from '@/stores/systemsettings'

const store = useSystemSettingsStore()

const fields = ref([])

onMounted(() => {
  
})

const drag = ref(false)

const onStart = () => {
  drag.value = true
}

const onEnd = () => {
  drag.value = false
  console.log(myArray.value)
}

const myArray = ref([{ name: 'John' }, { name: 'Joao' }, { name: 'Jean' }])
</script>
