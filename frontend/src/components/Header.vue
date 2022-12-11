<template>
  <div class="header">
    {{ title }}
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "HeaderComponent",
  props: {
    title: String,
  },
});
</script>

<style>
.header {
  width: 100%;
  height: 50px;
  background-color: rgb(30, 100, 110);
  display: flex;
  align-items: center;
  padding-left: 10px;
  font-size: larger;
  font-weight: bold;
  color: white;
}
</style>
