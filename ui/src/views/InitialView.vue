<template>
  <div class="min-w-[500px] border rounded-md bg-indigo-500 font-mono">
    <div class="p-7">
      <div class="h-11 text-4xl text-center text-white">MetaForge</div>
      <div class="text-xl pt-1 text-center text-white">数据库配置</div>
      <div class="text-sm text-center text-white pt-1">设置数据库连接</div>
    </div>
    <a-form class="p-7 bg-white" :labelCol="{ span: 5 }">
      <a-form-item label="数据库类型">
        <a-select v-model:value="formState.kind" placeholder="数据库类型">
          <a-select-option value="">--选择数据库类型--</a-select-option>
          <a-select-option value="mysql">Mysql</a-select-option>
          <a-select-option value="sqlite">Sqlite</a-select-option>
        </a-select>
      </a-form-item>

      <a-form-item label="主机(host)" :labelCol="{ span: 5 }">
        <a-input></a-input>
      </a-form-item>
      <a-form-item label="端口(port)" :labelCol="{ span: 5 }">
        <a-input></a-input>
      </a-form-item>
      <a-form-item label="用户" :labelCol="{ span: 5 }">
        <a-input></a-input>
      </a-form-item>

      <a-form-item label="密码" :labelCol="{ span: 5 }">
        <a-input></a-input>
      </a-form-item>

      <a-form-item label="数据库" :labelCol="{ span: 5 }">
        <a-input></a-input>
      </a-form-item>

      <a-form-item :wrapper-col="{ offset: 19 }">
        <a-button type="primary" html-type="submit" size="large" class="bg-indigo-500"
          >提交</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { reactive } from 'vue'

const formState = reactive({
  kind: ''
})
</script>
