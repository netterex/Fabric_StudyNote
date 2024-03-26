<template>
  <el-row>
    <el-col :span="8"></el-col>
    <el-col :span="8">
      <h1 style="text-align: center; padding: 10px;">区块链房屋租赁--登录</h1>
      <el-form ref="form" :model="form">
        <el-form-item>
          <el-radio-group v-model="form.role">
            <el-radio label="4">个人</el-radio>
            <el-radio label="1">公安局</el-radio>
            <el-radio label="2">房管局</el-radio>
            <el-radio label="3">征信中心</el-radio>
          </el-radio-group>
        </el-form-item>

        <el-form-item>
          <el-input v-model="form.username" style="width: 400px" placeholder="请输入用户名" />
        </el-form-item>

        <el-form-item>
          <el-input v-model="form.pwd" type="password" style="width: 400px" placeholder="请输入密码" show-password/>
        </el-form-item>

        <el-form-item>
          <el-button type="primary" @click="onSubmit">登录</el-button>
        </el-form-item>

        <div>
          <el-link type="info" style="float: right;" href="/register">没有账号？立即注册</el-link>
        </div>
      </el-form>
    </el-col>
  </el-row>
</template>

<script>
import qs from 'qs'

export default {
  name: "Login",
  data: function() {
    return {
      form: {
        role: "4",
        username: "",
        pwd: ""
      }
    }
  },

  methods: {
    onSubmit() {
      this.$axios.post("/user/login",qs.stringify({
        username: this.form.username,
        role: this.form.role,
        pwd: this.form.pwd
      })).then(response => {
        console.log(response)
        const role = response.data.data.role
        if (role == 4) {
          this.$router.push("/user/contractQuery")
        } else if (role == 1) {
          this.$router.push("/policeQuery")
        } else if (role == 2) {
          this.$router.push("/houseQuery")
        } else if (role == 3) {
          this.$router.push("/creditQuery")
        }

      }).catch(error => {
        if (error.response.status == 400) {
          alert('该用户不存在')
        } else {
          alert('网络错误')
        }
      })
    }
  }
}
</script>

<style>

</style>