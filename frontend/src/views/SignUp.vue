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
        <section class="SignUp">
            <div class="SignUp-container">
                <h3>Регистрация</h3>
                <form action="" class="SignUp-form">
                    <input class="form__input" type="text" placeholder="ФИО" />
                    <input class="form__input" type="text" placeholder="Логин" />
                    <input class="form__input" type="password" placeholder="Пароль" />
                    <input class="form__input" type="password" placeholder="Повторите пароль" />
                    <input class="form__submit" type="submit" value="Регистрация" @click="goOnTheMainPage()"/>
                    <RouterLink class="aue" to="/login">Есть аккаунт? Войдите в систему!</RouterLink>
                </form>
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