'use client'

import styles from './page.module.css'
import React from 'react'
import axios from 'axios'

export default function Home() {
  const [ name, setName ] = React.useState('')
  const [ response, setResponse ] = React.useState('')

  const sendRequest = () => {
    if (name !== '') {
      axios({
        method: 'post',
        url: 'http://localhost:3001/api/welcomePost',
        data: { name: name },
        headers: { 'Content-Type': 'application/json' }
      })
      .then((response) => setResponse(response.data))
      .catch((error) => console.log(error));
    }
  }

  return (
    <main className={styles.main}>
      <div style={{display:'flex', flexDirection: 'column', gap: '10px'}}>
        <input onChange={(e) => setName(e.target.value)} placeholder='write your name'></input>
        <button onClick={() => sendRequest()}>send request</button>
        <div>response: { response }</div>
      </div>
    </main>
  )
}
