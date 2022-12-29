<script setup>
</script>

<template>
  <h1>Моя компания</h1>
  <p>У нас самые лучшие услуги и команда в этой области!</p>
  <h1>Наша команда</h1>
	<div id="raiting">
		<div class="d-flex">
			<div v-for="(item, index) in items" :key="index">
				<button @click="game(item, index)" :class="{'btn btn-light btn-lg me-1': true, 'btn-dark': (item.active==true)}"><i class="bi-star"></i></button>
			</div>
		</div>
	</div>
  <div class="row">
	  <div class="col-md-4">
		  <div class="card shadow rounded my-3 th">
			  <div class="card-body">
				  <div><img src="assets/img/p1.jpg" class="avatar"></div>
				  <div class="avatar-post"><i class="bi bi-person"></i> Генеральный директор</div>
				  <div class="avatar-text">Самый целеустремленный человек</div>
			  </div>
		  </div>
	  </div>
	  <div class="col-md-4">
		  <div class="card shadow rounded my-3 th">
			  <div class="card-body">
				  <div><img src="assets/img/p2.jpg" class="avatar"></div>
				  <div class="avatar-post">Уборщик</div>
				  <div class="avatar-text">Моет полы</div>
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
			items: [{},{},{},{},{}]
		}
	},
	computed: {

	},
	async mounted() {

	},
	methods: {
		async game(item, index) {
			item.active=true
			for	(let i=0;i<this.items.length;i++) {
				if (i<=index)
					this.items[i].active=true
				else
					this.items[i].active=false

			}
		}
	}
}
</script>
