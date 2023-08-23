import React from 'react'
import { FiArrowLeft } from "react-icons/fi";

const Signup = () => {
    return (
        <div>
            <div className='flex flex-col gap-4 justify-center items-center h-screen'>
                <FiArrowLeft className='absolute top-0 left-0 m-5 cursor-pointer' size={25} color='#FFF' />
                <input type="text" placeholder='Enter email' className='w-52 h-12 indent-2 rounded-md' />
                <input type="text" placeholder='Enter password' className='w-52 h-12 indent-2 rounded-md' />
                <div className='flex bg-white w-52 h-12 rounded-md cursor-pointer'>
                    <h1 className='m-auto'>Sign up</h1>
                </div>
            </div>
        </div>
    )
}

export default Signup