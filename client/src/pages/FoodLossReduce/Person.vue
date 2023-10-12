<script setup>
import { ref, onMounted } from 'vue';
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie, getCookie } from '../../modules/module';
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const ReduceScore = ref();

onMounted(() => {
    checkCookie();
    const userData = {
        username: getCookie(),
    };
    axios.post("http://3.27.122.207:8080/foodlossreduce/person", userData)
    .then(response => {
        ReduceScore.value = response.data.reduce_score;
    })
    .catch(error => {
        console.log(error);
    })
});

</script>
<template>
    <h1>あなたのフードロス削減度</h1>
    <div class="rounded-full border-4 border-rakuten bg-gradient-to-br from-amber-200 via-amber-50 to-amber-200 w-80 h-80 relative mx-auto my-6">
        <div class="text-7xl font-black text-center text-rakuten mt-28">{{ ReduceScore }}</div>
    </div>
    <HomeBtn />
</template>
<style>
</style>