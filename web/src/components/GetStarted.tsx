import React from 'react'
import { redirect } from 'react-router-dom'

const GetStarted = () => {
    return (
        <div>
            <div className='flex flex-col gap-4 justify-center items-center h-screen'>
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer'>
                    <h1 className='m-auto'>Sign up</h1>
                </div>
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer'>
                    <h1 className='m-auto'>Log in</h1>
                </div>
            </div>
        </div>
    )
}

export default GetStarted