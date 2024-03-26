<template>
    <el-container style="margin-top: -60px">
        <el-header style="padding: 0;">
        </el-header>
        <el-main>
            <router-view></router-view>
        </el-main>
    </el-container>
    <el-row>
        <el-col :span="8"></el-col>
        <el-col :span="8">
            <h1 style="text-align: center; padding: 20px;">交易信息添加</h1>
            <el-form ref="form" :model="form">
                <el-form-item>
                <el-input v-model="form.id_order" style="width: 400px" placeholder="请输入订单编号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.period" style="width: 400px" placeholder="请输入交易期数" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.id_landlord" style="width: 400px" placeholder="请输入房东身份证号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.id_tenant" style="width: 400px" placeholder="请输入租户身份证号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.name_landlord" style="width: 400px" placeholder="请输入房东姓名" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.name_tenant" style="width: 400px" placeholder="请输入租户姓名" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.id_house" style="width: 400px" placeholder="请输入房屋编号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.rent" style="width: 400px" placeholder="请输入租金" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.time_tx" style="width: 400px" placeholder="请输入交易时间" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.id_contract" style="width: 400px" placeholder="请输入合同编号" />
                </el-form-item>

                <el-form-item>
                <el-input v-model="form.desc" style="width: 400px" placeholder="请输入备注" />
                </el-form-item>

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
                id_order: "",
                period: "",
                id_landlord: "",
                id_tenant: "",
                name_landlord: "",
                name_tenant: "",
                id_house: "",
                rent: "",
                time_tx: "",
                id_contract: "",
                desc: ""
            }
        }
    },

    methods: {
        onSubmit() {
            this.$axios.post("/transaction/set",qs.stringify({
                id_order: this.form.id_order,
                period: this.form.period,
                id_landlord: this.form.id_landlord,
                id_tenant: this.form.id_tenant,
                name_landlord: this.form.name_landlord,
                name_tenant: this.form.name_tenant,
                id_house: this.form.id_house,
                rent: this.form.rent,
                time_tx: this.form.time_tx,
                id_contract: this.form.id_contract,
                desc: this.form.desc
            })).then(response => {
                alert('添加成功')
                this.$router.push("/user/transactionQuery")
            }).catch(error => {
                if (error.response.status == 400) {
                    alert('该交易已存在')
                } else {
                    alert('网络错误')
                }
            })
        }
    }

}
</script>