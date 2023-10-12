<script setup>
import axios from 'axios';
import { ref } from 'vue';
import { useRouter } from 'vue-router';

const Router = useRouter();

const Name = ref('');
const Password1 = ref('');
const Password2 = ref('');
const passwordFlag = ref("");
const userNameFlag = ref("");

const register = () => {
    //登録する処理
    if (checkPassword() === true) {
        const userData = {
            username: Name.value,
            password: Password1.value,
        }
        axios.post("http://3.27.122.207:8080/signup", userData)
            .then(response => {
                if (response.data.result === true) {
                    Router.push('/login');
                } else {
                    userNameFlag.value = true;
                }
            })
            .catch(error => { 
                userNameFlag.value = true;
                console.error("faild!", error);
            });
    } else {
        passwordFlag.value = true;
    }
}

const login = () => {
    Router.push("/login");
}

const checkPassword = () => {
    if (Password1.value === Password2.value) {
        return true;
    }
    return false;
}

</script>
<template>
    <h1>Signup Page</h1>
    <div class="border-[#bf0000] w-80 h-auto border-2 rounded-lg text-xl login mx-auto mt-4 text-left flex justify-center">
        <div class="w-auto place-content-center h-auto">
            <form @submit.prevent="register">
                <div class="userName mb-3 mx-auto">
                    <p>ユーザ名</p>
                    <input required class="input border-2 mt-4" v-model="Name" type="text" placeholder="user name" />
                </div>
                <div class="password mb-3 mx-auto">
                    <p>パスワード</p>
                    <input required class="input border-2 mt-4" v-model="Password1" type="password" placeholder="password" />
                </div>
                 <div class="password mb-3 mx-auto">
                    <p>パスワード(再入力)</p>
                    <input required class="input border-2 mt-4" v-model="Password2" type="password" placeholder="password" />
                </div>
                <p v-if="passwordFlag === true" class="text-red-600 text-xs">パスワードが異なります</p>
                <p v-if="userNameFlag === true" class="text-red-600 text-xs">登録済みのユーザ名です</p>
                <div class="flex justify-end">
                    <button class="border-[#bf0000] bg-transparent hover:bg-[#bf0000] text-[#bf0000] font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded-lg mb-4 mt-4 float">register</button>
                </div>
            </form>
        </div>
    </div>
    <div class="flex justify-end mr-4 mt-4">
        <button class="border-[#bf0000] bg-transparent hover:bg-[#bf0000] text-[#bf0000] font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded-lg mb-4 mt-4 float" @click="login()">back to login page</button>
    </div>
</template>

<style>
</style>