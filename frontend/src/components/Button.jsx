export default function Button({ children, className = '', ...props }) {
  return (
    <button
      {...props}
      className={
        'inline-flex items-center justify-center px-4 py-2 rounded-md shadow-neon bg-neon-pink hover:brightness-95 text-black font-semibold transition ' +
        className
      }
    >
      {children}
    </button>
  )
}
