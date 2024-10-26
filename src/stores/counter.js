import { ref, computed } from 'vue'
import { defineStore } from 'pinia'

export const user = defineStore('user', () => {
  const user = ref({
    id: 1,
    name: 'John Doe',
  })
  return user.value
})
