<script setup>
import Card from '../components/Card.vue'
import {onMounted, ref, computed} from 'vue'

const tasks = ref({
    backlog:[
    {title: "Задача один", executor:"Коля Кравченко", date:'2021-01-01'},
    {title: "Задача один", executor:"Коля Кравченко", date:'2021-02-01'},
    {title: "Задача один", executor:"", date: '2022-03-04'},
    ],
    proccessing:[
    {title: "Задача два", executor:"Ваня усач", date: '2021-01-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2021-01-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2024-01-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2021-10-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2023-01-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2021-12-05'},
    {title: "Задача один", executor:"Коля Кравченко", date: '2021-01-09'},
    ],
    finished:[
    {title: "Задача три", executor:"Ваня усыч", date: '2021-01-05'}
    ],
})

// onMounted(() => {
//     console.log(tasks.value.backlog[0])
// })

const searchInput = document.querySelector('.search') 
 
const queryVal = ref('')

    function search(e) { 
        const tasksOnPage = document.querySelectorAll('.container-for-task')

        console.log(queryVal.value)
        console.log(tasksOnPage)
        Array.from(tasksOnPage).forEach(task => { 
            const taskText = task.textContent.toLowerCase(); 
            console.log(task.textContent.toLowerCase())
            if (taskText.includes(queryVal.value.toLowerCase())) { 
                task.style.display = 'block' 
            } else { 
                task.style.display = 'none' 
            } 
        }) 
    }
    const selectedOption = ref('')

</script>

<template>
    <header class="header">
        <nav class="header-nav">
            <a href="https://webpractik.ru/" target="_blank" title="Сайт вебпрактик"><img class="header-nav__logo" src="../../public/logo.svg" alt="webpractice logo"></a>
            <div>
            <select name="filter" id="filter" v-model="selectedOption" @onChange="filter()">
                <option value="default">По умолчанию</option>
                <option value="dateUp">По возрастанию</option>
                <option value="dateDown">По убыванию</option>
            </select>
            {{ selectedOption }}
        </div>
            <div><input @input="search" v-model="queryVal" class="search" type="text" placeholder="Найти"></div>
            <p>{{ searchInput }}</p>
            <h3 class="header-nav__exit">Выйти</h3>
        </nav>
    </header>
    <main>
        <div class="card-container" >
            <Card title="Беклог" :tasks="tasks.backlog" :selectedOption="selectedOption"/>
            <Card title="В процессе" :tasks="tasks.proccessing" :selectedOption="selectedOption"/>
            <Card title="Завершено" :tasks="tasks.finished" :selectedOption="selectedOption"/>
        </div>
    </main>

</template>

<style scoped>
#filter{
    height: 30px;
    border-radius: 30px;
    border: 1px solid #d1d1d1;
    padding-left: 15px;
}
.search{
    height: 30px;
    border-radius: 30px;
    border: 1px solid #d1d1d1;
    padding-left: 15px;
    width: 250px;

}
.search::placeholder{
    font-size: 16px;
}
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
    color: #555555;
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