<template>
    <section class="profile-page">
        <Card class="profile-card">
            <template #content>
                <div class="profile-header">
                    <Avatar
                        :image="userStore.avatar"
                        :label="userStore.avatar ? undefined : profile.nickname"
                        size="xlarge"
                        shape="circle"
                        class="profile-avatar"
                    />
                    <div class="profile-meta">
                        <h1 class="profile-name">{{ profile.nickname }}</h1>
                        <div class="profile-row">
                            <i class="pi pi-envelope"></i>
                            <span>{{ profile.email }}</span>
                        </div>
                        <p class="profile-bio">{{ profile.bio }}</p>
                        <div class="profile-tags">
                            <Tag
                                v-for="tag in profile.tags"
                                :key="tag"
                                severity="success"
                                rounded
                            >
                                {{ tag }}
                            </Tag>
                        </div>
                    </div>
                </div>

                <Divider layout="horizontal" class="profile-divider">
                    <span class="divider-chip">Profile</span>
                </Divider>

                <div class="profile-panels">
                    <Card class="info-card">
                        <template #title>
                            <div class="info-title">
                                <i class="pi pi-id-card"></i>
                                <span>基本信息</span>
                            </div>
                        </template>
                        <template #content>
                            <ul class="info-list">
                                <li>
                                    <span>昵称</span>
                                    <strong>{{ profile.nickname }}</strong>
                                </li>
                                <li>
                                    <span>邮箱</span>
                                    <strong>{{ profile.email }}</strong>
                                </li>
                                <li>
                                    <span>所在地</span>
                                    <strong>{{ profile.location }}</strong>
                                </li>
                                <li>
                                    <span>加入时间</span>
                                    <strong>{{ profile.joinedAt }}</strong>
                                </li>
                            </ul>
                        </template>
                    </Card>

                    <Card class="info-card">
                        <template #title>
                            <div class="info-title">
                                <i class="pi pi-phone"></i>
                                <span>联系方式</span>
                            </div>
                        </template>
                        <template #content>
                            <ul class="info-list">
                                <li>
                                    <span>联系电话</span>
                                    <strong>{{ profile.phone }}</strong>
                                </li>
                                <li>
                                    <span>联系地址</span>
                                    <strong>{{ profile.address }}</strong>
                                </li>
                                <li>
                                    <span>最近活跃</span>
                                    <strong>{{ userStore.lastActive }}</strong>
                                </li>
                            </ul>
                            <Button class="contact-btn" label="联系我" icon="pi pi-send" @click="onContact" />
                        </template>
                    </Card>
                </div>
            </template>
        </Card>
    </section>
</template>

<script setup lang="ts">
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Tag from 'primevue/tag'
import Divider from 'primevue/divider'
import Button from 'primevue/button'
import { useToast } from 'primevue/usetoast'
import { onMounted } from 'vue'

definePageMeta({ layout: 'profile-page-layout' })
const toast = useToast()
const userStore = useUserStore()
const profile = userStore.userinfo

function onContact() {
    window.open(`mailto:${profile.email}`, '_blank')
    toast.add({ severity: 'success', summary: '邮件已打开', detail: '请在邮件客户端继续交流', life: 2500 })
}

onMounted(() => {
    getUserProfile()
})
</script>

<style scoped>
.profile-page {
    min-height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    padding: 2.5rem 1rem;
    background: radial-gradient(circle at 20% 20%, rgba(99, 102, 241, 0.16), transparent 45%),
        radial-gradient(circle at 80% 0%, rgba(16, 185, 129, 0.18), transparent 42%),
        linear-gradient(135deg, rgba(15, 23, 42, 0.04), rgba(15, 23, 42, 0));
}

.profile-card {
    width: min(960px, 100%);
    border-radius: 24px;
    background-color: rgba(255, 255, 255, 0.9);
    backdrop-filter: blur(18px);
    box-shadow: 0 32px 60px -30px rgba(15, 23, 42, 0.4);
    border: 1px solid rgba(148, 163, 184, 0.14);
}

.profile-header {
    display: flex;
    flex-direction: column;
    gap: 1.5rem;
    align-items: flex-start;
}

.profile-avatar {
    border: 6px solid rgba(255, 255, 255, 0.8);
    box-shadow: 0 15px 35px -15px rgba(59, 130, 246, 0.5);
}

.profile-meta {
    display: flex;
    flex-direction: column;
    gap: 0.75rem;
}

.profile-name {
    font-size: clamp(1.75rem, 2.4vw, 2.4rem);
    font-weight: 700;
    color: #0f172a;
    margin: 0;
}

.profile-row {
    display: inline-flex;
    align-items: center;
    gap: 0.5rem;
    color: #475569;
    font-size: 0.95rem;
}

.profile-row .pi {
    color: #6366f1;
    font-size: 1rem;
}

.profile-bio {
    margin: 0;
    color: #475569;
    line-height: 1.7;
}

.profile-tags {
    display: flex;
    flex-wrap: wrap;
    gap: 0.5rem;
}

.profile-divider {
    margin-block: 2rem;
}

.divider-chip {
    display: inline-flex;
    align-items: center;
    gap: 0.35rem;
    background-color: rgba(59, 130, 246, 0.1);
    padding: 0.25rem 0.75rem;
    border-radius: 999px;
    color: #1d4ed8;
    font-weight: 600;
    letter-spacing: 0.02em;
}

.profile-panels {
    display: grid;
    gap: 1.25rem;
}

.info-card {
    border-radius: 18px;
    padding:20px;
    border: 1px solid rgba(148, 163, 184, 0.16) !important;
    box-shadow: 0 12px 32px -20px rgba(15, 23, 42, 0.35);
    background: linear-gradient(145deg, rgba(255, 255, 255, 0.92), rgba(248, 250, 252, 0.9));
}

.info-card :deep(.p-card-body) {
    padding: 0;
}

.info-card :deep(.p-card-content) {
    padding: 1.5rem;
}

.info-title {
    display: inline-flex;
    align-items: center;
    gap: 0.6rem;
    font-size: 1.1rem;
    font-weight: 600;
    color: #1e293b;
}

.info-title .pi {
    color: #10b981;
}

.info-list {
    list-style: none;
    margin: 0;
    padding: 0px;
    display: grid;
    gap: 0.9rem;
}

.info-list li {
    display: flex;
    flex-direction: column;
    gap: 0.25rem;
    font-size: 0.95rem;
    color: #475569;
}

.info-list span {
    text-transform: uppercase;
    letter-spacing: 0.08em;
    font-size: 0.7rem;
    font-weight: 600;
    color: rgba(100, 116, 139, 0.9);
}

.info-list strong {
    font-size: 1.05rem;
    color: #1f2937;
    letter-spacing: 0.01em;
}

.contact-btn {
    width: 100%;
    margin-top: 1.5rem;
    justify-content: center;
}

@media (min-width: 768px) {
    .profile-header {
        flex-direction: row;
        align-items: center;
    }

    .profile-panels {
        grid-template-columns: repeat(2, minmax(0, 1fr));
    }
}

@media (max-width: 767px) {
    .profile-card {
        padding-inline: 0.75rem;
    }
}
</style>
