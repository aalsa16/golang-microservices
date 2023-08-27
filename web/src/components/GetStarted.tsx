import React from 'react'
import { redirect, useNavigate } from 'react-router-dom'

const GetStarted = () => {
    let navigate = useNavigate();

    const goToSignUp = () => {
        navigate('/signup');
    };

    const goToLogin = () => {
        navigate('/login');
    };

    return (
        <div>
            <div className='flex flex-col gap-4 justify-center items-center h-screen'>
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer' onClick={goToSignUp}>
                    <h1 className='m-auto'>Sign up</h1>
                </div>
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer' onClick={goToLogin}>
                    <h1 className='m-auto'>Log in</h1>
                </div>
            </div>
        </div>
    )
}

export default GetStarted