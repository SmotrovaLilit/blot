<script setup lang="ts">
import {RouterLink, RouterView} from 'vue-router'
import UserInfo from "@/components/UserInfo.vue";

</script>

<template>
  <header>
    <div class="wrapper">
      <nav>
<!--        <RouterLink to="/gamesets">Games</RouterLink>-->
      </nav>
    </div>
  </header>
  <div class="dark-theme container">
    <UserInfo />
    <RouterView />
  </div>
</template>

<style scoped>
.light-theme {
  --card-background: #ccd3d9;
  --card-border: #e0e0e0;
  --card-shadow: 0 10px 10px rgba(0, 0, 0, 0.1),
}
.dark-theme {
  --card-background: #494b4d;
  --card-border: #333333;
  --card-shadow: 0 10px 10px rgba(255, 255, 255, 0.1),
  0 2px 4px rgba(255, 255, 255, 0.06);
}

header {
  line-height: 1.5;
  max-height: 100vh;
  position: fixed;
  width: 100%;
  top: 0;
  left: 0;
}

.logo {
  display: block;
  margin: 0 auto 2rem;
}

nav {
  width: 100%;
  font-size: 12px;
  text-align: center;
  margin-top: 2rem;
}

nav a.router-link-exact-active {
  color: var(--color-text);
}

nav a.router-link-exact-active:hover {
  background-color: transparent;
}

nav a {
  display: inline-block;
  padding: 0 1rem;
  border-left: 1px solid var(--color-border);
}

nav a:first-of-type {
  border: 0;
}

@media (min-width: 1024px) {
  nav {
    text-align: left;
    margin-left: -1rem;
    font-size: 1rem;

    padding: 1rem 0;
    margin-top: 1rem;
  }
}
</style>
