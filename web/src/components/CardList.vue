<script setup>
import Card from './Card.vue'
function show() {
  alert('Hello!')
}
</script>

<template>
  <div class="grid grid-cols-4 gap-5">
    <Card
      image-url="/sneakers/sneakers-1.jpg"
      title="Мужские кроссовки Nike Bazzar Mid"
      :price="5000"
      :onClickAdd="show"
    />
  </div>
</template>
