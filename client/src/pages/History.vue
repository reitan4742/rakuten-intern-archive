<script setup>
import { ref, onMounted } from "vue";
import axios from 'axios';
import { useRouter } from 'vue-router';
import { checkCookie, getCookie } from "../modules/module";
import HomeBtn  from "@/components/HomeBtn.vue";

const Router = useRouter();

const HistoryList = ref("")

onMounted(() => {
    checkCookie();
    const userData = {
        username: getCookie(),
    };
    axios.post("http://3.27.122.207:8080/history", userData)
    .then((response) => {
        formatCreatedAt(response.data.historylist);
        HistoryList.value = response.data.historylist;
    })
    .catch((error) => {
        console.log(error);
    })
});

const formatCreatedAt = (list) => {
    for (let i = 0; i < list.length; i ++) {
        list[i].createdat = list[i].createdat.slice(0,10);
    }
}
</script>

<template>
    <h1>履歴</h1>
    <table class="relative mx-auto border-4 border-rakuten mt-6 table-fixed w-80">
        <thead>
            <tr class="font-bold text-4xl bg-amber-200 text-center">
                <th class="">日付</th>
                <th>獲得Exp</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(item, index) in HistoryList" :key="index " class="text-xl font-normal">
                <td class="text-center" :class="[(index%2)==0?'bg-orange-100':'bg-amber-100']">{{ item.createdat }}</td>
                <td class="text-right" :class="[(index%2)==0?'bg-orange-50':'bg-amber-50']">{{ item.getexp }}</td>
            </tr>
        </tbody>
    </table>
    <HomeBtn />
</template>

<style>
</style>