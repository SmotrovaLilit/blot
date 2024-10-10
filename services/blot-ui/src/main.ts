import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import { createVuetify } from 'vuetify';
import 'vuetify/styles';
import { aliases, mdi } from 'vuetify/iconsets/mdi';
import '@mdi/font/css/materialdesignicons.css';
import '@fortawesome/fontawesome-free/css/all.css';
import DraggableResizableVue from 'draggable-resizable-vue3'

import App from './App.vue'
import router from './router'

const app = createApp(App)

app.use(createPinia())
app.use(router)

const vuetify = createVuetify({
    theme: {
        defaultTheme: 'light',
        themes: {
            light: {
                colors: {
                    primary: '#1976D2',
                    secondary: '#424242',
                    accent: '#82B1FF',
                    error: '#FF5252',
                    info: '#2196F3',
                    success: '#4CAF50',
                    warning: '#FFC107',
                    // cardBackground: '#ccd3d9',
                },
            },
            dark: {
                colors: {
                    primary: '#1E88E5',
                    secondary: '#FFCDD2',
                    accent: '#8C9EFF',
                    error: '#FF5252',
                    info: '#2196F3',
                    success: '#4CAF50',
                    warning: '#FFC107',
                    // cardBackground: '#494b4d',
                },
            }
        },
    },
    icons: {
        defaultSet: 'mdi',
        aliases,
        sets: { mdi },
    },
});

app.use(vuetify)
app.use(DraggableResizableVue)

app.mount('#app')

