<script setup>

import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import PopUp from '../components/PopUp.vue';
import axios from 'axios'
import { user } from '@/stores/counter';

const router = useRouter();


const error = ref(false)
function goOnTheMainPage(){
    if(loginUserData.value.login > 1 || loginUserData.value.password.length > 1) {
        axios.post('http://193.188.23.216/api/v1/token/', {
            username: loginUserData.value.login,
            password: loginUserData.value.password
        }) .then(res => {
            localStorage.setItem('access', res.data.access)
            localStorage.setItem('refresh', res.data.refresh)
            console.log(res)
            router.push('/')
        }) .catch(err => {
            console.log(err)
            error.value = true
        })
        error.value = false
    } else {
        error.value = true
    }
}
const loginUserData = ref({
    login: '',
    password: '',
})

</script>

<template>
    <PopUp :isActive="error"/>
    <main>
        <section class="LogIn">
            <div class="LogIn-container">
                <h3>Вход</h3>
                <form action="" class="LogIn-form">
                    <input class="form__input" type="text" placeholder="Логин" v-model="loginUserData.login"/>
                    <input class="form__input" type="password" placeholder="Пароль" v-model="loginUserData.password"/>
                    <input class="form__submit" type="submit" value="Войти" @click.prevent="goOnTheMainPage()"/>
                </form>
                <RouterLink to="/signup">Нет аккаунта? Зарегистрируйтесь!</RouterLink>
            </div>
        </section>
    </main>

</template>

<style scoped>
.form__submit{
    width: 100%;
    transition-duration: 0.3s;
}
.form__submit:hover{
    cursor:pointer;
    background-color: #914ee2;
    transition-duration: 0.3s;
}
.form__input{
    transition-duration: 0.3s;
}
.form__input:hover{
    border: 1px solid #b291db;
    transition-duration: 0.1s;
}
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
