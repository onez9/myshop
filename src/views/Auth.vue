<script setup>


</script>

<template>
  <h1>Страница авторизации</h1>
  <main>
    <div class="container">
      <div class="row">
        <div class="col-sm-4">
          <label for="email">Email:</label>
          <input type="text" id="email" class="form-control" placeholder="Введите email" required>
          <label for="pass">Пароль:</label>
          <input type="password" id="pass" class="form-control" placeholder="Введите пароль" required>
          
          <button class="btn btn-success mt-1 me-1">Войти</button>
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
