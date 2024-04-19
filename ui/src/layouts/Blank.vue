<template>
  <a-layout>
    <a-layout-content class="container">
      <router-view></router-view>
    </a-layout-content>
  </a-layout>
</template>

<style>
.container {
  display: flex;
  justify-content: center; /* 水平居中 */
  align-items: center; /* 垂直居中 */
  height: 100vh; /* 可选，使容器高度为视口高度以实现全屏垂直居中 */
}
</style>
