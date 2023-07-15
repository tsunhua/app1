import './assets/main.css'

import { createApp } from 'vue'
import App from './App.vue'
// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import "@mdi/font/css/materialdesignicons.css";
import { aliases, mdi } from "vuetify/lib/iconsets/mdi";

// Router
// import VueRouter from 'vue-router'
import { createRouter, createWebHashHistory, createWebHistory } from 'vue-router';
import Home from './pages/Home.vue'
import Login from './pages/Login.vue'

const app = createApp(App)

const global = {
	'BASE_URL': 'http://127.0.0.1:8888',
}
for (const key in global) {
	if (global.hasOwnProperty(key)) {
		const value = global[key];
		app.provide(key, value)
	}
}

const vuetify = createVuetify({
	theme: {
		defaultTheme: "dark",
	},
	icons: {
		defaultSet: "mdi",
		aliases,
		sets: {
			mdi,
		},
	},
	components,
	directives,
})
app.use(vuetify)

const router = createRouter({
	history: createWebHashHistory(),
	routes: [
		{ name: 'home', path: '/', component: Home },
		{ name: 'login', path: '/login', component: Login },
	],
})
app.use(router)

app.mount('#app')
