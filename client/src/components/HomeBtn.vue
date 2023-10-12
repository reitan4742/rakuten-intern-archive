<script setup>
import { useRouter } from "vue-router";

const Router = useRouter();

</script>

<template>
    <button @click="Router.push('/home')" class="rounded-full absolute bottom-5 inset-x-8 m-auto border-4 border-rakuten p-3 bg-white font-medium">
        ホーム画面に戻る
    </button>
</template>

<style>

</style>