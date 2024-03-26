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
          <el-link href="/user/contractAdd">
              <el-button style="margin-left: -600px" type="primary" icon="el-icon-plus">添加</el-button>
          </el-link>
          <div style="margin-top: 15px;">
              <el-input
              placeholder="请输入合同编号"
                  v-model="id_contract"
                  class="input-with-select"
                  prefix-icon="el-icon-search">
                  <template #append>
                      <el-button type="primary" icon="el-icon-search" @click="queryContract">搜索</el-button>
                  </template>
              </el-input>
          </div>
          <el-table :data="contractData" style="width: 100%">
              <el-table-column prop="id_contract" label="合同编号" width="180" />
              <el-table-column prop="id_tenant" label="租户身份证号" width="180" />
              <el-table-column prop="id_house" label="房屋编号" width="180" />
              <el-table-column prop="hash_contractImg" label="合同图片hash" width="180" />
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
          id_contract: "",
          contractData: []
      }
  },
  methods: {
      queryContract() {
          this.$axios.get("/contract/query",{
              params: {
                id_contract: this.id_contract
              }
          }).then(response => {
              console.log(response)
              this.contractData[0] = response.data.data
              this.contractData[0].id_contract = this.id_contract
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