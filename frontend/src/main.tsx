import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import Inbox from './Inbox'
import {BrowserRouter, Route, Routes} from "react-router-dom";
import Welcome from "./Welcome";
import Login from "./Login";
import Init from "./Init";

const container = document.getElementById('root')

const root = createRoot(container!)

root.render(
    <React.StrictMode>
        <BrowserRouter>
            <Routes>
                <Route path="/" element={<Init/>}/>
                <Route path="/welcome" element={<Welcome/>}/>
                <Route path="/login" element={<Login/>}/>
                <Route path="/inbox" element={<Inbox/>}/>
            </Routes>
        </BrowserRouter>
    </React.StrictMode>
)
