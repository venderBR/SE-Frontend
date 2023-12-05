import React from 'react';
import logo from './logo.svg';
import './App.css';
import { BrowserRouter, Link, Route, Router, Routes } from "react-router-dom";
import UserRequest from "./page/request/UserRequest";
import Request from './page/requested/Request';
import Donate from './page/Donate/Admindonate';

function App() {
  return (
    <div>
      <BrowserRouter>
      <Routes>
        <Route path='/' element={<UserRequest/>}/>
        <Route path='/request' element={<Request/>}/>
        <Route path='/donate' element={<Donate/>}/>
      </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
