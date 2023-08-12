<script setup>


</script>

<template>
  <div v-for="(item, i) in items" :key="i">
    <img :src="item.name" style="width: 10%;" alt="">
    {{ item.name.split('/')[2] }} 
  </div>
  

</template>

<script>
export default {
  data() {
    return {
      items: []
    }
  },
  async mounted() {
    await this.get_items()
    // console.log('Компонент примонтирован!');
  },
  methods: {
    async get_items() {
      const response = await fetch('/getg', {
        method: "POST",

      })
      this.items = await response.json();
    }
  }
}
</script>
