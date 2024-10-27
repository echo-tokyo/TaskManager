import { ref, computed } from 'vue'
import { defineStore } from 'pinia'
import axios from 'axios'

// axios.get('http://193.188.23.216/api/v1/users/1/')
// .then(res => {
//     console.log(res)
//     user.value.fio = res.data.fio;
// })
// .catch(err => {
//     console.log(err)
// })

export const user = defineStore('user', () => {
  const user = ref('Иванов Иван')
  return user.value
})
