export default function Button({ children, className = '', ...props }) {
  return (
    <button
      {...props}
      className={
        'inline-flex items-center justify-center px-4 py-2 rounded-md shadow-sm bg-blue-600 hover:bg-blue-700 text-white font-medium transition ' +
        className
      }
    >
      {children}
    </button>
  )
}
