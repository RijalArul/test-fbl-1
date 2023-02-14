import axios from 'axios'
import { useState } from 'react'
import { useNavigate } from 'react-router-dom'
import Swal from 'sweetalert2'

export default function LoginPage () {
  const [user, setUser] = useState({})
  const navigation = useNavigate()

  const handleChange = e => {
    setUser({
      ...user,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      const resp = await axios({
        method: 'POST',
        url: 'http://localhost:8080/users/register',
        data: user
      })

      const { data } = resp.data
      if (data) {
        navigation('/login')
      }
    } catch (err) {
      const { error } = err.response.data
      Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: `${error}`
      })
    }
  }
  return (
    <div>
      <h4>Ini Register Page Page</h4>
      <form>
        <div class='mb-3'>
          <label for='username' class='form-label'>
            username
          </label>
          <input
            type='text'
            class='form-control'
            id='username'
            name='username'
            onChange={e => handleChange(e)}
          />
        </div>
        <div class='mb-3'>
          <label for='exampleInputPassword1' class='form-label'>
            Password
          </label>
          <input
            type='password'
            class='form-control'
            id='exampleInputPassword1'
            name='password'
            onChange={e => handleChange(e)}
          />
        </div>
        <button type='button' class='btn btn-primary' onClick={handleSubmit}>
          Submit
        </button>
      </form>
    </div>
  )
}
