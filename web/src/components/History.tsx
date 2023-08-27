import React, { useState, useEffect } from 'react';
import customFetch from '../api/axios';
import toast from 'react-hot-toast';
import Navbar from './Navbar';

const History = () => {
  const [quotes, setQuotes] = useState([]);

  useEffect(() => {
    const fetchQuotes = async () => {
      try {
        const response = await customFetch.get(`quotes/getAllQuotes?uuid=${localStorage.getItem("uuid")}`);
        const { data } = response;
        setQuotes(data.message.reverse());
      } catch (err) {
        toast.error('Error while fetching quotes!');
      }
    };

    window.onload = () => {
      fetchQuotes();
    };
  }, []);

  return (
    <div>
      <Navbar />
      <div className='flex flex-col items-center justify-center min-h-screen'>
        {quotes.map((quote: any, index: any) => (
          <div key={index} className='bg-[#252525] rounded-xl text-white p-6 max-w-md w-full mb-4'>
            <div className='break-words mb-4'>
              <h1 className='text-center'>{quote.quote}</h1>
            </div>
            <div className='flex justify-between'>
              <h1>{quote.author}</h1>
              <h1>{new Date(quote.created_at).toLocaleString()}</h1>
            </div>
          </div>
        ))}
      </div>
    </div>
  );
};

export default History;
