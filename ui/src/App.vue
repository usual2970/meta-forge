<template>
  <component :is="this.$route.meta.layoutComponent">
    <slot />
  </component>
</template>

<script>
export default {
  name: 'AppLayout'
}
</script>

<style scoped></style>
