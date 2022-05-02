import styles from '../styles/Home.module.css'
import Todo_st from '../styles/Todo.module.css'
import schedule_st from '../styles/Todo.module.css'
import React, { useEffect, useState } from "react";
import { useCookies } from "react-cookie";
import axios from "axios";

export default function Home() {
    const [cookies, setCookie] = useCookies();

    useEffect(() => {
        (async() => {
            const data = await axios.post("http://localhost/user/check",{}, { withCredentials: true})
            console.log(data,cookies['Token'])
        })();
    }, []);

    return (
        <div className={styles.main}>
            <div className={styles.top_menu}><h1 className={styles.title_top}>Work Efficiency</h1></div>
            <div className={styles.down_menu}>
                <div className={styles.item}>
                    <div className={schedule_st.main}>
                        <h1 className={schedule_st.title}>일정</h1>

                    </div>
                </div>
                <div className={styles.item}>
                    <div className={Todo_st.main}>
                        <h1 className={Todo_st.title}>Todo</h1>
                    </div>
                </div>
            </div>
        </div>
    )
}