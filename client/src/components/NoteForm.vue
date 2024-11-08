<template>
  <form @submit.prevent="submitNote">
    <div class="form-group">
      <textarea v-model="text" placeholder="Note Content" required></textarea>
    </div>
    <div class="form-group">
      <input
        v-model="expire_date"
        type="datetime-local"
        placeholder="Expiration date"
        required
      />
    </div>
    <div class="form-group">
      <input
        v-model="remaining_views"
        type="number"
        placeholder="Max remaining views"
        required
      />
    </div>
    <button type="submit">Create Note</button>
  </form>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue'
import { useRouter } from 'vue-router'
import { createNote } from '../client'

export default defineComponent({
  name: 'NoteForm',
  setup() {
    const text = ref('')
    const expire_date = ref('')
    const remaining_views = ref(0)
    const router = useRouter()

    const submitNote = async () => {
      try {
        expire_date.value = expire_date.value.replace('T', ' ')
        const response = await createNote({
          content: text.value,
          remaining_views: remaining_views.value,
          expire_date: expire_date.value.replace('T', ' '),
        })
        console.log('id', response)
        router.push(`/${response.id}`)
      } catch (error) {
        console.error('Failed to create note', error)
      }
    }

    return {
      text,
      expire_date,
      remaining_views,
      submitNote,
    }
  },
})
</script>

<style scoped>
.form-group {
  margin-bottom: 1rem;
}

textarea {
  width: 100%;
  height: 100px;
  padding: 0.5rem;
}

input,
button {
  width: 100%;
  padding: 0.5rem;
  margin-top: 0.5rem;
  border-radius: 4px;
}

button {
  background-color: #007bff;
  color: white;
  cursor: pointer;
  border: none;
}

button:hover {
  background-color: #0056b3;
}
</style>
