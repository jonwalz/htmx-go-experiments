import React from 'react'
import ReactDOM from 'react-dom/client'

function ChildComponent() {
  const [state, setState] = React.useState(0)
  return (
    <div className="card bg-primary text-primary-content w-96 mt-4">
      <div className="card-body">
        <div>Number: {state}</div>
        <div className="card-actions justify-end">
          <button
            className="btn"
            onClick={() => setState((prevState) => prevState + 1)}
          >
            +
          </button>
        </div>
      </div>
    </div>
  )
}

const root = document.getElementById('about')

if (root) {
  ReactDOM.createRoot(root).render(
    <React.StrictMode>
      <ChildComponent />
    </React.StrictMode>
  )
}
