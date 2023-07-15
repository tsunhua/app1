<script setup>
import { ref } from 'vue'
import { login } from '../api'
import { useRouter } from 'vue-router'

const router = useRouter()

const name = ref('')
const password = ref('')

const nameRules = [
  value => {
    if (value?.length >= 3 && value?.length <= 20) return true
    return '用户名限定在 3 到 20 个字符之间'
  },
]

const passwordRules = [
  value => {
    if (value?.length >= 6 && value?.length <= 20) return true
    return '密码限定在 6 到 20 个字符之间'
  },
]

const onLoginClick = async () => {
  const ok = await login(name.value, password.value)
  if (ok) {
    router.replace({ name: 'home' })
  } else {
    //TODO
  }
}

</script>

<template>
  <h1 class="font-weight-bold mt-6 mb-8 text-basil mx-auto">
    Login
  </h1>
  <v-sheet width="360" class="mx-auto pa-4" rounded>
    <v-form fast-fail @submit.prevent>
      <v-text-field v-model="name" label="用户名" :rules="nameRules"></v-text-field>
      <v-text-field v-model="password" label="密码" :rules="passwordRules"></v-text-field>
      <v-btn prepend-icon="mdi-login" color="green" type="submit" size="large" block class="mt-4 pa-4"
        :onclick="onLoginClick">登录</v-btn>
    </v-form>
  </v-sheet>
</template>

<style scoped>
.text-basil {
  color: #356859 !important;
}
</style>
