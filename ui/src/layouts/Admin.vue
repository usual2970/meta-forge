<template>
  <a-layout id="mf-layout" class="font-mono">
    <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible>
      <div class="p-4 text-center text-white font-extrabold text-xl">{{ logo }}</div>
      <a-menu
        v-model:selectedKeys="selectedKeys"
        theme="dark"
        mode="inline"
        :items="menuItems"
        @click="onMenuClick"
      >
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <menu-unfold-outlined
          v-if="collapsed"
          class="trigger"
          @click="() => (collapsed = !collapsed)"
        />
        <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
      </a-layout-header>
      <a-layout-content
        :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
      >
        <router-view v-slot="{ Component }">
          <keep-alive>
            <component :is="Component" />
          </keep-alive>
        </router-view>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script setup>
import { computed, ref } from 'vue'
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons-vue'
import { useRouter } from 'vue-router'

import { useSystemSettingsStore } from '@/stores/systemsettings'
const selectedKeys = ref(['1'])
const collapsed = ref(false)

const logo = computed(() => {
  return collapsed.value ? 'MF' : 'MetaForge'
})

const store = useSystemSettingsStore()

const menuItems = store.menuItems

const router = useRouter()

const onMenuClick = (e) => {
  if (e.keyPath.length == 1) {
    router.push({ name: e.key })
  }
  if (e.keyPath.length == 2 && e.keyPath[0] == 'entity') {
    router.push('/' + e.keyPath[0] + '/' + e.keyPath[1])
    return
  }
}
</script>
<style>
#mf-layout {
  min-height: 100vh;
}

#mf-layout .trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#mf-layout .trigger:hover {
  color: #1890ff;
}

#mf-layout .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.3);
  margin: 16px;
}

.site-layout .site-layout-background {
  background: #fff;
}
</style>
