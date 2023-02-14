import React from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import LoginPage from './features/LoginPage'
import './App.css'
import HomePage from './features/HomePage'
import PublicRoutes from './Routes/PublicRoutes'
import PrivateRoutes from './Routes/PrivateRoutes'
import RegisterPage from './features/RegisterPage'
function App () {
  return (
    <Router>
      <div>
        <Routes>
          <Route
            path='/register'
            element={
              <PublicRoutes>
                <RegisterPage />
              </PublicRoutes>
            }
          />

          <Route
            path='/login'
            element={
              <PublicRoutes>
                <LoginPage />
              </PublicRoutes>
            }
          />
          <Route
            path='/'
            element={
              <PrivateRoutes>
                <HomePage />
              </PrivateRoutes>
            }
          />
        </Routes>
      </div>
    </Router>
  )
}

export default App
