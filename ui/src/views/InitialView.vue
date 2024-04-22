<template>
  <div class="min-w-[500px] border rounded-md bg-blue-500 font-mono">
    <div class="p-7">
      <div class="h-11 text-4xl text-center text-white">MetaForge</div>
      <div class="text-xl pt-1 text-center text-white">数据库配置</div>
      <div class="text-sm text-center text-white pt-1">设置数据库连接</div>
    </div>
    <a-form
      class="p-7 bg-white"
      :labelCol="{ span: 5 }"
      ref="formRef"
      :rules="rules"
      :model="formState"
    >
      <a-form-item label="数据库类型" has-feedback name="kind">
        <a-select v-model:value="formState.kind" placeholder="数据库类型">
          <a-select-option value="">--选择数据库类型--</a-select-option>
          <a-select-option value="mysql">Mysql</a-select-option>
          <a-select-option value="sqlite">Sqlite</a-select-option>
        </a-select>
      </a-form-item>
      <template v-if="showFile">
        <a-form-item label="数据库文件" :labelCol="{ span: 5 }" has-feedback name="file">
          <a-input v-model:value="formState.file"></a-input>
        </a-form-item>
      </template>
      <template v-if="showUrl">
        <a-form-item label="主机(host)" :labelCol="{ span: 5 }" has-feedback name="host">
          <a-input v-model:value="formState.host"></a-input>
        </a-form-item>
        <a-form-item label="端口(port)" :labelCol="{ span: 5 }" has-feedback name="port">
          <a-input v-model:value="formState.port"></a-input>
        </a-form-item>
        <a-form-item label="用户" :labelCol="{ span: 5 }" has-feedback name="user">
          <a-input v-model:value="formState.user"></a-input>
        </a-form-item>

        <a-form-item label="密码" :labelCol="{ span: 5 }" has-feedback name="password">
          <a-input-password v-model:value="formState.password"></a-input-password>
        </a-form-item>

        <a-form-item label="数据库" :labelCol="{ span: 5 }" has-feedback name="database">
          <a-input v-model:value="formState.database"></a-input>
        </a-form-item>
      </template>

      <a-form-item :wrapper-col="{ offset: 19 }" v-if="showSubmit">
        <a-button
          type="primary"
          html-type="submit"
          size="large"
          class="bg-blue-500"
          @click="onSubmit"
          >提交</a-button
        >
      </a-form-item>
    </a-form>
  </div>
</template>

<script setup>
import { reactive, computed, ref } from 'vue'
import { initialize } from '@/api/systemsettings'
import { message } from 'ant-design-vue'
import { useRouter } from 'vue-router'

const formRef = ref()
const formState = reactive({
  kind: '',
  file: '',
  host: '',
  port: '',
  user: '',
  password: '',
  database: ''
})

const showFile = computed(() => {
  return formState.kind === 'sqlite'
})

const showUrl = computed(() => {
  return ['mysql', 'postgresql'].includes(formState.kind)
})

const showSubmit = computed(() => {
  return showFile.value || showUrl.value
})

const rules = {
  kind: [
    {
      message: 'Please select a kind',
      required: true,
      trigger: 'change'
    }
  ],
  file: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'sqlite' && value == '') {
          return Promise.reject('Please input a file')
        }
        return Promise.resolve()
      },
      trigger: 'change'
    }
  ],
  host: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'mysql' && value == '') {
          return Promise.reject('Please input a host')
        }
        return Promise.resolve()
      },

      trigger: 'change'
    }
  ],
  port: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'mysql' && value == '') {
          return Promise.reject('Please input a port')
        }
        return Promise.resolve()
      },
      trigger: 'change'
    }
  ],
  user: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'mysql' && value == '') {
          return Promise.reject('Please input a user')
        }
        return Promise.resolve()
      },
      trigger: 'change'
    }
  ],
  password: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'mysql' && value == '') {
          return Promise.reject('Please input a password')
        }
        return Promise.resolve()
      },
      trigger: 'change'
    }
  ],
  database: [
    {
      validator: async (rule, value) => {
        if (formState.kind == 'mysql' && value == '') {
          return Promise.reject('Please input a database')
        }
        return Promise.resolve()
      },
      trigger: 'change'
    }
  ]
}

const router = useRouter()

const onSubmit = async () => {
  try {
    await formRef.value.validate()

    const resp = await initialize(formState)
    if (resp.code == 0) {
      message.success('Initialize success', 2)
      return router.push('/')
    } else {
      return message.error(resp.msg)
    }
  } catch (err) {
    message.error(err.message)
    console.log(JSON.stringify(err))
  }
}
</script>
