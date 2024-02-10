"use client"

import React from 'react'
import styles from './home.module.css'
import axios from 'axios';

export default function Home() {
    const [script, setScript] = React.useState('');
    const [RCodeOutput, setRCodeOutput] = React.useState('');

    const sendRequest = () => {
      axios.post('http://localhost:3001/api/runRcode', { script: script })
      .then((response) => setRCodeOutput(response.data))
      .catch((error) => console.log(error));
    }

    return (
      <div className={styles.homepage_main_wrapper}>
        <input onChange={(e) => {setScript(e.target.value)}} placeholder='R script here'></input>
        <button onClick={() => {sendRequest()}}>send</button>
        <div style={{color: 'black'}}>{ RCodeOutput }</div>
      </div>
    )
  }