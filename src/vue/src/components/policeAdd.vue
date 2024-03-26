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
                <el-menu-item>区块链房屋租赁--公安局用户信息添加</el-menu-item>
            </el-menu>
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="8"></el-col>
        <el-col :span="8">
            <h1 style="text-align: center; padding: 20px;">公安局用户信息添加</h1>
            <el-form ref="form" :model="form">
                <el-form-item>
                <el-input v-model="form.id_card" style="width: 400px" placeholder="请输入身份证号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.name" style="width: 400px" placeholder="请输入姓名" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.age" style="width: 400px" placeholder="请输入年龄" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.addr" style="width: 400px" placeholder="请输入住址" />
                </el-form-item>

                <!-- <el-form-item>
                    <el-upload
                        class="upload-demo"
                        drag
                        action="填写上传的url"
                        accept=".jpg, .jpeg, .png"
                        multiple>
                        <i class="el-icon-upload"></i>
                        <div class="el-upload_text">
                            将房屋照片拖到此处，或<em>点击上传</em>
                        </div>
                        <template #tip>
                            <div class="el-upload_tip">
                                只能上传 jpg/jpeg/png 文件，且不超过500KB
                            </div>
                        </template>
                    </el-upload>
                </el-form-item> -->

                <el-form-item>
                    <el-button type="primary" @click="onSubmit">提交</el-button>
                </el-form-item>
            </el-form>
        </el-col>
        <el-col :span="8"></el-col>
    </el-row>
</template>

<script>
import qs from 'qs'

export default {
    name: "Police",
    data: function() {
        return {
            form: {
                id_card: "",
                name: "",
                age: "",
                addr: ""
            }
        }
    },

    methods: {
        onSubmit() {
            this.$axios.post("/police/set",qs.stringify({
                id_card: this.form.id_card,
                name: this.form.name,
                age: this.form.age,
                addr: this.form.addr
            })).then(response => {
                alert('注册成功')
                this.$router.push("/policeQuery")
            }).catch(error => {
                if (error.response.status == 400) {
                    alert('该用户已存在')
                } else {
                    alert('网络错误')
                }
            })
        }
    }

}
</script>