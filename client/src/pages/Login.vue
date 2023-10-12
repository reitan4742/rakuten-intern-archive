<script setup>
import axios from "axios";
import VueCookie from "vue-cookie";
import { calcExpirationDate } from "../modules/module.js";
import { useRouter } from "vue-router";
import { ref } from "vue";

const Router = useRouter();

const userName = ref("");
const password = ref("");
const wrongFlag = ref("");

const login = () => {
    const userData = {
        username: userName.value,
        password: password.value,
    }
    axios.post("http://3.27.122.207:8080/login", userData)
        .then(response => {
            if (response.data.result === true) {
                const exp = calcExpirationDate();
                VueCookie.set("userName", userName.value, {
                    expires: exp
                });
                Router.push("/home");
            } else {
                wrongFlag.value = true;
            }
        })
        .catch(error => {
            wrongFlag.value = true;
            console.error("faild!", error);
        });
}
const signup = () => {
    Router.push("/signup");
}

</script>

<template>
    <h1>Login Page</h1>
    <div class="border-[#bf0000] w-80 h-auto border-2 rounded-lg text-xl login mx-auto mt-4 text-left flex justify-center">
        <div class="w-auto place-content-center h-auto">
            <form @submit.prevent="login">
                <div class="userName mb-3 mx-auto">
                    <label for="userName">ユーザ名</label><br>
                    <input type="text" class="border-2 mt-4" id="userName" placeholder="user name" v-model="userName" required>
                </div>
                <div class="password mb-3 mx-auto">
                    <label for="password">パスワード</label><br>
                    <input type="password" class="border-2 mt-4" placeholder="password" id="password" v-model="password" required>
                </div>
                <p v-if="wrongFlag === true" class="text-red-600 text-xs">ユーザ名またはパスワードが間違っています</p>
                <div class="flex justify-end">
                    <!-- <button class="border-[#bf0000] bg-transparent hover:bg-blue-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded-lg mb-4 mt-4 float">login</button> -->
                    <button class="border-[#bf0000] bg-transparent hover:bg-[#bf0000] text-[#bf0000] font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded-lg mb-4 mt-4 float">login</button>
                </div>
            </form>
        </div>
    </div>
    <div class="flex justify-end mr-4 mt-4">
        <button class="border-[#bf0000] bg-transparent hover:bg-[#bf0000] text-[#bf0000] font-semibold hover:text-white py-2 px-4 border hover:border-transparent rounded-lg mb-4 mt-4 float" @click="signup()">signup</button>
    </div>
</template>

<style>
</style>
