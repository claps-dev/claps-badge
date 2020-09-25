/* eslint-disable */
import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import UIkit from '@foxone/uikit'

import zh from 'vuetify/es5/locale/zh-Hans'
import en from 'vuetify/es5/locale/en'
import ja from 'vuetify/es5/locale/ja'
import enUIkit from '@foxone/uikit/src/locales/en'
import jaUIkit from '@foxone/uikit/src/locales/ja'
import zhUIkit from '@foxone/uikit/src/locales/zh-Hans'


Vue.use(Vuetify)
Vue.use(UIkit)
export default function (store) {
  const isDark = (store && store.state?.app?.dark) || false

  return new Vuetify({
    icons: {},
    theme: {
      dark: isDark
    },
    lang: {
      locales: {
        zh: {
          ...zh,
          ...zhUIkit
        },
        en: {
          ...en,
          ...enUIkit
        },
        ja: {
          ...ja,
          ...jaUIkit
        },
      },
    },
  });
}
