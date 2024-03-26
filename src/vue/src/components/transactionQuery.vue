<template>
    <el-container style="margin-top: -60px">
        <el-header style="padding: 0;">
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="6"></el-col>
        <el-col :span="12">
            <el-link href="/user/transactionAdd">
                <el-button style="margin-left: -600px" type="primary" icon="el-icon-plus">添加</el-button>
            </el-link>
            <div style="margin-top: 15px;">
                <el-input
                placeholder="请输入订单编号"
                    v-model="id_order"
                    class="input-with-select"
                    prefix-icon="el-icon-search">
                </el-input>
            </div>
            <div style="margin-top: 15px;">
                <el-input
                    placeholder="请输入交易期数"
                    v-model="period"
                    class="input-with-select"
                    prefix-icon="el-icon-search">
                    <template #append>
                        <el-button type="primary" icon="el-icon-search" @click="queryTransaction">搜索</el-button>
                    </template>
                </el-input>
            </div>
            <el-table :data="transactionData" style="width: 100%">
                <el-table-column prop="id_order" label="订单编号" width="180" />
                <el-table-column prop="period" label="交易期数" width="180" />
                <el-table-column prop="id_landlord" label="房东身份证号" width="180" />
                <el-table-column prop="id_tenant" label="租户身份证号" width="180" />
                <el-table-column prop="name_landlord" label="房东姓名" width="180" />
                <el-table-column prop="name_tenant" label="租户姓名" width="180" />
                <el-table-column prop="id_house" label="房屋编号" width="180" />
                <el-table-column prop="rent" label="租金" width="180" />
                <el-table-column prop="time_tx" label="交易时间" width="180" />
                <el-table-column prop="id_contract" label="合同编号" width="180" />
                <el-table-column prop="desc" label="备注" width="180" />
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
            id_order: "",
            period: "",
            transactionData: []
        }
    },
    methods: {
        queryTransaction() {
            this.$axios.get("/transaction/query",{
                params: {
                    id_order: this.id_order,
                    period: this.period
                }
            }).then(response => {
                console.log(response)
                this.transactionData[0] = response.data.data
                this.transactionData[0].id_order = this.id_order
                this.transactionData[0].period = this.period
            }).catch(error => {
                if (error.response.status == 400) {
                    console.log(error)
                    alert('该合同信息在系统中不存在，请先添加')
                } else {
                    alert('网络错误')
                }
            })
        }
    }
  }
  </script>