<script setup>
import { ref, onMounted, nextTick } from 'vue';
import router from '../router';
import { useRouter } from 'vue-router';
import  axios  from 'axios';
import VueCookie from "vue-cookie";
import { checkCookie, getCookie } from '../modules/module';
import LoadingAnime from '@/components/LoadingAnime.vue';

//ダミーデータ
const Character = ref('');
const Level = ref('');
const Exp = ref('');

const LoadingFlg = ref(true);

const GetData = async () => {
    const userData = {
        username: getCookie(),
    };
    await axios.post("http://3.27.122.207:8080/home", userData)
        .then(response => {
            Character.value = "data:image/gif;base64," + response.data.character;
            Level.value = response.data.level
            Exp.value = response.data.exp;
        })
        .catch(error => {
            console.log("faild", error);
        });
    LoadingFlg.value = false;
};

onMounted(() => {
    checkCookie();
    //apiの処理
    GetData();
});

const Router = useRouter();

const OpenMenuFlg = ref(false);

const MenuItems = ref([
    { name: 'あなたのフードロス貢献度', id: '/foodlossreduce/person'},
    { name: 'みんなのフードロス貢献度', id: '/foodlossreduce/all'},
    { name: '履歴', id: '/history'},
    { name: 'レシート送信', id: '/receipt'},
    { name: 'ログアウト', id: '/login'}
]);

const ClickedMenu = (targetIndex) => {
    if (MenuItems.value[targetIndex].id === "/login") {
        deleteCookie();
    }
    Router.push(MenuItems.value[targetIndex].id);
}

const deleteCookie = () => {
    VueCookie.delete("userName");
}

</script>
<template>
    <h1>HOME</h1>
    <div class="flex justify-center border-4 border-rakuten rounded-full bg-gradient-to-bl from-amber-200 via-amber-50 to-amber-200 mx-2 mt-4">
        <div class="w-16 h-16 border-4 border-amber-500 bg-white p-2 m-2 rounded-full font-bold drop-shadow text-center text-4xl text-rakuten">{{ Level }}</div>
        <p class="text-4xl m-auto text-center font-bold text-rakuten drop-shadow">{{ 'Exp. ' + Exp }}</p>
    </div>
    <div v-if="Character" class="flex">
        <img :src="Character" class="w-80 h-80 mx-auto mt-6" />
    </div>

        <button v-if="OpenMenuFlg==false" class="absolute inset-x-8 bottom-5 bg-rakuten p-3 rounded-full font-bold text-white" @click="OpenMenuFlg = !OpenMenuFlg">メニュー</button>
    <div v-if="OpenMenuFlg" class="absolute z-0 inset-x-8 mx-auto bottom-5 text-2xl bg-white border border-rakuten rounded-xl">
        <div v-for="(item, index) in MenuItems" :key="index">
            <div @click="ClickedMenu(index)" class="text-center border-b">{{ item.name }}</div>
        </div>
        <div @click="OpenMenuFlg=false" class="text-center border-t border-rakuten">閉じる</div>
    </div>
    <LoadingAnime v-show="LoadingFlg" />
</template>
<style>
.menubtn{
    position: absolute;
    bottom: 20px;
    right: 20px;
}
.menu{
    position: absolute;
    bottom: 40px;
    right: 20px;
}
</style>