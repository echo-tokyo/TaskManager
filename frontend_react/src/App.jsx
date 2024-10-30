import { useEffect, useRef, useState } from 'react'
import './App.css'
import ReactQuill from 'react-quill'
import axios from 'axios'

function App() {
  // FIXME: старая библа или код quill, заменить в будущем
  const urlParams = new URLSearchParams(window.location.search)
  const taskId = urlParams.get('task_id')
  const accessToken = localStorage.getItem('access') 
  console.log(accessToken)
  
  const [taskName, setTaskName] = useState('Название задачи')
  const [isEditing, setIsEditing] = useState(false)
  const [selectedData, setSelectedData] = useState('backlog')
  const [isAdding, setIsAdding] = useState(false)
  const [inpVal, setInpVal] = useState('')
  const [users, setUsers] = useState([])
  const allUsers = useRef(null)
  
  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      setUsers(prev => [...prev, {id: Date.now(), username:'username', fio: inpVal, is_staff: false, job_title: 'Не указано'}])
      setInpVal('')
    }
  }
  
  const modules = {
    toolbar: [
      [{ header: [1, 2, 3, 4, 5, 6, false] }],
      ["bold", "italic", "underline", "strike", "blockquote"],
      [{ list: "ordered" }, { list: "bullet" }],
      ["link"],
      [{ color: ["red", "#785412"] }],
      [{ background: ["red", "#785412"] }],
    ]
  };

  const formats = [
    "header",
    "bold",
    "italic",
    "underline",
    "strike",
    "blockquote",
    "list",
    "bullet",
    "link",
    "color",
    "image",
    "background",
    "align",
    "size",
    "font"
  ];

  useEffect(() => {
    axios.get(`http://193.188.23.216/api/v1/tasks/${taskId}/`, {
      headers: {'Authorization': `Bearer ${accessToken}`}
    })
    .then((res) => {
      console.log(res.data)
      setCode(res.data.task_desc);
      setUsers(res.data.executors);
      setTaskName(res.data.title);
      setSelectedData(res.data.status)
    })
    .catch((err) => {
      console.error(err);
    });

    axios.get('http://193.188.23.216/api/v1/users/', {
      headers: {'Authorization': `Bearer ${accessToken}`}
    })
    .then(res => {
      allUsers.current = res.data
      console.log(res.data)
    })
    .catch(err => {
      console.error(err);
    });
  }, [])

  // TODO: в инишиал ставить данные с сервера
  const [code, setCode] = useState('');
  // eslint-disable-next-line no-unused-vars
  const handleProcedureContentChange = (content, delta, source, editor) => {
    setCode(content);
  };

  const getSelectData = (e) => {
    setSelectedData(e.target.value)
  }
  
  const handleSubmit = () => {
    const formData = {
      id: Number(taskId),
      board_id: 1,
      title: taskName,
      task_desc: code,
      status: selectedData,
      executors: users,
    }
    
    axios.patch(`http://193.188.23.216/api/v1/tasks/${taskId}/`, formData, {headers: {'Authorization': `Bearer ${accessToken}`}
    })
    .then((res)=>{
      console.log(res)
    })
    .catch((err)=>{
      console.error("Ошибка", err)
    })
  }

  const checkUsers = (e) => {
    e.preventDefault()
    const allFios = allUsers.current.map(user => user.fio)
    const allMatch = users.every(user => allFios.includes(user.fio))

    if(allMatch === true){
      handleSubmit()
    }
    else{
      alert('Указанного пользователя нет')
      location.reload()
    }
  }
  
  const deleteTask = () => {
    axios.delete(`http://193.188.23.216/api/v1/tasks/${taskId}/`, {headers: {'Authorization': `Bearer ${accessToken}`}})
    .then(res => {
      console.log(res)
    })
    .catch(err=>{
      console.err(err)
      alert('Задача удалена или попробуйде позже')
    })
  }

  return (
    <>
      <div className="modalTask">
        <div className="modalTop">
          <h2 className=''>{taskName}</h2>
          <select name="status" id="" value={selectedData} onChange={(e) => getSelectData(e)}>
            <option value="backlog">Беклог</option>
            <option value="proccessing">В процессе</option>
            <option value="finished">Завершено</option>
          </select>
          <div className="modalTask__cards">
            <div className="modalTask__card-users">
              <div className="modalTask__card-header">
                <h3>Ответственный</h3>
                <img src="../public/plus-large-svgrepo-com 1.svg" alt="add" onClick={() => setIsAdding(!isAdding)}/>
              </div>
              <div className="modalTask__card-item">
                {users.map((el, index) => <p key={el.id}>{el.fio}{index !== users.length - 1 && ', '}</p>)}
                {isAdding &&
                <input type="text" placeholder='ФИО' value={inpVal} onChange={(e) => setInpVal(e.target.value)} onKeyDown={(e)=> handleKeyDown(e)}/>
                }
              </div>
            </div>
            {isEditing === false ? (
              <>
            <div className="modalTask__card-desc">
              <div className="modalTask__desc-header">
                <h3>Описание</h3>
                <img src="../public/edit-2-svgrepo-com 1.svg" alt="edit" onClick={() => setIsEditing(true)}/>
                {/* <button onClick={() => setIsEditing(true)}>/</button> */}
              </div>
              <div className="description-preview" dangerouslySetInnerHTML={{ __html: code }}/>
            </div>
              </>
            ) : (
            <div className="modalTask__editor">
              <p onClick={() => setIsEditing(false)} className='modalTask__editor-done'>Готово</p>
              <ReactQuill
                theme="snow"
                modules={modules}
                formats={formats}
                value={code}
                onChange={handleProcedureContentChange}
              />
            </div>)}
          </div>
        </div>
        <form action="" className='modalBottom' onSubmit={(e) => checkUsers(e)}>
          <input type="submit" value='Сохранить'/>
          <p onClick={() => deleteTask()}>Удалить задачу</p>
        </form>
      </div>
    </>
  )
}

export default App
