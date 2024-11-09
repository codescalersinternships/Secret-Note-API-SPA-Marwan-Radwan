import axios from 'axios'
const api = axios.create({
  baseURL: 'http://localhost:3000/api',
  headers: {
    'Content-Type': 'application/json',
  },
})

export const createNote = async data => {
  try {
    const response = await api.post('/notes/', data)
    return response.data
  } catch (error) {
    console.error('Failed to create note', error)
  }
}

export const getNote = async noteId => {
  try {
    const response = await api.get(`/notes/${noteId}`)
    return response.data
  } catch (error) {
    console.error('Failed to fetch note', error)
  }
}
