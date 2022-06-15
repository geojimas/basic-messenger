<template>
  <main>
    <div style="overscroll-behavior: none">
      <div
        class="fixed w-full bg-green-400 h-16 pt-2 text-white flex justify-center shadow-md"
        style="top: 0px; overscroll-behavior: none">
        <input
          class="my-3 text-black font-bold text-lg"
          type="text"
          v-model="username"
          placeholder="username" />
      </div>
      <div class="mt-20 mb-16">
        <div
          class="animate__animated animate__fadeInDown bg-gray-300 mx-12 my-2 p-2 rounded-lg"
          v-for="(msg, index) in messages"
          :key="index">
          <p class="font-bold">{{ msg.username }}</p>
          <hr />
          <p class="ml-8">{{ msg.message }}</p>
        </div>
      </div>
    </div>
    <div class="fixed w-full flex justify-between bg-green-100" style="bottom: 0px">
      <textarea
        class="flex-grow m-2 py-2 px-3 mr-1 rounded-full border border-gray-300 resize-none"
        rows="1"
        v-model="message"
        placeholder="Message..."
        style="outline: none"></textarea>
      <button @click="submit" class="m-2" style="outline: none">
        <SentBtn />
      </button>
    </div>
  </main>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import Pusher from 'pusher-js'
import SentBtn from './SentBtn.vue'

const username = ref('John Doe')
const message = ref('')
const messages = ref([])

onMounted(() => {
  Pusher.logToConsole = true
  const pusher = new Pusher(import.meta.env.VITE_PUSHER_ID, {
    cluster: import.meta.env.VITE_CLUSTER
  })
  const channel = pusher.subscribe('chat')
  channel.bind('message', data => {
    messages.value.push(data)
  })
})

const submit = async () => {
  axios
    .post(import.meta.env.VITE_API, {
      username: username.value,
      message: message.value
    })
    .then(response => {
      console.log(response)
      message.value = ''
    })
    .catch(error => {
      console.log(error)
    })
}
</script>
