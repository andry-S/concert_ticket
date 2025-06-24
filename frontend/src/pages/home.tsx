import { useNavigate } from 'react-router-dom';
import { useEffect } from 'react';
import Swal from 'sweetalert2';
import { motion } from 'framer-motion'; // ✨ import motion

import '../styles/main.css';

interface Event {
  id: number;
  name: string;
  date: string;
  location: string;
  price: number;
  image: string;
}

const events: Event[] = [
  {
    id: 1,
    name: 'Coldplay World Tour',
    date: '15 Juli 2025',
    location: 'GBK Jakarta',
    price: 1750000,
    image: '/images/coldplay.jpg',
  },
  {
    id: 2,
    name: 'Java Jazz Festival',
    date: '21 Agustus 2025',
    location: 'JIExpo Kemayoran',
    price: 950000,
    image: '/images/jazz.jpg',
  },
  {
    id: 3,
    name: 'LANY - A Beautiful Blur Tour',
    date: '10 September 2025',
    location: 'ICE BSD',
    price: 1200000,
    image: '/images/lany.jpg',
  },
  {
    id: 4,
    name: 'Coldplay World Tour',
    date: '15 Juli 2025',
    location: 'GBK Jakarta',
    price: 1750000,
    image: '/images/coldplay.jpg',
  },
  {
    id: 5,
    name: 'Java Jazz Festival',
    date: '21 Agustus 2025',
    location: 'JIExpo Kemayoran',
    price: 950000,
    image: '/images/jazz.jpg',
  },
  {
    id: 6,
    name: 'LANY - A Beautiful Blur Tour',
    date: '10 September 2025',
    location: 'ICE BSD',
    price: 1200000,
    image: '/images/lany.jpg',
  },
  {
    id: 7,
    name: 'Coldplay World Tour',
    date: '15 Juli 2025',
    location: 'GBK Jakarta',
    price: 1750000,
    image: '/images/coldplay.jpg',
  },
  {
    id: 8,
    name: 'Java Jazz Festival',
    date: '21 Agustus 2025',
    location: 'JIExpo Kemayoran',
    price: 950000,
    image: '/images/jazz.jpg',
  },
  {
    id: 9,
    name: 'LANY - A Beautiful Blur Tour',
    date: '10 September 2025',
    location: 'ICE BSD',
    price: 1200000,
    image: '/images/lany.jpg',
  },
  {
    id: 10,
    name: 'LANY - A Beautiful Blur Tour',
    date: '10 September 2025',
    location: 'ICE BSD',
    price: 1200000,
    image: '/images/lany.jpg',
  },
];

const Home = () => {
  const navigate = useNavigate();

  const handleBook = (id: number) => {
    const userId = localStorage.getItem('user_id');
    if (!userId) {
      Swal.fire({
        title: 'Kamu belum login!',
        text: 'Silakan login dulu untuk memesan tiket konser.',
        icon: 'warning',
        confirmButtonText: 'Login Sekarang',
        confirmButtonColor: '#2563eb',
      }).then((result) => {
        if (result.isConfirmed) {
          navigate('/login');
        }
      });
    } else {
      navigate(`/book/${id}`);
    }
  };

  return (
    <motion.div
      className="home-container"
      initial={{ opacity: 0, y: 20 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -20 }}
      transition={{ duration: 0.4 }}
    >
      <div className="banner-section">
        <img className="banner-image" src="/images/banner.jpg" alt="Banner" />
        <div className="banner-text">
          <h1>Cari Tiket Konser Favoritmu!</h1>
          <p>Booking mudah, cepat, dan terpercaya 🎵</p>
        </div>
      </div>

      <div className="section-wrapper">
        <h2 className="section-title">Konser Populer</h2>
        <div className="event-scroll">
          {events.map((e) => (
            <div key={e.id} className="event-card" onClick={() => handleBook(e.id)}>
              <img src={e.image} alt={e.name} className="event-image" />
              <div className="event-info">
                <h3>{e.name}</h3>
                <p>{e.location}</p>
                <p className="price">Rp {e.price.toLocaleString()}</p>
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="section-wrapper">
        <h2 className="section-title">Semua Konser</h2>
        <div className="event-grid">
          {events.map((e) => (
            <div key={e.id} className="event-grid-card">
              <img src={e.image} alt={e.name} className="event-image" />
              <div className="event-info">
                <h3>{e.name}</h3>
                <p>{e.date}</p>
                <p>{e.location}</p>
                <p className="price">Rp {e.price.toLocaleString()}</p>
                <button className="book-button" onClick={() => handleBook(e.id)}>Pesan Sekarang</button>
              </div>
            </div>
          ))}
        </div>
      </div>
    </motion.div>
  );
};

export default Home;
