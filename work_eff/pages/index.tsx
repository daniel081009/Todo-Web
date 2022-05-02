import type { NextPage } from 'next'
import React, { useState } from "react";
import styles from '../styles/Index.module.css'
import { useCookies } from "react-cookie";
import axios from "axios";
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import useRouter from "next/router";
import Heads from '../modules/Heads'

const Home: NextPage = () => {
  let [user, set_user] = useState({
    "email": "",
    "password": ""
  })
  const [cookie,setCookie] = useCookies()

  const login_buttonClick = () => {
    axios.post('http://localhost/user/login', { Email: user.email, Password: user.password }, { withCredentials: true })
        .then(res => {
          if (res.data.code === 200) {
            setCookie("Token", res.data.token)
            localStorage.setItem("Token", JSON.stringify(res.data.token))
            useRouter.push("/home")
            toast.success('로그인 성공!', {
              position: "top-right",
              autoClose: 1000,
              hideProgressBar: true,
              closeOnClick: true,
              pauseOnHover: true,
              draggable: true,
              progress: undefined,
            });
          } else {
            toast.error('로그인 실패!', {
              position: "top-right",
              autoClose: 1000,
              hideProgressBar: true,
              closeOnClick: true,
              pauseOnHover: true,
              draggable: true,
              progress: undefined,
            });
          }
        })
  }
  return (
      <>
      {Heads("Todo - Index","Todo index")}
      <ToastContainer/>
      <div className = { styles.main }>
          <div className = { styles.viewer } >
              <h1 className = { styles.item } > Todo </h1>
              <div className = { styles.item } > Email: <input type = { "email" } onChange = {(e) => { set_user({ email: e.target.value, password: user.password }) }}/>
              </div>
              <div className = { styles.item } > Password:
                  <input type = { "password" } onChange = {(e) => { set_user({ email: user.email, password: e.target.value }) }}/>
              </div>
              <div className = { styles.button }
                   onClick = { login_buttonClick }>
                  접속
              </div>
          </div>
      </div>
      </>
  )
}

export default Home
