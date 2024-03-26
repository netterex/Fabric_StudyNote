<template>
    <el-row>
        <el-col :span="8"></el-col>
        <el-col :span="8">
            <h1 style="text-align: center; padding: 20px;">合同信息添加</h1>
            <el-form ref="form" :model="form">
                <el-form-item>
                    <el-input v-model="form.id_tenant" class="input-with-select" placeholder="请输入身份证号">
                        <template #append>
                            <el-button icon="el-icon-search" type="primary" @click="policeVerify">认证</el-button>
                        </template>
                    </el-input>
                </el-form-item>

                <el-form-item>
                    <el-input v-model="form.id_house" class="input-with-select" placeholder="请输入房屋编号">
                        <template #append>
                            <el-button icon="el-icon-search" type="primary" @click="houseVerify">认证</el-button>
                        </template>
                    </el-input>
                </el-form-item>

                <el-form-item>
                    <el-input v-model="form.id_contract" class="input-with-select" placeholder="请输入合同编号" />
                </el-form-item>

                <el-form-item>
                    <el-upload
                        class="upload-demo"
                        drag
                        action="填写上传的url"
                        accept=".jpg, .jpeg, .png"
                        :before-upload="onSubmit">
                        <i class="el-icon-upload"></i>
                        <div class="el-upload_text">
                            将合同照片拖到此处，或<em>点击上传</em>
                        </div>
                        <template #tip>
                            <div class="el-upload_tip">
                                只能上传 jpg/jpeg/png 文件，且不超过500KB
                            </div>
                        </template>
                    </el-upload>
                </el-form-item>
            </el-form>
        </el-col>
        <el-col :span="8"></el-col>
    </el-row>
</template>

<script>
import qs from 'qs'

export default {
    name: "Contract",
    data: function() {
        return {
            form: {
                id_contract: "",
                id_tenant: "",
                id_house: ""
            },
            policeSuccess: false,
            houseSuccess: false
        }
    },

    methods: {
        onSubmit(file) {
            if (this.policeSuccess && this.houseSuccess) {
                const formData = new FormData()
                formData.append("id_contract", this.form.id_contract)
                formData.append("id_house", this.form.id_house)
                formData.append("id_tenant", this.form.id_tenant)
                formData.append("img_contract", file)
                this.$axios.post("/contract/set", formData).then(response => {
                    alert('添加成功')
                    this.$router.push("/user/contractQuery")
                }).catch(error => {
                    if (error.response.status == 400) {
                        alert('该合同编号已存在')
                    } else {
                        alert('网络错误')
                    }
                })
            } else {
                alert('认证未通过，请先通过认证')
            }
        },
        policeVerify() {
            this.$axios.get("/police/query",{
                params: {
                    id_card: this.form.id_tenant
                }
            }).then(response => {
                console.log(response)
                if (response.status == 200) {
                    alert('认证成功')
                    this.policeSuccess = true
                }
            }).catch(error => {
            if (error.response.status == 400) {
                console.log(error)
                alert('认证失败，该身份证号在公安局系统中不存在，请先注册')
            } else {
                alert('网络错误')
            }
            })
        },
        houseVerify() {
            this.$axios.get("/housemng/query",{
                params: {
                    id_house: this.form.id_house
                }
            }).then(response => {
                console.log(response)
                if (response.status == 200) {
                    alert('认证成功')
                    this.houseSuccess = true
                }
            }).catch(error => {
            if (error.response.status == 400) {
                console.log(error)
                alert('认证失败，该房屋编号在房管局系统中不存在，请先添加')
            } else {
                alert('网络错误')
            }
            })
        }
    }

}
</script>