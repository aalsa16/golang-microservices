import React, { useState, useEffect } from 'react';
import customFetch from '../api/axios';
import toast from 'react-hot-toast';
import Navbar from './Navbar';

const Home = () => {
  const [quote, setQuote] = useState('');
  const [author, setAuthor] = useState('');
  const [timestamp, setTimestamp] = useState('');

  const generateQuote = async () => {
    try {
      const response = await customFetch.get(`quotes/getQuote?uuid=${localStorage.getItem("uuid")}`);
      const { data } = response;

      setQuote(data.message.quote);
      setAuthor(data.message.author);
      setTimestamp(data.message.created_at);
    } catch (err) {
      toast.error('Error while generating quote!');
    }
  };

  return (
    <div>
      <Navbar />
      <div className='flex flex-col items-center justify-center min-h-screen'>
        <div className='bg-[#252525] rounded-xl text-white p-6 max-w-md w-full relative'>
          <div className='break-words mb-4'>
            <h1 className='text-center'>{quote}</h1>
          </div>
          <div className='flex justify-center'>
            <h1>{author}</h1>
          </div>
          <div className='flex justify-between items-center'>
            <button
              className='bg-blue-500 text-white px-4 py-2 rounded mt-4 hover:bg-blue-600 transition duration-300'
              onClick={generateQuote}
            >
              Generate New Quote
            </button>
            <span className='text-gray-400'>
              {new Date(timestamp).toLocaleString()}
            </span>
          </div>
        </div>
      </div>
    </div>
  );
};

export default Home;
