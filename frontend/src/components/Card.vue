<script setup>
import { defineProps, ref, onMounted, computed, watch } from 'vue'
import draggable from "vuedraggable"
import Task from '../components/Task.vue'
import axios from 'axios';
const props = defineProps({
    getRows: Function,
    required: true,
    title: String,
    tasks: Object,
    selectedOption: String,
})
function log(event) {
    let res = props.tasks
    if (event.added) {
        if (props.title == 'Беклог') {
            event.added.element.status = 'backlog'
            axios.patch('http://193.188.23.216/api/v1/tasks/' + event.added.element.id, event.added.element)
                .then(function (response) {
                    console.log(response)
                })
                .catch(function (error) {
                    console.log(error)
                })
        } else if (props.title == 'В процессе') {
            event.added.element.status = 'proccessing'
            axios.patch('http://193.188.23.216/api/v1/tasks/' + event.added.element.id, event.added.element)
        } else if (props.title == 'Завершено') {
            event.added.element.status = 'finished'
            axios.patch('http://193.188.23.216/api/v1/tasks/' + event.added.element.id, event.added.element)
        }


        // console.log(props.title)
        // console.log(event.added.element.id)
        // console.log(props.tasks) 
        // props.tasks[0].status
        // for(let i = 0; i < props.tasks.length; i++) {
        //     if(props.tasks[i].id !== event.added.element.id) {
        //         event.added.element.status = props.tasks[i].status;
        //         console.log(props.tasks[i].status)
        //         axios.patch('http://193.188.23.216/api/v1/tasks/' + event.added.element.id, event.added.element)
        //         .then(function (response) {
        //             console.log(response)
        //         })
        //         .catch(function (error) {
        //             console.log(error)
        //         })
        //         break
        //     }   else{
        //         continue;
        //     }
        // }
    }
}
watch(props.tasks, () => {
    console.log(props.tasks)
})

const nameOfTheTaskInModal = ref('')
function addTask() {
    // if(nameOfTheTaskInModal.value == ''){
    //     props.tasks.push({title: "Без названия", executor:"", date: Date.now()})
    // }   else {
    //     props.tasks.push({title: nameOfTheTaskInModal.value, executor:"", date: Date.now()})
    // }
    if (props.title == 'Беклог') {
        axios.post('http://193.188.23.216/api/v1/tasks/', {
            board_id: 1,
            title: nameOfTheTaskInModal.value,
            task_desc: "",
            executors:[],
            status: "backlog",
        })
            .then(function (response) {
                console.log(response)
                props.tasks.push(response.data)
                props.getRows()
            })
            .catch(function (error) {
                console.log(error)
            })
    } else if (props.title == 'В процессе') {
        axios.post('http://193.188.23.216/api/v1/tasks/', {
            board_id: 1,
            title: nameOfTheTaskInModal.value,
            task_desc: "",
            executors: [],
            status: "proccessing",
        })
            .then(function (response) {
                console.log(response)
                props.tasks.push(response.data)
                props.getRows()
            })
            .catch(function (error) {
                console.log(error)
            })
    } else if (props.title == 'Завершено') {
        axios.post('http://193.188.23.216/api/v1/tasks/', {
            board_id: 1,
            title: nameOfTheTaskInModal.value,
            task_desc: "",
            executors: [],
            status: "finished",
        })
            .then(function (response) {
                console.log(response)
                // props.tasks.push({title: nameOfTheTaskInModal.value, executors:[], taskId: response.data.id, date: Date.now()})
                props.tasks.push(response.data)
                props.getRows()
            })
            .catch(function (error) {
                console.log(error)
            })
    }

}
const currentTask = ref()
function onTaskClicked(e) {
    isIframeActivated.value = true
    currentTask.value = e.srcElement.id
    console.log(currentTask.value)
}
const access = ref(localStorage.getItem('access'))

