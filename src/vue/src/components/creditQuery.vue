<template>
    <el-container style="margin-top: -60px">
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
                <el-menu-item>区块链房屋租赁--征信中心信息管理</el-menu-item>
            </el-menu>
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="6"></el-col>
        <el-col :span="12">
            <el-link href="/creditAdd">
                <el-button style="margin-left: -600px" type="primary" icon="el-icon-plus">添加</el-button>
            </el-link>
            <div style="margin-top: 15px;">
                <el-input
                placeholder="请输入身份证号"
                    v-model="id_card"
                    class="input-with-select"
                    prefix-icon="el-icon-search">
                    <template #append>
                        <el-button type="primary" icon="el-icon-search" @click="queryCredit">搜索</el-button>
                    </template>
                </el-input>
            </div>
            <el-table :data="creditData" style="width: 100%">
                <el-table-column prop="id_card" label="身份证号" width="180" />
                <el-table-column prop="creditLevel" label="征信级别" width="180" />
            </el-table>
        </el-col>
        <el-col :span="6"></el-col>
    </el-row>
</template>

<script>
export default {
    name: "Index",
    data: function() {
        return {
            id_card: "",
            creditData: []
        }
    },
    methods: {
        queryCredit() {
            this.$axios.get("/credit/query",{
                params: {
                    id_card: this.id_card
                }
            }).then(response => {
                console.log(response)
                this.creditData[0] = response.data.data
                this.creditData[0].id_card = this.id_card
            }).catch(error => {
                if (error.response.status == 400) {
                    console.log(error)
                    alert('该用户信息在征信中心系统中不存在，请先添加')
                } else {
                    alert('网络错误')
                }
            })
        }
    }
}
</script>