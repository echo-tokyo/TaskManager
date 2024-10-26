import { useEffect, useState } from 'react'
import './assets/App.css'

function Search() {
  const initialTaskList = [
    {
      id: 1,
      taskName: 'имя1',
      date: '2021-01-01'
    },
    {
      id: 2,
      taskName: 'имя2',
      date: '2020-01-01'
    },
    {
      id: 3,
      taskName: 'имя3',
      date: '2023-01-01'
    }
  ]

  const [taskList, setTaskList] = useState(initialTaskList)

  // поиск
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

  // сортировка
  const dateSort = (e) => {
    if(e.target.value === 'dateUp'){
      setTaskList([...taskList.sort((a, b) => new Date(b.date) - new Date(a.date))])
    }
    else if(e.target.value === 'dateDown'){
      setTaskList([...taskList.sort((a, b) => new Date(a.date) - new Date(b.date))])
    }
    else{
      setTaskList(initialTaskList)
    }
  }

  return (
    <>
    <select name="" id="" onChange={(e) => dateSort(e)}>
      <option value="default">по умолчанию</option>
      <option value="dateUp">по дате (возврастанию)</option>
      <option value="dateDown">по дате (убыванию)</option>
    </select>
    <input type="text" className='search'/>
    <div className="cards">
      {taskList.map(el => <div key={el.id} className='card'><p className='task'>{el.taskName}</p></div>)}
    </div>
    </>
  )
}

export default Search
