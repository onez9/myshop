<script setup>
</script>

<template>
  <main class="container" id="elements">
    <h1>Моя компания</h1>
    <p>Тут представлены наши основные товары</p>



    <nav aria-label="Page navigation example">
      <ul class="pagination">
        <li v-if="currentPage>0" class="page-item"><a class="page-link" href="#" @click="changePage(currentPage-1)">Пред</a></li>
        <li :class="{'page-item': true, 'active': (page-1==currentPage)}" v-for="page in totalpages" :key="page">
          <a class="page-link" @click="changePage(page-1)" href="#">{{ page }}</a>
        </li>
        <li v-if="currentPage<totalpages-1" class="page-item"><a class="page-link" href="#" @click="changePage(currentPage+1)">След</a></li>
      </ul>
      <button v-if="!hide" @click="hide=true" class="btn btn-success">Добавить товар</button>
      <div v-if="hide">
        <label for="name">Name: {{ name_product }}</label>
        <input v-model="name_product" type="text" id="name" class="form-control" required>
        <label for="description">Description: {{ description_product }}</label>
        <input v-model="description_product" type="text" id="description" class="form-control" required>
        <label for="price">Price:</label>
        <input v-model.number="price_product" type="number" min="100" max="1000" id="price" class="form-control mb-1" required>
        <label for="Avatar">Avatar:</label>
        <input type="file" id="Avatar" class="form-control mb-1">
        <!-- <label for="name">name</label> -->
        <!-- <input type="text" id="name" class="form-control"> -->

        <div class="d-flex">
          <button @click="addProduct(name_product, description_product, price_product)" class="btn btn-success me-1">Добавить</button>
          <button @click="hide=false" class="btn btn-danger">Отмена</button>
        </div>
      </div>
    </nav>
    <!-- {{currentPage}} -->
    <table class="table table-striped table-hover">
      <thead>
        <tr>
          <!-- <th scope="col">#</th>
          <th scope="col">First</th>
          <th scope="col">Last</th>
          <th scope="col">Handle</th> -->
        </tr>
      </thead>
      <tbody>
        
        
       

        <tr @click="activeElem = element" v-for="(element, index) in elements" :key="index">
          <div class="card shadow rounded my-1">
            <div class="card-body">
              <div class="row">
                <div class="col-md-2">
                  <!--Здесь будет фото товара-->
                  <img v-bind:src="element.imgpath" class="product-image">
                
                </div>
                <div class="col-md-10">

                  <div v-if="!element.editmode" class="product-name">{{ element.name }}</div>
                  <label v-else>Имя товара</label>
                  <input class="form-control mb-1" type="text" v-if="element.editmode" v-model="element.name">

                  <div v-if="!element.editmode" class="product-description">{{ element.description }}</div>
                  <label v-else>Описание товара</label>
                  <input class="form-control mb-1" type="text" v-if="element.editmode" v-model="element.description">


                  <label v-if="element.editmode">Цена</label>
                  <input type="number" class="form-control mb-1" v-model="element.price" v-if="element.editmode">
                  <div class="flex-grow-1"></div>
                  <div class="d-flex justify-content-end">
                    
                    
                    <div v-show="!element.editmode" class="product-price">Цена: <strong>{{ element.price }}</strong></div>

                    
                    <button v-show="!element.editmode" @click="element.editmode=true" class="btn btn-success me-1" title="Выполнить какие-нибудь изменения">
                      <i class="bi bi-pen"></i> <span class="d-none d-md-inline">Редактировать</span>
                    </button>
                    <button v-show="element.editmode" @click="element.editmode=false" class="btn btn-info me-1" title="Выполнить какие-нибудь изменения">
                      <i class="bi bi-pen"></i> <span class="d-none d-md-inline">Сохранить</span>
                    </button>



                    <button @click="delProduct(element)" class="btn btn-danger me-1" title="Удалить">
                      <i class="bi bi-x-circle"></i> <span class="d-none d-md-inline">Удалить</span>
                    </button>
                    <button @click="t1" class="btn btn-warning" title="В корзину">
                      <i class="bi bi-cart"></i> <span class="d-none d-md-inline">В корзину</span>
                    </button>

                  </div>
                </div>
              </div>
            </div>
          </div>
        </tr>
      </tbody>
    </table>

   





  </main>

</template>

<script>
const produtcsPerPage = 3
// document.getElementById("b1").addEventListener('click', () => alert("Товар добавлен в корзину"), false)
// document.getElementById("b2").addEventListener('click', () => alert("Товар добавлен в корзину"), false)
// document.getElementById("b3").addEventListener('click', () => alert("Товар добавлен в корзину"), false)
export default {
  el: '#elements',
  data() {
    return {
      currentRating: 0,
      totalpages: 0,
      elements: [],
      currentPage: 0,
      hide: false,
      name_product: "",
      price_product: 0,
      description_product: "",
      
    }
  },
  computed: { 
  },
  async mounted() {
    await this.getProducts()
    await this.getProductsCount()
    // console.log('Компонент примонтирован!');
  },
  methods: {
    t1() {
      alert('Товар добавлен в корзину')
    },

    async addProduct(name, description, price) {
      // alert(name + description)
      const response = await fetch('/sendProduct', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          name: name,
          description: description ,
          price: price
        }) 

      })
      await this.getProducts()
      await this.getProductsCount()
      this.elements = await response.json();
    },

    async updateProduct() {

    },

    async delProduct(element) {
      this.elements.splice(this.elements.indexOf(element), 1)
      // alert(this.elements.indexOf(element))
    },

    async changePage(page) {
      this.currentPage = page
      await this.getProducts()
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
      const response = await fetch(`/testget?name=Василий`, {
        method: 'GET',      
      });

      this.elements = await response.text();

      console.log(elements);
    },

    async getProductsCount() {
      const response = await fetch('getproductscount', {
        method: 'GET'
      })
      const productscount = parseInt(await response.text(), 10)
      this.totalpages = Math.ceil(productscount / produtcsPerPage)

    },

    async getProducts() {
      const response = await fetch(`/getproducts?p=${this.currentPage}&limit=${produtcsPerPage}`, {
        method: 'GET'
      })

      this.elements = await response.json();
      // console.log(elements);
    }



  }

}
</script>
