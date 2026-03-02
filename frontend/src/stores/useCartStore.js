import { create } from 'zustand'

export const useCartStore = create((set) => ({
  items: [],
  add(item) {
    set((s) => ({ items: [...s.items, item] }))
  },
  remove(id) {
    set((s) => ({ items: s.items.filter((i) => i.id !== id) }))
  },
  clear() {
    set({ items: [] })
  },
}))

export default useCartStore
