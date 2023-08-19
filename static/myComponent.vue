<script setup>
import { ref } from 'vue'
// import ButtonCounter from './buttonCounter.vue'

const greeting = ref('Hello World! from vue')
// const posts = [1, 2, 3]
</script>

<template>
    <p class="greeting">{{ greeting }}</p>
    <!-- <span>this is span</span>
    <ul>
        <li v-for="post in posts">
            <ButtonCounter />
        </li>
    </ul>
    <ButtonCounter></ButtonCounter> -->
</template>

<style>
.greeting {
    color: red;
    font-weight: bold;
}
</style>