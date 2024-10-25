<script setup>
import Card from '../components/Card.vue'
import {onMounted, ref} from 'vue'

const tasks = ref({
    backlog:[
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:""},
    ],
    proccessing:[
    {title: "Задача два", executor:"Ваня усач"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    {title: "Задача один", executor:"Коля Кравченко"},
    ],
    finished:[
    {title: "Задача три", executor:"Ваня усыч"}
    ],
})

// onMounted(() => {
//     console.log(tasks.value.backlog[0])
// })

</script>

<template>
    <header class="header">
        <nav class="header-nav">
            <a href="https://webpractik.ru/" target="_blank" title="Сайт вебпрактик"><img class="header-nav__logo" src="../../public/logo.svg" alt="webpractice logo"></a>
            <h3 class="header-nav__exit">Выйти</h3>
        </nav>
    </header>
    <main>
        <div class="card-container" >
            <Card title="Беклог" :tasks="tasks.backlog"/>
            <Card title="В процессе" :tasks="tasks.proccessing"/>
            <Card title="Завершено" :tasks="tasks.finished"/>
        </div>
    </main>

</template>

<style scoped>

.header-nav {
    margin-top: 35px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-left: 200px;
    margin-right: 200px;
}
.header-nav__logo{
    width: 200px;
    user-select: none;
    pointer-events: none;
    -moz-user-select: none;
     -webkit-user-select: none;
      -ms-user-select: none;
       user-select: none;
}

.header-nav__logo:hover{
    cursor:pointer;
}

.header-nav__exit {
    font-size: 18px;
    user-select: none;
}
.header-nav__exit:hover {
    cursor: pointer;
}

.card-container {
    display: flex;
    justify-content: space-between;
    margin-left: 200px;
    margin-right: 200px;
}
main {
    margin-top: 20px;
}
</style>