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
                <el-menu-item>区块链房屋租赁--房管局房屋信息添加</el-menu-item>
            </el-menu>
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="8"></el-col>
        <el-col :span="8">
            <h1 style="text-align: center; padding: 20px;">房管局房屋信息添加</h1>
            <el-form ref="form" :model="form">
                <el-form-item>
                <el-input v-model="form.id_house" style="width: 400px" placeholder="请输入房屋编号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.owner" style="width: 400px" placeholder="请输入房屋所有人" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.addr_house" style="width: 400px" placeholder="请输入房屋地址" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.type_house" style="width: 400px" placeholder="请输入房屋类型" />
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
    name: "House",
    data: function() {
        return {
            form: {
                id_house: "",
                owner: "",
                addr_house: "",
                type_house: ""
            }
        }
    },

    methods: {
        onSubmit() {
            this.$axios.post("/housemng/set",qs.stringify({
                id_house: this.form.id_house,
                owner: this.form.owner,
                addr_house: this.form.addr_house,
                type_house: this.form.type_house
            })).then(response => {
                alert('添加成功')
                this.$router.push("/houseQuery")
            }).catch(error => {
                if (error.response.status == 400) {
                    alert('该房屋信息已存在')
                } else {
                    alert('网络错误')
                }
            })
        }
    }

}
</script>