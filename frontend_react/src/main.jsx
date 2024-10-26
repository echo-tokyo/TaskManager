import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './assets/reset.css'
import './assets/global.css'
import Router from './Router.jsx'

createRoot(document.getElementById('root')).render(
  <StrictMode>
    <Router />
  </StrictMode>,
)
