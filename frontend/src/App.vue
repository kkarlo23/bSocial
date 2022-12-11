<script setup>
import { RouterView } from "vue-router";
import HeaderComponent from "./components/Header.vue";
</script>

<template>
  <div class="main-container">
    <HeaderComponent title="bSocial" />
    <div class="main-content">
      <RouterView />
    </div>
  </div>
</template>

<style scoped>
.main-container {
  min-height: 100vw;
  min-width: 100vw;
  background-color: rgb(235, 235, 235);
}
.main-content {
  padding: 20px;
}
</style>
