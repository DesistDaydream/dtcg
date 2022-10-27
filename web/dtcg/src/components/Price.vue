

<script setup>
import { ref } from 'vue'

// 初始化表单中的变量，设为空
let message = ref('')
let resp = ref({})

function commit(params) {
    console.log(message)
    let xhr = new XMLHttpRequest()
    xhr.open("POST", "http://localhost:2205/api/v1/deck/price")
    xhr.send(
        JSON.stringify(
            ({
                deck: message.value,
                envir: "chs",
            })
        )
    )

    xhr.onload = function () {
        resp.value = JSON.parse(xhr.responseText)
        console.log(resp.value)
    }
}

</script>

<template>
    <DirectivesRoute />

    卡组：<textarea v-model="message" placeholder="输入内容" cols="45" rows="5"></textarea>
    <button @click="commit">提交</button>
    <p>卡组：{{ message }}</p>

    <table border="1">
        <thead>
            <tr>
                <th>最低价</th>
                <th>集换价</th>
            </tr>
        </thead>
        <tr>
            <th>{{ resp.min_price }}</th>
            <th>{{ resp.avg_price }}</th>
        </tr>
    </table>

    <table border="1">
        <thead>
            <tr>
                <th>名称</th>
                <th>数量</th>
                <th>编号</th>
                <th>最低价</th>
                <th>集换价</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="(item, _) in resp.data">
                <th>{{ item.sc_name }}</th>
                <th>{{ item.count }}</th>
                <th>{{ item.serial }}</th>
                <th>{{ item.min_price }}</th>
                <th>{{ item.avg_price }}</th>
            </tr>
        </tbody>
    </table>
</template>

<style>

</style>