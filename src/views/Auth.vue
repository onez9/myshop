<script setup>


</script>

<template>
  <h1>Страница авторизации</h1>
  <main>
    <div class="container">
      <div class="row">
        <div class="col-sm-4">
          <form action="/authentication1" method="POST" enctype="multipart/form-data">
            <label for="email">Email:</label>
            <input name="email" type="text" id="email" class="form-control" placeholder="Введите email" required>
            <label for="pass">Пароль:</label>
            <input name="password" type="password" id="pass" class="form-control" placeholder="Введите пароль" required>
            
            <button type="submit" class="btn btn-success mt-1 me-1"><i class="bi-box-arrow-in-left"></i> Войти</button>

            <router-link to="/reg" customv-slot="{ navigate }">
              <button class="btn btn-success mt-1" @click="navigate" role="link">
                <i class="bi-person-plus"></i> Регистрация
              </button>
            </router-link>
            <!-- <button class="btn btn-success mt-1">Регистрация</button> -->
          </form>

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
