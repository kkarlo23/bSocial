<template>
  <div>
    <div class="search-container">
      <input
        class="search-input"
        placeholder="Search for user..."
        type="text"
      />
    </div>
  </div>
</template>

<script>
import { defineComponent } from "vue";

export default defineComponent({
  name: "MainView",
  props: {
    title: String,
  },
  setup() {},
});
</script>

<style>
.search-container {
  width: 100%;
  display: flex;
  justify-content: center;
}
.search-input {
  width: 300px;
  height: 35px;
  box-sizing: border-box;
  border-radius: 10px;
  padding: 10px;
  outline: none;
  border: 0;
}
</style>
