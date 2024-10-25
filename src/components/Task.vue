<script setup>
import { defineProps, onMounted } from 'vue'


const props = defineProps({
    task: Object,
})

onMounted(() => {
    console.log(props.task)
})
</script>

<template>
    <div class="container-for-task">
        <div class="task-text">
            <p class="task-title">{{ props.task.title }}</p>
            <p class="task-executor" v-if="props.task.executor">{{ props.task.executor }}</p>
            <p class="task-noexecutor" v-else>Взять задачу</p>
        </div>

    </div>

</template>

<style scoped>
.task-text{
    display: flex;
    flex-direction: column;
    justify-content: space-between;
    height: 45px;
}
.task-title{
    font-size: 18px;
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
    margin-top: 10px;
    border: 1px solid #D1D1D1;
    height: 50px;
    width: 15vw;
    border-radius: 20px;
    padding-left: 20px;
    padding-top: 10px;
    padding-right: 20px;
    padding-bottom: 10px;
}
</style>