import { useState } from 'react'

function App() {
  const [count, setCount] = useState(0)

  return (
    <div className="min-h-screen bg-gray-100 flex items-center justify-center">
      <div className="bg-white p-8 rounded-lg shadow-md max-w-md w-full">
        <h1 className="text-3xl font-bold text-gray-800 mb-6 text-center">
          Simple React App
        </h1>
        
        <div className="text-center">
          <p className="text-gray-600 mb-4">
            Welcome to your React + Vite + Tailwind application
          </p>
          
          <div className="mb-6">
            <button 
              onClick={() => setCount((count) => count + 1)}
              className="bg-blue-500 hover:bg-blue-600 text-white font-semibold py-2 px-4 rounded transition-colors"
            >
              Count: {count}
            </button>
          </div>
          
          <p className="text-sm text-gray-500">
            Click the button to increment the counter
          </p>
        </div>
      </div>
    </div>
  )
}

export default App
