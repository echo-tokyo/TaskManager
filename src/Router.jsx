import { BrowserRouter, Route, Routes } from 'react-router-dom'
import Search from './Search'
import Editor from './editor'

const Router = () => {
	return (
		<BrowserRouter>
			<Routes>
				<Route path='/' element={<Search />} />
				<Route path='/editor' element={<Editor />} />
			</Routes>
		</BrowserRouter>
	)
}

export default Router