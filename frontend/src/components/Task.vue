<script setup>
import { defineProps, onMounted, ref } from 'vue'
import { user } from '../stores/counter.js'
import axios from 'axios'

axios.defaults.headers.common['Authorization'] = 'Bearer ' + localStorage.getItem('access')

const props = defineProps({
    task: Object,
})


const tasks = ref(props.task)

const store = user()
function setExecutor(user) {
    console.log(props.task.id)
    axios.post('http://193.188.23.216/api/v2/tasks/add-executor/', {taskId: props.task.id})
    .then(function (response) {
        console.log(props.task.executors)
        props.task.executors.push({fio: 'Вы'})
        console.log(props.task.executors)
    })
}   
function deleteTask(){
    axios.delete('http://193.188.23.216/api/v1/tasks/' + props.task.id)
        .then((res) => {
            console.log(res)
            const willRemoved = document.getElementById(props.task.id)
            willRemoved.remove()
        })
        .catch(error=>{
            alert('Ошибка при удалении')
            console.log(error)
        })
}

const executor = ref('Иван Иванович')

</script>

<template>
    
    <div class="container-for-task" :id="props.task.id">
        <div class="task-text" >
            <img @click.stop="deleteTask" id="cross" src="../../public/cross.svg" alt="delete task button">
            <p class="task-title">{{ props.task.title }}</p>
            <p class="task-executor" v-if="props.task.executors[0]"  v-for="executor in props.task.executors">{{ executor.fio }}</p>
            <p class="task-noexecutor" v-else @click.stop="setExecutor(user()), console.log(props.task.executors)">Взять задачу</p>
        </div>

    </div>

</template>

<style scoped>

#cross{
    position: absolute;
    right: 10px;
    width: 15px;
}
#cross:hover{
    height: 18px;
    transition-duration: 0.2s;
}
.container-for-task:hover{
    cursor:pointer;
    box-shadow: 0 2px 4px 0 rgba(201, 194, 194, 0.7);
    transition-duration: 0.1s;
}
.task-text{
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 45px;
}
.task-title{
    font-size: 18px;
    width:90%;
    white-space: nowrap; /* Запрет переноса текста */
    overflow: hidden; /* Скрытие переполненного текста */
    text-overflow: ellipsis; /* Добавление троеточия */
}
.task-executor {
    font-size: 13px;
    color: #566369;
}
.task-noexecutor {
    font-size: 13px;
    color: #7927E0;
    width: 80px;
}
.task-noexecutor:hover {
    cursor:pointer;
}
.container-for-task{
    position: relative;
    margin-top: 10px;
    border: 1px solid #D1D1D1;
    height: 50px;
    width: 15vw;
    border-radius: 20px;
    padding-left: 20px;
    padding-top: 10px;
    padding-right: 20px;
    padding-bottom: 10px;
    transition-duration: 0.3s;
}
</style>