<template>
    <el-container style="margin-top: -60px;">
        <el-header style="padding: 0;">
            <el-menu 
                :default-active="2"
                class="el-menu-demo"
                mode="horizontal"
                @select="handleSelect"
                background-color="#545c64"
                text-color="#fff"
                active-text-color="#ffd04b"
                router="true">
                <el-menu-item>区块链房屋租赁--个人中心</el-menu-item>
                <el-menu-item index="1" :route="{name: 'contractQuery'}">合同管理</el-menu-item>
                <el-menu-item index="2" :route="{name: 'transactionQuery'}">交易管理</el-menu-item>
            </el-menu>
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
</template>

<script>
export default {
    name: "Frame",
}
</script>