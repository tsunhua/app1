<script setup>
import Card from '../components/Card.vue'
import { onMounted, ref } from 'vue'
import { querySites, openApp } from '../api'

const groups = ref({});

onMounted(async () => {
  groups.value = await querySites()
})

const handleClickEvent = (path) => {
  openApp(path)
}

const colors = ['#3F51B5', '#FF9800', '#4CAF50', '#FF5722', '#03A9F4']

</script>

<template>
  <h1 class="font-weight-bold mt-6 mb-4 text-basil">
    App1
  </h1>
  <v-expansion-panels multiple>
    <v-expansion-panel v-for="x in groups" :key="x" :value="x">
      <v-expansion-panel-title>
        <v-icon :icon="x.icon" size="small" color="#BDBDBD" class="me-1" /> {{ x.name.toUpperCase() }}
      </v-expansion-panel-title>
      <v-expansion-panel-text>
        <v-row justify="start" no-gutters>
          <v-col cols="2" no-gutters class="ml-3 mt-3 mb-3" v-for="(y, index) in x.sites">
            <Card :bg-color="colors[Math.floor(index % colors.length)]" action="进入" :title="y['name']"
              :sub-title="y['desc']" :icon="y['icon']" @click="handleClickEvent(y['path'])" />
          </v-col>
        </v-row>
      </v-expansion-panel-text>
    </v-expansion-panel>
  </v-expansion-panels>
</template>

<style scoped>
.text-basil {
  color: #356859 !important;
}
</style>
