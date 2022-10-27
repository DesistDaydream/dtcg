<script setup>
</script>

<template>
    <link rel="stylesheet" href="https://cdn.staticfile.org/font-awesome/4.7.0/css/font-awesome.css">

    <p>这里搜索卡牌</p>
    <form class="example" action="这里写在哪搜索">
        <input type="text" placeholder="搜索.." name="search">
        <button type="submit"><i class="fa fa-search"></i></button>
    </form>
</template>

<style scoped>
* {
    box-sizing: border-box;
}

/* 设置搜索框 */
form.example input[type=text] {
    padding: 10px;
    font-size: 17px;
    border: 1px solid grey;
    float: left;
    width: 80%;
    background: #f1f1f1;
}

/* 设置提交按钮 */
form.example button {
    float: left;
    width: 20%;
    padding: 10px;
    background: #2196F3;
    color: white;
    font-size: 17px;
    border: 1px solid grey;
    border-left: none;
    /* 防止双边框 */
    cursor: pointer;
}

form.example button:hover {
    background: #0b7dda;
}

/* 清除浮动 */
form.example::after {
    content: "";
    clear: both;
    display: table;
}
</style>