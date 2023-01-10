<script setup>
</script>

<template>
  <h1>Моя компания</h1>
  <p>У нас самые лучшие услуги и команда в этой области!</p>
  <h1>Наша команда</h1>

      <!--Здесь представление компонента (код на HTML)-->
      <div class="col mt-1">

			<div class="d-flex">
				<div class="" v-for="(item, index) in items" :key="index">
					<button class="b1" @click="game(index)"><i :class="{'bi-star-fill hamIcon': true, 'bi-star hamIcon': (item.field==false)}"></i></button>
				</div>
			</div>

			</div>
			<h5>{{label_status}}</h5>
			<!-- <svg class="bi" width="32" height="32" fill="currentColor">
			<use xlink:href="bootstrap-icons.svg#heart-fill"/>
			</svg> -->

			<!-- <i class="bi-star hamIcon"></i> -->




  <div class="row">
	  <div class="col-md-4">
		  <div class="card shadow rounded my-3 th">
			  <div class="card-body">
				  <div><img src="assets/img/p1.jpg" class="avatar"></div>
				  <div class="avatar-post"><i class="bi bi-person"></i> Владелец магазина</div>
				  <div class="avatar-text">Елена Карловна</div>
			  </div>
		  </div>
	  </div>
	  <div class="col-md-4">
		  <div class="card shadow rounded my-3 th">
			  <div class="card-body">
				  <div><img src="assets/img/p2.jpg" class="avatar"></div>
				  <div class="avatar-post">Уборщик</div>
				  <div class="avatar-text">Екатерин Кубикиова</div>
			  </div>
		  </div>
	  </div>
	  <div class="col-md-4">
		  <div class="card shadow rounded my-3 th">
			  <div class="card-body">
				  <div><img src="assets/img/p3.jpg" class="avatar"></div>
				  <div class="avatar-post">Рабочий</div>
				  <div class="avatar-text">Самый главный продавец</div>
			  </div>
		  </div>
	  </div>
 </div>
</template>

<script>
export default {
	el: "#raiting",
	data() {
		return {
			// items: [{a:1},{b:2},{c:3},{d:4},{e:5}]
			// items: [1,2,3,4,5]
			items: [{},{},{},{},{}],
			label_status: ""
		}
	},
	computed: {

	},
	async mounted() {
		await start()
	},
	methods: {
		async game(index) {
			// item.active=true

			switch (index) {
				case 0:
					this.label_status = "Всё плохо"
					break;
				case 1:
					this.label_status = "Не удовлетворительно"
					break;
				case 2:
					this.label_status = "Удовлетворительно"
					break;
				case 3:
					this.label_status = "Хорошо"
					break;
				case 4:
					this.label_status = "Отлично"
					break;

				default:
					break;
			}
			for	(let i=0;i<this.items.length;i++) {
				// if (i<=index)
				// 	this.items[i].active=true
				// else
				// 	this.items[i].active=false
				this.items[i].field = (i<=index)? true:false;

			}
		},
		async start() {
      for (let index = 0; index < this.items.length; index++) {
        this.items[index].field=true;

      }
    },

	}
}
</script>


<style scoped>
.btn-success {
  color: #fff;
  background-color: #20ad00;
  border-color: #20ad00;
}

.btn-warning {
  color: #fff;
  background-color: #f36a02d1;
  border-color: #ffa04d;
}
.hamIcon{
  font-size: 10em;  

}
.b1,.b1:hover,.b1:focus,.b1::-moz-focus-inner  {
  color:#fcf805;
  background-color: #ffffff;
  border: 0;
  outline: 0 !important;
}

</style>