<script setup>

import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';

const router = useRouter();

function goOnTheMainPage(){
    router.push("/")
}
</script>

<template>
    <main>
        <section class="LogIn">
            <div class="LogIn-container">
                <h3>Вход</h3>
                <form action="" class="LogIn-form">
                    <input class="form__input" type="text" placeholder="Логин" />
                    <input class="form__input" type="password" placeholder="Пароль" />
                    <input class="form__submit" type="submit" value="Войти" @click="goOnTheMainPage()"/>
                </form>
                <RouterLink to="/signup">Нет аккаунта? Зарегистрируйтесь!</RouterLink>
            </div>
        </section>
    </main>

</template>

<style scoped>
body {
    height: 100vh;
    width: 100%;
}

a{
    color:rgb(61, 61, 61);
}

.LogIn-form {
    background-color: #fff;
    box-shadow: 5px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 20px;
}

.LogIn {
    display: flex;
    align-items: center;
    width: 100%;
    height: 95vh;
}

.LogIn-container {
    background-color: #fff;
    box-shadow: 0 4px 9px 0 rgba(201, 194, 194, 0.7);
    display: flex;
    flex-direction: column;
    margin: 0 auto;
    align-items: center;
    gap: 20px;
    width: 30vw;
    padding: 30px;
    border-radius: 15px;
}
</style>
