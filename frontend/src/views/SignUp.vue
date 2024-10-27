<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import PopUp from '../components/PopUp.vue';
import axios from 'axios'
const router = useRouter();
const error = ref(false)
function goOnTheMainPage(){
    const fioRegex = /^[\u0400-\u04FF\s-]{1,50}$/;

const isFioValid = fioRegex.test(signupUserData.value.fio);
const isLoginValid = signupUserData.value.login.length <= 20 && signupUserData.value.login !== '';
const isPasswordValid = signupUserData.value.password.length <= 20 && signupUserData.value.password !== '';
const isRepeatPasswordValid = signupUserData.value.repeatPassword.length <= 20 && signupUserData.value.repeatPassword !== '';
const arePasswordsEqual = signupUserData.value.password === signupUserData.value.repeatPassword;

if (isFioValid && isLoginValid && isPasswordValid && isRepeatPasswordValid && arePasswordsEqual) {
    axios.post('http://193.188.23.216/api/v1/users/', {
        username: signupUserData.value.login,
        password: signupUserData.value.password,
        fio: signupUserData.value.fio,
        job_title: 'Не указано',
    })
    .then(res => {
        console.log(res)
        localStorage.setItem("access", res.data.access);
        localStorage.setItem("refresh", res.data.access);
        console.log(localStorage.getItem("access"))
        console.log(localStorage.getItem("refresh"))
    }) .catch(err => {
        error.value = true
        console.log(err)
    })
    error.value = false
    router.push('/')
} else {
    error.value = true
}
}

const signupUserData = ref({
    fio: '',
    login: '',
    password: '',
    repeatPassword: '',
})
</script>

<template>
    <PopUp :isActive="error"/>
    <main>
        <section class="SignUp">
            <div class="SignUp-container">
                <h3>Регистрация</h3>
                <form action="" class="SignUp-form">
                    <input class="form__input" type="text" placeholder="ФИО" v-model="signupUserData.fio"/>
                    <input class="form__input" type="text" placeholder="Логин" v-model="signupUserData.login"/>
                    <input class="form__input" type="password" placeholder="Пароль" v-model="signupUserData.password"/>
                    <input class="form__input" type="password" placeholder="Повторите пароль" v-model="signupUserData.repeatPassword"/>
                    <input class="form__submit" type="submit" value="Регистрация" @click.prevent="goOnTheMainPage()"/>
                    <RouterLink class="aue" to="/login">Есть аккаунт? Войдите в систему!</RouterLink>
                </form>
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

.SignUp-form {
    background-color: #fff;
    box-shadow: 5px;
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 20px;
}

.SignUp {
    display: flex;
    align-items: center;
    width: 100%;
    height: 95vh;
}

.SignUp-container {
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