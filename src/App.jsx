import { useEffect } from 'react'
import './App.css'

function App() {
  useEffect(() => {
    const searchInput = document.querySelector('.search')
    const tasks = document.querySelectorAll('.task')

    searchInput.addEventListener('input', function() {
        const query = this.value.toLowerCase()
  
        Array.from(tasks).forEach(task => {
            const taskText = task.textContent.toLowerCase();
            if (taskText.includes(query)) {
                task.style.display = 'block'
            } else {
                task.style.display = 'none'
            }
        })
    });
  }, [])

  return (
    <>
    <input type="text" className='search'/>
    <div className="cards">
      <div className="card">
        <p className='task'>имя</p>
      </div>
      <div className="card">
        <p className='task'>имя2</p>
      </div>
      <div className="card">
        <p className='task'>имя3</p>
      </div>
    </div>
    </>
  )
}

export default App
