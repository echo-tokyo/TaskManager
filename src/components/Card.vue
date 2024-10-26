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

function addTask() {
    props.tasks.push({title: "Без названия", executor:"", date: Date.now()})
}
</script>
<template>
    <div class="card">
        <h3 class="card__title">{{ props.title }}</h3>

        <draggable         
        class="container-for-tasks"
        :list="props.tasks"
        group="people"
        @change="log"
        :itemKey="props.title" >
            <template  #item="{ element: tasks }" tag="div">
                <Task :task="tasks"/>
            </template>
        </draggable>
        <!-- <div :class="'container-for-tasks'">
        </div> -->
        <div class="add-task" @click="addTask()">
            <p>Добавить задачу</p>
        </div>
    </div>
</template>

<style>
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
    overflow: scroll;
    overflow-x: hidden;
    height: 71%;
    padding-right: 5px;
}

::-webkit-scrollbar {
    position: absolute;
    width: 4px;
    height: 8px;
    background-color: #D9D9D9;
}

::-webkit-scrollbar-thumb {
    border-radius: 9em;
    background-color: #474747;
}

::-webkit-scrollbar-thumb:hover {
    background-color: #005DFF;
    width: 20px;
}

.card {
    padding-top: 20px;
    display: flex;
    flex-direction: column;
    align-items: center;
    width: 22vw;
    height: 75vh;
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