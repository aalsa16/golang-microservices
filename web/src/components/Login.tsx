import React, { useState } from 'react'
import { FiArrowLeft } from 'react-icons/fi'
import { useNavigate } from 'react-router-dom';
import toast from 'react-hot-toast';
import customFetch from '../api/axios';
import { useAuth } from '../context/AuthProvider';

const Login = () => {
    let navigate = useNavigate();
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    // @ts-ignore
    const { setToken } = useAuth();

    const handleBack = () => {
        navigate('/getstarted');
    };

    const handleSignIn = async (e: any) => {
        e.preventDefault();

        if (email.length !== 0 && password.length !== 0) {
            try {
                const response = await customFetch.post('auth/signin', {
                    email: email,
                    password: password
                }, {
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });

                setToken(response.data.message.access_token);

                localStorage.setItem("refresh_token", response.data.message.refresh_token);
                localStorage.setItem("uuid", response.data.message.user.uuid);
    
                toast.success('Successfully logged in!');

                navigate('/');
            } catch (err: any) {
                console.log(err);
                toast.error('Error while signing in!');
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
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer' onClick={handleSignIn}>
                    <h1 className='m-auto'>Log in</h1>
                </div>
            </div>
        </div>
    )
}

export default Login