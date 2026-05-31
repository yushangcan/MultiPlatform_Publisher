<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { getHealth } from './api/client'

const loading = ref(true)
const status = ref('checking')
const error = ref('')

async function loadHealth() {
  loading.value = true
  error.value = ''
  try {
    const result = await getHealth()
    status.value = result.status
  } catch (err) {
    status.value = 'offline'
    error.value = err instanceof Error ? err.message : '无法连接后端服务'
  } finally {
    loading.value = false
  }
}

onMounted(loadHealth)
</script>

<template>
  <main class="app-shell">
    <header class="topbar">
      <div>
        <p class="eyebrow">MultiPlatform Publisher</p>
        <h1>多平台内容发布工作台</h1>
      </div>
      <button class="ghost-button" type="button" @click="loadHealth">刷新状态</button>
    </header>

    <section class="workspace">
      <div class="panel">
        <h2>后端连接</h2>
        <div class="status-row">
          <span class="status-dot" :class="{ online: status === 'ok' }" />
          <span>{{ loading ? '检查中' : status }}</span>
        </div>
        <p v-if="error" class="error-text">{{ error }}</p>
        <p v-else class="muted">前端已连接到后端健康检查接口，下一步会接入模糊输入和平台改写流程。</p>
      </div>
    </section>
  </main>
</template>
