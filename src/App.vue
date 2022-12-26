<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <header>
    <nav class="navbar navbar-expand-lg bg-light">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">
          <i class="bi-diagram-3-fill"></i>
          Мой сайт
        </a>
        <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
          <span class="navbar-toggler-icon"></span>
        </button>
        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav me-auto mb-2mb-lg-0">
            <li class="nav-item">
              <!-- <a class="nav-link active" aria-current="page" href="index.html"><i class="bi-house"></i> Главная</a> -->
              <router-link class="nav-link" to="/"><i class="bi-house"></i> Главная</router-link>
            </li>
            <li class="nav-item">
              <!-- <a class="nav-link" href="products.html"><i class="bi-bag-check"></i> Товары</a> -->
              <router-link class="nav-link" to="/products"><i class="bi-star"></i> Товары</router-link>
            </li>
            <li class="nav-item">
              <!-- <a class="nav-link" href="about.html"><i class="bi-file-earmark-person"></i> О нас</a> -->
              <router-link class="nav-link" to="/about"><i class="bi-person-circle"></i> О нас</router-link>
            </li>


          </ul>

          <div class="d-flex">
            <!-- <div>
              <a class="nav-link" href="about.html"><i class="bi-file-earmark-person"></i> О нас</a> 
              <router-link class="nav-link" to="/auth"><i class="bi-box-arrow-in-left"></i> Войти</router-link>
            </div> -->
            <div> 
              <router-link class="nav-link" to="/reg"><i class="bi-box-arrow-in-right"></i> Выйти</router-link>
            </div>
          </div>
        </div>
      </div>
    </nav> 

  </header>
  <main class="container">
    <RouterView /> <!--Именно в этом месте vue-router будет загружать компоненты в соответствие с навигацией-->
  </main>
  <footer>
  </footer>
</template>


<script>
</script>
<style scoped>
</style>
