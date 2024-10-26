import { useState } from 'react'
import './App.css'
import ReactQuill from 'react-quill'

function App() {
  // FIXME: старая библа или код quill, заменить в будущем
  const urlParams = new URLSearchParams(window.location.search)
  const taskId = urlParams.get('task_id')
  console.log(taskId)
  
  const [isEditing, setIsEditing] = useState(false)
  const [isAdding, setIsAdding] = useState(false)
  const [inpVal, setInpVal] = useState('')
  const [users, setUsers] = useState([
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
  ])
  
  const handleKeyDown = (e) => {
    if (e.key === 'Enter') {
      setUsers(prev => [...prev, {id: Date.now(), name: inpVal}])
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
  const [code, setCode] = useState("hellllo");
  // eslint-disable-next-line no-unused-vars
  const handleProcedureContentChange = (content, delta, source, editor) => {
    setCode(content);
  };
  
  const getSelectData = (e) => {
    console.log(e.target.value)
  }
  
  return (
    <>
      <div className="modalTask">
        <div className="modalTop">
          <h2 className=''>Заголовок задачи</h2>
          <img className='cross' src="../public/cross-svgrepo-com (1) 2.svg" alt="" />
          <select name="status" id="" onChange={(e) => getSelectData(e)}>
            <option value="backlog">Беклог</option>
            <option value="processing">В процессе</option>
            <option value="finished">Завершено</option>
          </select>
          <div className="modalTask__cards">
            <div className="modalTask__card-users">
              <div className="modalTask__card-header">
                <h3>Ответственный</h3>
                <button onClick={() => setIsAdding(!isAdding)}>+</button>
              </div>
              <div className="modalTask__card-item">
                {users.map((el, index) => <p key={el.id}>{el.name}{index !== users.length - 1 && ', '}</p>)}
                {isAdding &&
                <input type="text" placeholder='ФИО' value={inpVal} onChange={(e) => setInpVal(e.target.value)} onKeyDown={(e)=> handleKeyDown(e)}/>
                }
              </div>
            </div>
            {isEditing === false ? (
            <div className="modalTask__card-desc">
              <div className="modalTask__desc-header">
                <h3>Описание</h3>
                <button onClick={() => setIsEditing(true)}>/</button>
              </div>
              <div className="description-preview" dangerouslySetInnerHTML={{ __html: code }}/>
            </div>
            ) : (
            <div className="modalTask__editor">
              <p onClick={() => setIsEditing(false)}>Готово</p>
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
        <form action="" className='modalBottom'>
          <input type="submit" value='Сохранить'/>
          <p>Удалить задачу</p>
        </form>
      </div>
    </>
  )
}

export default App
