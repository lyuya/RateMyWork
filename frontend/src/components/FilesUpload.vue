<template>
  <div class="container">
    <div class="title">{{ title }}</div>
    <div class="upload">
      Work in progress
    </div>
  </div>
</template>

<script>
export default {
  name: 'FilesUpload',
  props: {
    title: String
  }
}
</script>

<style scoped>
.upload {
  height: 15rem;
  width: 40rem;
  border: 0.2rem dotted rgba(0,0,0,0.5);
  display: flex;
  align-items: center;
  justify-content: center;
}
.title {
  text-align: start;
}
</style>
