import { useParams, useNavigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import axios from 'axios';

interface Event {
  id: number;
  name: string;
  date: string;
  location: string;
  price: number;
  image: string;
}

const Book = () => {
  const { id } = useParams<{ id: string }>();
  const navigate = useNavigate();
  const [event, setEvent] = useState<Event | null>(null);
  const [quantity, setQuantity] = useState<number>(1);
  const [loading, setLoading] = useState<boolean>(false);

  // Dummy data sebagai pengganti API /events/:id
  const dummyEvents: Event[] = [
    {
      id: 1,
      name: 'Coldplay Music of the Spheres Tour',
      date: '2025-07-15',
      location: 'Stadion Utama GBK, Jakarta',
      price: 1750000,
      image: '/images/coldplay.jpg',
    },
    {
      id: 2,
      name: 'Java Jazz Festival',
      date: '2025-08-21',
      location: 'JIExpo Kemayoran',
      price: 950000,
      image: '/images/jazz.jpg',
    },
    {
      id: 3,
      name: 'LANY A Beautiful Blur Tour',
      date: '2025-09-10',
      location: 'ICE BSD City',
      price: 1200000,
      image: '/images/lany.jpg',
    },
    {
      id: 4,
      name: 'Coldplay Music of the Spheres Tour',
      date: '2025-07-15',
      location: 'Stadion Utama GBK, Jakarta',
      price: 1750000,
      image: '/images/coldplay.jpg',
    },
    {
      id: 5,
      name: 'Java Jazz Festival',
      date: '2025-08-21',
      location: 'JIExpo Kemayoran',
      price: 950000,
      image: '/images/jazz.jpg',
    },
    {
      id: 6,
      name: 'LANY A Beautiful Blur Tour',
      date: '2025-09-10',
      location: 'ICE BSD City',
      price: 1200000,
      image: '/images/lany.jpg',
    },
    {
      id: 7,
      name: 'Coldplay Music of the Spheres Tour',
      date: '2025-07-15',
      location: 'Stadion Utama GBK, Jakarta',
      price: 1750000,
      image: '/images/coldplay.jpg',
    },
    {
      id: 8,
      name: 'Java Jazz Festival',
      date: '2025-08-21',
      location: 'JIExpo Kemayoran',
      price: 950000,
      image: '/images/jazz.jpg',
    },
    {
      id: 9,
      name: 'LANY A Beautiful Blur Tour',
      date: '2025-09-10',
      location: 'ICE BSD City',
      price: 1200000,
      image: '/images/lany.jpg',
    },
    {
      id: 10,
      name: 'LANY A Beautiful Blur Tour',
      date: '2025-09-10',
      location: 'ICE BSD City',
      price: 1200000,
      image: '/images/lany.jpg',
    },
  ];

  useEffect(() => {
    const selected = dummyEvents.find((e) => e.id === Number(id));
    setEvent(selected || null);
  }, [id]);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!event || quantity < 1) return;

    const userId = localStorage.getItem('user_id');
    if (!userId) {
      alert('Silakan login terlebih dahulu');
      navigate('/login');
      return;
    }

    const total_price = event.price * quantity;

    try {
      setLoading(true);
      await axios.post('http://localhost:8081/bookings', {
        user_id: Number(userId),
        event_id: event.id,
        ticket_quantity: quantity,
        total_price,
        status: 'menunggu',
      });
      alert('Tiket berhasil dipesan!');
      navigate('/dashboard');
    } catch (error) {
      alert('Gagal memesan tiket');
    } finally {
      setLoading(false);
    }
  };

  if (!event) return <p>Event tidak ditemukan</p>;

  return (
    <div className="book-page" style={{ padding: '30px' }}>
      <div className="book-card" style={{ maxWidth: 600, margin: 'auto', background: '#fff', padding: 20, borderRadius: 10 }}>
        <img src={event.image} alt={event.name} style={{ width: '100%', borderRadius: 10 }} />
        <h2 style={{ marginTop: 20 }}>{event.name}</h2>
        <p>{event.date}</p>
        <p>{event.location}</p>
        <p>
          Harga Tiket: <strong>Rp {event.price.toLocaleString()}</strong>
        </p>

        <form onSubmit={handleSubmit} style={{ marginTop: 20 }}>
          <label>Jumlah Tiket:</label>
          <input
            type="number"
            value={quantity}
            min={1}
            onChange={(e) => setQuantity(Number(e.target.value))}
            style={{ width: '100%', padding: 8, marginTop: 5, marginBottom: 15 }}
            required
          />

          <p>Total Harga: <strong>Rp {(event.price * quantity).toLocaleString()}</strong></p>

          <button
            type="submit"
            disabled={loading}
            style={{
              padding: 10,
              background: '#2563eb',
              color: 'white',
              border: 'none',
              borderRadius: 6,
              width: '100%',
              cursor: 'pointer',
              fontWeight: 600,
            }}
          >
            {loading ? 'Memesan...' : 'Pesan Tiket'}
          </button>
        </form>
      </div>
    </div>
  );
};

export default Book;
