<!--
    registerform.vue
    说明：用户注册表单组件，包含邮箱验证码发送与校验逻辑
    Emits:
        - submit: 校验并提交注册数据（payload: RegisterRequest）
        - login: 点击已有账号跳转到登录
    依赖：调用了 useRegister 组合函数中的 onSendCode / onVerifyCode，请确保实现存在。
-->
<template>
        <el-card class="login-card" :body-style="{ padding: '20px' }">
        <h3 class="title">注册</h3>

        <el-form :model="form" :rules="rules" ref="formRef" label-position="top">
            <el-form-item label="用户名" prop="username">
                <el-input v-model="form.username" placeholder="请输入用户名" clearable/>
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
                <el-input v-model="form.email"  placeholder="请输入邮箱" clearable/>
            </el-form-item>
            <el-form-item label="密码" prop="password">
                <el-input v-model="form.password" type="password" placeholder="请输入密码" show-password />
            </el-form-item>
            <el-form-item label="确认密码" prop="compassword">
                <el-input v-model="form.compassword" type="password" placeholder="请确认密码" show-password />
            </el-form-item>
            <el-form-item label="验证邮箱" prop="comemail">
                <div class="flex justify-between w-full">
                <el-input v-model="comemail" style="width: 70%;" placeholder="请输入验证码"/>
                <el-button type="success" :loading="LodingStore.sendLoading" @click="onSendCode(form.email),Countdown(60)" plain>发送验证码</el-button>
                </div>
                <div v-if="LodingStore.sendLoading" class="flex items-center">请在{{ count }}秒后重试</div>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" :loading="loading" @click="onSubmit" style="width:100%">注册</el-button>
            </el-form-item>

            <div class="register">
                已有账号？ <a @click.prevent="$emit('login')">登录</a>
            </div>
        </el-form>
    </el-card>
</template>
<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { onSendCode , onVerifyCode} from '~/composables/useRegister'
import type { RegisterRequest } from '~/types/register'
import type { FormRules } from 'element-plus'
import {useLoadingStore} from '../utils/loadingStore'
const count = ref(0)
const emit = defineEmits<{
    (e: 'submit', payload: RegisterRequest): void
    (e: 'register'): void
    (e: 'forgot'): void
    (e: 'login'): void
}>()
const LodingStore=useLoadingStore()
const formRef = ref<any>(null)
const form = ref({ username: '',email:'', password: '',compassword:'',isverified:false })
const comemail=ref<string>('')
const loading = ref(false)
const rules:  FormRules = {
    username: [{ required: true, message: '请输入用户名或邮箱', trigger: 'blur' }],
    email: [
        { required: true, message: '请输入邮箱', trigger: 'blur' },
        { type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] },
    ],
    password: [{ required: true, message: '请输入密码', trigger: 'blur' }],
    compassword: [
        { required: true, message: '请确认密码', trigger: 'blur' },
        {
            validator: (rule, value, callback) => {
                if (!value) {
                callback(new Error('请确认密码'))
            } else if (value !== form.value.password) {
                callback(new Error('两次输入的密码不一致'))
            } else {
                callback()
            }
        },
            trigger: ['blur', 'change'],
        }
    ]
}
async function onSubmit() {
    if (!formRef.value) return
        try {
                await formRef.value.validate()
        } catch (err) {
                ElMessage.error('注册失败')
                 return
        }
        loading.value = true
        try {   
                const response=await onVerifyCode(form.value.email,comemail.value)
                if(response!==200){
                        ElMessage.error('验证码错误或已过期')
                        loading.value = false
                        throw new Error('验证码错误或已过期')
                }
                else{
                        form.value.isverified=true
                }
                emit('submit', { ...form.value })
        } catch (err) {
                ElMessage.error('注册失败')
        } finally {
                loading.value = false
        }
}
function Countdown(seconds: number) {
    count.value = seconds
    const interval = setInterval(() => {
        if (count.value > 0) {
            count.value--
        } else {
            clearInterval(interval)
            LodingStore.cancelLoading(false)
        }
    }, 1000)
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
