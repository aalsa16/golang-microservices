import React, { useState, useContext } from 'react'
import AuthContext from '../context/AuthProvider';
import { FiArrowLeft } from "react-icons/fi";
import { useNavigate } from 'react-router-dom';
import customFetch from '../api/axios';
import toast from 'react-hot-toast';

const Signup = () => {
    let navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');

    const handleBack = () => {
        navigate('/getstarted');
    };

    const handleSignUp = async (e: any) => {
        e.preventDefault();

        if (email.length !== 0 && password.length !== 0) {
            try {
                await customFetch.post('auth/signup', {
                    email: email,
                    password: password
                }, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
    
                toast.success('Successfully signed up! Please login.');
            } catch (err: any) {
                toast.error('Error while signing up');
            }
        } else {
            toast.error('Email or password cannot be blank!');
        }
    };

    return (
        <div>
            <div className='flex flex-col gap-4 justify-center items-center h-screen'>
                <FiArrowLeft className='absolute top-0 left-0 m-5 cursor-pointer' size={25} color='#FFF' onClick={handleBack} />
                <input type="text" placeholder='Enter email' className='w-52 h-12 indent-2 rounded-md' onChange={(e: any) => setEmail(e.target.value)} value={email} />
                <input type="text" placeholder='Enter password' className='w-52 h-12 indent-2 rounded-md' onChange={(e: any) => setPassword(e.target.value)} value={password} />
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer' onClick={handleSignUp}>
                    <h1 className='m-auto'>Sign up</h1>
                </div>
            </div>
        </div>
    )
}

export default Signup