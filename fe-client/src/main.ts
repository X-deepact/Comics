import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from './router'

// I18n
import { createI18n } from 'vue-i18n'
import en from './locales/en.json'
import vi from './locales/vi.json'
import cn from './locales/cn.json'

// PrimeVue
import PrimeVue from 'primevue/config';
import Aura from '@primevue/themes/aura';
import ToastService from 'primevue/toastservice';
import 'primeicons/primeicons.css'
import Tooltip from 'primevue/tooltip';

const messages = {
  en: en,
  vi: vi,
  zh: cn
};

export const i18n = createI18n({
	legacy: false,
	globalInjection: true,
  locale: import.meta.env.VITE_DEFAULT_LOCALE || "en",
  fallbackLocale: import.meta.env.VITE_FALLBACK_LOCALE || "vi",
  messages,
});

const app = createApp(App);

app.directive('tooltip', Tooltip);

app
.use(createPinia())
.use(i18n)
.use(ToastService)
.use(PrimeVue, {
  theme: {
      preset: Aura,
      options: {
          cssLayer: {
              name: 'primevue',
              order: 'tailwind-base, primevue, tailwind-utilities'
          }
      }
  }
});
app.use(router);

app.mount('#app')
