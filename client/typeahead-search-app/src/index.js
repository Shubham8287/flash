import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import { BrowserRouter,Routes,Route } from "react-router-dom";
import reportWebVitals from './reportWebVitals';
import NavbarTop from './components/Navbar';
import MainContent  from './components/MainContent';
import About from './components/About';

ReactDOM.render(
  <React.StrictMode>
    <NavbarTop />
    <BrowserRouter>
    <Routes>
        <Route path="/" element={<MainContent />} />
        <Route path="about" element={<About />} />
    </Routes>
  </BrowserRouter>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
