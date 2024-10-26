import './App.css'

function App() {
  const users = [
    {
      id: 1,
      name: 'John II',
    },
    {
      id: 2,
      name: 'Johna II',
    },
    {
      id: 3,
      name: 'John I',
    }
  ]

  return (
    <>
      <div className="modalTask">
        <div className="modalTop">
          <h2 className=''>Заголовок задачи</h2>
          <img className='cross' src="../public/cross-svgrepo-com (1) 2.svg" alt="" />
          <select name="" id="">
            <option value="">Беклог</option>
            <option value="">В процессе</option>
            <option value="">Завершено</option>
          </select>
          <div className="modalTask__cards">
            <div className="modalTask__card-users">
              <h3>Ответственный</h3>
              <div className="modalTask__card-item">
                <p>Иванов И.И.</p>
                <input type="text" placeholder='ФИО'/>
              </div>
            </div>
            <div className="modalTask__card-desc">
              <h3>Описание</h3>
              <p>Описание Описание Описание</p>
            </div>
          </div>
        </div>
        <form action="" className='modalBottom'>
          <input type="submit" value='Сохранить'/>
          <p>Удалить задачу</p>
        </form>
      </div>
    </>
  )
}

export default App
