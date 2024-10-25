<script setup>
import { defineProps, ref, onMounted } from 'vue'
import Task from '../components/Task.vue'
const props = defineProps({
    title: String,
    tasks: Object,
})
onMounted(()=> {
    console.log(props.tasks)
})
</script>
<template>
    <div class="card">
        <h3 class="card__title">{{ props.title }}</h3>
        <div class="container-for-tasks">
            <Task v-for="task in props.tasks" :task="task"/>
        </div>
    </div>
</template>

<style scoped>
.container-for-tasks{
    overflow: scroll;
    overflow-x:hidden;
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
}::-webkit-scrollbar-thumb:hover {
 background-color: #005DFF;
 width: 20px;}
    .card {
        padding-top: 20px;
        display: flex;
        flex-direction: column;
        align-items: center;
        width:22vw;
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