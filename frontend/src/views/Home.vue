<template>
  <div class="home">
<!--    <img alt="RateMyWorkLogo" src="../assets/logo.png">-->
    <div class="home__container">
      <h1>Rate My Work</h1>
      <div class="home__contest">
        <h2 class="home__subtitle">Start a new contest right now !</h2>
        <FilesUpload title="Upload some work !"></FilesUpload>
      </div>

      <div class="home__share">
        <div>Partager le lien !</div>
      </div>
    </div>
  </div>
</template>

<script>
// @ is an alias to /src

import FilesUpload from '@/components/FilesUpload'

export default {
  name: 'Home',
  components: { FilesUpload }
}
</script>

<style>
.home {
  display: flex;
  justify-content: center;
}
.home__container {
  text-align: center;
}
.home__contest {
  padding: 1.5rem;
  background-color: #f2cc8f;
  border-radius: 2rem;
}
.home__subtitle {
  margin-top: 0;
}
.home__share {
  padding: 1.5rem;
  margin: 2rem 0;
  border-radius: 2rem;
  background-color: #81b29a;
}
</style>
