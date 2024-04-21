
import axios from 'axios';
import { message } from 'ant-design-vue';
// import { useAuthStore } from '@/stores/auth'; // 假设有个auth store存放token

const instance = axios.create({
  baseURL: import.meta.env.VUE_APP_API_BASE_URL,
  timeout: 10000,
});

instance.interceptors.request.use(
  (config) => {
    // const authStore = useAuthStore();
    // if (authStore.token) {
    //   config.headers.Authorization = `Bearer ${authStore.token}`;
    // }
    return config;
  },
  (error) => Promise.reject(error),
);

instance.interceptors.response.use(
  (response) => response.data,
  (error) => {
    if (error.response.status === 401) {
      // 处理未授权
    } else {
        message.error(error.message || '未知错误');
    }
    return Promise.reject(error);
  },
);

export default instance;