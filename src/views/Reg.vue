<script setup>


</script>

<template>
  <!-- <div class="d-flex">
    <button class="btn btn-warning m-3" @click="testGet">Тест GET запроса</button>
    <button class="btn btn-danger m-3" @click="testPost">Тест POST запроса</button>
  </div> -->

  <h1>Страница регистрации</h1>
  <main>
    <div class="container">
      <div class="row">
        <div class="col-sm-4">
          <label for="firstname">Имя:</label>
          <input type="text" id="firstname" class="form-control" required>
          <label for="lastname">Фамилия:</label>
          <input type="text" id="lastname" class="form-control" required>
          <label for="email">Email:</label>
          <input type="email" id="email" class="form-control" required>
          <label for="pass1">Введите пароль:</label>
          <input type="password" id="pass1" class="form-control" required>
          <label for="pass2">Повторите пароль:</label>
          <input type="password" id="pass2" class="form-control" required>
          <button class="btn btn-success mt-1 me-1">Войти</button>
          <button class="btn btn-success mt-1">Авторизация</button>

        </div>
      </div>
    </div>
  </main>


</template>

<script>
export default {
  data() {
    return {
    }
  },
  methods: {
    async testPost() {
      // Отправляем запрос типа POST
      const response = await fetch('/testpost', {
        method: 'POST', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: "Иван"
        }) 
      });

      const jsonResult = await response.json();

      console.log(jsonResult);
    },
    async testGet() {
      // Отправляем запрос типа GET
      const response = await fetch('/testget', {
        method: 'GET',
        // 'Access-Control-Allow-Origin': '*'
      });

      const textResult = await response.text();

      console.log(textResult);
    }
  }
}
</script>
