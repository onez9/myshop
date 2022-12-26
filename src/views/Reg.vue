<script setup>

// import sha256 from 'sha256'
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
          <input v-model="firstname" type="text" id="firstname" class="form-control" required>
          <label for="lastname">Фамилия:</label>
          <input v-model="lastname" type="text" id="lastname" class="form-control" required>
          <label for="email">Email:</label>
          <input v-model="email" type="email" id="email" class="form-control" required>
          <label for="pass1">Введите пароль:</label>
          <input v-model="password" type="password" id="pass1" class="form-control" required>
          <label for="pass2">Повторите пароль:</label>
          <input type="password" id="pass2" class="form-control" required>
          <button @click="addUser(firstname,lastname,email, password)" class="btn btn-success mt-1 me-1">Создать</button>
          <button class="btn btn-success mt-1">Авторизация</button>

        </div>
      </div>
    </div>
  </main>


</template>

<script>
// import { crypto } from 'crypto'
export default {
  data() {
    return {
      firstname: "",
      lastname: "",
      email: "",
      password: "",
    }
  },
  methods: {

    async addUser(firstname,lastname,email,password) {
      // Отправляем запрос типа POST


      // const sha256 = crypto.createHash('sha256');
      // const hash = sha256.update(password).digest('base64');


      const response = await fetch('/addUser', {
        method: 'POST', 
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          firstname: firstname,
          lastname: lastname,
          email: email,
          password: password,
        }) 
      });

      const jsonResult = await response.json();

      console.log(jsonResult);
    },
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
