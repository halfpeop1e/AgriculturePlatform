<template>
    <section class=" bg-gradient-to-tr from-indigo-500/10 to-emerald-500/10 flex items-center justify-center py-8 px-4 h-full">
        <el-card
            class="w-full h-full  border-none shadow-xl rounded-3xl bg-white/80 dark:bg-slate-900/80 backdrop-blur-lg"
            shadow="hover"
            :body-style="{ padding: '32px', display: 'flex', flexDirection: 'column', gap: '24px' }"
        >
            <div class="flex flex-col gap-6 md:flex-row md:items-center">
                <el-avatar :size="96" :src="userStore.avatar" alt="用户头像">
                    {{ profile.nickname }}
                </el-avatar>
                <div class="flex-1 min-w-[220px] space-y-3">
                    <h2 class="text-3xl font-semibold text-slate-900 dark:text-slate-100">{{ profile.nickname }}</h2>
                    <div class="flex items-center gap-2 text-slate-500 dark:text-slate-300">
                        <el-icon><Message /></el-icon>
                        <span>{{ profile.email }}</span>
                    </div>
                    <p class="text-slate-600 dark:text-slate-300 leading-relaxed">{{ profile.bio }}</p>
                    <div class="flex flex-wrap gap-2">
                        <el-tag v-for="tag in profile.tags" :key="tag" type="success" effect="plain">
                            {{ tag }}
                        </el-tag>
                    </div>
                </div>
            </div>

            <el-divider />

            <el-row :gutter="16" class="gap-y-4">
                <el-col :xs="24" :md="12">
                    <el-card
                        class="rounded-2xl border border-slate-200/80 dark:border-slate-700/60 shadow-none"
                        shadow="never"
                        :body-style="{ padding: '20px' }"
                    >
                        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-3">基本信息</h3>
                        <el-descriptions :column="1" border size="small">
                            <el-descriptions-item label="昵称">{{ profile.nickname }}</el-descriptions-item>
                            <el-descriptions-item label="邮箱">{{ profile.email }}</el-descriptions-item>
                            <el-descriptions-item label="所在地">{{ profile.location }}</el-descriptions-item>
                            <el-descriptions-item label="加入时间">{{ profile.joinedAt }}</el-descriptions-item>
                        </el-descriptions>
                    </el-card>
                </el-col>
                <el-col :xs="24" :md="12">
                    <el-card
                        class="rounded-2xl border border-slate-200/80 dark:border-slate-700/60 shadow-none"
                        shadow="never"
                        :body-style="{ padding: '20px' }"
                    >
                        <h3 class="text-lg font-semibold text-slate-800 dark:text-slate-100 mb-3">联系方式</h3>
                        <ul class="grid gap-3 list-none p-0 mb-4">
                            <li class="flex items-center gap-2 text-slate-600 dark:text-slate-300 text-sm">
                                <el-icon><Phone /></el-icon>
                                <span>{{ profile.phone }}</span>
                            </li>
                            <li class="flex items-center gap-2 text-slate-600 dark:text-slate-300 text-sm">
                                <el-icon><Position /></el-icon>
                                <span>{{ profile.address }}</span>
                            </li>
                            <li class="flex items-center gap-2 text-slate-600 dark:text-slate-300 text-sm">
                                <el-icon><Timer /></el-icon>
                                <span>最近活跃：{{ userStore.lastActive }}</span>
                            </li>
                        </ul>
                        <el-button class="w-full" type="primary" @click="onContact">联系我</el-button>
                    </el-card>
                </el-col>
            </el-row>
        </el-card>
    </section>
</template>

<script setup lang="ts">
import { reactive } from 'vue'
import { ElMessage } from 'element-plus'
import { Message, Phone, Position, Timer } from '@element-plus/icons-vue'
definePageMeta({ layout: 'profile-page-layout' })
const userStore = useUserStore()
const profile = userStore.userinfo

function onContact() {
    window.open(`mailto:${profile.email}`, '_blank')
    ElMessage.success('已打开邮件客户端~')
}
</script>
