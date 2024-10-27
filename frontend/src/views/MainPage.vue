<script setup>
import Card from '../components/Card.vue'
import {onMounted, ref, computed} from 'vue'
import { useRouter } from 'vue-router';
import axios from 'axios'

const router = useRouter();

axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('access')

const rowTasks = ref({
    backlog: [],
    proccessing: [],
    finished: []
})

function getTasks(){
    console.log('get tasks')
    axios.get('http://193.188.23.216/api/v1/tasks/')
        .then(res => {
            console.log(res)
            rowTasks.value.backlog = res.data.backlog
            rowTasks.value.proccessing = res.data.proccessing
            rowTasks.value.finished = res.data.finished
        })
        .catch(err => {
            console.log(err)
            checkAuthorization()
        })
}

onMounted(() => {
    checkAuthorization()
    getTasks()
})

function checkAuthorization() {
  axios
    .post('http://193.188.23.216/api/v1/token/verify/', {
      token: localStorage.getItem('access')
    })
    .then(function (response) {
      console.log(response)
    })
    .catch(function (error) {
      console.log(error)
      refreshAccessToken()

    })
}
function refreshAccessToken() {
  axios
    .post('http://193.188.23.216/api/v1/token/refresh/', {
      refresh: localStorage.getItem('refresh')
    })
    .then(function (response) {
      localStorage.setItem('access', response.data.access)
    })
    .catch(function (error) {
      console.log(error)
      router.push('/LogIn')
    })
}


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

    function search() { 
        const tasksTitles = document.querySelectorAll('.task-title')
        tasksTitles.forEach(el => {
            if (el.textContent.toLowerCase().includes(queryVal.value.toLowerCase())) {
                el.parentElement.parentElement.style.display = 'block'
            } else {
                el.parentElement.parentElement.style.display = 'none'
            }
        })
    }
    const selectedOption = ref('')
    const minMaxDate = ref({
        min: '',
        max: ''
    })
    function inputDateFrom(e){
        console.log(e.target.value)
        minMaxDate.value.min = e.target.value
    }
    function inputDateTo(e){
        console.log(e.target.value)
        minMaxDate.value.max = e.target.value
    }
</script>

<template>
    <header class="header">
        <nav class="header-nav">
            <a href="https://webpractik.ru/" target="_blank" title="Сайт вебпрактик"><img class="header-nav__logo" src="../../public/logo.svg" alt="webpractice logo"></a>
            <div class="header-nav__right">
                <div class="header-nav__date">
                    <div class="header-nav__item">
                        <p>С</p>
                        <input class="header-nav__date-input" @input="inputDateFrom" type="date" :max="minMaxDate.max"/>
                    </div>
                    <div class="header-nav__item">
                        <p>По</p>
                        <input class="header-nav__date-input" @input="inputDateTo" type="date" :min="minMaxDate.min"/>
                    </div>
                </div>
                <div><input @input="search" v-model="queryVal" class="search" type="text" placeholder="Найти"></div>
                <h3 class="header-nav__exit">Выйти</h3>
            </div>
        </nav>
    </header>
    <main>
        <div class="card-container" >
            <Card :getRows="getTasks" title="Беклог" :tasks="rowTasks.backlog" :selectedOption="selectedOption"/>
            <Card :getRows="getTasks" title="В процессе" :tasks="rowTasks.proccessing" :selectedOption="selectedOption"/>
            <Card :getRows="getTasks" title="Завершено" :tasks="rowTasks.finished" :selectedOption="selectedOption"/>
        </div>
    </main>

</template>

<style scoped>
.header-nav__date{
    display: flex;
    align-items: center;
    gap: 10px;
}
input[type='date']::-webkit-datetime-edit {
    color: #566369;
}
input[type='date']::-webkit-calendar-picker-indicator {
    cursor: pointer;
    padding-right: 20px;
}
.header-nav__right{
    display: flex;
    gap: 20px;
    align-items: center;
}
.header-nav__item{
    display: flex;
    align-items: center;
    gap: 5px;
}
.header-nav__date input{
    height: 30px;
    border-radius: 30px;
    border: 1px solid #d1d1d1;
    padding-left: 20px;
}
#filter{
    height: 30px;
    border-radius: 30px;
    border: 1px solid #d1d1d1;
    padding-left: 20px;
}
.search{
    height: 30px;
    border-radius: 30px;
    border: 1px solid #d1d1d1;
    padding-left: 20px;
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
    margin-left: 230px;
    margin-right: 230px;
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
    margin-left: 230px;
    margin-right: 230px;
}
main {
    margin-top: 20px;
}
</style>