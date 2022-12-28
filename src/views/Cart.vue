<script setup>
</script>

<template>
  <div class="container" id="elements">
    <!-- <p>Список товаров</p> -->
    <!-- <form action="/sendProduct" method="POST" enctype="multipart/form-data"> -->
    <!-- {{currentPage}} -->
    <table class="table table-striped table-hover">
      <tbody>
        <tr @click="activeElem = element" v-for="(element, index) in cart_items" :key="index">
          <div class="card shadow rounded my-1">
            <div class="card-body">
              <div class="row">
                <div class="col-md-2">
                  <!--Здесь будет фото товара-->
                  <img v-bind:src="element.imgpath" class="product-image">
    
                </div>
                <div class="col-md-10">
                  {{ element.id }}

                  <div class="product-name">{{ element.name }}</div>
                  <div class="product-description">{{ element.description }}</div>

                  <div class="flex-grow-1"></div>
                  <div class="d-flex justify-content-end">    
                    <div class="product-price">{{ temp }}: <strong>{{ element.price }}</strong></div>
    
                    <button @click="delItemCart(element)" class="btn btn-danger me-1" title="Удалить">
                      <i class="bi bi-x-circle"></i> <span class="d-none d-md-inline">Удалить</span>
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
export default {
  el: '#elements',
  data() {
    return {
      // основные элементы корзины
      cart_items: [],
      temp: "Пиздатость"
      
    }
  },
  computed: { 
  },
  async mounted() {
    await this.getItemsCart()
    // console.log('Компонент примонтирован!');
  },
  methods: {
    // кнопка удалить в корзине
    async delItemCart(element) {
      this.cart_items.splice(this.cart_items.indexOf(element), 1)
      // alert(this.elements.indexOf(element))
      const response = await fetch(`/del_rec_in_cart?id=${element.id}`, {
        method: 'POST',

      })
      this.cart_items = await response.json();
    },
    // выполняется при загрузке страницы
    async getItemsCart() {
      const response = await fetch(`/getcartitems`, {
        method: 'GET'
      })

      this.cart_items = await response.json();
      // console.log(elements);
    },
  }

}
</script>
