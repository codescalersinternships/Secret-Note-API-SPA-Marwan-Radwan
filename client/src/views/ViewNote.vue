<template>
  <div class="view-note-container">
    <h1>View Note</h1>
    <NoteView :noteId="id" />
    <button @click="copyToClipboard" class="share-button">Share Link</button>
  </div>
</template>

<script lang="ts">
import { defineComponent } from 'vue'
import NoteView from '../components/NoteView.vue'
import { useRoute } from 'vue-router'

export default defineComponent({
  name: 'ViewNote',
  components: {
    NoteView,
  },
  props: {
    id: {
      type: String,
      required: true,
    },
  },
  setup(props) {
    const route = useRoute()

    const copyToClipboard = async () => {
      try {
        const url = `${window.location.origin}/view/${props.id}`
        await navigator.clipboard.writeText(url)
        alert('Link copied to clipboard!')
      } catch (error) {
        console.error('Failed to copy link:', error)
        alert('Failed to copy the link. Please try again.')
      }
    }

    return {
      copyToClipboard,
    }
  },
})
</script>

<style scoped>
.view-note-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 1rem;
  border: 1px solid #ccc;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

.share-button {
  margin-top: 1rem;
  padding: 0.5rem 1rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.share-button:hover {
  background-color: #0056b3;
}
</style>
