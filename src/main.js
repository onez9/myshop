import 'vite/modulepreload-polyfill'
import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import '../public/assets/bootstrap/js/bootstrap.bundle.min'
import crypto from 'crypto'

import './assets/main.css'

const app = createApp(App)

app.use(router)

app.mount('#app')
