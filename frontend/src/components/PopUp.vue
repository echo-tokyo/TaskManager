<script setup>
import { ref, defineProps } from 'vue'

const props = defineProps({
    isActive: Boolean,
})
</script>

<template>
    <div class="pop-up" v-if="props.isActive || isActive">
        <div class="pop-up-text">
            <h3>Ошибка входа</h3>
            <p>Проверьте вводимые данные🤓</p>
        </div>

    </div>

</template>

<style scoped>

@keyframes fall {
  0% {
    right: 0px; /* Начальная позиция */
  }
  100% {
    right: 4vw; /* Конечная позиция */
  }
}
.pop-up-text{
    display: flex;
    flex-direction: column;
    gap: 10px;
    color:#fff;
}
.pop-up{
    padding-left: 2vw;
    padding-top: 4vh;
    padding-right: 2vw;
    top: 3vh;
    display: flex;
    justify-content: space-between;
    position: absolute;

    background-color: #7927E0;
    width: 20vw;
    height: 11vh;
    border-radius: 15px;
    border-top-right-radius: 0px;
    animation: fall 0.3s forwards; /* Запускаем анимацию */
}
.pop-up p {
    color: #ffffff;
}
.pop-up svg:hover{
    cursor: pointer
}
</style>