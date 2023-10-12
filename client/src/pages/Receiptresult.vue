<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie, getCookie } from '../modules/module';
import HomeBtn  from "../components/HomeBtn.vue";

const Router = useRouter();

const result = ref();

onMounted(() => {
    checkCookie();
    const userData = {
        username: getCookie(),
    };
    axios.post('http://3.27.122.207:8080/receiptresult', userData)
    .then((response) => {
        console.log(response.data.getexp)
        result.value = response.data.getexp;
    })
    .catch(error => {
        console.log(error);
    })
});

</script>
<template>
    <h1>送信完了</h1>
    
    <div class="rounded-3xl w-80 h-80 border-4 border-rakuten text-center mx-auto mt-6 bg-amber-100">
        <h2 class="font-bold text-2xl text-rakuten">獲得Exp</h2>
    <div class="text-9xl text-rakuten font-black m-10">{{ result }}</div>
    </div>
    <HomeBtn />
</template>
<style>
.Result{
    text-align: center;
    font-size: 100px;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>