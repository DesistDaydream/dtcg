<script setup>
import { ref } from 'vue'
import { RouterLink, RouterView } from 'vue-router'

// 生成顶部路由导航栏
let menuList = ref([
  { menu: "首页", link: "/" },
  { menu: "卡组价格", link: "/card" },
])

// 判断导航栏中哪个按钮被激活了
// 点击某一项的时候拿到相应的index 。然后把这个index赋值给isActiveNum 。再用 RouterLink 里面的 index 去对比。如果一样的话就加上 isActive 类 否则类名为空。
let isActiveNum = ref(0)
let btnMenu = (index) => {
  isActiveNum.value = index
}
</script>

<template>
  <header>
    <div>
      <nav>
        <ul @click="() => btnMenu(index)" v-for="(item, index) in menuList">
          <li>
            <RouterLink :class="isActiveNum == index ? 'isActive' : ''" :to="item.link">{{ item.menu }}
            </RouterLink>
          </li>
        </ul>
      </nav>
    </div>
  </header>

  <RouterView />
</template>

<style scoped>
.isActive {
  color: #f23051;
}

header {
  /* margin: 50px auto; */
  height: 60px;
  background-color: #690;
}

ul {
  list-style: none;
  display: inline;
  margin-left: 50px;
}

li {
  display: inline;
}

/* 链接的样式 */
a {
  line-height: 60px;
  color: #fff;
  text-decoration: none;
  padding: 10px 10px;
}

/* 当鼠标放在链接上时 */
a:hover {
  background-color: #060;
}
</style>
