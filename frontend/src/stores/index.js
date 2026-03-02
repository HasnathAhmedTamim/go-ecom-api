export { useAuthStore } from './useAuthStore'
export { useCartStore } from './useCartStore'
export { useToastStore } from './useToastStore'

// Also provide default exports for convenience
import defaultAuth from './useAuthStore'
import defaultCart from './useCartStore'
import defaultToast from './useToastStore'

export default {
  useAuthStore: defaultAuth,
  useCartStore: defaultCart,
  useToastStore: defaultToast,
}
