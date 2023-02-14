import axios from 'axios'
import { useEffect, useState } from 'react'
import Swal from 'sweetalert2'
import fileDownload from 'js-file-download'

export default function HomePage () {
  const [transactions, setTransactions] = useState([])
  const [newTransaction, setNewTransaction] = useState({})
  const [companies, setCompanies] = useState([])

  const accessToken = localStorage.getItem('access_token')

  useEffect(() => {
    const getTransactions = async () => {
      const resp = await axios({
        method: 'GET',
        url: 'http://localhost:8080/transactions',
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      })

      const { data } = resp.data
      setTransactions(data)
    }

    const getCompanies = async () => {
      const resp = await axios({
        method: 'GET',
        url: 'http://localhost:8080/companies'
      })

      const { data } = resp.data
      setCompanies(data)
    }

    getTransactions()
    getCompanies()
  }, [])

  const handleChange = e => {
    setNewTransaction({
      ...newTransaction,
      [e.target.name]: e.target.value
    })
  }

  const handleSubmit = async e => {
    e.preventDefault()
    try {
      newTransaction.total_quantity = parseInt(newTransaction.total_quantity)
      newTransaction.company_id = parseInt(newTransaction.company_id)
      newTransaction.product_id = parseInt(newTransaction.product_id)
      const resp = await axios({
        method: 'POST',
        url: 'http://localhost:8080/transactions',
        data: newTransaction,
        headers: {
          Authorization: `Bearer ${accessToken}`
        }
      })

      const { data } = resp.data
      setTransactions(...transactions, data)
    } catch (err) {
      const { error } = err.response.data
      Swal.fire({
        icon: 'error',
        title: 'Oops...',
        text: `${error}`
      })
    }
  }

  const handleDownload = (url, filename, e) => {
    e.preventDefault()
    axios
      .post(
        url,
        { path: 'records.csv' },
        {
          responseType: 'blob'
        }
      )
      .then(res => {
        fileDownload(res.data, filename)
      })
  }

  return (
    <div class='container'>
      <div>
        <h3>Transaction Table</h3>;
        <button
          onClick={e => {
            handleDownload('http://localhost:8080/downloads', 'records.csv', e)
          }}
        >
          Download Image
        </button>
        <table class='table'>
          <thead>
            <tr>
              <th scope='col'>ID</th>
              <th scope='col'>Created At</th>
              <th scope='col'>Company Name</th>
              <th scope='col'>Product Name</th>
              <th scope='col'>Product Price</th>
              <th scope='col'>Total Quantity</th>
              <th scope='col'>Total Price</th>
              <th scope='col'>Rest Stock</th>
            </tr>
          </thead>
          <tbody>
            {transactions?.map(transaction => {
              return (
                <tr>
                  <td>{transaction.id}</td>
                  <td>{transaction.created_at}</td>
                  <td>{transaction.company_name}</td>
                  <td>{transaction.product_name}</td>
                  <td>{transaction.price}</td>
                  <td>{transaction.total_quantity}</td>
                  <td>{transaction.total_price}</td>
                  <td>{transaction.rest_stock}</td>
                </tr>
              )
            })}
          </tbody>
        </table>
      </div>
      <div>
        <h3>Companies Table</h3>
        <table class='table'>
          <thead>
            <tr>
              <th scope='col'>ID</th>
              <th scope='col'>Company Name</th>
              <th scope='col'>Company Code</th>
            </tr>
          </thead>
          <tbody>
            {companies.map(company => {
              return (
                <tr>
                  <td>{company.id}</td>
                  <td>{company.company_name}</td>
                  <td>{company.company_code}</td>
                </tr>
              )
            })}
          </tbody>
        </table>
      </div>

      <div>
        Add Transaction
        <form>
          <div class='mb-3'>
            <label for='total_quantity' class='form-label'>
              Quantity
            </label>
            <input
              type='number'
              class='form-control'
              id='total_quantity'
              name='total_quantity'
              onChange={e => handleChange(e)}
            />
            <div className='col-span-6 sm:col-span-3'>
              <select
                id='company_id'
                name='company_id'
                autoComplete='company_id'
                className='mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm'
                onChange={e => handleChange(e)}
              >
                <option selected>-- Please Select Your Company --</option>
                {companies.map(company => {
                  return (
                    <>
                      <option value={company.id}>{company.company_name}</option>
                    </>
                  )
                })}
              </select>
            </div>

            <div className='col-span-6 sm:col-span-3'>
              <select
                id='product_id'
                name='product_id'
                autoComplete='product_id'
                className='mt-1 block w-full py-2 px-3 border border-gray-300 bg-white rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm'
                onChange={e => handleChange(e)}
              >
                <option selected>-- Please Select Your Produc --</option>
                {newTransaction.company_id === undefined
                  ? null
                  : companies.map(company => {
                      if (parseInt(newTransaction.company_id) == company.id) {
                        return company?.products?.map(product => {
                          console.log(product)
                          return (
                            <>
                              <option value={product.id}>{product.Name}</option>
                            </>
                          )
                        })
                      }
                    })}
              </select>
            </div>
          </div>
          <div class='mb-3'></div>
          <button type='button' class='btn btn-primary' onClick={handleSubmit}>
            Submit
          </button>
        </form>
      </div>
    </div>
  )
}
