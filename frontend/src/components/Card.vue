<script setup>
import { defineProps, ref, onMounted, computed, watch } from 'vue'
import draggable from "vuedraggable"
import Task from '../components/Task.vue'
const props = defineProps({
    title: String,
    tasks: Object,
    selectedOption: String,
})
function log(event) {
}

watch(props.tasks, () => {
    console.log(props.tasks)
})

const nameOfTheTaskInModal = ref('')
function addTask() {
    if(nameOfTheTaskInModal.value == ''){
        props.tasks.push({title: "Без названия", executor:"", date: Date.now()})
    }   else {
        props.tasks.push({title: nameOfTheTaskInModal.value, executor:"", date: Date.now()})
    }
}

function onTaskClicked(e){
    console.log(e)
}
const isModalActivated = ref(false)
const isIframeActivated = ref(false)
</script>
<template>
    <div class="modal-container" v-if="isModalActivated || isIframeActivated" @click="isModalActivated = false, isIframeActivated = false"></div>
    <iframe class="iframe" src="http://192.168.30.4:5173/" frameborder="0" v-if="isIframeActivated"></iframe>
    <div class="modal" v-if="isModalActivated">
        <h2>Добавить задачу</h2>
        <input type="text" name="" id="" placeholder="Название задачи" v-model="nameOfTheTaskInModal">
        <input @click="addTask(), isModalActivated = false" type="button" value="Добавить">
    </div>
    <div class="card">
        <h3 class="card__title">{{ props.title }}</h3>

        <draggable     
            
        class="container-for-tasks"
        :list="props.tasks"
        group="people"
        @change="log"
        :itemKey="props.title" >
            <template  #item="{ element: tasks }" tag="div">
                <Task @click="(event) => onTaskClicked(event), isIframeActivated = true" :task="tasks"/>
            </template>
        </draggable>
        <!-- <div :class="'container-for-tasks'">
        </div> -->
        <div class="add-task" @click="isModalActivated = true">
            <p>Добавить задачу</p>
        </div>
    </div>
</template>

<style>
.iframe {
    border-radius: 20px;
    left: 15vw;
    top: 10vh;
    position: absolute;
    background-color: #fff;
    width: 70vw;
    height: 80vh;
    z-index: 11;
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
.modal{
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
    cursor:pointer;
    background-color: #914ee2;
    transition-duration: 0.3s;
}
.add-task {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 10px;
    width: 85%;
    height: 50px;
    background-color: #ffffff;
    border: 1px solid #D9D9D9;
    border-radius: 15px;
    font-size: 16px;
    transition-duration: 0.3s;
}

.add-task:hover{
    cursor: pointer;
    box-shadow: 0 4px 9px 0 rgba(201, 194, 194, 0.7);
    transition-duration: 0.3s;
}


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
    position: absolute;
    width: 3px;
    height: 3px;
    background-color: #D9D9D9;
}

::-webkit-scrollbar-thumb {
    border-radius: 9em;
    background-color: #474747;
}

::-webkit-scrollbar-thumb:hover {
    background-color: #9e9e9e;
}

.card {
    padding-top: 20px;
    padding-bottom: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 22vw;
    height: 75vh;
    justify-content: space-between;
    background-color: #fff;
    border-radius: 20px;
    box-shadow: 0 4px 9px 0 rgba(201, 194, 194, 0.7);
}

.card__title {
    margin-bottom: 10px;
    font-size: 16px;
    color: #7927E0;
}
</style>