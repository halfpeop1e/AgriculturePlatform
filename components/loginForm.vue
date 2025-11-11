<template>
    <el-card class="login-card" :body-style="{ padding: '20px' }">
        <h3 class="title">登录</h3>

        <el-form :model="form" :rules="rules" ref="formRef" label-position="top">
            <el-form-item label="邮箱" prop="useremail">
                <el-input v-model="form.useremail" placeholder="请输入邮箱" clearable class="h-10"/>
            </el-form-item>

            <el-form-item label="密码" prop="password">
                <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
            </el-form-item>

            <el-form-item>
                <div class="flex-row">
                    <el-checkbox v-model="form.remember">记住我</el-checkbox>
                    <a class="forgot" @click.prevent="onForgot">忘记密码?</a>
                </div>
            </el-form-item>

            <el-form-item>
                <el-button type="primary" :loading="loading" @click="onSubmit" style="width:100%">登录</el-button>
            </el-form-item>

            <div class="register">
                还没有账号？ <a @click.prevent="toRegister">注册</a>
            </div>
        </el-form>
        <el-divider>或者</el-divider>

        <div class="flex-col items-center">
            <el-button type="success" :loading="loading" @click="onWechatLogin" style="width:100%;margin: 0;" plain>
                <Wechat :size="20" class="mr-2"/>
                使用微信登录
            </el-button>
            <div class="m-4"></div>
            <el-button color="#303133" :loading="loading" @click="onWechatLogin" style="width:100%; margin: 0;" plain>
                <Apple :size="20" class="mr-2"/>
                使用Apple登录
            </el-button>
        </div>
    </el-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import Wechat from './icons/Wechat.vue';
import Apple from './icons/Apple.vue';
const emit = defineEmits<{
    (e: 'register'): void;
    (e: 'forgot'): void
}>()

const formRef = ref<any>(null)
const form = ref({ useremail: '', password: '', remember: false })
const loading = ref(false)

const rules = {
    useremail: [{ required: true, message: '请输入邮箱', trigger: 'blur' }],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
}

async function onSubmit() {
    if (!formRef.value) return
    formRef.value.validate(async (valid: boolean) => {
        if (!valid) return
        loading.value = true
        try {
            const response=await loginUser({
                username: form.value.useremail,
                password: form.value.password,
            })
            getUserProfile()
            if(form.value.remember){
                setCookie("AuthToken" ,response?.tokens,1)
            }
            ElMessage.success('登录成功（模拟）')
        } catch (err) {
            ElMessage.error('登录失败')
        } finally {
            loading.value = false
        }
    })
}

function onForgot() {
    emit('forgot')
}
function onWechatLogin() {
    // 处理微信登录逻辑
}
const toRegister=()=>{
    emit('register')
}
</script>

<style scoped>
.login-card {
    width: 400px;
    height: 600px;
    border-radius: 20px;
    background-color: rgba(255, 255, 255, 0.244);
}
.title {
    font-size: 1.25rem;
    margin-bottom: 12px;
    text-align: center;
}
.flex-row {
    width: 100%;
    display: flex;
    align-items: center;
    justify-content: space-between;
}
.forgot {
    color: #059669;
    cursor: pointer;
    font-size: 0.9rem;
}
.forgot:hover {
    text-decoration: underline;
}
.register {
    margin-top: 8px;
    text-align: center;
    color: #666;
}
.register a {
    color: #059669;
    cursor: pointer;
}
</style>
