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
                <el-menu-item>区块链房屋租赁--房管局房屋信息管理</el-menu-item>
            </el-menu>
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="6"></el-col>
        <el-col :span="12">
            <el-link href="/houseAdd">
                <el-button style="margin-left: -600px" type="primary" icon="el-icon-plus">添加</el-button>
            </el-link>
            <div style="margin-top: 15px;">
                <el-input
                placeholder="请输入房屋编号"
                    v-model="id_house"
                    class="input-with-select"
                    prefix-icon="el-icon-search">
                    <template #append>
                        <el-button type="primary" icon="el-icon-search" @click="queryHouse">搜索</el-button>
                    </template>
                </el-input>
            </div>
            <el-table :data="houseData" style="width: 100%">
                <el-table-column prop="id_house" label="房屋编号" width="180" />
                <el-table-column prop="owner" label="房屋所有人" width="180" />
                <el-table-column prop="addr_house" label="房屋地址"/>
                <el-table-column prop="type_house" label="房屋类型"/>
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
            id_house: "",
            houseData: []
        }
    },
    methods: {
        queryHouse() {
            this.$axios.get("/housemng/query",{
                params: {
                    id_house: this.id_house
                }
            }).then(response => {
                console.log(response)
                this.houseData[0] = response.data.data
                this.houseData[0].id_house = this.id_house
            }).catch(error => {
            if (error.response.status == 400) {
                console.log(error)
                alert('该房屋编号在房管局系统中不存在，请先添加')
            } else {
                alert('网络错误')
            }
            })
        }
    }
}
</script>