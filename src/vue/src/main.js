import { createApp } from 'vue'
import App from './App.vue'
import ElementPlus from 'element-plus';
import 'element-plus/dist/index.css';
import router from './router/index';
import instance from './axios/index';

const app = createApp(App)
app.use(ElementPlus)
app.use(router)

app.mount('#app')

app.config.globalProperties.$router = router
app.config.globalProperties.$axios = instance