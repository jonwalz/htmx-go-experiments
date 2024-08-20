import React from 'react'
import ReactDOM from 'react-dom/client'

function ChildComponent() {
  const [state, setState] = React.useState(0)
  return (
    <div>
      <div>Child: {state}</div>
      <button onClick={() => setState((prevState) => prevState + 1)}>+</button>
    </div>
  )
}

const root = document.getElementById('about')

if (root) {
  // @ts-ignore
  ReactDOM.createRoot(root).render(
    <React.StrictMode>
      <div className="card bg-primary text-primary-content w-96 mt-4">
        <div className="container mx-auto">React component</div>
        <ChildComponent />
      </div>
    </React.StrictMode>
  )
}
