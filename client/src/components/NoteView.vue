<template>
  <div v-if="note">
    <p>
      <strong>Note Content:</strong>
      {{ note.content }}
    </p>
    <p><strong>Remaining Views:</strong> {{ note.remaining_views }}</p>
  </div>
  <div v-else-if="error">
    <p>{{ error }}</p>
  </div>
  <div v-else>
    <p>Loading...</p>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import axios from 'axios'

export default defineComponent({
  name: 'NoteView',
  props: {
    noteId: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const note = ref(null)
    const error = ref('')

    onMounted(async () => {
      try {
        const response = await axios.get(
          `http://localhost:3000/api/notes/${props.noteId}`,
        )
        note.value = response.data
      } catch (err) {
        if (err.response && err.response.status === 404) {
          error.value = 'Note is not found or maybe deleted.'
        } else {
          console.error('Error fetching the note:', err)
          error.value = 'An error occurred while fetching the note.'
        }
      }
    })

    return {
      note,
      error,
    }
  },
})
</script>
