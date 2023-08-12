<script setup>



</script>

<template>
  <!-- <div class="d-flex">
    <button class="btn btn-warning m-3" @click="testGet">Тест GET запроса</button>
    <button class="btn btn-danger m-3" @click="testPost">Тест POST запроса</button>
  </div> -->

  <!-- <img src="/assets/img/main.jpg" class="img-fluid"> -->
  <!-- <img src="assets/img/main.jpg" class="img-fluid"> -->
  <div v-if="true" id="carouselExampleIndicators" class="carousel slide" data-bs-ride="carousel">
    <div class="carousel-indicators">
      <!-- <button 
        type="button" 
        data-bs-target="#carouselExampleIndicators" 
        data-bs-slide-to="0" 
        class="active" 
        aria-current="true" 
        aria-label="Slide 1">
      </button> -->
      <button
        :class="{'active': (index==0)}"
        v-for="(element, index) in elements" :key="index" 
        type="button" 
        data-bs-target="#carouselExampleIndicators" 
        :data-bs-slide-to="index" 
        :aria-label="index">
      
      </button>
    </div>
    <div class='carousel-inner' >
      <div 
        :class="{'carousel-item': true, 'active': (index==0)}" 
        v-for="(element, index) in elements" :key="index">
        <!-- <div class="carousel-item active"> -->
        <img src="assets/img/c4c78f0e-2b99-4928-aff8-c0f31c781d2c.jpeg" class="d-block w-100" alt="...">
      </div>
    </div>
    <button 
      class="carousel-control-prev" 
      type="button" 
      data-bs-target="#carouselExampleIndicators" 
      data-bs-slide="prev">
      <span class="carousel-control-prev-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Previous</span>
    </button>
    <button 
      class="carousel-control-next" 
      type="button" 
      data-bs-target="#carouselExampleIndicators" 
      data-bs-slide="next">
      <span class="carousel-control-next-icon" aria-hidden="true"></span>
      <span class="visually-hidden">Next</span>
    </button>
  </div>


</template>

<script>
export default {
  data() {
    return {
      elements: [1, 2, 3, 4, 5]
    }
  },
  methods: {
  
  }
}
</script>
