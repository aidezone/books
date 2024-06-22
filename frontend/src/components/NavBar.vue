<template>
  <el-menu
    :default-active="activeIndex"
    class="el-menu-demo"
    mode="horizontal"
    :ellipsis="false"
    @select="handleSelect"
  >
    <el-menu-item index="0">
      图书管理系统
    </el-menu-item>
    <div class="flex-grow" />
    <el-menu-item index="1">图书查询</el-menu-item>
    <el-sub-menu index="2">
      <template #title>个人中心</template>
      <el-menu-item index="2-1">已借阅图书</el-menu-item>
      <el-menu-item index="2-2">退出登陆</el-menu-item>
    </el-sub-menu>
  </el-menu>
</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router';

const activeIndex = ref('1')
const router = useRouter();

const logout = () => {
  localStorage.setItem("token", null);
  router.push({name: "login"})
}

const handleSelect = (key: string, keyPath: string[]) => {
  console.log(key, keyPath)
  switch(key) {
  case "2-2":
    logout();
    break;
  }
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>