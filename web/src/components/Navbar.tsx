import React from 'react'
import { FiUser } from "react-icons/fi";

const Navbar = () => {

  return (
    <div className='p-5 top-0'>
        <div className='flex justify-between'>
            <nav className="justify-center space-x-4 m-auto">
                {[
                    ['Home', '/home'],
                    ['History', '/history'],
                ].map(([title, url]) => (
                    <a href={url} className="rounded-lg px-3 py-2 text-white font-medium hover:bg-slate-100 hover:text-slate-900">{title}</a>
                ))}
            </nav>
            <div className='p-3 bg-white rounded-full cursor-pointer'>
                <FiUser size={25} />
            </div>
        </div>
    </div>
  )
}

export default Navbar