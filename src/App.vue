<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>

<template>
  <header>
    <nav class="navbar navbar-expand-lg bg-light">
      <div class="container-fluid">
        <a class="navbar-brand" href="#">
          <i class="bi-diagram-3-fill"></i>
          {{ name_site }}
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
              <router-link class="nav-link" to="/products"><i class="bi-star"></i> {{ elements }}</router-link>
            </li>
            <li class="nav-item">
              <!-- <a class="nav-link" href="about.html"><i class="bi-file-earmark-person"></i> О нас</a> -->
              <router-link class="nav-link" to="/about"><i class="bi-person-circle"></i> О нас</router-link>
            </li>
            <li class="nav-item">
              <!-- <a class="nav-link" href="about.html"><i class="bi-file-earmark-person"></i> О нас</a> -->
              <router-link class="nav-link" to="/girls"><i class="bi-person-circle"></i> Бабы</router-link>
            </li>


          </ul>

          <div class="d-flex">
            <!-- <div>
              <a class="nav-link" href="about.html"><i class="bi-file-earmark-person"></i> О нас</a> 
              <router-link class="nav-link" to="/auth"><i class="bi-box-arrow-in-left"></i> Войти</router-link>
            </div> -->
            <span class="navbar-text">
              {{ my_name }}
            </span>
            <div v-if="my_name!=''"> 
              <router-link @click="logout" class="nav-link" to="/auth"><i class="bi-box-arrow-in-right"></i> Выйти</router-link>
            </div>
            <div v-else> 
              <router-link @click="logout" class="nav-link" to="/auth"><i class="bi-box-arrow-in-right"></i> Войти</router-link>
            </div>
          </div>
        </div>
      </div>
    </nav> 

  </header>
  <main class="container" >
    <RouterView 
    :get_collection="get_collection" 
    :temp="temp" 
    :new_element="new_element" 
    :my_name="my_name"/> 
    <!--Именно в этом месте vue-router будет загружать компоненты в соответствие с навигацией-->
  </main>
  <footer>
  </footer>
</template>


<script>

export default {
  data() {
    return {
      firstname: "",
      lastname: "",
      email: "",
      password: "",
      pass1: "",
      pass2: "",
      elements: "Товары",
      my_name: "",
      temp: "Важность",
      new_element: "Добавить задачу",
      get_collection: "Моя корзина",
      name_site: "Сувенирная"
    }
  },
  async mounted() {
    await this.whoami()
  },
  methods: {
    async whoami() {
      const response = await fetch("/whoami", {
        method: "GET"
      })

      this.my_name = await response.json()
      // this.my_name="666"
    },
    async logout() {
      this.my_name = ""
      const response = await fetch("/logout", {
        method: "POST",

      })
      console.log(await response.json())
    }
  }
}


</script>
<style scoped>
</style>