const isModalActivated = ref(false)
const isIframeActivated = ref(false)
</script>
<template>
    <div class="modal-container" v-if="isModalActivated || isIframeActivated" @click="isModalActivated = false, isIframeActivated = false"></div>
    <div class="iframe-wrapper" v-if="isIframeActivated">
        <iframe class="iframe" :src="'http://192.168.29.60:5173/modal/?task_id=' + currentTask +'&access='+  access" frameborder="0" ></iframe>
        <img src="../../public/cross.svg" alt="" @click="isIframeActivated = false">
    </div>
    <div class="modal" v-if="isModalActivated">
        <h2>Добавить задачу</h2>
        <input type="text" name="" id="" placeholder="Название задачи" v-model="nameOfTheTaskInModal">
        <input @click="addTask(), isModalActivated = false" type="button" value="Добавить">
    </div>
    <div class="card">
        <div class="card-top">
            <h3 class="card__title">{{ props.title }}</h3>
    
            <draggable class="container-for-tasks" :list="props.tasks" group="people" @change="log" :itemKey="props.title">
            <template #item="{ element: tasks }" tag="div">
                <Task @click="onTaskClicked($event)" :task="tasks" />
            </template>
        </draggable>
            <!-- <div :class="'container-for-tasks'">
            </div> -->
        </div>
        <div class="add-task" @click="isModalActivated = true">
            <p>Добавить задачу</p>
        </div>
    </div>
</template>

<style>
.iframe-wrapper{
    border-radius: 20px;
    left: 15vw;
    top: 10vh;
    position: absolute;
    background-color: #fff;
    width: 70vw;
    height: 80vh;
    z-index: 11;
}

.iframe {
    position: relative;
    height: 100%;
    width: 100%;
}
.iframe-wrapper img{
    position: absolute;
    right: 50px;
    top: 50px;
    width: 40px;
    height: 40px;
    cursor: pointer;
}
.card-top{
    display: flex;
    flex-direction: column;
    align-items: center;
    height: 100%;
    overflow: scroll;
}
.card-top::-webkit-scrollbar {
    /* position: absolute;
    width: 3px;
    height: 3px; */
    background-color: #f4f3f3;
    width: 0px;
    height: 0px;
}
.modal-container{
    display: flex;
    justify-content: center;
    align-items: center;
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 10;
    background-color: rgba(0, 0, 0, 0.3);
}

.modal {
    padding-top: 40px;
    padding-bottom: 20px;
    display: flex;
    align-items: center;
    justify-content: space-between;
    flex-direction: column;
    border-radius: 20px;
    left: 20vw;
    top: 10vh;
    position: absolute;
    background-color: #fff;
    width: 60vw;
    height: 55vh;
    z-index: 11;
}

.modal input[type="text"] {
    height: 60px;
    width: 85%;
    border: 1px solid #D9D9D9;
    border-radius: 15px;
    font-size: 16px;
    padding-left: 20px;
    margin-bottom: 20px;
}

.modal input[type="button"] {
    height: 60px;
    width: 20%;
    font-family: "proximaNova bold";
    color: white;
    border: 0;
    border-radius: 15px;
    font-size: 16px;
    margin-bottom: 20px;
    background-color: #7927E0;
    transition-duration: 0.3s;
}

.modal input[type="button"]:hover {
    cursor: pointer;
    background-color: #914ee2;
    transition-duration: 0.3s;
}

.add-task {
    cursor: pointer;
    display: flex;
    justify-content: center;
    align-items: center;
    /* margin-top: 10px; */
    /* width: calc(100% - 15vw); */
    width: calc(15vw + 40px);
    margin-right: 5px;
    height: 50px;
    background-color: #ffffff;
    border: 1px solid #D9D9D9;
    border-radius: 15px;
    font-size: 16px;
    transition-duration: 0.3s;
}

/* 
.add-task:hover{
>>>>>>> ed306a57c222588650b902c09d3aa5506c5babc9
    cursor: pointer;
    box-shadow: 0 4px 9px 0 rgba(201, 194, 194, 0.7);
    background-color: #7927E0;
    color: white;
    transition-duration: 0.3s;
} */


.container-for-tasks:empty {
    overflow: hidden;
    overflow-x: hidden;

}

.container-for-tasks {
    display: flex;
    flex-direction: column;
    gap: 8px;
    overflow: scroll;
    overflow-x: hidden;
    height: 100%;
    padding-right: 5px;
}

::-webkit-scrollbar {
    /* position: absolute;
    width: 3px;
    height: 3px; */
    background-color: #f4f3f3;
    width: 2px;
    height: 2px;
}

::-webkit-scrollbar-thumb {
    border-radius: 9em;
    background-color: rgb(214, 214, 214);
}

/* ::-webkit-scrollbar-thumb:hover {
    background-color: #9e9e9e;
}  */

.card {
    gap: 10px;
    padding-top: 50px;
    padding-bottom: 50px;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 22vw;
    height: 65vh;
    justify-content: space-between;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 0 4px 9px 0 rgba(201, 194, 194, 0.7);
}

.card__title {
    margin-bottom: 20px;
    font-size: 16px;
    color: #7927E0;
}
</style>