<script setup>
import { ref, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import  axios  from 'axios';
import { checkCookie, getCookie } from '../../modules/module';
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const AllReduce = ref();

onMounted(() => {
    checkCookie();
    const userData = {
        username: getCookie(),
    }
    axios.post("http://3.27.122.207:8080/foodlossreduce/all", userData)
    .then((response) => {
        AllReduce.value = response.data.allreduce;
    })
    .catch((error) => {
        console.log(error);
    })
});

</script>
<template>
    <h1>みんなのフードロス削減度</h1>
    <div class="rounded-full border-4 border-rakuten bg-gradient-to-br from-amber-200 via-amber-50 to-amber-200 w-80 h-80 relative mx-auto mt-6">
        <div class="text-7xl font-black text-center text-rakuten mt-28">{{ AllReduce }}</div>
    </div>
    <HomeBtn />
</template>
<style>
.AllReduce{
    text-align: center;
    font-size: 100px;
}
.HomeBtn{
    position: absolute;
    bottom: 20px;
}
</style>