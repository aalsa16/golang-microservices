import React from 'react';
import ReactDOM from "react-dom/client";
import { BrowserRouter, Route } from "react-router-dom";
import Navbar from './components/Navbar';
import GetStarted from './components/GetStarted';
import Signup from './components/Signup';
import Login from './components/Login';
import { Toaster } from 'react-hot-toast';
import Home from './components/Home';
import AuthProvider from './context/AuthProvider';
import Routes from './components/Routes';

function App() {
  return (
    <AuthProvider>
      <Toaster
        position="top-right"
        reverseOrder={false}
      />
      <Routes />
    </AuthProvider>
    // <BrowserRouter>
    //   <Toaster
    //     position="top-right"
    //     reverseOrder={false}
    //   />
    //   <Routes>
    //     <Route path="/" element={<GetStarted />} />
    //     <Route path="signup" element={<Signup />} />
    //     <Route path="login" element={<Login />} />
    //     <Route path="home" element={<Home />} />
    //   </Routes>
    // </BrowserRouter>
  );
}

export default App;
