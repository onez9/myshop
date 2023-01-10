<script setup>
import Swal from 'sweetalert2'
</script>

<template>
  <div class="container" id="elements">
    <!-- <p>Список товаров</p> -->
  
    <!-- навигация по страницам -->
    <div class="mt-3 d-flex justify-content-center">
      <nav aria-label="Page navigation mt-1 example">
        <ul class="pagination">
          <li v-if="currentPage>0" class="page-item"><a class="page-link" href="#" @click="changePage(currentPage-1)">Пред</a></li>
          <li :class="{'page-item': true, 'active': (page-1==currentPage)}" v-for="page in totalpages" :key="page">
            <a class="page-link" @click="changePage(page-1)" href="#">{{ page }}</a>
          </li>
          <li v-if="currentPage<totalpages-1" class="page-item"><a class="page-link" href="#" @click="changePage(currentPage+1)">След</a></li>
        </ul>
      </nav>
    </div>
    
    <div class="d-flex">
      <button v-if="!hide && access" @click="hide=true" class="btn btn-info me-1"><i class="bi-plus-square"></i> {{ new_element }}</button>
      <!-- <button @click="" class="btn btn-warning"><i class="bi-plus-square"></i> Перейти в корзину</button> -->
      <!-- {{ my_name }} -->
      <router-link v-if="my_name!=''" to="/cart" customv-slot="{ navigate }">
      <button class="btn btn-warning" @click="navigate" role="link">
        <i class="bi-link"></i> {{ get_collection }}
      </button>
      </router-link>


    </div>

    <!-- добавить товар -->
    <div v-if="hide">
      <!-- <form action="/sendProduct" method="POST" enctype="multipart/form-data"> -->
      <form method="POST">
        <label for="name">Название: {{ name_product }}</label>
        <input name="name" v-model="name_product" type="text" id="name" class="form-control" required>
        <label for="description">Описание: {{ description_product }}</label>
        <input name="description" v-model="description_product" type="text" id="description" class="form-control" required>
        <label for="price">{{ temp }}:</label>
        <input name="price" v-model.number="price_product" type="number" min="100" max="1000" id="price" class="form-control mb-1" required>
        <label for="Avatar">Аватар:</label>
        <input name="file" type="file" id="Avatar" class="form-control mb-1" required>
        <!-- <input name="file" ref="file" type="file" id="Avatar" class="form-control mb-1" required> -->
        <!-- <label for="name">name</label> -->
        <!-- <input type="text" id="name" class="form-control"> -->
        <div class="d-flex mt-3 mb-3">
          <button @click="sendProduct" type="button" class="btn btn-success btn-sm me-1"><i class="bi-plus-square"></i> Добавить</button>
          <button @click="canceling" class="btn btn-warning btn-sm me-1"><i class="bi-x-square"></i> Отмена</button>
          <button @click="ending" class="btn btn-danger btn-sm"><i class="bi-badge-cc"></i> Завершить</button>
        </div>
      </form>
    </div>
    <!-- {{currentPage}} -->
    <!-- список товаров -->
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
                  {{ element.id }}
                  <div v-if="!element.editmode" class="product-name">{{ element.name }}</div>
                  <label v-else>Название:</label>
                  <input class="form-control mb-1" type="text" v-if="element.editmode" v-model="element.name">
                  <div v-if="!element.editmode" class="product-description">{{ element.description }}</div>
                  <label v-else>Описание:</label>
                  <input class="form-control mb-1" type="text" v-if="element.editmode" v-model="element.description">
                  <label v-if="element.editmode">{{ temp }}</label>
                  <input type="number" class="form-control mb-1" v-model="element.price" v-if="element.editmode">
                  <div v-if="element.editmode" class="product-description">Изображение</div>
                  <input type="file" :id="element.id" class="form-control mb-1" v-if="element.editmode">
                  
                  
                  <div class="flex-grow-1 mb-3"></div>
                  <div class="d-flex justify-content-end">
    
    
                    <div v-if="!element.editmode" class="product-price">{{ temp }}: <strong>{{ element.price }}</strong></div>
    
                    <button v-if="!element.editmode && access" @click="enableEditMode(element)" class="btn btn-info btn-sm me-1" title="Выполнить какие-нибудь изменения">
                      <i class="bi bi-pen"></i> <span class="d-none d-md-inline"> Редактировать</span>
                    </button>
                    <button v-if="element.editmode && access" @click="updateProduct(element)" class="btn btn-success btn-sm me-1" title="Выполнить какие-нибудь изменения">
                      <i class="bi-journals"></i> <span class="d-none d-md-inline"> Сохранить</span>
                    </button>
                    <button v-if="access" @click="delProduct(element)" class="btn btn-danger btn-sm me-1" title="Удалить">
                      <i class="bi-x-circle"></i> <span class="d-none d-md-inline">Удалить</span>
                    </button>
                    <button v-if="my_name!=''" @click="addToCart(element)" class="btn btn-warning btn-sm me-1" title="В корзину">
                      <i class="bi-cart"></i> <span class="d-none d-md-inline">В корзину</span>
                    </button>
                    <button v-if="element.editmode" @click="disableEditMode(element)" class="btn btn-secondary btn-sm" title="Отменить">
                      <i class="bi-x-octagon"></i> <span class="d-none d-md-inline"> Отмена</span>
                    </button>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
