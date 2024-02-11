"use client"

import React, { useState } from 'react'
import styles from './home.module.css'
import axios from 'axios';

export default function Home() {
    const [helloWorldResponse, setHelloWorldResponse] = useState('');

    const sendRequest = () => {
      axios.get('http://localhost:8000/api/helloWorld')
        .then((response) => setHelloWorldResponse(response.data))
        .catch((error) => console.log(error));
    }

    return (
      <div className={styles.homepage_main_wrapper}>
        <button onClick={() => {sendRequest()}}>send</button>
        <div style={{color: 'black'}}>{ helloWorldResponse }</div>
      </div>
    )
  }