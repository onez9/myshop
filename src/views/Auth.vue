<script setup>


</script>

<template>
  <h1>Страница авторизации</h1>
  <main>
    <div class="container">
      <div class="row">
        <div class="col-sm-4">
          <label for="email">Email:</label>
          <input v-model="email" type="text" id="email" class="form-control" placeholder="Введите email" required>
          <label for="pass">Пароль:</label>
          <input v-model="password" type="password" id="pass" class="form-control" placeholder="Введите пароль" required>
          
          <button @click="authCheck(email, password)" class="btn btn-success mt-1 me-1">Войти</button>
          <button class="btn btn-success mt-1">Регистрация</button>

        </div>
      </div>
    </div>
  </main>

</template>

<script>
export default {
  data() {
    return {
      email: "",
      password: "",
    }
  },
  methods: {
    async authCheck(email, password) {
      // Отправляем запрос типа POST
      const response = await fetch('/authentication', {
        method: 'POST', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          email: email,
          password: password
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