const productsPerPage = 3
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
      file: "",
      // temp: "Цена",
      // new_element: "Создать товар",
      access: false,
    }
  },
  props: {
    temp: String,
    new_element: String,
    my_name: String,
    get_collection: String,
  },
  computed: { 
    a() {
      this.elements.sort(function (a, b) {
        if (parseInt(a.id) > parseInt(b.id)) {
          return 1;
        }
        if (parseInt(a.id) < parseInt(b.id)) {
          return -1;
        }
        // a должно быть равным b
        return 0;
      });
    }
  },
  async mounted() {
    await this.getProducts()
    await this.getProductsCount()
    await this.is_access()
    await this.sorting()
    // this.='cast '
    // console.log('Компонент примонтирован!');
  },
  methods: {
    async sorting() {
      this.elements.sort(function (a, b) {
        if (parseInt(a.id) > parseInt(b.id)) {
          return 1;
        }
        if (parseInt(a.id) < parseInt(b.id)) {
          return -1;
        }
        // a должно быть равным b
        return 0;
      });
    },
    async t1() {
      // alert('Товар добавлен в корзину')
      // Swal.fire({
      //   title: 'Вы уверены что хотит добавить?',
      //   showDenyButton: true,
      //   showCancelButton: true,
      //   confirmButtonText: 'Добавить',
      //   denyButtonText: `Don't save`,
      // }).then((result) => {
      //   /* Read more about isConfirmed, isDenied below */
      //   if (result.isConfirmed) {
      //     Swal.fire('Saved!', '', 'success')
      //   } else if (result.isDenied) {
      //     Swal.fire('Changes are not saved', '', 'info')
      //   }
      // })
      Swal.fire('Товар добавлен в корзину!')
    },
    async uploadFile(e) {
      const files = e.target.files
      const data = new FormData()
      data.append('file', files[0])
    
      const response = await fetch('/t2', {
        method: 'POST',
        body: data
      })


      // .then(response => response.json())
      // .then(data => {
      //   console.log(data)
      // })
      // .catch(error => {
      //   console.error(error)
      // })
      console.log(response.json())


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
      this.price_product=100
      this.name_product=""
      this.description_product=""

      this.elements = await response.json();
      await this.getProducts()
      await this.getProductsCount()

    },
    async sendProduct() {

      let formData = new FormData()

      // const files = e.target.files // работает с @change=""
      // const data = new FormData()
      let input = document.getElementById('Avatar')
      formData.append('file', input.files[0])
      // this.file = this.$refs.file.files[0]
      // formData.append('file', this.file)
      // formData.append('file', files[0])
      
      formData.append('name', this.name_product)
      formData.append('description', this.description_product)
      formData.append('price', this.price_product)


      // const data = new URLSearchParams();
      // for (const pair of formData) {
      //     data.append(pair[0], pair[1]);
          
      // }
      
      
      const response = await fetch(`/sendProduct?p=${this.currentPage}&limit=${productsPerPage}`, {
        method: 'POST',
        // headers: {
        //   'Content-Type': 'application/x-www-form-urlencoded'
        // },
        body: formData

      })
      // let product = document.getElementsByName("name").text
      // let description = document.getElementsByName("description").Values
      // let price = document.getElementsByName("price").Value

      this.elements = await response.json();

      await this.getProductsCount()


      // this.price_product=100
      // this.name_product=""
      // this.description_product=""
      // document.querySelector('input[type=file]').value = '';
      await this.canceling()

      // await this.getProducts()
      // await this.getProductsCount()
      // console.log(product)
      // console.log(description)
      // console.log(price)
    },
    async updateProduct(element) {
      element.editmode=false

      let formData = new FormData()
      let input = document.getElementById(element.id)
      formData.append('file', input.files[0])
      formData.append('name', element.name)
      formData.append('description', element.description)
      formData.append('price', element.price)
      formData.append('id', element.id)
      console.log('element.id: ', element.id)

      const response = await fetch(`/updaterec?p=${this.currentPage}&limit=${productsPerPage}`, {
        method: 'POST',
        headers: {
          // 'Content-Type': 'application/json'
          // 'Content-Type': 'multipart/form-data'
        },
        // body: JSON.stringify({
        //   id: element.id,
        //   name: element.name,
        //   description: element.description,
        //   price: element.price
        // })
        body: formData

      })

      this.elements = await response.json();
      this.elements.sort(function (a, b) {
        if (parseInt(a.id) > parseInt(b.id)) {
          return 1;
        }
        if (parseInt(a.id) < parseInt(b.id)) {
          return -1;
        }
        // a должно быть равным b
        return 0;
      });
      await this.getProductsCount()
    },
    async delProduct(element) {
      this.elements.splice(this.elements.indexOf(element), 1)
      // alert(this.elements.indexOf(element))
      const response = await fetch(`/delrec?name=${element.name}&p=${this.currentPage}&limit=${productsPerPage}`, {
        method: 'POST',

      })
      this.elements = await response.json();
      await this.getProductsCount()
    },
    async addToCart(element) {
      const response = await fetch("/addtocart", {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({
          product_id: element.id
        })
      })
      this.t1()
      console.log(await response.json())
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
      this.totalpages = Math.ceil(productscount / productsPerPage)

    },
    async getProducts() {
      const response = await fetch(`/getproducts?p=${this.currentPage}&limit=${productsPerPage}`, {
        method: 'GET'
      })

      this.elements = await response.json();
      // console.log(elements);
    },
    async ending() {
      this.price_product=100
      this.name_product=""
      this.description_product=""
      document.querySelector('input[type=file]').value = "";
      this.hide=false
    },
    async canceling(){
      this.price_product=100
      this.name_product=""
      this.description_product=""
      document.querySelector('input[type=file]').value = "";
    },
    async is_access() {
      const response = await fetch("/is_access", {
        method: "POST"
      })

      this.access = await response.json()
    },
    async enableEditMode(element) {
      element.editmode=true
      element.name_back=element.name
      element.description_back=element.description
      element.price_back=element.price
    },
    async disableEditMode(element) {
      element.editmode=false
      element.name=element.name_back
      element.description=element.description_back
      element.price=element.price_back
    }
  }

}
</script>
