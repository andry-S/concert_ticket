import { useEffect, useState } from 'react';
import axios from 'axios';
import { motion } from 'framer-motion';
import '../styles/main.css';

interface Booking {
  id: number;
  event_id: number;
  ticket_quantity: number;
  total_price: number;
  status: string;
  created_at: string;
}

interface User {
  id: number;
  name?: string;
  email: string;
}

const Dashboard = () => {
  const [bookings, setBookings] = useState<Booking[]>([]);
  const [user, setUser] = useState<User>({
    id: Number(localStorage.getItem('user_id')),
    email: localStorage.getItem('user_email') || '-',
  });

  useEffect(() => {
    const userId = localStorage.getItem('user_id');
    if (!userId) return;

    axios
      .get(`http://localhost:8081/bookings/user/${userId}`)
      .then((res) => {
        setBookings(res.data.bookings);
      })
      .catch((err) => {
        console.error('Gagal fetch data booking:', err);
      });
  }, []);

  return (
    <motion.div
      className="dashboard-traveloka"
      initial={{ opacity: 0, y: 30 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -30 }}
      transition={{ duration: 0.4 }}
    >
      <div className="header">
        <h2>🎟️ Tiket Saya</h2>
        <p className="user">📧 {user.email}</p>
      </div>

      <div className="booking-list">
        {bookings.length === 0 ? (
          <p style={{ textAlign: 'center', color: '#999' }}>Belum ada data booking</p>
        ) : (
          bookings.map((b) => (
            <div className="booking-card" key={b.id}>
              <div className="card-header">
                <h3>Event #{b.event_id}</h3>
                <span className={`status ${b.status.toLowerCase()}`}>
                  {b.status === 'SUKSES' && '✅ Berhasil'}
                  {b.status === 'MENUNGGU' && '⏳ Menunggu'}
                  {b.status === 'GAGAL' && '❌ Gagal'}
                  {!['SUKSES', 'MENUNGGU', 'GAGAL'].includes(b.status) && b.status}
                </span>
              </div>
              <div className="card-body">
                <p><strong>Tiket:</strong> {b.ticket_quantity}</p>
                <p><strong>Total:</strong> Rp {b.total_price.toLocaleString()}</p>
                <p><strong>Tanggal:</strong> {new Date(b.created_at).toLocaleDateString()}</p>
                <p><strong>ID Booking:</strong> #{b.id}</p>
              </div>
            </div>
          ))
        )}
      </div>
    </motion.div>
  );
};

export default Dashboard;
