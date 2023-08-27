import React from 'react'
import { FiUser } from "react-icons/fi";
import { useAuth } from '../context/AuthProvider';
import { useNavigate } from 'react-router-dom';

const Navbar = () => {
    // @ts-ignore
    const { setToken } = useAuth();
    const navigate = useNavigate();
  
    const handleLogout = () => {
      setToken();

      localStorage.removeItem('uuid');
      localStorage.removeItem('refresh_token');

      setTimeout(() => {
        navigate("/getstarted", { replace: true });
      }, 1000);
    };

  return (
    <div className='p-5 top-0'>
        <div className='flex justify-between'>
            <nav className="space-x-4">
                {[
                    ['Home', '/'],
                    ['History', '/history'],
                ].map(([title, url]) => (
                    <a href={url} className="rounded-lg px-3 py-2 text-white font-medium hover:bg-slate-100 hover:text-slate-900">{title}</a>
                ))}
            </nav>
            <div className='p-3 bg-white rounded-full cursor-pointer' onClick={handleLogout}>
                <h1>Logout</h1>
            </div>
        </div>
    </div>
  )
}

export default Navbar