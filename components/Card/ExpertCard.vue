<template>
  <Card
    class="h-full border border-slate-100 shadow-sm transition-all duration-150 hover:-translate-y-1 hover:shadow-lg cursor-pointer"
    @click="handleClick"
  >
    <template #content>
      <div class="flex flex-col h-full">
        <div class="flex items-center gap-4">
          <Avatar :image="expert.avatar || defaultAvatar" shape="circle" class="w-16 h-16" />
          <div class="flex flex-col gap-1">
            <div class="flex items-center gap-2">
              <h3 class="text-lg font-semibold text-gray-800">{{ expert.name }}</h3>
              <Tag :value="expert.field" severity="info" class="text-xs" />
            </div>
            <div class="flex items-center gap-3 text-xs text-gray-500">
              <span class="flex items-center gap-1">
                <i class="pi pi-users text-primary-500"></i>
                {{ expert.consultCount }} 次咨询
              </span>
              <span class="flex items-center gap-1">
                <i class="pi pi-star-fill text-yellow-500"></i>
                {{ expert.rating.toFixed(1) }} 分
              </span>
              <span v-if="expert.responseTime" class="flex items-center gap-1">
                <i class="pi pi-clock text-cyan-500"></i>
                {{ expert.responseTime }}
              </span>
            </div>
          </div>
        </div>

        <p class="mt-3 text-sm text-gray-600 leading-relaxed min-h-[60px]">
          {{ expert.introduction }}
        </p>

        <div v-if="displaySkills.length" class="mt-4 flex flex-wrap gap-2">
          <Tag
            v-for="skill in displaySkills"
            :key="skill"
            :value="skill"
            severity="secondary"
            class="text-xs bg-slate-100 text-slate-700"
          />
          <Tag v-if="remainingSkillCount > 0" :value="`+${remainingSkillCount}`" class="text-xs" />
        </div>

        <Button
          label="查看详情"
          class="w-full mt-5 p-button-outlined"
          icon="pi pi-arrow-right"
          iconPos="right"
          @click.stop="handleClick"
        />
      </div>
    </template>
  </Card>
</template>

<script setup lang="ts">
import { computed } from 'vue'
import Card from 'primevue/card'
import Avatar from 'primevue/avatar'
import Tag from 'primevue/tag'
import Button from 'primevue/button'
import type { Expert } from '~/types/expert'

const props = defineProps<{
  expert: Expert
}>()

const emit = defineEmits<{
  (e: 'click'): void
  (e: 'select', id: string): void
}>()

const defaultAvatar = '/ioanImage/default-avatar.png'

const displaySkills = computed(() => props.expert.skills?.slice(0, 3) || [])
const remainingSkillCount = computed(
  () => Math.max(0, (props.expert.skills?.length || 0) - displaySkills.value.length)
)

const handleClick = () => {
  emit('click')
  emit('select', props.expert.id)
}
</script>
